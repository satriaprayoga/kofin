import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { cloneDeep } from "lodash";
import { apiImportKegiatanBudget} from "services/BudgetService";
import { apiGetAllPrograms } from "services/ProgramService";
import { apiGetKegiatanProgram } from "src/services/KegiatanService";

export const getKegiatans=createAsyncThunk(
    'importKegiatan/getKegiatans',
    async(data)=>{
       // console.log(data)
        const response = await apiGetKegiatanProgram(data)
       
        return response.data
    }
)

export const getPrograms=createAsyncThunk(
    'importKegiatan/getPrograms',
    async()=>{
        const response = await apiGetAllPrograms()
        const optData=[]
        response.data.forEach((d)=>{
            optData.push({
                value:d.id,
                label:d.program_name
            })
        })
        return optData
    }
)

export const importKegiatan=async(rows)=>{
        rows.forEach(async (row,idx)=>{
            const data = cloneDeep(row.original)
            data.included=true
            const success = await apiImportKegiatanBudget(data)
            if (!success){
                return false
            }
        })
        
       return true
    }


const dataSlice = createSlice({
    name:'importKegiatan/data',
    initialState:{
        loading:false,
        kegiatansData:[],
        programData:[],
        programId:0
    },
    reducers:{
        updateKegiatansData:(state,action)=>{
            state.kegiatansData=action.payload
        },
        setProgramId:(state,action)=>{
            state.programId=action.payload
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
        [getPrograms.fulfilled]:(state,action)=>{
            state.programData=action.payload
            state.loading=false
        },
        [getPrograms.pending]:(state)=>{
            state.loading=true
        },
       
    }
})

export const{
    updateKegiatansData,
    setProgramId
}=dataSlice.actions

export default dataSlice.reducer


