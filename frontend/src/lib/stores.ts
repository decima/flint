import { writable } from 'svelte/store';

export const authToken = writable(localStorage.getItem('authToken'));

authToken.subscribe(token => {
  if (token) {
    localStorage.setItem('authToken', token);
  } else {
    localStorage.removeItem('authToken');
  }
});