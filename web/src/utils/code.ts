
export function ArrayBufferToStr(arrayBuffer: string, func: any) {
    let blob = new Blob([arrayBuffer])
    let reader = new FileReader()
    reader.readAsText(blob, 'utf-8')
    reader.onload = () => {
        if (func) {
            func.call(null, reader.result)
        }
    }
}

export function StrToArrayBuffer(str: string, func: any) {
    let blob = new Blob([str], {type: 'text/plain'})
    let reader = new FileReader()
    reader.readAsArrayBuffer(blob)
    reader.onload = () => {
        if (func) {
            func.call(null, reader.result)
        }
    }
}

export function Uint16ToBytes(value: any) {
    let a = new Uint8Array(2)
    a[1] = (value >> 8) & 0xFF
    a[0] = value & 0xFF
    return a
}


export function BytesToUint16(value: any) {
    let a = new Uint16Array(1)
    a[0] = a[0] | ((value[1] & 0xFF) << 8)
    a[0] = a[0] | (value[0] & 0xFF)
    return a[0]
}
