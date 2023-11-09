import ApiService from "./ApiService";

export default function apiGetKegiatans(data){
    return ApiService.fetchData({
        url:'/kegiatans',
        method:'post',
        data
    })
}