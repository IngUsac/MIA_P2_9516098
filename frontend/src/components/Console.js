import { useState } from "react";
import { executeCommand } from "../services/commandService";

function Console({ onCommandExecuted }) {

    const [command, setCommand] = useState("");

    const [output, setOutput] = useState("");

    async function ejecutar() {

        if (command.trim() === "") return;

        try {

            const response = await executeCommand(command);

            if (response.success) {

                setOutput(response.data.output);

                if (onCommandExecuted) {
                    await onCommandExecuted();
                }

            } else {

                setOutput(response.message);

            }

        } catch (error) {

            setOutput("Error conectando con el backend.");

        }

    }
        


    return (

        <div>

            <textarea

                value={command}

                onChange={(e)=>setCommand(e.target.value)}

                rows={8}

                style={{
                    width:"100%"
                }}

            />

            <br/><br/>

            <button onClick={ejecutar}>

                Ejecutar

            </button>

            <pre

                style={{
                    marginTop:15,
                    background:"#111",
                    color:"#00ff00",
                    padding:15,
                    minHeight:180,
                    overflow:"auto"
                }}

            >

                {output}

            </pre>

        </div>

    );

}

export default Console;