import React, { useEffect } from "react";
import { AdaptableCard, Container, Loading } from "components/shared";
import { useDispatch, useSelector } from "react-redux";
import { useNavigate } from "react-router-dom";
import { Notification, toast } from "src/components/ui";
import NewForm from "./NewForm";
import { apiCreateBudget } from "src/services/BudgetService";



const options=[
    {value:2024,label:'2024'},
    {value:2025,label:'2025'}
]

const BudgetNew=()=>{
    const navigate = useNavigate()

    const handleDiscard=()=>{
        navigate('/budget/setup')
    }

    const updateBudget=async(data)=>{
        const response = await apiCreateBudget(data)
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
            navigate('/budget/setup')
        }
    }


    return (
        <Container>
        <AdaptableCard>
        <NewForm 
       options={options} 
       type="new"
       onFormSubmit={handleFormSubmit}
       onDiscard={handleDiscard}/>
        </AdaptableCard>
    </Container>
    )
}

export default BudgetNew