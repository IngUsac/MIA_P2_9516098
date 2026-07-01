/*
Home.jsx

Pantalla principal del Proyecto 1.

- Verifica la conexión con el Backend REST.
- Obtiene la lista de discos disponibles.
- Muestra el estado de conexión.
*/


import { useNavigate } from "react-router-dom";
import { useEffect, useState } from "react";
import { getStatus, getDisks, getPartitions, getReports, getSession } from "../api/api";
import "../styles/home.css";
import DiskList from "../components/DiskList";
import PartitionList from "../components/PartitionList";
import Panel from "../components/Panel";
import FileTree from "../components/FileTree";
import Console from "../components/Console";
import ReportList from "../components/ReportList";
import { executeCommand } from "../services/commandService";
import LoginModal from "../components/LoginModal";


function Home() {

    const navigate = useNavigate();

    // Estado de conexión con el backend.
    const [backend, setBackend] = useState(false);

    // Lista de discos disponibles.
    const [disks, setDisks] = useState([]);

    const [selectedDisk, setSelectedDisk] = useState(null);

    const [partitions, setPartitions] = useState([]);

    const [selectedPartition,setSelectedPartition] = useState(null);

    const [reports, setReports] = useState([]);

    const [selectedReport, setSelectedReport] = useState(null);

    const [session, setSession] = useState({ logged: false, user: "", id: "",  partition: "" });

    const [showLogin, setShowLogin] = useState(false);



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

            const discos = Array.isArray(lista) ? lista : [];

            setDisks(discos);

            // Si desapareció el disco seleccionado
            if (
                selectedDisk &&
                !discos.find(d => d.name === selectedDisk)
            ) {
                setSelectedDisk(null);
                setPartitions([]);
                setSelectedPartition(null);
            }

            const listaReportes = await getReports();

            setReports(
                Array.isArray(listaReportes)
                    ? listaReportes
                    : []
            );

            const sesion = await getSession();

            setSession(sesion);   


            if (
                selectedReport &&
                !listaReportes.find(
                    r => r.name === selectedReport.name
                )
            ) {

                setSelectedReport(null);

            }




        } catch (error) {

            console.error("Error:", error);

            setBackend(false);

        }



    }

    async function cerrarSesion() {

        try {

            const respuesta =
                await executeCommand("logout");

            console.log(respuesta);

            await cargarInformacion();

        } catch (error) {

            console.error(error);

        }

    }


    return (

        <div className="home">

            <div
    style={{
        display: "flex",
        justifyContent: "space-between",
        alignItems: "center",
        background: "#f5f5f5",
        padding: "10px 15px",
        borderRadius: 8,
        marginBottom: 20
    }}
>

    <div>

        {

            session.logged

            ?

            <>

                <b>👤 Usuario:</b>

                {" "}

                {session.user}

                {" | "}

                <b>🆔 ID:</b>

                {" "}

                {session.id}

            </>

            :

            <b>No hay sesión activa</b>

        }

    </div>

    <div>

        {

            session.logged

            ?

            <button

                onClick={cerrarSesion}

                style={{

                    padding:"8px 16px",

                    cursor:"pointer"

                }}

            >

                Cerrar sesión

            </button>

            :

            <button

                onClick={()=>setShowLogin(true)}

                style={{

                    padding:"8px 16px",

                    cursor:"pointer"

                }}

            >

                Iniciar sesión

            </button>

        }

    </div>

</div>

<h2>

    Backend:

    {

        backend

        ? " 🟢 Conectado"

        : " 🔴 Desconectado"

    }

</h2>

            <div className="grid">

                <Panel title="📀 Discos">

                    <DiskList

                        disks={disks}

                        selectedDisk={selectedDisk}

                        onSelectDisk={async (disk) => {

                            setSelectedDisk(disk);

                            const lista =
                                await getPartitions(disk);

                            setPartitions(lista);

                        }}

                    />

                </Panel>

                <Panel title="💽 Particiones">

                    <PartitionList

                        partitions={partitions}

                        selectedPartition={

                            selectedPartition?.name

                        }

                        onSelectPartition={

                            (partition)=>{

                                setSelectedPartition(
                                    partition
                                );

                            }

                        }

                    />

                </Panel>

            </div>

            <Panel title="🌳 Árbol del Sistema de Archivos">

                <FileTree />
                    <hr />

                        <h3>

                            Partición seleccionada

                        </h3>

                        {

                            selectedPartition

                            ?

                            (

                                <div>

                                    <p>

                                        <b>Nombre:</b>

                                        {" "}

                                        {selectedPartition.name}

                                    </p>

                                    <p>

                                        <b>Tipo:</b>

                                        {" "}

                                        {selectedPartition.typeName}

                                    </p>

                                    <p>

                                        <b>Tamaño:</b>

                                        {" "}

                                        {selectedPartition.size}

                                        {" "}bytes

                                    </p>

                                    <p>

                                        <b>Inicio:</b>

                                        {" "}

                                        {selectedPartition.start}

                                    </p>

                                </div>

                            )

                            :

                            (

                                <p>

                                    Ninguna partición seleccionada.

                                </p>

                            )

                        }

            </Panel>

            <Panel title="📄 Reportes">

                <ReportList

                    reports={reports}

                    selected={selectedReport?.name}

                    onSelect={setSelectedReport}

                />

            </Panel>

            <Panel title="👁 Vista previa">

                {

                    selectedReport

                    ?

                    (

                        selectedReport.name.endsWith(".png")

                        ?

                        <img

                            src={
                                "http://localhost:8080/reportes/" +
                                selectedReport.name
                            }

                            alt={selectedReport.name}

                            style={{
                                width:"100%"
                            }}

                        />

                        :

                        <iframe

                            title={selectedReport.name}

                            src={
                                "http://localhost:8080/reportes/" +
                                selectedReport.name
                            }

                            style={{
                                width:"100%",
                                height:"700px",
                                border:"none"
                            }}

                        />

                    )

                    :

                    (

                        <p>

                            Seleccione un reporte.

                        </p>

                    )

                }

            </Panel>


            <Panel title="💻 Consola">

                <Console
                    onCommandExecuted={cargarInformacion}
                />

            </Panel>

            <LoginModal

                open={showLogin}

                onClose={()=>

                    setShowLogin(false)

                }

                onSuccess={

                    cargarInformacion

                }

                partitions={

                    partitions

                }

            />


        </div>

    );

}

export default Home;