/**
 * 对象数组深拷贝
 * @param {Array,Object} source 需要深拷贝的对象数组
 * @param {Array} noClone 不需要深拷贝的属性集合
 */
export function deepClone(source: any, noClone: string[] = []): any {
    if (!source && typeof source !== 'object') {
        throw new Error('error arguments deepClone')
    }
    const targetObj: any = source.constructor === Array ? [] : {}
    Object.keys(source).forEach((keys: string) => {
        if (source[keys] && typeof source[keys] === 'object' && noClone.indexOf(keys) === -1) {
            targetObj[keys] = deepClone(source[keys], noClone)
        } else {
            targetObj[keys] = source[keys]
        }
    })
    return targetObj
}



/**
 * @param {date} time 需要转换的时间 毫秒级
 * @param {String} fmt 需要转换的格式 如 yyyy-MM-dd、yyyy-MM-dd HH:mm:ss
 */
export function formatTime(time: any, fmt: string) {
    if (!time) return ''
    else {
        const date = new Date(time)
        const o = {
            'M+': date.getMonth() + 1,
            'd+': date.getDate(),
            'H+': date.getHours(),
            'm+': date.getMinutes(),
            's+': date.getSeconds(),
            'q+': Math.floor((date.getMonth() + 3) / 3),
            S: date.getMilliseconds()
        }
        if (/(y+)/.test(fmt)) {
            fmt = fmt.replace(RegExp.$1, (date.getFullYear() + '').substr(4 - RegExp.$1.length))
        }
        for (const k in o) {
            if (new RegExp('(' + k + ')').test(fmt)) {
                fmt = fmt.replace(
                    RegExp.$1,
                    RegExp.$1.length === 1 ? o[k] : ('00' + o[k]).substr(('' + o[k]).length)
                )
            }
        }
        return fmt
    }
}