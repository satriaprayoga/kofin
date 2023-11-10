import ApiService from "./ApiService";

export default function apiGetKegiatans(data){
    return ApiService.fetchData({
        url:'/kegiatans',
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