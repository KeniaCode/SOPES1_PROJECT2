package main

import (
	context "context"
	"encoding/json"
	"fmt"
	_ "github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"
)

type Persona struct {
	Nombre string `json:"nombre"`
	Departamento string `json:"departamento"`
	Edad int32 `json:"edad"`
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
	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusCreated)
	//w.Write([]byte(`{"message": "Post Called"}`))

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "ERROR: datos incorrectos en el body de la solicitud POST")
		return
	}

	println("Recibido: ")
	println(string(reqBody))

	//Parseamos el JSON a la estructura Persona
	objPersona := Persona{}
	err = json.Unmarshal([]byte(reqBody), &objPersona)
	if err != nil {
		log.Panic(err)
	}

	nuevos := &CasoRequest{
		state:         protoimpl.MessageState{},
		sizeCache:     0,
		unknownFields: nil,
		Casos:         &CasoItem{},
	}
	cass := &CasoItem{
		Nombre : objPersona.Nombre,
		Departamento : objPersona.Departamento,
		Estado : objPersona.Estado,
		FormaContagio : objPersona.Forma_contagio,
		Edad : objPersona.Edad,
	}
	nuevos.Casos = cass;

	host := "localhost"
	port := 5000
	str := net.JoinHostPort(host, strconv.Itoa(port))

	conn, err := grpc.Dial(str, grpc.WithInsecure())
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
		fmt.Fprintf(w, "ERROR: No se ha podido enviar el caso.")
		return
	}
	message := response.GetMensaje()
	w.WriteHeader(http.StatusOK)
	println("Respuesta python: "+ message)
	fmt.Fprintf(w, "Mensaje: %s", message)
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Servidor inicializado")
}
