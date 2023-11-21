import ApiService from "./ApiService";

export default function apiGetKegiatans(params,data){
    return ApiService.fetchData({
        url:'/kegiatans/paginate/'+params,
        method:'post',
        data
    })
}

export async function apiGetKegiatan(params){
    return ApiService.fetchData({
        url:"/kegiatans/get",
        method:'get',
        params
    })
}

export async function apiCreateKegiatanData(data){
    return ApiService.fetchData({
        url:'/kegiatans/create',
        method:'post',
        data
    })
}

export async function apiPutKegiatan(data){
    return ApiService.fetchData({
        url:'/kegiatans/update',
        method:'put',
        data,
    })
}

export async function apiDeleteKegiatan(data){
    return ApiService.fetchData({
        url:'/kegiatans/delete',
        method:'delete',
        data
    })
}

export async function apiGetKegiatanProgram(params){
    return ApiService.fetchData({
        url:"/kegiatans/program/budget",
        method:'get',
        params
    })
}