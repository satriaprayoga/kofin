import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { apiGetProgramData } from "services/ProgramService";

export const getPrograms=createAsyncThunk(
    'programs/data/getProgram',
    async()=>{
        const response = await apiGetProgramData()
        return response.data
    }
)

const dataSlice = createSlice({
    name:'program/data',
    initialState:{
        loading:false,
        programsData:[]
    },
    reducers:{},
    extraReducers:{
        [getPrograms.pending]:(state)=>{
            state.loading=true
        },
        [getPrograms.fulfilled]:(state,action)=>{
            state.loading=true,
            state.programsData=[action.payload]
        }
    }
})

export default dataSlice.reducer