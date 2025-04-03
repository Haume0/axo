package routes

import (
	"axo/axo"
	"axo/mail"
	"net/http"
)

func MailTest(w http.ResponseWriter, r *http.Request) {
	var mailadr string = r.URL.Query().Get("mail")
	if mailadr == "" {
		axo.Error(w, "Mail is required", http.StatusBadRequest)
		return
	}
	template, err := mail.LoadTemplate("verification_code")
	if err != nil {
		println(err.Error())
		axo.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	template = axo.MultiReplace(template, map[string]string{
		"{{.login_code}}":  "test-test-test-test",
		"{{.base_url}}":    "http://localhost:3000",
		"{{.title}}":       "Welcome to Axo!",
		"{{.description}}": "This is a test email to verify your account.",
		"{{.sub_text}}":    "If you did not request this, please ignore this email.",
		"{{.warning}}":     "Do not share this code with anyone.",
	})
	mail.Send(mailadr, "Welcome to Axo!", template, true)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Message sent!"}`))
}
