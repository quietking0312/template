
export const isServer = typeof window === 'undefined'

export const isDef = <T = unknown>(val?: T): val is T => {
    return typeof val !== 'undefined'
}

export const isFunction = (val: unknown): val is Function => typeof val === 'function'

export const isUnDef = <T = unknown>(val?: T): val is T => {
    return !isDef(val)
}
