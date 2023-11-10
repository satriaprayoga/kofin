import ApiService from "./ApiService";

export async function apiGetBudget(){
    return ApiService.fetchData({
        url:'/budget',
        method:'get'
    })
}

export async function apiPutBudget(data){
    return ApiService.fetchData({
        url:'/budget/setup',
        method:'put',
        data
    })
}