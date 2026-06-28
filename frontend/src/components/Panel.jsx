/*
Panel.jsx

Contenedor reutilizable para mostrar
secciones del sistema.
*/

import "./Panel.css";

function Panel({ title, children }) {

    return (

        <div className="panel">

            <div className="panel-header">

                <h3>{title}</h3>

            </div>

            <div className="panel-body">

                {children}

            </div>

        </div>

    );

}

export default Panel;