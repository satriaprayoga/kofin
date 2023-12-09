import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { cloneDeep } from "lodash";
import { apiGetAvailablePrograms, apiGetBudgetList } from "src/services/BudgetService";

export const getPrograms=createAsyncThunk(
    'importBudgetProgram/getPrograms',
    async(data)=>{
       // console.log(data)
        const response = await apiGetAvailablePrograms(data)
       
        return response.data.data
    }
)

export const getBudgets=createAsyncThunk(
    'importBudgetProgram/getBudgets',
    async()=>{
        const response = await apiGetBudgetList()
        const optData=[]
        //console.log(response.data.data)
        response.data.data.forEach((d)=>{
           // console.log(d)
            const b = cloneDeep(d)
            optData.push({
                value:b.budget_id,
                label:b.budget_year+" "+b.budget_desc
            })
            
        })
      // console.log(optData)
        
        return optData
    }
)




const dataSlice = createSlice({
    name:'importBudgetProgram/data',
    initialState:{
        loading:false,
        programsData:[],
        budgetsData:[],
        budgetId:1
    },
    reducers:{
        updateProgramsData:(state,action)=>{
            state.programsData=action.payload
        },
        setBudgetId:(state,action)=>{
            state.budgetId=action.payload
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
        [getBudgets.fulfilled]:(state,action)=>{
            state.budgetsData=action.payload
            state.loading=false
        },
        [getBudgets.pending]:(state)=>{
            state.loading=true
        },
       
    }
})

export const{
    updateProgramsData,
    setBudgetId
}=dataSlice.actions

export default dataSlice.reducer


