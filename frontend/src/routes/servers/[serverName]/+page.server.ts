import api from "../../../store/api.js";

export async function load({params}) {
    return {
        server: await api.getServer(params.serverName)
    };
}
