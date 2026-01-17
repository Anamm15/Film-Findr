import axios from "axios";
import { BASE_API_URL } from "../utils/constant";
const BASE_URL = `${BASE_API_URL}/users`;

export const getAllUser = async () => {
    const response = await axios.get(`${BASE_URL}/`, {
        withCredentials: true,
    });
    return response;
};

export const getUserByUsername = async (username) => {
    const response = await axios.get(`${BASE_URL}?username=${username}`, {
        withCredentials: true,
    });
    return response.data;
};

export const getUserById = async (id) => {
    const response = await axios.get(`${BASE_URL}/${id}`, {
        withCredentials: true,
    });
    return response;
};

export const getMe = async () => {
    const response = await axios.get(`${BASE_URL}/me`, {
        withCredentials: true,
    });
    console.log("response me: " + response);
    return response.data;
};

export const registerUser = async (data) => {
    const response = await axios.post(`${BASE_URL}/`, data);
    return response;
};

export const loginUser = async (data) => {
    const response = await axios.post(`${BASE_URL}/login`, data, {
        withCredentials: true,
    });
    return response;
};

export const logoutUser = async () => {
    const response = await axios.post(`${BASE_URL}/logout`, {}, {
        withCredentials: true,
    });
    return response;
};

export const updateUser = async (id, data) => {
    const response = await axios.patch(`${BASE_URL}/${id}`, data, {
        withCredentials: true,
    });
    return response;
};

export const deleteUser = async (id) => {
    const response = await axios.delete(`${BASE_URL}/${id}`, {
        withCredentials: true,
    });
    return response;
};
