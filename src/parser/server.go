package parser

import (
	"log"
	"net/http"
	"os"
)

func handler(response http.ResponseWriter, request *http.Request) {
	jsonResponse, statusCode := parser(request.URL)
	response.WriteHeader(statusCode)
	response.Write(jsonResponse)
}

func ServerStart() {
	port := os.Getenv("PORT")
	server := &http.Server{
		Addr:    ":" + port,
		Handler: http.HandlerFunc(handler),
	}

	log.Println("Starting parser server at port" + port + "...")
	err := server.ListenAndServe()

	if err != nil {
		log.Println("Error starting the server.", err)
	}
}
