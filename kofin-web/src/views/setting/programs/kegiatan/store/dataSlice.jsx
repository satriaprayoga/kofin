import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import apiGetKegiatans from "services/KegiatanService";
import { apiGetProgram } from "services/ProgramService";

export const getProgramKegiatans=createAsyncThunk(
    'kegiatans/data/getProgramKegiatans',
    async(params,data)=>{
        const response = await apiGetKegiatans(params,data)
        return response.data
    }

)

export const getProgram=createAsyncThunk(
    'kegiatans/data/getProgram',
    async(data)=>{
        const response = await apiGetProgram(data)
        return response.data.data
    }
)

export const initialTableData={
    total:0,
    pageIndex:1,
    pageSize:10,
    query:'',
    sort:{
        order:'',
        key:''
    }
}

const dataSlice = createSlice({
    name:'kegiatans/data',
    initialState:{
        loading:false,
        programData:{},
        tableData:initialTableData,
        kegiatansData:[]
    },
    reducers:{
        updateKegiatansData:(state,action)=>{
            state.kegiatansData=action.payload
        },
        setTableData:(state,action)=>{
            state.tableData=action.payload
        }
    },
    extraReducers:{
        [getProgramKegiatans.pending]:(state)=>{
            state.loading=true
        },
        [getProgramKegiatans.fulfilled]:(state,action)=>{
            state.loading=false
            state.kegiatansData=action.payload.data
            state.tableData.total=action.payload.total
        },
        [getProgram.pending]:(state)=>{
            state.loading=true
        },
        [getProgram.fulfilled]:(state,action)=>{
            state.loading=false
            state.programData=action.payload
        }
    }
})

export const{
    updateKegiatansData,
    setTableData,
}=dataSlice.actions

export default dataSlice.reducer
