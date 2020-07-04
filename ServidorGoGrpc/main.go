package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Persona struct {
	Nombre string `json:"nombre"`
	Departamento string `json:"departamento"`
	Edad int `json:"edad"`
	Forma_contagio string `json:"forma_contagio"`
	Estado string `json:"estado"`
}

func main() {
	http.HandleFunc("/", indexRoute)
	http.HandleFunc("/postData", postData)
	log.Println("Starting server. Listening on port 7070.")
	http.ListenAndServe(":7070", nil)
}

func postData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Post Called"}`))

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "ERROR: datos incorrectos en el body de la solicitud POST")
		return
	}

	println("Recibido: ")
	println(string(reqBody))

	nuevos := &CasoRequest{}
	if err := json.Unmarshal(reqBody, nuevos); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "ERROR: no se puede convertir JSON a la estructura")
		return
	}

	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "ERROR: no se pudo conectar al servidor GRPC")
		return
	}
	defer conn.Close()

	nuevoCliente := NewCasoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := nuevoCliente.CrearCasos(ctx, nuevos)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "ERROR: No se ha podido enviar el caso.")
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
