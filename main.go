package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func mirrorHandler(writer http.ResponseWriter, request *http.Request) {
	request.Write(writer)
}

func cookieHandler(writer http.ResponseWriter, request *http.Request) {
	cookie := &http.Cookie{}
	cookie.Name = "test"
	cookie.Value = time.Now().Format(time.RFC3339)
	http.SetCookie(writer, cookie)
}

func main() {
	// Port is last argument
	port := "8080"
	for _, arg := range os.Args[1:] {
		port = arg
	}

	log.Println("Listening on :" + port)
	http.HandleFunc("/", mirrorHandler)
	http.HandleFunc("/cookie", cookieHandler)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
