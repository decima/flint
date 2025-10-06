import type {ClientInit} from "@sveltejs/kit";
import {authToken} from "$lib/store/auth";
import {beforeNavigate} from "$app/navigation";

const AUTH_TOKEN_KEY = "authToken";
const AUTH_REFRESH_TOKEN_KEY = "refreshToken";
export const init: ClientInit = async () => {
        authToken.set({
            token: localStorage.getItem(AUTH_TOKEN_KEY),
            refresh: localStorage.getItem(AUTH_REFRESH_TOKEN_KEY)
        });

        authToken.subscribe(value => {
            if (value?.token) {
                localStorage.setItem(AUTH_TOKEN_KEY, value.token);
            } else {
                localStorage.removeItem(AUTH_TOKEN_KEY);
            }

            if (value?.refresh) {
                localStorage.setItem(AUTH_REFRESH_TOKEN_KEY, value.refresh);
            } else {
                localStorage.removeItem(AUTH_REFRESH_TOKEN_KEY);
            }
        });


        authToken.subscribe((value) => {
            if (window.location.pathname !== '/login' && !(value?.token)) {
                window.location.href = '/login';
                return;
            }
        });

    }
;