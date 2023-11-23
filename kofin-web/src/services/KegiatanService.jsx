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
        url:"/kegiatans/"+params,
        method:'get',
    })
}

export async function apiCreateKegiatanData(data){
    return ApiService.fetchData({
        url:'/kegiatans/',
        method:'post',
        data
    })
}

export async function apiPutKegiatan(params,data){
    return ApiService.fetchData({
        url:'/kegiatans/'+params,
        method:'put',
        data,
    })
}

export async function apiDeleteKegiatan(data){
    return ApiService.fetchData({
        url:'/kegiatans/'+data,
        method:'delete',
    })
}

export async function apiGetKegiatanProgram(params){
    return ApiService.fetchData({
        url:"/kegiatans/program/budget",
        method:'get',
        params
    })
}