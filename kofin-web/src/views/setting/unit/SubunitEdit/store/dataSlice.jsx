import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { apiDeleteSubunit, apiGetSubunit, apiPutSubunit } from "services/UnitService";

export const getSubunit = createAsyncThunk(
    'subunitEdit/data/getSubunit',
    async (data) => {
        const response = await apiGetSubunit(data)
        return response.data
    }
)

export const updateSubunit= async(data)=>{
    const response = await apiPutSubunit(data)
    return response.data
}

export const deleteSubunit= async(data)=>{
    const response = await apiDeleteSubunit(data)
    return response.data
}

const dataSlice = createSlice({
    name:'subunitEdit/data',
    initialState:{
        loading:false,
        subunitData:[],
    },
    reducers:{},
    extraReducers:{
        [getSubunit.fulfilled]:(state,action)=>{
            state.subunitData=action.payload
            state.loading=false
        },
        [getSubunit.pending]:(state)=>{
            state.loading=true
        },
    },
})

export default dataSlice.reducer