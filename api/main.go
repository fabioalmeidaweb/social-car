package main

import (
	"fmt"
	"log"
	"net/http"
	"social-car/src/config"
	"social-car/src/database"
	"social-car/src/router"
)

func init() {
	// key := make([]byte, 64)
	// if _, err := rand.Read(key); err != nil {
	// 	log.Fatal(err)
	// }
	// stringBase64 := base64.StdEncoding.EncodeToString(key)
	// fmt.Println("SECRET_KEY: " + stringBase64)
}

func main() {
	config.LoadConfig()
	_, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	r := router.Generate()

	fmt.Printf("Server running on http://localhost:%d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
