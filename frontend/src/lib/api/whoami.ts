import {makeApiCall} from "$lib/api/makeApiCall";

export function whoAmI() {
    return makeApiCall("/whoami", "GET", null);
}