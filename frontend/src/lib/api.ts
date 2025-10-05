import {get} from 'svelte/store';
import {authToken} from './stores';

async function request(path: string, options: RequestInit = {}) {
    const token = get(authToken);

    const headers = new Headers(options.headers);
    if (token) {
        headers.set('Authorization', `Bearer ${token}`);
    }
    headers.set('Content-Type', 'application/json');

    const response = await fetch(`/api${path}`, {
        ...options,
        headers,
    });

    if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.message || `API request to ${path} failed`);
    }

    return response.json().then(data => data.result);
}

export default {
    get: (path: string) => request(path, {method: 'GET'}),
    post: (path: string, data: any) => request(path, {method: 'POST', body: JSON.stringify(data)}),
    // Add other methods like put, delete as needed
    login: (username: string, password: string) => request('/login', {
        method: 'POST',
        body: JSON.stringify({username, password}),
    }),
    listServers: () => request('/servers', {method: 'GET'}),

};