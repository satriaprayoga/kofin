import ApiService from "./ApiService";

export async function apiGetBudgetList(){
    return ApiService.fetchData({
        url:'/budgets/',
        method:'get'
    })
}

export async function apiGetActiveBudget(){
    return ApiService.fetchData({
        url:'/budgets/setup/active',
        method:'get',
    })
}

export async function apiGetBudget(){
    return ApiService.fetchData({
        url:'/budgets/1',
        method:'get'
    })
}

export async function apiCreateBudget(data){
    return ApiService.fetchData({
        url:'/budgets/setup/',
        method:'post',
        data,
    })
}

export async function apiPutBudget(data){
    return ApiService.fetchData({
        url:'/budget/setup',
        method:'put',
        data
    })
}

export async function apiGetAvailablePrograms(params){
    return ApiService.fetchData({
        url:'/expend_programs/available/'+params,
        method:'get',
    })
}

export async function apiGetAvailableKegiatans(data){
    return ApiService.fetchData({
        url:'/expend_kegiatans/available',
        method:'get',
        data
    })
}

export async function apiImportProgramBudget(params,data){
    return ApiService.fetchData({
        url:'/expend_programs/import/'+params,
        method:'put',
        data
    })
}

export async function apiImportKegiatanBudget(data){
    return ApiService.fetchData({
        url:'/kegiatans/import',
        method:'put',
        data
    })
}
