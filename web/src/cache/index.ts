
import WebStorageCache from 'web-storage-cache'

const wsCache: WebStorageCache = new WebStorageCache({
    storage: 'sessionStorage'
})

export default wsCache
