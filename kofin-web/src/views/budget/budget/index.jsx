import React, { useEffect } from "react";
import { AdaptableCard, Loading } from "components/shared";
import { Button, Timeline } from "components/ui";
import { injectReducer } from "src/store";
import reducer from "./store";
import { useDispatch, useSelector } from "react-redux";
import { getBudgets } from "./store/dataSlice";
import { isEmpty } from "lodash";
import { useNavigate } from "react-router-dom";

injectReducer("budgets",reducer)

const Budget=()=>{
    const dispatch = useDispatch()
    const navigate = useNavigate()
    const budgets=useSelector((state)=>state.budgets.data.budgetsData)
    const loading=useSelector((state)=>state.budgets.data.loading)

    const fetchData=async()=>{
        dispatch(getBudgets())
    }

    useEffect(()=>{
        fetchData()
      
    },[])

    return(
        <AdaptableCard className="h-full" bodyClass="h-full">
             <div className="lg:flex items-center justify-between mb-4">
                <h3 className="mb-4 lg:mb-0">Tahapan Anggaran</h3>
            </div>
            <Loading loading={loading}>
            {!isEmpty(budgets) &&
                
                <div>
                     <div className="mt-8 max-w-[350px] mx-auto">
                        <Button className="mb-2" variant="solid" onClick={()=>{navigate('/budget/setup/new')}}>
                            Pindah Tahapan
                        </Button>
               
                    </div>
                    <Timeline>
                        {budgets.map((b,i)=>{
                            return  <Timeline.Item>Anggaran - {b.budget_year} Murni - {b.budget_status==0?"Draft":"Final"}</Timeline.Item>
                        })}
                    
                    </Timeline>
            </div>
            }
            {
                isEmpty(budgets) && 
                <div className="text-center">
                     <h5 className="mb-2">
                        Tahapan Penganggaran Belum Ada, Silahkan buat Tahapan Baru
                    </h5>
                    <div className="mt-8 max-w-[350px] mx-auto">
                        <Button className="mb-2" variant="solid" onClick={()=>{navigate('/budget/setup/new')}}>
                            Buat Tahapan
                        </Button>
               
                    </div>
                </div>
            }
            </Loading>
        </AdaptableCard>
    )
}

export default Budget