# Sistemas Operativos 1 - Proyecto 2
El proyecto consiste en crear un sitio web que muestre los datos enviados en tiempo real con ayuda de dos balanceador de carga, Docker Compose, kubernetes y las bases de datos de Mongodb y Redis.
La aplicación utilizara el concepto de colas para el manejo de la concurrencia. Se debera de transmitir la información mediante diferentes nodos, cada uno con implementado con tecnología diferente.

# ClienteGo

Requiere Golang 1.14  

# Build Go webapp and start the server  
<pre>  cd /SOPES1_PROJECT2/ClieteGo</pre>     
<pre>  go run main.go </pre>  

#Host and port  
<pre>  localhost:8080</pre> 


Esta parte consiste en enviar los datos de un archivo de entrada a los balanceadores de carga, los balanceadores de carga serán
implementados en: Contour y Nginx

Cuando se ejecute el programa de Go se pregunta lo siguiente:
1) Url del balanceador de carga que se desea enviar
2) Cantidad de hilos que se desean para enviar
3) Cantidad de solicitudes que tiene el archivo (la cantidad de casos que se desean enviar)
4) Ruta del archivo que se desea cargar


# ServidorWebGo

Requiere Golang 1.14  

# Build Go web server and start 
<pre>  cd /SOPES1_PROJECT2/ServidorWebGo </pre>     
<pre>  go run main.go </pre>  

#Host and port  
<pre>  localhost:7070</pre>  

Esta parte contiene un servicio POST que recibe un JSON del cliente. El JSON tiene la siguiete estructura: 

<pre>
{	
"url":"localhost:7070/postDatos",
"solicitudes": "5",
"hilos":"5"
"datos":
[
	{	"Nombre":"Nombre1",
		"Departamento":"Guatemala",
		"Edad":111,
		"Forma de contagio":"comunitario",
		"Estado": "Activo"
	},
	{	"Nombre":"Nombre2",
		"Departamento":"Guatemala",
		"Edad":222,
		"Forma de contagio":"comunitario",
		"Estado": "Activo"
	},
	{	"Nombre":"Nombre3",
		"Departamento":"Guatemala",
		"Edad":333,
		"Forma de contagio":"comunitario",
		"Estado": "Activo"
	},
	{	"Nombre":"Nombre4",
		"Departamento":"Guatemala",
		"Edad":444,
		"Forma de contagio":"comunitario",
		"Estado": "Activo"
	}
]

}

</pre>

