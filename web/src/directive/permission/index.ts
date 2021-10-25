import {CheckPermission} from "@/utils/permission";


function checkPermission(el: HTMLElement, binding: any) {
    const { value } = binding
    console.log(value)
    if (value && value instanceof Array) {
        if (value.length > 0) {
            const hasPermission = CheckPermission(value)
            console.log(hasPermission)
            if (!hasPermission) {
                console.log(el)
                el.parentNode && el.parentNode.removeChild(el)
            }
        }
    }
}

export default {
    mounted(el: any, binding: any) {
        checkPermission(el, binding)
    },
    updated(el: any, binding: any) {
        checkPermission(el, binding)
    }
}
