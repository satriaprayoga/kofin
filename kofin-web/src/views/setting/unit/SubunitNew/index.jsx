import React, { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { Notification, toast } from "components/ui";
import { apiCreateSubunitData } from "services/UnitService";
import SubunitForm from "../SubunitForm";
import { AdaptableCard, Container } from "src/components/shared";
import { injectReducer } from "store";
import reducer from "./store";
import { useDispatch, useSelector } from "react-redux";
import { getUnit } from "./store/dataSlice";

injectReducer("subunitNew",reducer)

const SubunitNew=()=>{
    const navigate = useNavigate()

    const dispatch = useDispatch()
    const unitData = useSelector((state)=>state.subunitNew.data.unitData)

    const addSubunit=async(data)=>{
        const response = await apiCreateSubunitData(data)
        return response.data
    }

    const fetchData=async()=>{
        dispatch(getUnit())
    }

    useEffect(()=>{
        fetchData()
    },[])

    const handleFormSubmit = async (values,setSubmitting)=>{
        setSubmitting(true)
        const unit=unitData.data
        values.unit_head_id="000003"
        values.unit_induk_id=unit.unit_id
        values.unit_induk_head=unit.unit_head
        values.unit_induk_head_id=unit.unit_induk_head_id
        values.root=false
        console.log(values)
        const success = await addSubunit(values)
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