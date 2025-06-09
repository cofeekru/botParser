package parser

import (
	"log"
	"net/http"
)

func handler(response http.ResponseWriter, request *http.Request) {
	jsonResponse, statusCode := parser(request.URL)
	response.WriteHeader(statusCode)
	response.Write(jsonResponse)
}

func ServerStart(HOST string, PORT string) {

	server := &http.Server{
		Addr:    PORT,
		Handler: http.HandlerFunc(handler),
	}

	log.Println("Starting parser server at port" + PORT + "...")
	err := server.ListenAndServe()

	if err != nil {
		log.Println("Error starting the server.", err)
	}
}
