import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { apiGetBudget } from "src/services/BudgetService";

export const getBudget=createAsyncThunk(
    'budget/data/getBudget',
    async()=>{
        const response = await apiGetBudget()
        return response.data
    }
)

const dataSlice = createSlice({
    name:'budget/data',
    initialState:{
        loading:false,
        budgetData:{},
    },
    extraReducers:{
        [getBudget.pending]:(state)=>{
            state.loading=true
        },
        [getBudget.fulfilled]:(state,action)=>{
            state.loading=false
            state.budgetData=action.payload
        }
    }
})

export default dataSlice.reducer