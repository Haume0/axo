package main

import (
	"axo/auth"
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

/*
🪸 Welcome to Axo 🌊
Axo is a Restful API scaffold for Go, built on top of stdlib and gorm.
It is designed to be simple, fast, and easy to use.
For more information, please visit: https://haume.me/axo

License: MIT
Copyright (c) 2025 Haume
It's not neccesary but i'll be greatful if you give me a star on GitHub and mention me in your project.
*/

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
	router.HandleFunc("POST /auth/register", auth.RegisterRoute)
	router.HandleFunc("POST /auth/login", auth.LoginRoute)
	//!DEV
	if !*flags.IsProduction {
		router.HandleFunc("GET /auth/users", func(w http.ResponseWriter, r *http.Request) {
			var users []models.User
			database.DB.Preload("Role").Find(&users)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(users)
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
	img.Init(router, "/image", staticPath, "enabled")

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
