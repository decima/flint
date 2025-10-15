import {error} from '@sveltejs/kit';
import type {PageLoad} from './$types';
import {getServerSummary} from "$lib/api/servers";

export const load: PageLoad = async ({params}) => {

    return {
        server: params.name,
    }

    error(404, 'Not found');
};