
class Socket {
    public url: string // 链接地址
    private isLock: boolean // 防止重复链接
    public ws: WebSocket | undefined
    constructor(url: string, ) {
        this.url = url
        this.isLock = false
        this.__init__()
    }
    // 初始化
    private __init__(): void {
        let url = this.url
        if (url == undefined || url == '') {
            throw new Error("websocket链接地址不能为空")
        } else {
            this.wsInit()
        }
    }
    // 链接创建
    public wsInit(): void {
        if (this.isLock) {
            return
        }
        this.ws = new WebSocket(this.url)
        this.isLock = true
    }
    // 关闭链接
    public Close() {
        if (this.ws && (this.ws as WebSocket).readyState == 1) {
            this.ws?.close()
            this.isLock = false
        }
    }
}
