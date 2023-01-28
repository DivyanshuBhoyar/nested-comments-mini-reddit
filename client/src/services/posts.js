import { makeRequest } from "./makeRequest";

export const getPosts = () => makeRequest("/posts");

export const fetchSinglePost = (id) => makeRequest(`/posts/${id}`);
