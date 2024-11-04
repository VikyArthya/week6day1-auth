package main

import (
	"fmt"
	"net/http"

	authController "github.com/VikyArthya/go-auth/controllers"
)

func main() {
	http.HandleFunc("/", authController.Index)
	http.HandleFunc("/login", authController.Login)

	fmt.Println("server berjalan di http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
