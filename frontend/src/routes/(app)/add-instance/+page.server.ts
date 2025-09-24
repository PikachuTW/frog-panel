import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({fetch}) => {
    return {
        versionTypes: ['fabric', 'vanilla'],
    };
};