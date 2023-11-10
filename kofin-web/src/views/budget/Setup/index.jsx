import React, { useEffect } from "react";
import { AdaptableCard, Container, Loading } from "components/shared";
import { injectReducer } from "store";
import reducer from "../store";
import { useDispatch, useSelector } from "react-redux";
import { getBudget } from "../store/dataSlice";
import { redirect, useNavigate } from "react-router-dom";
import { isEmpty } from "lodash";
import SetupForm from "./SetupForm";
import { apiPutBudget } from "src/services/BudgetService";
import { Notification, toast } from "src/components/ui";

injectReducer("budget",reducer)


const options=[
    {value:2024,label:'2024'},
    {value:2025,label:'2025'}
]

const Setup=()=>{
    const dispatch = useDispatch()
    const navigate = useNavigate()
    const budget = useSelector((state)=>state.budget.data.budgetData)
    const loading = useSelector((state)=>state.budget.data.loading)

    const fetchData=async()=>{
        dispatch(getBudget())
    }

    const handleDiscard=()=>{
        navigate('/home')
    }

    const updateBudget=async(data)=>{
        const response = await apiPutBudget(data)
        return response.data
    }

    const handleFormSubmit = async (values,setSubmitting)=>{
        setSubmitting(true)
        const success = await updateBudget(values)
        setSubmitting(false)
        if (success){
            toast.push(
                <Notification
                    title={'Setup Berhasil!'}
                    type="success"
                    duration={2500}>
                       Setup Anggaran Berhasil! 
                </Notification>,
                {
                    placement:'top-center'
                }
            )
            navigate('/setting/programs')
        }
    }

    useEffect(()=>{
        fetchData()
       
    },[])

    return (
       <Loading loading={loading}>
        <Container>
            <AdaptableCard>
            {!isEmpty(budget) && (
           
           <SetupForm 
           options={options} 
           initialData={budget} 
           type="edit"
           onFormSubmit={handleFormSubmit}
           onDiscard={handleDiscard}/>
            )}  
            </AdaptableCard>
        </Container>
       </Loading>
    )
}

export default Setup