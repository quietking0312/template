import {fetch} from "@/request/axios";


export const DownloadFile = (params: any) => {
    return fetch({
        url: "v1/download",
        method: "get",
        responseType: 'blob',
        params
    })
}