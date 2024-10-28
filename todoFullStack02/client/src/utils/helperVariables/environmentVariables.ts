const mode = import.meta.env.VITE_MODE

export const apiBaseUrl = mode === "development" ? import.meta.env.VITE_API_BASE_URL : "/api"