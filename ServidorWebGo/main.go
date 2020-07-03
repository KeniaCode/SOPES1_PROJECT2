package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"

	//"log"
	"net/http"
	//"fmt"
)


type Welcome struct {
	Name string
	Time string
}


type Persona struct {
	Nombre string
	Departamento string
	Edad int
	Forma string `json:"Forma de contagio"`
	Estado string
}

func main() {
	http.HandleFunc("/", postData)
	http.HandleFunc("/postData", postData)
	http.ListenAndServe(":7070", nil)
}

func inicio(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message": "Servidor inicializado"}`))
}


func postData(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
    w.Write([]byte(`{"message": "servicio post inicializado"}`))

    var persona Persona
    //Leemos el JSON que recibimos
	jsonPost, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Panic(err)
	}
	//Convertimos el JSON a la estructura Datos
	json.Unmarshal(jsonPost, &persona)

	//Imprimimos en la consola para ver el Json que recibimos
	println("Recibido: "+persona.Nombre)
	println(string(jsonPost) + "\n")


	//Hacemos petición POST a Phyton
	respuesta, err := http.Post("http://localhost:5000/postData", "application/json", bytes.NewBuffer(jsonPost))
	if err != nil {
		log.Panic("Error creando petición a Phyton: %v", err)
	}

	//Leemos la respuesta de la petición POST
	defer respuesta.Body.Close()
	cuerpoRespuesta, err := ioutil.ReadAll(respuesta.Body)
	if err != nil {
		log.Panic("Error leyendo respuesta: %v", err)
	}
	println("Petición a Phyton realizada: ")
	println(string(cuerpoRespuesta))
}


