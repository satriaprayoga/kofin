import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { apiGetSubunitData } from "services/UnitService";

export const getSubunits=createAsyncThunk(
    'programsOption/getSubunits',
    async()=>{
        const response = await apiGetSubunitData()
        const optData=[]
        console.log(response.data)
        response.data.forEach((d)=>{
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

const dataSlice = createSlice({
    name:'programsOptions/data',
    initialState:{
        loading:false,
        subunitsData:[]
    },
    reducers:{},
    extraReducers:{
        [getSubunits.pending]:(state)=>{
            state.loading=true
        },
        [getSubunits.fulfilled]:(state,action)=>{
            state.loading=false,
            state.subunitsData=action.payload
        }
    }
})

export default dataSlice.reducer