import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { apiGetSubunitData } from "services/UnitService";
import { apiGetProgramUnit } from "src/services/ProgramService";

export const getPrograms=createAsyncThunk(
    'importProgram/getPrograms',
    async(data)=>{
        console.log(data)
        const response = await apiGetProgramUnit(data)
       
        return response.data
    }
)

export const getSubunits=createAsyncThunk(
    'importProgram/getSubunits',
    async()=>{
        const response = await apiGetSubunitData()
        const optData=[]
        response.data.forEach((d)=>{
            optData.push({
                value:d.unit_id,
                label:d.unit_name
            })
        })
        return optData
    }
)

const dataSlice = createSlice({
    name:'importProgram/data',
    initialState:{
        loading:false,
        programsData:[],
        subunitData:[],
        subunitId:0
    },
    reducers:{
        updateProgramsData:(state,action)=>{
            state.programsData=action.payload
        },
        setSubunitId:(state,action)=>{
            state.subunitId=action.payload
        }
    },
    extraReducers:{
        [getPrograms.pending]:(state)=>{
            state.loading=true
        },
        [getPrograms.fulfilled]:(state,action)=>{
            state.loading=false
            state.programsData=action.payload
        },
        [getSubunits.fulfilled]:(state,action)=>{
            state.subunitData=action.payload
            state.loading=false
        },
        [getSubunits.pending]:(state)=>{
            state.loading=true
        }
    }
})

export const{
    updateProgramsData,
    setSubunitId
}=dataSlice.actions

export default dataSlice.reducer


