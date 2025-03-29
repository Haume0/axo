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
	mail.Send(mailadr, "Test Mail", "This is a test mail from Axo", false)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Message sent!"}`))
}
