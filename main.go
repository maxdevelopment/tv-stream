package main

import (
	"github.com/maxdevelopment/tv-stream/config"
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"log"
	"io/ioutil"
	"github.com/maxdevelopment/tv-stream/ws"
	"github.com/maxdevelopment/tv-stream/api"
)

func main() {
	config.ReadConfig()

	hub := ws.H
	go hub.Run()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", GetIndex).Methods("GET")
	router.HandleFunc("/stream", GetIndex).Methods("GET")
	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("web/dist/"))))

	router.HandleFunc("/api/start-stream", api.Handler).Methods("GET")
	router.HandleFunc("/stream/{id}", ws.Handler).Methods("GET")

	router.NotFoundHandler = http.HandlerFunc(GetIndex)

	srv := &http.Server{
		Handler:      router,
		Addr:         config.Get.IP + ":" + config.Get.Port,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func GetIndex(writer http.ResponseWriter, request *http.Request) {
	indexFile, err := ioutil.ReadFile("web/index.html")
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.Write(indexFile)
}
