import {Action, getModule, Module, Mutation, VuexModule} from "vuex-module-decorators";
import store from "@/store";
import {userInfoApi} from "@/api/login";
import wsCache, {cacheKey} from "@/cache";




export interface UserInfoState {
    permissionIdList: number[]
}

@Module({dynamic: true, namespaced: true, store, name: "userinfo"})
class UserInfo extends VuexModule implements UserInfoState {
    public permissionIdList = [] as number[]

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
                    const permission_id = (data as object)['permission_id']
                    this.SET_PERMISSIONIDLIST(permission_id as number[])
                    resolve(permission_id)
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
