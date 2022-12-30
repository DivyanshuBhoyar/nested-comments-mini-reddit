import axios from "axios";

const api = axios.create({
  baseURL: import.meta.env.VITE_API_SERVER_URL,
  withCredentials: true,
});

export const makeRequest = (url, options) => {
  return api(url, options)
    .then((res) => res.data)
    .catch((error) =>
      Promise.reject(error?.response?.data?.message ?? "Error")
    );
};
