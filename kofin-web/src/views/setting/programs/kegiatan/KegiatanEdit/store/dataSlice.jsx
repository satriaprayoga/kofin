import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { apiDeleteKegiatan, apiGetKegiatan, apiPutKegiatan } from "src/services/KegiatanService";


export const getKegiatan=createAsyncThunk(
    'kegiatanEdit/data/getKegiatan',
    async(data)=>{
        const response = await apiGetKegiatan(data)
        return response.data.data
    }
)

export const updateKegiatan=async (params,data)=>{
    const response = await apiPutKegiatan(params,data)
    return response.data
}

export const deleteKegiatan= async(data)=>{
    const response = await apiDeleteKegiatan(data)
    return response.data
}

const dataSlice = createSlice({
    name:'kegiatanEdit/data',
    initialState:{
        loading:false,
        kegiatan:{},
    },
    reducers:{},
    extraReducers:{
        [getKegiatan.fulfilled]:(state,action)=>{
            state.kegiatan=action.payload
            state.loading=false
        },
        [getKegiatan.pending]:(state)=>{
            state.loading=true
        },
    },
})

export default dataSlice.reducer