import {fetch} from "@/request/axios";

export const loginApi = (data: any) => {
    return fetch({url: "v1/login", method: "post", data: data, headersType: 'application/x-protobuf'})
}
