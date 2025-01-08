import axios from 'axios'

export function books(){
    var token = window.sessionStorage.getItem("jwt-token");
    return axios.create({
        baseURL: "http://localhost:3001",
        headers: {"Authorization": token}
    });
}

export function users(){
    var token = "" || window.sessionStorage.getItem("jwt-token");
    return axios.create({
        baseURL: "http://localhost:3000",
        headers: {"Authorization": token}
    });
}