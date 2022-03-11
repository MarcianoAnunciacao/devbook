package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"
)

/* func init() {
	hashKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
	fmt.Println(hashKey)

	blockKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
	fmt.Println(blockKey)
} */

func main() {
	config.Load()
	cookies.ConfigCookie()
	utils.LoadTemplates()
	r := router.GenerateRoutes()

	fmt.Printf("Running WebApp on Port %d\n", config.Port)
	log.Fatal(http.ListenAndServe(":3000", r))
}
