import {userInfoStore} from "@/store/modules/userInfo";

export function CheckPermission(value: any):boolean {
    if (value && value instanceof Array && value.length > 0) {
        const permissionIdList = userInfoStore.permissionIdList
        return permissionIdList.some(permission_id => {
            if (permission_id === 100000) {
                return true
            }
            return value.includes(permission_id)
        })
    } else {
        return false
    }
}


export function PermissionListToTree(data: any) {
    if (!Array.isArray(data)) {
        return data
    }
    let map = {}
    data.forEach(item => {
        delete item.children
        if (item.permission_id % 100 === 1) {
            item.pid = 0
        } else {
            item.pid = Math.floor(item.permission_id / 100) * 100 + 1
        }
        map[item.permission_id] = item
    })
    for (let i = 0; i < data.length; i++) {
        if (data[i].pid && map[data[i].pid]) {
            if (!map[data[i].pid].children) {
                map[data[i].pid].children = []
            }
            map[data[i].pid].children.push(data[i])
            data.splice(i, 1)
            i--
        }
    }
    return data
}