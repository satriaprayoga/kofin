import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { apiGetBudget, apiGetBudgetList } from "services/BudgetService";

export const getBudgets=createAsyncThunk(
    'budgets/data/getBudgets',
    async()=>{
        const response = await apiGetBudgetList()
       // console.log(response.data.data)
        return response.data.data
    }
)

const dataSlice = createSlice({
    name:'budgets/data',
    initialState:{
        loading:false,
        budgetsData:[],
    },
    extraReducers:{
        [getBudgets.pending]:(state)=>{
            state.loading=true
        },
        [getBudgets.fulfilled]:(state,action)=>{
            state.loading=false
            state.budgetsData=action.payload
        }
    }
})

export default dataSlice.reducer