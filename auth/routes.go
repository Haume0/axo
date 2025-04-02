package auth

import "net/http"

func LoginRoute(w http.ResponseWriter, r *http.Request)    {}
func RegisterRoute(w http.ResponseWriter, r *http.Request) {}
func VerifyRoute(w http.ResponseWriter, r *http.Request)   {}
func LogoutRoute(w http.ResponseWriter, r *http.Request)   {}
func RefreshRoute(w http.ResponseWriter, r *http.Request)  {}

// On mail verification i'm planing to hold verification token in lru cache.
