import { useEffect, useState } from "react";
import { getFiles } from "../api/api";

function FileTree() {

    const [ruta, setRuta] = useState("/");

    const [files, setFiles] = useState([]);

    useEffect(() => {

        cargar(ruta);

    }, [ruta]);

    async function cargar(path) {

        try {

            const lista = await getFiles(path);

            setFiles(
                Array.isArray(lista)
                    ? lista
                    : []
            );

        } catch (e) {

            console.error(e);

            setFiles([]);

        }

    }

    return (

        <div>

            <h3>

                Ruta: {ruta}

            </h3>

            {

                ruta !== "/" &&

                <button

                    onClick={() => {

                        const partes =
                            ruta.split("/").filter(Boolean);

                        partes.pop();

                        if (partes.length === 0) {

                            setRuta("/");

                        } else {

                            setRuta(
                                "/" +
                                partes.join("/")
                            );

                        }

                    }}

                >

                    ⬅ Regresar

                </button>

            }

            <ul>

                {

                    Array.isArray(files) &&
                    files.map((f) => (

                        <li
                            key={f.name}
                        >

                            {

                                f.isDirectory

                                ?

                                <button

                                    onClick={() =>

                                        setRuta(

                                            ruta === "/"

                                            ?

                                            "/" + f.name

                                            :

                                            ruta + "/" + f.name

                                        )

                                    }

                                >

                                    📁 {f.name}

                                </button>

                                :

                                <span>

                                    📄 {f.name}

                                </span>

                            }

                        </li>

                    ))

                }

            </ul>

        </div>

    );

}

export default FileTree;