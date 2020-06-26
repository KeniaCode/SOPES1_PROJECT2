package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)


type Welcome struct {
	Name string
	Time string
}

type Datos struct {
	Url string
	Hilos int
	Solicitudes int
	Ruta string
	ArrayDatos []Persona
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


func postData(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte(`{"message": "post called"}`))

    var datos Datos
    //Leemos el JSON que recibimos
	jsonPost, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//Convertmos el JSON a la estructura Datos
	json.Unmarshal(jsonPost, &datos)

	//Imprimimos en la consola para ver el Json que recibimos
	println(string(jsonPost))
	println(datos.Ruta)
}


