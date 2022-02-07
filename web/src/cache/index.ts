import WebStorageCache from 'web-storage-cache'
import {appInfoApi} from "@/api/login";
import config from "@/request/config";


const wsCache: WebStorageCache = new WebStorageCache({
    // storage: 'sessionStorage'
    storage: "localStorage",
    exp: 3600 * 12
})

export enum cacheKey {
    lang = "language",
    userInfo = "user_info",
    conf = "conf"
}

export default wsCache

export const GetConf = () => {
    let cacheConf = wsCache.get(cacheKey.conf)
    if (cacheConf) {
        return cacheConf
    } else {
        try {
            appInfoApi().then(res => {
                const {code, data} = res as any
                if (code == config.result_code) {
                    wsCache.set(cacheKey.conf, data)
                    return data
                }
            })
        } catch (err) {
            console.log(err)
            return null
        }
    }
}
