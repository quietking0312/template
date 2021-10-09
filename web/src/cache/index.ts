
import WebStorageCache from 'web-storage-cache'

const wsCache: WebStorageCache = new WebStorageCache({
    storage: 'sessionStorage'
})

/**
 * wsCache key
 */
export enum cacheKey {
    userInfo = "user_info",
    lang = "language"
}

export default wsCache
