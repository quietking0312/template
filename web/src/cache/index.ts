import WebStorageCache from 'web-storage-cache'


const wsCache: WebStorageCache = new WebStorageCache({
    // storage: 'sessionStorage'
    storage: "localStorage"
})

export enum cacheKey {
    lang = "language",
    userInfo = "user_info",
    conf = "conf"
}

export default wsCache
