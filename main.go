/*
	initialise the API
*/

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"gitlab.com/safesurfer/minio-file-server/pkg/common"
	"gitlab.com/safesurfer/minio-file-server/pkg/routes"
)

func handleWebserver() {
	// bring up the API
	port := common.GetAppPort()
	router := mux.NewRouter().StrictSlash(true)
	for _, endpoint := range routes.GetEndpoints("/") {
		router.HandleFunc(endpoint.EndpointPath, endpoint.HandlerFunc).Methods(endpoint.HTTPMethods...)
	}

	router.HandleFunc("/{.*}", routes.APIUnknownEndpoint)

	router.Use(common.Logging)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET"},
		AllowCredentials: true,
	})

	srv := &http.Server{
		Handler:      c.Handler(router),
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Listening on", port)
	log.Fatal(srv.ListenAndServe())
}

func main() {
	// initialise the app
	handleWebserver()
}
