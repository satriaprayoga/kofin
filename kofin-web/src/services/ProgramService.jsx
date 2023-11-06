import ApiService from "./ApiService";

export function apiGetProgramsData(){
    return ApiService.fetchData({
        url:'/programs',
        method:'get'
    })
}

export function apiGetProgramData(params){
    return ApiService.fetchData({
        url:'/programs/get',
        method:'get',
        params
    })
}

export function apiGetProgramDataByUnit(params){
    return ApiService.fetchData({
        url:'/programs/unit/',
        method:'get',
        params
    })
}
