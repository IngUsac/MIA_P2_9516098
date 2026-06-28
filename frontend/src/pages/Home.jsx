/*
Home.jsx

Pantalla principal del Proyecto 1.

- Verifica la conexión con el Backend REST.
- Obtiene la lista de discos disponibles.
- Muestra el estado de conexión.
*/

import { useEffect, useState } from "react";
import { getStatus, getDisks } from "../api/api";
import "../styles/home.css";

function Home() {

    // Estado de conexión con el backend.
    const [backend, setBackend] = useState(false);

    // Lista de discos disponibles.
    const [disks, setDisks] = useState([]);

    /*
    Carga la información inicial de la aplicación.
    */
    useEffect(() => {

        cargarInformacion();

    }, []);

    /*
    Consulta el backend y obtiene los discos disponibles.
    */
    async function cargarInformacion() {

        try {

            console.log("Verificando Backend...");

            const status = await getStatus();

            console.log("Respuesta Backend:", status);

            if (status.success) {

                setBackend(true);

            } else {

                setBackend(false);
                return;

            }

            console.log("Obteniendo discos...");

            const lista = await getDisks();

            console.log("Discos:", lista);

            setDisks(lista);

        } catch (error) {

            console.error("Error:", error);

            setBackend(false);

        }

    }

    return (

        <div className="home">

            <h1>Proyecto 1 - MIA </h1>

            <h2>
                Backend:
                {" "}
                {backend
                    ? "🟢 Conectado"
                    : "🔴 Desconectado"}
            </h2>

            <hr />

            <h2>Discos disponibles</h2>

            {
                disks.length === 0
                    ?
                    (
                        <p>No hay discos disponibles.</p>
                    )
                    :
                    (
                        <ul>

                            {
                                disks.map((disk) => (

                                    <li key={disk.path}>

                                        📀 {disk.name}

                                    </li>

                                ))
                            }

                        </ul>
                    )
            }

        </div>

    );

}

export default Home;