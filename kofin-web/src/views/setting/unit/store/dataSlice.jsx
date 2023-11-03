import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { apiGetUnitData } from "services/UnitService";

export const getUnit = createAsyncThunk(
    'unit/data/getUnit',
    async()=>{
        const response = await apiGetUnitData()
        return response.data
    }
)
const dataSlice = createSlice({
    name:'unit/data',
    initialState:{
        loading:false,
        unitData:{}
    },
    reducers:{},
    extraReducers:{
        [getUnit.pending]:(state)=>{
            state.loading=true
        },
        [getUnit.fulfilled]:(state,action)=>{
            state.loading=false,
            state.unitData=[action.payload]
        }
    }
})

export default dataSlice.reducer