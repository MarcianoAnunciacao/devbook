package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	config.Load()
	utils.LoadTemplates()
	r := router.GenerateRoutes()

	fmt.Printf("Running WebApp on Port %d\n", config.Port)
	log.Fatal(http.ListenAndServe(":3000", r))
}
