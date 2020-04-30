package main

import (
	"auction_system/controller"
	"auction_system/model"
	"fmt"
	"log"
	"net/http"
	"os"
	"trip-planer/initializers"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("=======Auction=======")
	if initializers.InitiateEnv() == true {
		host := os.Getenv("SOCKET")
		router := controller.Register()
		db := model.Connection()
		err := db.Ping()
		if err != nil {
			fmt.Println("connection is not active")
			log.Fatal(err.Error)
		}
		defer db.Close() // defer is used to execute the statement end of the scope here last line of main()
		http.ListenAndServe(host, router)
	} else {
		log.Fatal("error in loading env")
	}

}
