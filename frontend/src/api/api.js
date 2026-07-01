/*
api.js

Centraliza la comunicación con el Backend REST.
Todas las peticiones HTTP del Frontend deberán
realizarse desde este archivo.
*/

// URL base del Backend REST.
const API_URL = "http://localhost:8080";

/*
request

Realiza una petición HTTP al Backend,
valida la respuesta y devuelve el JSON.
*/
async function request(url) {

    console.log("GET:", url);

    const response = await fetch(url, {

        method: "GET",

        headers: {

            "Content-Type": "application/json"

        }

    });

    console.log("STATUS:", response.status);

    if (!response.ok) {

        throw new Error(
            `HTTP ${response.status} - ${response.statusText}`
        );

    }

    const data = await response.json();

    console.log("DATA:", data);

    return data;

}

/*
getStatus

Obtiene el estado del Backend REST.
*/
export async function getStatus() {

    return request(
        `${API_URL}/api/status`
    );

}

/*
getDisks

Obtiene la lista de discos disponibles.
*/
export async function getDisks() {

    return request(
        `${API_URL}/api/disks`
    );

}

/*
getPartitions

Obtiene las particiones de un disco.
*/
export async function getPartitions(disk) {

    return request(
        `${API_URL}/api/partitions?disk=${disk}`
    );

}


// getTree: Obtiene el árbol del sistema de archivos de la partición seleccionada.

export async function getTree(id) {

    return request(
        `${API_URL}/api/tree?id=${id}`
    );

}

// getReports: muestra los reportes en el frontend
export async function getReports() {

    return request(
        `${API_URL}/api/reports`
    );

}


export async function getSession() {

    return request(
        `${API_URL}/api/session`
    );

}

export async function getMounts() {

    return request(
        `${API_URL}/api/mounts`
    );

}

export async function getFiles(path = "/") {

    const response = await fetch(

        `${API_URL}/api/files?path=${encodeURIComponent(path)}`

    );

    if (!response.ok) {

        throw new Error(await response.text());

    }

    return await response.json();

}