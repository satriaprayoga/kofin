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

export async function apiGetProgram(params){
    return ApiService.fetchData({
        url:'/programs/get',
        method:'get',
        params,
    })
}

export async function apiGetProgramKegiatans(params){
    return ApiService.fetchData({
        url:'/kegiatans/program',
        method:'get',
        params
    })
}

export async function apiPutProgram(data){
    return ApiService.fetchData({
        url:'/programs/update',
        method:'put',
        data,
    })
}

export async function apiDeleteProgram(data){
    return ApiService.fetchData({
        url:'/programs/delete',
        method:'delete',
        data
    })
}
