import axios from "axios";
import { BASE_API_URL } from "../utils/constant";

const BASE_URL = `${BASE_API_URL}/reviews`;

export const getReviewByUserId = async (id, page) => {
    const response = await axios.get(`${BASE_URL}/user/${id}?page=${page}`, {
        withCredentials: true,
    });
    return response.data;
};

export const getReviewByFilmId = async (id, page) => {
    const response = await axios.get(`${BASE_URL}/film/${id}?page=${page}`, {
        withCredentials: true,
    });
    return response.data.data;
};

export const createReview = async (data) => {
    const response = await axios.post(`${BASE_URL}/`, data, {
        withCredentials: true,
    });
    return response.data;
};

export const updateReview = async (id, data) => {
    const response = await axios.put(`${BASE_URL}/${id}`, data, {
        withCredentials: true,
    });
    return response;
};

export const updateReaksiReview = async (id, data) => {
    const response = await axios.patch(`${BASE_URL}/${id}/reaction`, data, {
        withCredentials: true,
    });
    return response.data;
};

export const deleteReview = async (id) => {
    const response = await axios.delete(`${BASE_URL}/${id}`, {
        withCredentials: true,
    });
    return response.data;
};
