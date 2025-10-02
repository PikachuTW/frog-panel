import createClient from 'openapi-fetch';
import type { paths } from './schema';

export const api = createClient<paths>({ 
  baseUrl: 'http://localhost:5173/api'
});
