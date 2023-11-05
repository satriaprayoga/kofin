import ApiService from "./ApiService";

export async function apiGetUnitData(){
    return ApiService.fetchData({
        url:'/units',
        method:'get'
    })
}

export async function apiGetSubunitData(){
    return ApiService.fetchData({
        url:'/units/subunits',
        method:'get'
    })
}

export async function apiCreateSubunitData(data){
    return ApiService.fetchData({
        url:'/units/subunits/create',
        method:'post',
        data,
    })
}

export async function apiGetSubunit(params){
    return ApiService.fetchData({
        url:'/units/subunit',
        method:'get',
        params,
    })
}

export async function apiPutSubunit(data){
    return ApiService.fetchData({
        url:'/units/subunit/update',
        method:'put',
        data,
    })
}

export async function apiDeleteSubunit(data){
    return ApiService.fetchData({
        url:'/units/subunit/delete',
        method:'delete',
        data
    })
}