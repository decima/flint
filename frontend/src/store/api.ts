import {get} from 'svelte/store';
import {authToken} from "./stores";

async function request(fetcher: typeof fetch, path: string, options: RequestInit = {}) {
    console.log("pass 1")
    const token = get(authToken);
    console.log("pass 2", token)
    const headers = new Headers(options.headers);
    console.log("pass 3")
    if (token) {
        headers.set('Authorization', `Bearer ${token}`);
    }
    headers.set('Content-Type', 'application/json');

    const response = await fetcher(`/api${path}`, {
        ...options,
        headers,
    });


    if (response.status === 401) {
        authToken.set(null);
        window.location.href = '/login';
    }

    if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.message || `API request to ${path} failed`);
    }

    return response.json().then(data => data.result);
}

export default {
    get: (fetcher: typeof fetch, path: string) => request(fetcher, path, {method: 'GET'}),
    post: (fetcher: typeof fetch, path: string, data: any) => request(fetcher, path, {
        method: 'POST',
        body: JSON.stringify(data)
    }),
    // Add other methods like put, delete as needed
    login: async (fetcher: typeof fetch, username: string, password: string) => {
        const response = await request(fetcher, '/login', {
            method: 'POST',
            body: JSON.stringify({username, password}),
        })
        if (!response.token) {
            throw new Error('No token received');
        }
        return response;

    },
    listServers: (fetcher: typeof fetch) => request(fetcher, '/servers', {method: 'GET'}),

    getServer: (fetcher: typeof fetch, serverName: string) => request(fetcher, '/servers/' + serverName, {method: 'GET'}),
}
;