package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/EleMint/CONTd/service/controllers"
)

func init() {
	ensureEnvironmentVariables()
}

func main() {
	log.Println("Starting CONTd")

	log.Println("Initializing controllers...")
	controllers.Init()
	log.Println("Controllers initialized")

	port := os.Getenv("PORT")
	host := fmt.Sprintf("localhost:%s", port)
	log.Printf("Server listening on '%s'...\n", host)
	log.Fatalf(
		"An error occurred while listening: %v\n",
		http.ListenAndServe(host, nil),
	)
}

func ensureEnvironmentVariables() {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		os.Setenv("ENVIRONMENT", "dev")
	}

	port := os.Getenv("PORT")
	if port == "" {
		os.Setenv("PORT", "8080")
	}

	access := os.Getenv("AWS_ACCESS_KEY")
	if access == "" {
		log.Fatalln("env AWS_ACCESS_KEY missing")
	}
	secret := os.Getenv("AWS_SECRET_KEY")
	if secret == "" {
		log.Fatalln("env AWS_SECRET_KEY missing")
	}
	region := os.Getenv("AWS_REGION")
	if region == "" {
		log.Fatalln("env AWS_REGION missing")
	}

	poolID := os.Getenv("POOL_ID")
	if poolID == "" {
		log.Fatalln("env POOL_ID missing")
	}

	clientID := os.Getenv("CLIENT_ID")
	if clientID == "" {
		log.Fatalln("env CLIENT_ID missing")
	}
}
