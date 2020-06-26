
function functionSenData() {
    let url = document.getElementById("inputUrlBalanceador").value;
    let hilos = document.getElementById("inputHilos").value;
    let solicitudes = document.getElementById("inputSolicitudes").value;
    let ruta = document.getElementById("inputRutaArchivo").value;


    const headers = new Headers();
   // headers.append('Content-Type', 'application/json');
    const init = {
        method: 'GET'
     //   headers
    };

    fetch('http://localhost:8080/getData?url='+url+'&hilos='+hilos+'&solicitudes='+solicitudes+'&ruta='+ruta, init)
        .catch((e) => {
            console.log("ERROR: " + e.toString());
        });

}

