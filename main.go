package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gomovie/Routes"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Listening...")
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), Routes.Router()))
}
