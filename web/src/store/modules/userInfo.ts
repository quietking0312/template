import {Action, getModule, Module, Mutation, VuexModule} from "vuex-module-decorators";
import store from "@/store";
import {userInfoApi} from "@/api/login";
import wsCache, {cacheKey} from "@/cache";
import internal from "stream";



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
                const permission_id = (res.data as object)['permission_id']
                this.SET_PERMISSIONIDLIST(permission_id as number[])
                resolve(permission_id)
            }).catch(err => {
                reject(err)
            })
        })

    }
}

export const userInfoStore = getModule<UserInfo>(UserInfo)
