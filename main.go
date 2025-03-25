package main

import (
	"axo/axo"
	"axo/axo/frontends"
	"axo/database"
	"axo/middlewares"
	"axo/routes"
	"fmt"
	"net/http"
	"os"
)

/*
🪐 Welcome to Axo ✨
AxoScaffold is a Restful API scaffold for Go, built on top of stdlib and gorm.
It is designed to be simple, fast, and easy to use.
For more information, please visit: https://haume.me/axo

License: MIT
Copyright (c) 2025 Haume
It's not neccesary but i'll be greatful if you give me a star on GitHub and mention me in your project.
*/

func main() {
	// 🔐 Getting the environment variables !! Dont put any print operation above .env initialization. !!
	InitDotenv()

	// 🏁 Initializations !! Please do not change the order of the initialization operations. !!
	database.Init()

	// 🏗️ Creating the router
	router := http.NewServeMux()

	// 🌐 Registering the routes
	router.HandleFunc("GET /error", routes.GetError)
	router.HandleFunc("GET /hello", routes.GetHello)

	// 🌍 Serving the Single Page Application
	frontends.ServeSPA(router, "npm run dev", "5173", "./site", "./site/dist")

	// // 🏙️ Image Optimization
	// if os.Getenv("IMG_OPTIMIZE") == "true" {
	// 	img.Init()
	// 	router.HandleFunc("/image", img.Optimize)
	// }

	// 🏗️ Static File Server
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// ⚙️ Adding middlewares
	handler := middlewares.Logger(router)

	// 💢 Adding cors setup
	handler = middlewares.Cors(handler)

	// 🚀 Starting the server
	var port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fmt.Println("🪸 Axo is live! 🌊")
	fmt.Printf("👀 You can see it on:\n")
	for _, ip := range axo.HostIPs() {
		fmt.Printf("\033[1;34mhttp://%v:%v\033[0m\n", ip, port)
	}
	http.ListenAndServe(":"+port, handler)

}
