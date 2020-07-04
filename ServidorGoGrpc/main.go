package main

import (
	"./casos"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute).Methods("GET")
	router.HandleFunc("/", postData).Methods("POST")
	router.Use(mux.CORSMethodMiddleware(router))
	log.Println("Starting server. Listening on port 4000.")
	http.ListenAndServe(":7070", nil)
}

func postData(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Se han enviado datos incorrectos para crear un caso")
		return
	}

	nuevos := &casos.CasoRequest{}
	if err := json.Unmarshal(reqBody, nuevos); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Se han enviado datos incorrectos para crear un caso")
		return
	}

	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "No se pudo conectar al servidor GRPC")
		return
	}
	defer conn.Close()

	nuevoCliente := casos.NewCasoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := nuevoCliente.CrearCasos(ctx, nuevos)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "No se ha podido enviar el caso.")
		return
	}
	message := response.GetMensaje()
	w.WriteHeader(http.StatusOK)
	log.Println(message)
	fmt.Fprintf(w, "Mensaje: %s", message)
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Servidor inicializado")
}
