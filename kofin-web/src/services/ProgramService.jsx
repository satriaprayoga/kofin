import ApiService from "./ApiService";

export async function apiGetProgramsData(data){
    return ApiService.fetchData({
        url:'/programs',
        method:'post',
        data
    })
}

