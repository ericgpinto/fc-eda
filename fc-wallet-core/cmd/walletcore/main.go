package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/database"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/event"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/event/handler"
	createaccount "github.com.br/devfullcycle/fc-ms-wallet/internal/usecase/create_account"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/usecase/create_client"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/usecase/create_transaction"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/web"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/web/webserver"
	"github.com.br/devfullcycle/fc-ms-wallet/pkg/events"
	"github.com.br/devfullcycle/fc-ms-wallet/pkg/kafka"
	"github.com.br/devfullcycle/fc-ms-wallet/pkg/uow"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql-go", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	createTables((db))
	seedDatabase(db)

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}
	kafkaProducer := kafka.NewKafkaProducer(&configMap)

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("TransactionCreated", handler.NewTransactionCreatedKafkaHandler(kafkaProducer))
	eventDispatcher.Register("BalanceUpdated", handler.NewUpdateBalanceKafkaHandler(kafkaProducer))
	transactionCreatedEvent := event.NewTransactionCreated()
	balanceUpdatedEvent := event.NewBalanceUpdated()

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
	createTransactionUseCase := create_transaction.NewCreateTransactionUseCase(uow, eventDispatcher, transactionCreatedEvent, balanceUpdatedEvent)
	createClientUseCase := create_client.NewCreateClientUseCase(clientDb)
	createAccountUseCase := createaccount.NewCreateAccountUseCase(accountDb, clientDb)

	webserver := webserver.NewWebServer(":8080")

	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)

	fmt.Println("Server is running")
	webserver.Start()
}

func createTables(db *sql.DB) {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS clients (
			id VARCHAR(36) PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`,
		`CREATE TABLE IF NOT EXISTS accounts (
			id VARCHAR(36) PRIMARY KEY,
			client_id VARCHAR(36) NOT NULL,
			balance INT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (client_id) REFERENCES clients(id)
		);`,
		`CREATE TABLE IF NOT EXISTS transactions (
			id VARCHAR(36) PRIMARY KEY,
			account_id_from VARCHAR(36) NOT NULL,
			account_id_to VARCHAR(36) NOT NULL,
			amount INT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (account_id_from) REFERENCES accounts(id),
			FOREIGN KEY (account_id_to) REFERENCES accounts(id)
		);`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			log.Fatalf("Erro ao criar tabela: %v\nQuery: %s", err, query)
		}
	}

	fmt.Println("Tabelas criadas com sucesso!")
}

func seedDatabase(db *sql.DB) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM clients").Scan(&count)
	if err != nil {
		log.Fatalf("Erro ao verificar dados existentes: %v", err)
	}

	if count == 0 {
		fmt.Println("Populando o banco de dados...")

		_, err := db.Exec("INSERT INTO clients (id, name, email, created_at) VALUES (?, ?, ?, NOW())",
			"4d2cd6d2-7879-4656-8ef4-03257339e51c", "João Silva", "joao@email.com")
		if err != nil {
			log.Fatalf("Erro ao inserir cliente: %v", err)
		}
		
		_, err = db.Exec("INSERT INTO clients (id, name, email, created_at) VALUES (?, ?, ?, NOW())",
			"850b9f3b-3cf2-48ba-80c7-702bd2a11d6f", "Wesley Will", "wwill@email.com")
		if err != nil {
			log.Fatalf("Erro ao inserir conta: %v", err)
		}
		
		_, err = db.Exec("INSERT INTO accounts (id, client_id, balance, created_at) VALUES (?, ?, ?, NOW())",
			"1b94b998-1f92-4897-a5e2-24bde6685b5d", "850b9f3b-3cf2-48ba-80c7-702bd2a11d6f", 1000)
		if err != nil {
			log.Fatalf("Erro ao inserir conta: %v", err)
		}

		_, err = db.Exec("INSERT INTO accounts (id, client_id, balance, created_at) VALUES (?, ?, ?, NOW())",
			"bb835285-769c-439f-b1cb-a8788bdf8e72", "4d2cd6d2-7879-4656-8ef4-03257339e51c", 50)
		if err != nil {
			log.Fatalf("Erro ao inserir conta: %v", err)
		}


		fmt.Println("Banco de dados populado com sucesso!")
	} else {
		fmt.Println("Banco de dados já contém dados, pulando seed.")
	}
}
