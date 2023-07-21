package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID        string
	Name      string
	Email     string
	Accounts  []*Account
	CreatedAt time.Time
	UpdateAt  time.Time
}

func NewClient(name string, email string) (*Client, error) {
	client := &Client{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}

	err := client.Validate()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (c *Client) Validate() error {
	if c.Name == "" {
		return errors.New("Name is required")
	}
	if c.Email == "" {
		return errors.New("Email is required")
	}
	return nil
}

func (c *Client) Update(name string, email string) error {
	c.Name = name
	c.Email = email
	c.UpdateAt = time.Now()
	err := c.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) addAccount(account *Account) error {
	if account.Client.ID != c.ID {
		return errors.New("Account does not belong to client")
	}
	c.Accounts = append(c.Accounts, account)
	return nil
}
