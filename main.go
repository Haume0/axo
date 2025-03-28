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
🪸 Welcome to Axo 🌊
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
	site := http.NewServeMux()

	// 🌐 Registering the routes
	router.HandleFunc("GET /error", routes.GetError)
	router.HandleFunc("GET /hello", routes.GetHello)
	router.HandleFunc("GET /testmail", routes.MailTest)

	// 🌍 Serving the Single Page Application
	frontends.ServeSPA(site, "npm run dev", "5173", "./site", "./site/dist")

	// 🏙️ Image Optimization
	//?[1] Comment out if you don't want image optimization!
	// img.Init(router, "/image")
	//?[2] and open this.
	router.HandleFunc("/image", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Image optimization is disabled!"}`))
	})
	// 🏙️ Image Optimization

	// 🏗️ Static File Server
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

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
