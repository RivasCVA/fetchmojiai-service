package main

import (
	"log"
	"net/http"

	"github.com/RivasCVA/fetchmojiai-service/api"
	"github.com/RivasCVA/fetchmojiai-service/server"
	"github.com/RivasCVA/fetchmojiai-service/server/handler/imagine"
	"github.com/RivasCVA/fetchmojiai-service/server/middleware"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// load the environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("unable to load .env variables")
	}

	// setup the server
	s := server.NewServer(
		imagine.NewHandler(),
	)

	// setup the handler
	h := api.HandlerWithOptions(s, api.GorillaServerOptions{
		BaseRouter:       mux.NewRouter(),
		Middlewares:      []api.MiddlewareFunc{middleware.CommonHeaders},
		BaseURL:          "/v1",
		ErrorHandlerFunc: nil,
	})

	// run the server
	srv := &http.Server{
		Handler: h,
		Addr:    ":8080",
	}
	log.Println("server running")
	log.Fatal(srv.ListenAndServe())
}
