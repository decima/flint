import {writable} from 'svelte/store';

export const authToken = writable(null);

authToken.subscribe(token => {
    console.log("token is set")
});