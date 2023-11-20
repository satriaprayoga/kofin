import ApiService from "./ApiService";

export async function apiGetUnitData(){
    return ApiService.fetchData({
        url:'/units/1',
        method:'get'
    })
}

export async function apiGetSubunitData(){
    return ApiService.fetchData({
        url:'/units/sub',
        method:'get'
    })
}

export async function apiCreateSubunitData(data){
    return ApiService.fetchData({
        url:'/units/',
        method:'post',
        data,
    })
}

export async function apiGetSubunit(params){
    return ApiService.fetchData({
        url:'/units/'+params,
        method:'get',
       
    })
}

export async function apiPutSubunit(params,data){
    return ApiService.fetchData({
        url:'/units/'+params,
        method:'put',
        data,
    })
}

export async function apiDeleteSubunit(data){
    return ApiService.fetchData({
        url:'/units/'+data,
        method:'delete',
    })
}