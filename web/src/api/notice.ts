import { fetch } from "@/request/axios";

export const getNotice = () => {
    return fetch({ url: "v1/notice", method: "get"} )
}
