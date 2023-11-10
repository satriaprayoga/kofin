import React from "react";
import { useNavigate } from "react-router-dom";
import { Notification, toast } from "components/ui";
import { apiCreateSubunitData } from "services/UnitService";
import SubunitForm from "../SubunitForm";
import { AdaptableCard, Container } from "src/components/shared";

const SubunitNew=()=>{
    const navigate = useNavigate()

    const addSubunit=async(data)=>{
        const response = await apiCreateSubunitData(data)
        return response.data
    }

    const handleFormSubmit = async (values,setSubmitting)=>{
        setSubmitting(true)
        const success = await addSubunit(values)
        console.log(values)
        setSubmitting(false)
        if (success){
            toast.push(
                <Notification
                    title={'Berhasil ditambah!'}
                    type="success"
                    duration={2500}>
                       Sub unit Berhasil Ditambahkan! 
                </Notification>,
                {
                    placement:'top-center'
                }
            )
            navigate('/setting/unit')
        }
    }

    const handleDiscard=()=>{
        navigate('/setting/unit')
    }

    return (
        <>
            <Container>
                <AdaptableCard>
                <SubunitForm
                type="new"
                onFormSubmit={handleFormSubmit}
                onDiscard={handleDiscard}/>
                </AdaptableCard>
            </Container>
        </>
    )
}

export default SubunitNew