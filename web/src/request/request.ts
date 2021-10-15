import axios, {AxiosError, AxiosInstance, AxiosRequestConfig, AxiosResponse} from "axios";
import { Message } from "@/components/Message";
import qs from 'qs'
import config from "@/request/config";

const { result_code, base_url} = config

export const PATH_URL: string = base_url[import.meta.env.VITE_MODE as string]

const service: AxiosInstance = axios.create({
    baseURL: PATH_URL,
    timeout: 500
})

service.interceptors.request.use(
    (config: AxiosRequestConfig) => {
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

interface respType {
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
        } else {
            if ((response.data as respType).code === result_code) {
                return response.data
            } else {
                Message.error((response.data as respType).message)
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
