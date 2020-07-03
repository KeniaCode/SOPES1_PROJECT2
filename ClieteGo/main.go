package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)



type Welcome struct {
	Name string
	Time string
}

type Persona struct {
	Nombre string `json:"nombre"`
	Departamento string `json:"departamento"`
	Edad int `json:"edad"`
	Forma_contagio string `json:"forma_contagio"`
	Estado string `json:"estado"`
}

var mensaje string

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
	//Leemos los parametros que vienen en la URL que recibimos del JS
	w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusOK)
	query := r.URL.Query()
	url := query.Get("url") //filters="color"
	hilos,_ := strconv.Atoi( query.Get("hilos") )
	solicitudes,_ := strconv.Atoi( query.Get("solicitudes") )
	ruta := query.Get("ruta")

	//Leemos el archivo JSON
	manejadorDeArchivo, err := ioutil.ReadFile(ruta)
	if err != nil {
		log.Panic(err)
	}

	newContents := strings.Replace(string(manejadorDeArchivo), "Forma de contagio", "forma_contagio", -1)


	//Parseamos el JSON a la estructura Persona
	listaObjPersona := []Persona{}
	err = json.Unmarshal([]byte(newContents), &listaObjPersona)
	if err != nil {
		log.Panic(err)
	}

	makeThreads(listaObjPersona,solicitudes, hilos, url);
	//w.Write([]byte(`{"message":"`+ mensaje +`"}`))

}


func makeThreads(listaPersonas []Persona, solicitudes int, hilos int, url string){
	cant := solicitudes / hilos
	var syncc sync.WaitGroup
	syncc.Add(hilos)
	println("INICA HILOS")
	for i := 0; i < hilos; i++ {
		if(i == hilos-1){
			go makePost(listaPersonas, cant*i, solicitudes, url)
		}else{
			go makePost(listaPersonas, cant*i, cant*(i+1), url)
		}
		syncc.Done()
	}
	println("Termina Hilos")
	syncc.Wait()

}

func makePost(persons []Persona, init int, cant int,url string){
	println("Entro al FOR, init: " + strconv.Itoa(init) + ", cant: "+strconv.Itoa(cant))
	for i := init; i < cant; i++ {
		jsonData, _ := json.Marshal(persons[i])
		_, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			mensaje = "ERROR: no se realizó POST correctamente: " + err.Error()
			log.Panic("Error al momento de enviar la información: %s\n", err)
		}else{
			println("Enviando hilo, persona: " + persons[i].Nombre)
			mensaje = "Hilos enviados correctamente "

		}
		time.Sleep(time.Millisecond * 10)
	}
}

