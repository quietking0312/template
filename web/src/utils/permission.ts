import {userInfoStore} from "@/store/modules/userInfo";

export function CheckPermission(value: any) {
    if (value && value instanceof Array && value.length > 0) {
        const permissionIdList = userInfoStore.permissionIdList
        console.log(permissionIdList)
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
