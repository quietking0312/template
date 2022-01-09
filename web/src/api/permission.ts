import {fetch} from "@/request/axios";

export const getUserListApi = (params:any) => {
    return fetch({url:"v1/permission/user/list", method: "get", params:params})
}

export const postUserApi = (data: any) => {
    return fetch({url:"v1/permission/user", method: "post", data:data})
}

export const updateUserApi = (data:any) => {
    return fetch({url: "v1/permission/user", method: "put", data: data})
}

export const deleteUserApi = (data:any) => {
    return fetch({url: "v1/permission/user", method: "delete", data: data})
}


export const getRoleListApi = (params: any) => {
    return fetch({url: "v1/permission/role/list", method: "get", params:params})
}

export const postRoleApi = (data:any) => {
    return fetch({url: "v1/permission/role", method: "post", data:data})
}

export const updateRoleApi = (data:any) => {
    return fetch({url: "v1/permission/role", method: "put", data: data})
}

export const deleteRoleApi =(data: any) => {
    return fetch({url: "v1/permission/role", method: "delete", data:data})
}

export const getPermissionListApi = () => {
    return fetch({url: "v1/permission", method: "get"})
}

export const postUserPermissionApi = (data: any) => {
    return fetch({url: "v1/permission/user/permission", method: "post", data:data})
}

export const postRolePermissionApi = (data: any) => {
    return fetch({url: "v1/permission/role/permission", method: "post", data: data})
}