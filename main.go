package main

import (
	"log"
	"net/http"
	"time"

	cart "github.com/DanielBcnicode/apigo/Cart"
	"github.com/gorilla/mux"
)

func main() {

	postItemHandle := DefaultContainer.GetPostItemHandleService()
	//new(cart.PostItemHandleService)

	router := mux.NewRouter()
	router.HandleFunc("/v1/cart/{cartId}", cart.GetCartHandle).Methods("GET")
	router.HandleFunc("/v1/cart/item", postItemHandle.PostItemHandle).Methods("POST")

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("Server Started.")

	log.Fatal(srv.ListenAndServe())
}
