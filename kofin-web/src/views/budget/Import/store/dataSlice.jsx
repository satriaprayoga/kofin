import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { apiGetProgramsData } from "services/ProgramService";
import { apiGetSubunitData } from "services/UnitService";

export const getPrograms=createAsyncThunk(
    'programs/data/getPrograms',
    async(data)=>{
        const response = await apiGetProgramsData(data)
        return response.data
    }
)

export const getSubunits=createAsyncThunk(
    'program/data/getSubunits',
    async()=>{
        const response = await apiGetSubunitData()
        const optData=[]
        response.data.forEach((d)=>{
            optData.push({
                value:d.unit_name,
                label:d.unit_name
            })
        })
        return optData
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
    name:'program/data',
    initialState:{
        loading:false,
        programsData:[],
        tableData:initialTableData,
        subunitData:[],
    },
    reducers:{
        updateProgramsData:(state,action)=>{
            state.programsData=action.payload
        },
        setTableData:(state,action)=>{
            state.tableData=action.payload
        }
    },
    extraReducers:{
        [getPrograms.pending]:(state)=>{
            state.loading=true
        },
        [getPrograms.fulfilled]:(state,action)=>{
            state.loading=false
            state.programsData=action.payload.data
            state.tableData.total=action.payload.total
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
    setTableData,
}=dataSlice.actions

export default dataSlice.reducer