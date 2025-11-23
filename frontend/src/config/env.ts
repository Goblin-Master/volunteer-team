export const API_BASE_URL: string = (import.meta.env.VITE_API_BASE_URL as string) ?? 'http://127.0.0.1:9000';
export const UPLOAD_BASE_URL: string = (import.meta.env.VITE_UPLOAD_BASE_URL as string) ?? 'http://127.0.0.1:9000/uploads/';
export const REQUEST_TIMEOUT: number = Number(import.meta.env.VITE_REQUEST_TIMEOUT ?? 10000);