import {fetch} from "@/request/axios";

export const appInfoApi = () => {
    return fetch({url: "v1/appinfo", method: "get"})
}

export const loginApi = (data: any) => {
    return fetch({url: "v1/login", method: "post", data: data})
}

export const userInfoApi = (params: any) => {
    return fetch({url: "v1/userinfo", method: "get", params: params})
}

export const resetPassApi = (data:any) => {
    return fetch({url: "v1/user/password", method: "put", data:data})
}
