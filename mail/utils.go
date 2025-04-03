package mail

import (
	"fmt"
	"os"
)

func LoadTemplate(templateName string) (string, error) {
	// Load the template file
	templatePath := fmt.Sprintf("./mail/templates/%s.html", templateName)
	content, err := os.ReadFile(templatePath)
	if err != nil {
		return "We cannot load template! >:(", err
	}

	// Return the content as a string
	return string(content), nil
}
