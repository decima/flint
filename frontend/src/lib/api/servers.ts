import {makeApiCall} from "$lib/api/makeApiCall";

export function getServers() {
    // @ts-ignore
    return makeApiCall("GET", "/api/servers").then(data => data.servers)
}