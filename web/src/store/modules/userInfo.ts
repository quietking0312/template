import {Action, getModule, Module, Mutation, VuexModule} from "vuex-module-decorators";
import store from "@/store";
import {userInfoApi} from "@/api/login";
import wsCache, {cacheKey} from "@/cache";


export interface UserInfoState {
    uid: string
    name: string
    userName: string
    lastLoginTime: number
    roles: any[]
    permissionIdList: number[]
}

@Module({dynamic: true, namespaced: true, store, name: "userinfo"})
class UserInfo extends VuexModule implements UserInfoState {
    public uid = ""
    public name = ""
    public userName = ""
    public lastLoginTime = 0
    public roles = [] as any[]
    public permissionIdList = [] as number[]

    @Mutation
    private SET_UID(uid: string): void {
        this.uid = uid
    }

    @Mutation
    private SET_NAME(name: string): void {
        this.name = name
    }

    @Mutation
    private SET_USERNAME(username: string): void {
        this.userName = username
    }

    @Mutation
    private SET_LAST_LOGIN_TIME(lastLoginTime: number): void {
        this.lastLoginTime = lastLoginTime
    }

    @Mutation
    private SET_ROLES(roles: any[]): void {
        this.roles = roles
    }

    @Mutation
    private SET_PERMISSIONIDLIST(permissionIdList: number[]): void {
        this.permissionIdList = permissionIdList
    }

    @Action
    public SetUserInfo(): Promise<unknown> {
        return new Promise<unknown>((resolve, reject) => {
            const params = { token: wsCache.get(cacheKey.userInfo)}
            userInfoApi(params).then(res => {
                const {code, data} = res as any
                if (code === 0) {
                    let permission: number[]
                    permission = data.permissionIds? data.permissionIds: []
                    this.SET_UID(data.uid)
                    this.SET_NAME(data.name)
                    this.SET_USERNAME(data.userName)
                    this.SET_ROLES(data.roles)
                    this.SET_LAST_LOGIN_TIME(data.lastLoginTime)
                    this.SET_PERMISSIONIDLIST(permission)
                    resolve(permission)
                } else if (code === 501) {
                    reject()
                }
            }).catch(err => {
                reject(err)
            })
        })
    }
    @Action
    public resetToken(): Promise<unknown> {
        return new Promise<unknown>((resolve) => {
            wsCache.delete(cacheKey.userInfo)
            resolve("")
        })
    }
}

export const userInfoStore = getModule<UserInfo>(UserInfo)
