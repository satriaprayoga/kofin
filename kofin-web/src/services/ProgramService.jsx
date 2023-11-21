import ApiService from "./ApiService";

export async function apiGetProgramsData(data){
    return ApiService.fetchData({
        url:'/programs/paginate',
        method:'post',
        data
    })
}

export async function apiGetAllPrograms(){
    return ApiService.fetchData({
        url:'/programs/budget',
        method:'get',
    })
}

export async function apiCreateProgramData(data){
    return ApiService.fetchData({
        url:'/programs/',
        method:'post',
        data
    })
}

export async function apiGetProgram(params){
    return ApiService.fetchData({
        url:'/programs/'+params,
        method:'get',
    })
}

export async function apiGetProgramUnit(params){
    return ApiService.fetchData({
        url:"/programs/unit",
        method:'get',
        params
    })
}

export async function apiGetProgramKegiatans(params){
    return ApiService.fetchData({
        url:'/kegiatans/program',
        method:'get',
        params
    })
}

export async function apiPutProgram(params,data){
    return ApiService.fetchData({
        url:'/programs/'+params,
        method:'put',
        data,
    })
}

export async function apiDeleteProgram(params){
    return ApiService.fetchData({
        url:'/programs/'+params,
        method:'delete',
    })
}
