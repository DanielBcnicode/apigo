package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	postItemHandle := DefaultContainer.GetPostItemHandleService()
	getCartHandle := DefaultContainer.GetGetCartHandle()

	router := mux.NewRouter()
	router.HandleFunc("/v1/cart/{cartId}", getCartHandle.Handle).Methods("GET")
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
