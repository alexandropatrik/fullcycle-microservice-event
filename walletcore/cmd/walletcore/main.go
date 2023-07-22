package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/alexandropatrik/fc-ms-wallet/internal/database"
	"github.com/alexandropatrik/fc-ms-wallet/internal/event"
	"github.com/alexandropatrik/fc-ms-wallet/internal/event/handlers"
	"github.com/alexandropatrik/fc-ms-wallet/internal/usecase/create_account"
	"github.com/alexandropatrik/fc-ms-wallet/internal/usecase/create_client"
	"github.com/alexandropatrik/fc-ms-wallet/internal/usecase/create_transaction"
	"github.com/alexandropatrik/fc-ms-wallet/internal/web"
	"github.com/alexandropatrik/fc-ms-wallet/internal/web/webserver"
	"github.com/alexandropatrik/fc-ms-wallet/pkg/events"
	"github.com/alexandropatrik/fc-ms-wallet/pkg/kafka"
	"github.com/alexandropatrik/fc-ms-wallet/pkg/uow"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&multiStatements=true", "root", "root", "mysql", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// migrations
	err = executeMigrations(db)
	if err != nil {
		panic(err)
	}

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka-fc:29092",
		"group.id":          "wallet",
	}
	kafkaProducer := kafka.NewKafkaProducer(&configMap)

	eventDispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTransactionCreated()
	eventDispatcher.Register("TransactionCreated", handlers.NewTransactionCreatedKafkaHandler(kafkaProducer))
	balanceUpdatedEvent := event.NewBalanceUpdated()
	eventDispatcher.Register("BalanceUpdated", handlers.NewBalanceUpdatedKafkaHandler(kafkaProducer))

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)

	ctx := context.Background()
	uow := uow.NewUow(ctx, db)

	uow.Register("AccountDB", func(tx *sql.Tx) interface{} {
		return database.NewAccountDB(db)
	})

	uow.Register("TransactionDB", func(tx *sql.Tx) interface{} {
		return database.NewTransactionDB(db)
	})

	createClientUseCase := create_client.NewCreateClientUseCase(clientDb)
	createAccountUseCase := create_account.NewCreateClientUseCase(accountDb, clientDb)
	createTransactionUseCase := create_transaction.NewCreateTransactionUseCase(uow, eventDispatcher, transactionCreatedEvent, balanceUpdatedEvent)

	webserver := webserver.NewWebServer(":3000")
	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)

	webserver.Start()
}

func executeMigrations(db *sql.DB) error {
	db.Exec("CREATE TABLE IF NOT EXISTS clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("CREATE TABLE IF NOT EXISTS accounts (id varchar(255), client_id varchar(255), balance int, created_at date)")
	db.Exec("CREATE TABLE IF NOT EXISTS transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount int, created_at date)")
	// se nao tem registros faz o insert
	stmt, err := db.Prepare("SELECT COUNT(*) FROM accounts ")
	if err != nil {
		return err
	}
	defer stmt.Close()

	var count int = 0
	row := stmt.QueryRow()
	err = row.Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		db.Exec("INSERT INTO clients (id, name, email, created_at) VALUES ('4dbd4406-d6cd-4e27-88bd-a593b7ab5faa', 'JOHN DOE', 'jo@jo.com', current_date)")
		db.Exec("INSERT INTO clients (id, name, email, created_at) VALUES ('a64fa165-a9e0-4bad-ac83-ea3d6fe2aa39', 'JANE DOE', 'ja@ja.com', current_date)")
		// john's account
		db.Exec("INSERT INTO accounts (id, client_id, balance, created_at) VALUES ('a8d715f8-9ba4-42d8-975a-ac878a3f6a6d', '4dbd4406-d6cd-4e27-88bd-a593b7ab5faa', 1000, current_date)")
		// jane's account
		db.Exec("INSERT INTO accounts (id, client_id, balance, created_at) VALUES ('1e4b9007-9537-46cf-b5d2-19973b5e7711', 'a64fa165-a9e0-4bad-ac83-ea3d6fe2aa39', 1000, current_date)")
	}

	return nil
}
