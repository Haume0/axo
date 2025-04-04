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
		fmt.Println("👋 Wanna create a .env file? (Y/n)")
		tty, err := tty.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer tty.Close()
		char, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}

		if char == 'y' || char == 'Y' || char == '\n' || char == '\r' {
			fmt.Println("Creating default .env file...")
			os.WriteFile(".env", []byte(defaultDotenv), 0644)
			fmt.Print("\033[H\033[2J")
			fmt.Println("✅ Default .env file created")
			fmt.Println("⚙️ Please edit values in the .env file")
			fmt.Println("👋 Bye!")
			os.Exit(0)
		} else {
			fmt.Print("\033[H\033[2J")
			fmt.Println("\n👌 Continue without .env file!")
			os.WriteFile(".env", []byte("# For getting '👋 Wanna create a .env file? (y/n)' question remove the .env file."), 0644)
			return
		}
	} else {
		// fmt.Print("\033[H\033[2J")
		fmt.Println("✅ .env file found")
	}
}

const defaultDotenv = `
# .env
# This is a default .env file
# You don't have and .env file? No problem! We're creating one for you.
# You can change the values as you like.

# Server Values
PORT=3000
HOST=localhost

# TLS Values * If you don't use TLS, you can remove these values or leave them empty.
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

# JWT Values * This project using JWT for authentication.
JWT_SECRET=secret
JWT_EXPIRATION=1h

# Email Values * This project using SMTP for sending emails.
# EMAIL_FROM=your-email@example.com
# EMAIL_PASSWORD=your-email-password
# SMTP_HOST=smtp.example.com
# SMTP_PORT=587
# SMTP_NOSSL=false

# Image Optimization Values
USE_BREAKPOINTS=true
BREAKPOINT_WIDTHS=440,640,800,960,1120,1280,1440
BREAKPOINT_HEIGHTS=440,640,800,960,1120,1280,1440
CACHE_DIR=memory
MAX_IMAGE_WIDTH=1440
MAX_IMAGE_HEIGHT=990
CACHE_EXPIRATION=14400
MAX_CACHE_SIZE=1024
`
