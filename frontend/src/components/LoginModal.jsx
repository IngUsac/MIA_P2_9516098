import { useEffect, useState } from "react";
import "./LoginModal.css";
import { executeCommand } from "../services/commandService";
import { getMounts } from "../api/api";

function LoginModal({

    open,

    onClose,

    onSuccess,

    partitions

}) {

    const [user, setUser] = useState("");

    const [pass, setPass] = useState("");

    const [id, setId] = useState("");

    const [mounts, setMounts] = useState([]);

    useEffect(() => {

        async function cargarMontajes() {

            try {

                const lista =
                    await getMounts();

                setMounts(
                    Array.isArray(lista)
                        ? lista
                        : []
                );

            } catch {

                setMounts([]);

            }

        }

        if (open) {

            cargarMontajes();

        }

    }, [open]);

    if (!open) return null;

    
    async function iniciarSesion() {

    if (
        user === "" ||
        pass === "" ||
        id === ""
    ) {
        alert("Complete todos los campos.");
        return;
    }

    const comando =
        `login -user=${user} -pass=${pass} -id=${id}`;

    const respuesta =
        await executeCommand(comando);

    console.log(respuesta);

    // Si el backend devolvió error
    if (
        respuesta?.success === false ||
        respuesta?.error
    ) {

        alert(
            respuesta.error ??
            "No fue posible iniciar sesión."
        );

        return;
    }

    await onSuccess();

    onClose();

}



    


    return (

        <div className="modal-overlay">

            <div className="modal">

                <h2>

                    Iniciar sesión

                </h2>

                <input

                    placeholder="Usuario"

                    value={user}

                    onChange={(e)=>

                        setUser(e.target.value)

                    }

                />

                <input

                    type="password"

                    placeholder="Contraseña"

                    value={pass}

                    onChange={(e)=>

                        setPass(e.target.value)

                    }

                />

               {
                    mounts.length === 0

                    ?

                    <p
                        style={{
                            color: "red",
                            marginBottom: "15px"
                        }}
                    >
                        No existen particiones montadas.
                    </p>

                    :

                    <select

                        value={id}

                        onChange={(e)=>setId(e.target.value)}

                    >

                        <option value="">

                            Seleccione una partición

                        </option>

                        {

                            mounts.map((m)=>(

                                <option

                                    key={m.id}

                                    value={m.id}

                                >

                                    {m.name} ({m.id})

                                </option>

                            ))

                        }

                    </select>
                }

                <div className="buttons">

                    <button

                        onClick={onClose}

                    >

                        Cancelar

                    </button>

                    <button

                        onClick={iniciarSesion}

                    >

                        Ingresar

                    </button>

                </div>

            </div>

        </div>

    );

}

export default LoginModal;