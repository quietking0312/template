import { store } from "@/store";
import {userInfoApi} from "@/api/login";
import wsCache, {cacheKey} from "@/cache";
import {defineStore} from "pinia";


export interface UserInfoState {
    uid: string
    name: string
    userName: string
    lastLoginTime: number
    roles: any[]
    permissionIdList: number[]
}
export const useUserInfoStore = defineStore({
  id: "userInfo",
    state: (): UserInfoState => ({
        uid: "",
        name: "",
        userName: "",
        lastLoginTime: 0,
        roles: [],
        permissionIdList: []
    }),
    getters: {
      getUid(): string {
          return this.uid
      }
    },
    actions: {
        SetUserInfo(): Promise<unknown> {
            return new Promise<unknown>((resolve, reject) => {
                const params = { token: wsCache.get(cacheKey.userInfo)}
                userInfoApi(params).then(res => {
                    const {code, data} = res as any
                    if (code === 0) {
                        let permission: number[]
                        permission = data.permissionIds? data.permissionIds: []
                        this.uid = data.uid
                        this.name = data.name
                        this.userName = data.userName
                        this.roles = data.roles
                        this.lastLoginTime = data.lastLoginTime
                        this.permissionIdList = permission
                        resolve(permission)
                    } else if (code === 501) {
                        reject()
                    }
                }).catch(err => {
                    reject(err)
                })
            })
        },

        resetToken(): Promise<unknown> {
            return new Promise<unknown>((resolve) => {
                wsCache.delete(cacheKey.userInfo)
                resolve("")
            })
        }
    }
})

export function useUserInfoStoreWithOut() {
    return useUserInfoStore(store)
}
