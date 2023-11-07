import ApiService from "./ApiService";

export async function apiGetProgramsData(data){
    return ApiService.fetchData({
        url:'/programs',
        method:'post',
        data
    })
}

export async function apiCreateProgramData(data){
    return ApiService.fetchData({
        url:'/programs/create',
        method:'post',
        data
    })
}

