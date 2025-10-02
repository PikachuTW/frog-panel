import { api } from '$lib/api/client';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async () => {
    return {
        versionTypes: (await api.GET('/info/version-types')).data ?? [],
    };
};