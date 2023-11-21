import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { apiGetProgram } from "services/ProgramService";

export const getProgram=createAsyncThunk(
    'kegiatansNew/data/getProgram',
    async(data)=>{
        const response = await apiGetProgram(data)
        return response.data.data
    }
)

const dataSlice = createSlice({
    name:'kegiatansNew/data',
    initialState:{
        loading:false,
        programData:{},
    },
    reducers:{},
    extraReducers:{
        [getProgram.pending]:(state)=>{
            state.loading=true
        },
        [getProgram.fulfilled]:(state,action)=>{
            state.loading=false
            state.programData=action.payload
        }
    }
})


export default dataSlice.reducer