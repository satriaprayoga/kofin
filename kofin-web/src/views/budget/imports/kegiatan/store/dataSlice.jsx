import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { cloneDeep } from "lodash";
import { apiGetAvailableKegiatans, apiGetAvailablePrograms, apiGetBudgetList, apiGetImportedPrograms, apiImportProgramBudget } from "src/services/BudgetService";

export const getPrograms=createAsyncThunk(
    'importBudgetKegiatan/getPrograms',
    async(data)=>{
       // console.log(data)
        const response = await apiGetImportedPrograms(data)
        const optData=[]
       
        response.data.data.forEach((d)=>{
            // console.log(d)
             const b = cloneDeep(d)
             optData.push({
                 value:b.expend_program_id,
                 label:b.program_name
             })
             
         })
         return optData
    }
)

export const getKegiatans=createAsyncThunk(
    'importBudgetKegiatan/getKegiatans',
    async(data)=>{
       // console.log(data)
        const response = await apiGetAvailableKegiatans(data)
       
        return response.data.data
    }
)

export const getBudgets=createAsyncThunk(
    'importBudgetKegiatan/getBudgets',
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

export const importProgram=async(rows)=>{
    rows.forEach(async (row,idx)=>{
        const data = cloneDeep(row.original)
        //console.log(data)
        data.included=true
        const success = await apiImportProgramBudget(data.expend_program_id,data)
        if (!success){
            return false
        }
        
    })
    
   return true
}




const dataSlice = createSlice({
    name:'importBudgetKegiatan/data',
    initialState:{
        loading:false,
        kegiatansData:[],
        budgetsData:[],
        budgetId:1
    },
    reducers:{
        updateKegiatansData:(state,action)=>{
            state.kegiatansData=action.payload
        },
        setBudgetId:(state,action)=>{
            state.budgetId=action.payload
        }
    },
    extraReducers:{
        [getKegiatans.pending]:(state)=>{
            state.loading=true
        },
        [getKegiatans.fulfilled]:(state,action)=>{
            state.loading=false
            state.kegiatansData=action.payload
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
    updateKegiatansData,
    setBudgetId
}=dataSlice.actions

export default dataSlice.reducer


