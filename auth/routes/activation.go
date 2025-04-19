package routes

import (
	"axo/auth"
	"axo/axo"
	"net/http"
)

func Deavcivate(w http.ResponseWriter, r *http.Request) {
	user, err := auth.GetUserRequest(r)
	if err != nil {
		axo.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	if err := auth.DeactivateUser(user.ID); err != nil {
		axo.Error(w, "Failed to deactivate user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "User deactivated successfully"}`))
}

func Activate(w http.ResponseWriter, r *http.Request) {
	user, err := auth.GetUserRequest(r)
	if err != nil {
		axo.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	if err := auth.ActivateUser(user.ID); err != nil {
		axo.Error(w, "Failed to activate user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "User activated successfully"}`))
}
