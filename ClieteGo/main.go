package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
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
	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}

	templates := template.Must(template.ParseFiles("static/index.html"))
	http.Handle("/static/", http.FileServer(http.Dir(".")))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}

		if err := templates.ExecuteTemplate(w, "index.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/getData", getData)

	http.ListenAndServe(":8080", nil)
}

func getData(w http.ResponseWriter, r *http.Request) {
	//Leemos los parametros que vienen en la URL
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	query := r.URL.Query()
	url := query.Get("url") //filters="color"
	hilos,_ := strconv.Atoi( query.Get("hilos") )
	solicitudes,_ := strconv.Atoi( query.Get("solicitudes") )
	ruta := query.Get("ruta")

	//Leemos el archivo JSON
	manejadorDeArchivo, err := ioutil.ReadFile(ruta)
	if err != nil {
		log.Fatal(err)
	}

	//Parseamos el JSON a la estructura Persona
	objPersona := []Persona{}
	err = json.Unmarshal(manejadorDeArchivo, &objPersona)
	if err != nil {
		log.Fatal(err)
	}

	//Creamos un objeto con los datos a enviar al balanceador
	infoObj := Datos{
		Url:   url,
		Hilos:  hilos,
		Solicitudes: solicitudes,
		Ruta: ruta,
		ArrayDatos: objPersona,
	}

	//Convertimos la estructura a un JSON
	jsonResponse, errorjson := json.Marshal(infoObj)
	if errorjson != nil {
		http.Error(w, errorjson.Error(), http.StatusInternalServerError)
		return
	}

	//imprimimos el JSON en consola
	println(string(jsonResponse))

	//Hacemos petición POST al balanceador
	clienteHttp := &http.Client{}
	peticion, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonResponse)))
	if err != nil {
		log.Fatalf("Error creando petición: %v", err)
	}

	peticion.Header.Add("Content-Type", "application/json")
	respuesta, err := clienteHttp.Do(peticion)
	if err != nil {
		log.Fatalf("Error haciendo petición: %v", err)
	}

	//Leemos la respuesta de la petición POST
	defer respuesta.Body.Close()
	cuerpoRespuesta, err := ioutil.ReadAll(respuesta.Body)
	if err != nil {
		log.Fatalf("Error leyendo respuesta: %v", err)
	}
	println(string(cuerpoRespuesta))
}


