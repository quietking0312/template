import {fetch} from "@/request/axios";

export const appInfoApi = () => {
    return fetch({url: "v1/appinfo", method: "get"})
}

export const loginApi = (data: any) => {
    return fetch({url: "v1/login", method: "post", data: data})
}
