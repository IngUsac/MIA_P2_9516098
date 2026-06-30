import { useState } from "react";
import { executeCommand } from "../services/commandService";

export default function Dashboard() {

    const [command, setCommand] = useState("");
    const [output, setOutput] = useState("");

    async function ejecutar() {

        try {

            const response = await executeCommand(command);

            setOutput(response.data.output);

        } catch {

            setOutput("Error conectando con la API");

        }

    }

    return (

        <div style={{ padding: 20 }}>

            <h2>Proyecto 1 - MIA</h2>

            <textarea
                rows={8}
                cols={100}
                value={command}
                onChange={(e) => setCommand(e.target.value)}
            />

            <br /><br />

            <button onClick={ejecutar}>
                Ejecutar
            </button>

            <pre
                style={{
                    marginTop:20,
                    background:"#111",
                    color:"#00ff00",
                    padding:20,
                    minHeight:250
                }}
            >
                {output}
            </pre>

        </div>

    );

}