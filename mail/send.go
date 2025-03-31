package mail

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"net/smtp"
	"os"
	"strconv"
)

// Send sends an email to a specified recipient.
//
// This function creates and sends an email with support for both plain text and HTML content.
// It handles different SMTP ports (587 for STARTTLS and 465 for SSL/TLS) and supports optional
// TLS verification skipping. The email configuration is loaded from environment variables.
//
// Environment Variables Required:
//   - EMAIL_FROM: Sender email address
//   - EMAIL_PASSWORD: Password for the sender email account
//   - SMTP_HOST: SMTP server hostname
//   - SMTP_PORT: SMTP server port (typically 587 for STARTTLS or 465 for SSL)
//   - SMTP_NOSSL: Set to "true" to skip TLS certificate verification (optional)
//
// Parameters:
//   - target string: The recipient's email address
//   - title string: The subject line for the email
//   - content string: The body content of the email
//   - html bool: If true, content will be sent as HTML; otherwise, as plain text
//
// Returns:
//   - error: An error if the sending process fails, or nil if successful
//
// Example:
//
//	err := mail.Send("recipient@example.com", "Hello", "<h1>Hello World</h1>", true)
//	if err != nil {
//	    log.Fatalf("Failed to send email: %v", err)
//	}
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
