package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mattn/go-tty"

	"github.com/joho/godotenv"
)

func InitDotenv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("❌ .env file not found!")
		fmt.Println("👋 Wanna create a .env file? (y/n)")
		tty, err := tty.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer tty.Close()
		char, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}

		if char == 'y' || char == 'Y' {
			fmt.Println("Creating default .env file...")
			os.WriteFile(".env", []byte(defaultDotenv), 0644)
			fmt.Print("\033[H\033[2J")
			fmt.Println("✅ Default .env file created")
			fmt.Println("⚙️ Please edit values in the .env file")
		} else {
			fmt.Print("\033[H\033[2J")
			fmt.Println("\n👌 Continue without .env file!")
			os.WriteFile(".env", []byte("# For getting '👋 Wanna create a .env file? (y/n)' question remove the .env file."), 0644)
			return
		}
	} else {
		fmt.Print("\033[H\033[2J")
		fmt.Println("✅ .env file founded.")
	}
}

const defaultDotenv = `
# .env
# This is a default .env file
# You don't have and .env file? No problem! We're creating one for you.
# You can change the values as you like.

# Server Values
PORT=3000
HOST=0.0.0.0

# TLS Values *If you don't use TLS, you can remove these values or leave them empty.
# If you're using TLS, you need to set the values for these variables.
# CERT_FILE=cert.pem
# KEY_FILE=key.pem

# Database Values (!) Warning this scaffold using postgres as default database.
# If you don't use postgres, you can free to change codes in axo/database folder.
# DB_HOST=localhost
# DB_PORT=5432
# DB_USER=root
# DB_PASSWORD=123
# DB_NAME=app
# DB_SSLMODE=disable
# DB_TIMEZONE=Europe/Istanbul

# JWT Values *This project using JWT for authentication.
JWT_SECRET=secret
JWT_EXPIRATION=1h
`
