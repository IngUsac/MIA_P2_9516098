import { useState } from "react";
import { useNavigate } from "react-router-dom";

function Login() {

    const navigate = useNavigate();

    const [carnet, setCarnet] = useState("");

    const [pass, setPass] = useState("");

    function ingresar() {

        if (
            carnet === "9516098" &&
            pass === "12345"
        ) {

            sessionStorage.setItem(
                "app_login",
                "true"
            );

            navigate("/");

            return;
        }

        alert("Credenciales incorrectas.");

    }

    return (

        <div
            style={{
                width:400,
                margin:"120px auto",
                padding:30,
                border:"1px solid #ccc",
                borderRadius:10
            }}
        >

            <h2>

                Acceso a la aplicación

            </h2>

            <input

                placeholder="Carnet"

                value={carnet}

                onChange={(e)=>

                    setCarnet(e.target.value)

                }

                style={{
                    width:"100%",
                    marginBottom:10
                }}

            />

            <input

                type="password"

                placeholder="Contraseña"

                value={pass}

                onChange={(e)=>

                    setPass(e.target.value)

                }

                style={{
                    width:"100%",
                    marginBottom:20
                }}

            />

            <button

                onClick={ingresar}

            >

                Ingresar

            </button>

        </div>

    );

}

export default Login;