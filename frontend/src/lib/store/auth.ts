import {writable} from "svelte/store";

export const authToken = writable<{ token: string | null, refresh: string | null } | null>(null);


export function clearAuthToken() {
    authToken.set(null);
}

export function isAuthenticated() {
    let isAuth = false;
    authToken.subscribe(value => {
        isAuth = !!(value?.token);
    })();
    return isAuth;
}

export async function login(username: string, password: string) {
    const response = await fetch("/api/login", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({username, password})
    })
    console.log(response)

    if (!response.ok) {
        throw (await response.json()).result.error;
    }
    const data = (await response.json()).result;
    authToken.set({token: data.token, refresh: data.refresh_token});
}

export async function refresh(refresh: string) {
    const response = await fetch("/api/login/refresh", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({refresh_token: refresh})
    })

    console.log("Refreshing token", response)
    if (!response.ok) {
        throw new Error("Failed to refresh token");
    }
    const data = (await response.json()).result;
    const formattedToken = {token: data.token, refresh: data.refresh_token};
    authToken.set(formattedToken);
    return formattedToken
}

export async function logout() {
    clearAuthToken();
}
