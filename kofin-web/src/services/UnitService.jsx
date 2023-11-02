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