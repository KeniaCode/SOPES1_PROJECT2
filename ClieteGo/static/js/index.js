
function functionSenData() {
    let url = document.getElementById("inputUrlBalanceador").value;
    let hilos = document.getElementById("inputHilos").value;
    let solicitudes = document.getElementById("inputSolicitudes").value;
    let ruta = document.getElementById("inputRutaArchivo").value;


    const headers = new Headers();
    headers.append('Content-Type', 'application/json');
    const init = {
        method: 'GET',
        headers
    };

    alert("Enviando Datos Al Balanceador")

    fetch('http://localhost:8080/getData?url='+url+'&hilos='+hilos+'&solicitudes='+solicitudes+'&ruta='+ruta, init)
      /*  .then(response => response.json())
        .then(data => {
            info = data
            alert(info.message)
            // text is the response body
        })*/
        .catch((e) => {
            alert("ERROR AL ENVIAR HILOS")
            console.log("ERROR: " + e.toString());
        });
}

