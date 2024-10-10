package main

import (
	"log"
	"meetdicer/configs"
	"meetdicer/data"
	"meetdicer/routes"
	"net/http"
)

func main() {
	prepareConfigs()

	prepareDB()

	prepareServer()
}

func prepareConfigs() {
	configs.LoadDotEnv()
}
func prepareDB() {
	data.ConnectToDB()
}

func prepareServer() {
	serverAddress := prepareServerAddress()

	router := routes.PrepareRouting()
	server := &http.Server{
		Addr:         serverAddress,
		ReadTimeout:  configs.DefaultReadTimeout(),
		WriteTimeout: configs.DefaultWriteTimeout(),
		Handler:      router,
	}

	log.Println("Starting server on " + server.Addr)
	log.Fatal(server.ListenAndServe())
}

func prepareServerAddress() string {
	serverPort := configs.EnvPort()
	hostname := configs.EnvHostname()
	return hostname + ":" + string(serverPort)
}
