import {makeApiCall} from "$lib/api/makeApiCall";

export function getServers() {
    // @ts-ignore
    return makeApiCall("GET", "/api/servers").then(data => data.servers)
}

export function getServerSummary(id: string): Promise<any> {
    // @ts-ignore
    return makeApiCall("GET", `/api/servers/${id}/summary`).then(data => data)
}

export function getFiles(id: string, path: string = "/", raw: boolean = false): Promise<any> {
    // @ts-ignore
    return makeApiCall("GET", `/api/servers/${id}/files${path}`, null, raw).then(data => data)
}

export function writeFile(id: string, path: string = "/", content: string): Promise<any> {
    return makeApiCall("PUT", `/api/servers/${id}/files${path}`, content,true).then(data => data)
}

export function deleteFile(id: string, path: string = "/"): Promise<any> {
    return makeApiCall("DELETE", `/api/servers/${id}/files${path}`).then(data => data)
}