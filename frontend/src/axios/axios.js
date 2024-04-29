import axios from "axios";

export const axiosInstance = axios.create({
    baseURL: "http://localhost:8000",
})

export const MINIO = "http://localhost:9000/images/"
