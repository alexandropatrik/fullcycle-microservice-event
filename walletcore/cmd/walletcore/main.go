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
	fmt.Println("0")
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka-fc:29092",
		"group.id":          "wallet",
	}
	kafkaProducer := kafka.NewKafkaProducer(&configMap)

	fmt.Println("1")

	eventDispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTransactionCreated()
	eventDispatcher.Register("TransactionCreated", handlers.NewTransactionCreatedKafkaHandler(kafkaProducer))
	balanceUpdatedEvent := event.NewBalanceUpdated()
	eventDispatcher.Register("BalanceUpdated", handlers.NewBalanceUpdatedKafkaHandler(kafkaProducer))

	fmt.Println("2")

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

	fmt.Println("3")

	createClientUseCase := create_client.NewCreateClientUseCase(clientDb)
	createAccountUseCase := create_account.NewCreateClientUseCase(accountDb, clientDb)
	createTransactionUseCase := create_transaction.NewCreateTransactionUseCase(uow, eventDispatcher, transactionCreatedEvent, balanceUpdatedEvent)

	fmt.Println("4")

	webserver := webserver.NewWebServer(":3000")
	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	fmt.Println("5")

	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)

	fmt.Println("6")

	webserver.Start()

	fmt.Println("7")
}
