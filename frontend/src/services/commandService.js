import api from "./api";

export async function executeCommand(command) {

    const response = await api.post("/api/execute", {
        command
    });

    return response.data;
}