// src/routes/login/+page.server.js

import {fail, redirect} from '@sveltejs/kit';
import api from '../../store/api'; // Assurez-vous que le chemin est correct

// ===================================================================
// FONCTION LOAD : S'exécute en GET pour préparer la page
// ===================================================================
/** @type {import('./$types').PageServerLoad} */
export const load = async ({cookies}) => {
    // Rôle : Si l'utilisateur a déjà un cookie de session valide,
    // on le redirige vers le dashboard sans même afficher la page de login.
    const token = cookies.get('session');
    if (token) {
        // Note : Il serait encore mieux de valider le token ici avant de rediriger
        throw redirect(303, '/dashboard');
    }

    // Si pas de token, on ne fait rien et on laisse la page s'afficher.
    return {};
};

// ===================================================================
// ACTIONS : S'exécute en POST quand le formulaire est soumis
// ===================================================================
/** @type {import('./$types').Actions} */
export const actions = {
    // Le nom 'default' peut être utilisé si vous n'avez qu'une seule action.
    // Si vous l'appelez 'login', le formulaire doit avoir <form action="?/login" ...>
    login: async ({request, cookies, fetch}) => {
        const data = await request.formData();
        const username = data.get('username');
        const password = data.get('password');

        if (!username || !password) {
            return fail(400, {error: 'Veuillez remplir tous les champs.'});
        }

        // Appel à votre API externe pour valider les identifiants
        const response = await api.login(fetch, username.toString(), password.toString());

        // Si c'est bon, on extrait le token
        const {token, refreshToken} = response; // ou await response.json() selon ce que votre API retourne

        // On configure le cookie de session
        cookies.set('session', token, {
            path: '/',
            httpOnly: true,
            secure: false,
            sameSite: 'lax',
            maxAge: 60 * 60 * 2
        });

        // Et on redirige vers le dashboard
        throw redirect(303, '/dashboard');
    }
};