package main

import (
	"fmt"
	"log"
	"os"

	vault "github.com/hashicorp/vault/api"
)

func main() {
	// Инициализация клиента Vault
	config := vault.DefaultConfig()
	config.Address = "http://127.0.0.1:8200" // Замените на ваш адрес Vault

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("Error creating Vault client: %v", err)
	}

	// Автотокенизация (используйте метод, соответствующий вашей конфигурации)
	client.SetToken(os.Getenv("VAULT_TOKEN"))

	// Чтение секрета из Vault
	secret, err := client.Logical().Read("secret/data/myapp/credentials") // Замените на ваш путь
	if err != nil {
		log.Fatalf("Error reading secret: %v", err)
	}

	// Обработка данных секрета
	if secret == nil {
		log.Fatal("No secret found at the specified path")
	}

	// Извлечение данных из секрета
	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		log.Fatalf("Invalid secret data at the specified path")
	}

	username, ok := data["username"].(string)
	if !ok {
		log.Fatalf("Username not found in secret data")
	}

	password, ok := data["password"].(string)
	if !ok {
		log.Fatalf("Password not found in secret data")
	}

	// Использование credentials в приложении
	fmt.Printf("Username: %s\n", username)
	fmt.Printf("Password: %s\n", password)

	// Пример: использование credentials для подключения к базе данных (MySQL)
	/*
		dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/dbname", username, password)
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			log.Fatalf("Error connecting to the database: %v", err)
		}

		defer db.Close()

		// Теперь мы можем выполнять запросы к базе данных с использованием полученных credentials
	*/
}
