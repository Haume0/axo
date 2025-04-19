package main

import (
	"axo/auth"
	auth_routes "axo/auth/routes"
	"axo/axo"
	"axo/axo/frontends"
	"axo/database"
	"axo/flags"
	"axo/img"
	"axo/middlewares"
	"axo/models"
	"axo/routes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// 🔒 Command line flags
	flags.Init()

	// 🔐 Getting the environment variables !! Dont put any print operation above .env initialization. !!
	InitDotenv()

	// 🏁 Initializations !! Please do not change the order of the initialization operations. !!
	database.Init()
	auth.Init()

	// 🏗️ Creating the router
	router := http.NewServeMux()
	site := http.NewServeMux()

	// ⚠️ Axo Rest API Routes ⚠️
	// 🎭 Auth Routes
	router.HandleFunc("POST /auth/register", auth_routes.Register)
	router.HandleFunc("POST /auth/login", auth_routes.Login)
	router.HandleFunc("/auth/logout", auth_routes.Logout)
	router.HandleFunc("/auth/refresh", auth_routes.Refresh)
	router.HandleFunc("/auth/verify", auth_routes.Verify)
	router.HandleFunc("/auth/reset-password", auth_routes.ResetPassword)

	//!DEV
	if !*flags.IsProduction {
		router.HandleFunc("GET /auth/users", func(w http.ResponseWriter, r *http.Request) {
			var users []models.User
			database.DB.Preload("Role").Find(&users)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(users)
		})
		router.HandleFunc("GET /tokens", func(w http.ResponseWriter, r *http.Request) {
			var del = r.URL.Query().Get("del")
			var tokens []models.RefreshToken
			database.DB.Find(&tokens)
			if del == "true" {
				for _, token := range tokens {
					database.DB.Delete(&token)
				}
			}
			json.NewEncoder(w).Encode(tokens)
		})
	}
	//!DEV
	// 🌐 Registering the routes
	router.HandleFunc("GET /error", routes.GetError)
	router.HandleFunc("GET /hello", routes.GetHello)
	//Mail test route
	router.HandleFunc("GET /testmail", routes.MailTest)
	//Demo Note App
	router.HandleFunc("GET /notes", routes.GetNotes)
	router.HandleFunc("POST /notes", routes.PostNote)
	router.HandleFunc("DELETE /notes", routes.DeleteNote)

	// 🌍 Serving the Single Page Application (SPA)
	frontends.ServeSPA(
		site, "/",
		"5173",
		"./site", "./site/dist",
		[2]string{
			"bun install",
			"bun dev",
		},
		[]string{
			"bun install",
			"bun run build",
		},
	)
	// 🌍 Serving the Multi Page Application (MPA)
	// frontends.ServeStatic(site, "/", "./site/dist",
	// 	[]string{ // Last argument is optional,
	// 		// if your static site needs'a build step please add build steps here
	// 		"./site", // Target folder
	// 		"bun install", //build command
	// 		"bun run build", //build command
	// 	})

	// 🏗️ Static File Server
	var staticPath = "static"
	router.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(staticPath))))

	// 🏙️ Image Optimization
	img.Init(router, "/image", staticPath, os.Getenv("ENABLE_IMAGE_OPTIMIZATION"))

	// ⚙️ Adding middlewares to router
	routerWithMiddlewares := middlewares.Logger(router)

	// 💢 Adding cors setup to router
	routerWithMiddlewares = middlewares.Cors(routerWithMiddlewares)

	// Combining router and site
	handler := http.NewServeMux()
	handler.Handle("/api/", http.StripPrefix("/api", routerWithMiddlewares))
	handler.Handle("/", site)

	// 🚀 Starting the server
	var port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fmt.Println("🪸 Axo is live! 🌊")
	fmt.Printf("👀 You can see it on:\n")
	for _, ip := range axo.HostIPs() {
		if os.Getenv("HOST") == "localhost" {
			fmt.Println("\033[1;90m🌐 Running on localhost! Set HOST=0.0.0.0 to publish on all IPs.\033[0m")
			fmt.Printf("\033[1;34mhttp://localhost:%v\033[0m\n", port)
			break
		} else {
			fmt.Printf("\033[1;34mhttp://%v:%v\033[0m\n", ip, port)
		}
	}
	http.ListenAndServe(os.Getenv("HOST")+":"+port, handler)

}
