import {authToken, refresh} from "$lib/store/auth";
import {ERR_NEEDS_RE_LOGIN, ERR_NEEDS_REFRESH_TOKEN, ROUTE_LOGIN} from "$lib/constants";
import type {Unsubscriber} from "svelte/store";

export function makeApiCall(method: string, path: string, body?: any) {
    return new Promise((resolve, reject) => {
        let unsub: Unsubscriber;
        unsub = authToken.subscribe(async uToken => {
            const response = await fetch(path, {
                method,
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": `Bearer ${uToken?.token}` // Assuming token is managed globally
                },
                body: body ? JSON.stringify(body) : undefined
            })
            unsub();

            if (response.status == 401) {
                await new Promise(r => setTimeout(r, 1000)); // wait for the authToken store to update
                try{
                    await refresh(`${uToken?.refresh}`)
                }catch(e){
                    if(e==ERR_NEEDS_RE_LOGIN){
                        window.location.href=ROUTE_LOGIN;
                        reject(ERR_NEEDS_RE_LOGIN);
                        return;
                    }
                }
                resolve(await makeApiCall(method, path, body))
                reject(ERR_NEEDS_REFRESH_TOKEN);
                return;

            }
            if (!response.ok) {
                reject(new Error("Failed to load resource"))
            }
            const data = (await response.json()).result;
            resolve(data);
        })
    })
}