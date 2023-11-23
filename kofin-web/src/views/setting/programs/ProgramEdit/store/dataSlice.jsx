import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { apiDeleteProgram, apiGetProgram, apiPutProgram } from "src/services/ProgramService";
import { apiGetSubunitData } from "src/services/UnitService";


export const getProgram=createAsyncThunk(
    'programEdit/data/getProgram',
    async(data)=>{
        const response = await apiGetProgram(data)
        return response.data.data
    }
)

export const getSubunits=createAsyncThunk(
    'programEdit/data/getSubunits',
    async()=>{
        const response = await apiGetSubunitData()
        const optData=[]
      
        response.data.data.forEach((d)=>{
            optData.push({
                value:d.unit_id,
                label:d.unit_name,
                kode:d.unit_kode
            })
        })
        console.log(optData)
        return optData
    }
)

export const updateProgram=async (params,data)=>{
    const response = await apiPutProgram(params,data)
    return response.data
}

export const deleteProgram= async(data)=>{
    const response = await apiDeleteProgram(data)
    return response.data
}

const dataSlice = createSlice({
    name:'programEdit/data',
    initialState:{
        loading:false,
        program:{},
        subunitData:[],
    },
    reducers:{},
    extraReducers:{
        [getSubunits.fulfilled]:(state,action)=>{
            state.subunitData=action.payload
            state.loading=false
        },
        [getSubunits.pending]:(state)=>{
            state.loading=true
        },
        [getProgram.fulfilled]:(state,action)=>{
            state.program=action.payload
            state.loading=false
        },
        [getProgram.pending]:(state)=>{
            state.loading=true
        },
    },
})

export default dataSlice.reducer