package mail

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"net/smtp"
	"os"
	"strconv"
)

// Send sends an email to the target with the title and content
func Send(target string, title string, content string, html bool) error {
	from := os.Getenv("EMAIL_FROM")
	password := os.Getenv("EMAIL_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	skipTLSVerify := os.Getenv("SMTP_NOSSL") == "true" // Yeni ayar

	if from == "" || password == "" || smtpHost == "" || smtpPort == "" {
		return fmt.Errorf("missing required environment variables")
	}

	to := []string{target}

	// Create MIME headers
	headers := make(map[string]string)
	headers["From"] = from
	headers["To"] = target
	if title != "" {
		encodedTitle := base64.StdEncoding.EncodeToString([]byte(title))
		headers["Subject"] = fmt.Sprintf("=?UTF-8?B?%s?=", encodedTitle)
	}
	headers["MIME-Version"] = "1.0"
	if html {
		headers["Content-Type"] = "text/html; charset=UTF-8"
	} else {
		headers["Content-Type"] = "text/plain; charset=UTF-8"
	}

	// Create the email body
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + content

	// Set up authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Define SMTP server and port
	serverAddr := fmt.Sprintf("%s:%s", smtpHost, smtpPort)

	port, err := strconv.Atoi(smtpPort)
	if err != nil {
		return fmt.Errorf("invalid SMTP port: %w", err)
	}

	tlsConfig := &tls.Config{
		InsecureSkipVerify: skipTLSVerify, // Yeni ayarı burada kullanıyoruz
		ServerName:         smtpHost,
	}

	if port == 587 {
		// Connect to the SMTP server without TLS
		client, err := smtp.Dial(serverAddr)
		if err != nil {
			return fmt.Errorf("failed to connect to SMTP server: %w", err)
		}
		defer client.Quit()

		// Upgrade to STARTTLS if possible
		if err = client.StartTLS(tlsConfig); err != nil {
			return fmt.Errorf("failed to start TLS: %w", err)
		}

		// Authenticate
		if err = client.Auth(auth); err != nil {
			return fmt.Errorf("failed to authenticate: %w", err)
		}

		// Set the sender and recipient
		if err = client.Mail(from); err != nil {
			return fmt.Errorf("failed to set sender: %w", err)
		}
		for _, recipient := range to {
			if err = client.Rcpt(recipient); err != nil {
				return fmt.Errorf("failed to set recipient: %w", err)
			}
		}

		// Send the email body
		wc, err := client.Data()
		if err != nil {
			return fmt.Errorf("failed to send email body: %w", err)
		}
		defer wc.Close()

		if _, err = wc.Write([]byte(message)); err != nil {
			return fmt.Errorf("failed to write email body: %w", err)
		}
	} else if port == 465 {
		// Connect to the SMTP server with direct TLS
		conn, err := tls.Dial("tcp", serverAddr, tlsConfig)
		if err != nil {
			return fmt.Errorf("failed to connect to SMTP server: %w", err)
		}
		defer conn.Close()

		client, err := smtp.NewClient(conn, smtpHost)
		if err != nil {
			return fmt.Errorf("failed to create SMTP client: %w", err)
		}
		defer client.Quit()

		// Authenticate
		if err = client.Auth(auth); err != nil {
			return fmt.Errorf("failed to authenticate: %w", err)
		}

		// Set the sender and recipient
		if err = client.Mail(from); err != nil {
			return fmt.Errorf("failed to set sender: %w", err)
		}
		for _, recipient := range to {
			if err = client.Rcpt(recipient); err != nil {
				return fmt.Errorf("failed to set recipient: %w", err)
			}
		}

		// Send the email body
		wc, err := client.Data()
		if err != nil {
			return fmt.Errorf("failed to send email body: %w", err)
		}
		defer wc.Close()

		if _, err = wc.Write([]byte(message)); err != nil {
			return fmt.Errorf("failed to write email body: %w", err)
		}
	}

	return nil
}
