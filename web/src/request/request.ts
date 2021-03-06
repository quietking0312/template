import axios, {AxiosError, AxiosInstance, AxiosRequestConfig, AxiosResponse} from "axios";
import { Message } from "@/components/Message";
import qs from 'qs'
import config from "@/request/config";
import JSON_BIG from 'json-bigint';
import wsCache, {cacheKey} from "@/cache";
import {useUserInfoStoreWithOut} from "@/store/modules/userInfo";


const { result_code, base_url, request_timeout} = config
console.log(import.meta.env)
export const PATH_URL: string = base_url[import.meta.env.VITE_API_BASE_PATH as string]

const service: AxiosInstance = axios.create({
    baseURL: PATH_URL,
    timeout: request_timeout,
    transformResponse: data => {
        try {
            return JSON_BIG.parse(data)
        }catch (err) {
            console.log(err)
            try {
                return JSON.parse(data);
            }catch (err) {
                return data
            }
        }
    }
})

service.interceptors.request.use(
    (config: AxiosRequestConfig) => {
        if (wsCache.get(cacheKey.userInfo)) {
            (config.headers as object)["auth"] = wsCache.get(cacheKey.userInfo)
        }
        if (config.method === 'post' && config.headers && config.headers['Content-Type'] ==  'application/x-www-form-urlencoded') {
            config.data = qs.stringify(config.data)
        }
        if (config.method === 'get') {
            config.paramsSerializer = params => {
                return qs.stringify(params, {indices: false})
            }
        }
        return config
    },
    (error: AxiosError) => {
        console.log(error)
        return Promise.reject(error)
    }
)

export interface respType {
    code: number,
    message: string,
    data: any
}

service.interceptors.response.use(
    (response: AxiosResponse) => {
        const contextType = response.headers['content-type'] ? response.headers['content-type'] : response.headers['Content-Type']
        if (contextType.indexOf('application/octet-stream') !== -1) {
            let fileName = response.headers['Content-Disposition'] ? response.headers['Content-Disposition'] : response.headers['content-disposition']
            if (fileName && fileName.length >= 2) {
                fileName = fileName.split('=')[1]
            }
            fileName = decodeURIComponent(fileName)

            const url = window.URL.createObjectURL(new Blob([response.data as BlobPart]))
            const link = document.createElement('a')
            link.style.display = 'none'
            link.href = url
            link.setAttribute('download', fileName)
            document.body.appendChild(link)
            link.click()
        } else if (contextType.indexOf('application/x-protobuf') !== -1) {
            return response.data
        }else {
            if ((response.data as respType).code === result_code) {
                return response.data
            } else {
                Message.error((response.data as respType).message)
                if ((response.data as respType).code >= 600) {
                    return response.data
                } else if ((response.data as respType).code == 501) {
                    const userInfoStore = useUserInfoStoreWithOut()
                    userInfoStore.resetToken().then()
                }else {
                    return response.data
                }
            }
        }
    },
    (error: AxiosError) => {
        console.log("err:" + error)
        Message.error(error.message)
        return Promise.reject(error)
    }
)

export default service
