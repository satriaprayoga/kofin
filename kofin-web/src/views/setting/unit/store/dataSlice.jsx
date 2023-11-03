import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { apiGetUnitData } from "services/UnitService";
import { apiGetSubunitData } from "src/services/UnitService";

export const getUnit = createAsyncThunk(
    'unit/data/getUnit',
    async()=>{
        const response = await apiGetUnitData()
        return response.data
    }
)

export const getSubunits=createAsyncThunk(
    'unit/data/getSubunits',
    async()=>{
        const response = await apiGetSubunitData()
        return response.data
    }
)

const dataSlice = createSlice({
    name:'unit/data',
    initialState:{
        loading:false,
        unitData:{},
        subunitsData:[]
    },
    reducers:{},
    extraReducers:{
        [getUnit.pending]:(state)=>{
            state.loading=true
        },
        [getUnit.fulfilled]:(state,action)=>{
            state.loading=false,
            state.unitData=[action.payload]
        },
        [getSubunits.pending]:(state)=>{
            state.loading=true
        },
        [getSubunits.fulfilled]:(state,action)=>{
            state.loading=false,
            state.subunitsData=[action.payload]
        }
    }
})

export default dataSlice.reducer