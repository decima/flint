import { authToken } from './store/stores';
import api from "./store/api"; // Une fonction utilitaire à créer
export async function handle({ event, resolve }) {
    // 1. Lire le cookie de session
    const token = event.cookies.get('session');


    if (token) {
        authToken.set(token)
    }

    // Continue le traitement de la requête
    const response = await resolve(event);
    return response;
}