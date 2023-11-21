import React, { useEffect, useState } from "react";
import { Notification, toast } from "components/ui";
import ProgramForm from "../ProgramForm";
import { useDispatch, useSelector } from "react-redux";
import { getSubunits } from "./store/dataSlice";
import { injectReducer } from "store";
import reducer from "./store";
import { useNavigate } from "react-router-dom";
import { apiCreateProgramData } from "src/services/ProgramService";
import { AdaptableCard, Container, Loading } from "components/shared";
import { isEmpty } from "lodash";


injectReducer('programsOption',reducer)
const ProgramNew=()=>{
    const navigate = useNavigate()

    const dispatch = useDispatch()
    const data = useSelector((state)=>state.programsOption.data.subunitsData)
    const loading = useSelector((state)=>state.programsOption.data.loading)
    
    const fetchData= async ()=>{
     dispatch(getSubunits())
    }

    const handleDiscard=()=>{
        navigate('/setting/programs')
    }

    const addProgram=async(data)=>{
        const response = await apiCreateProgramData(data)
        return response.data
    }

    const handleFormSubmit = async (values,setSubmitting)=>{
        setSubmitting(true)
        const success = await addProgram(values)
        console.log(values)
        setSubmitting(false)
        if (success){
            toast.push(
                <Notification
                    title={'Berhasil ditambah!'}
                    type="success"
                    duration={2500}>
                       Program Berhasil Ditambahkan! 
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
 
   return(
        <Loading loading={loading}>
            {!isEmpty(data) && (
                <>
                <Container>
                    <AdaptableCard>
                    <ProgramForm type="new" 
                options={data}
                onDiscard={handleDiscard}
                onFormSubmit={handleFormSubmit}/>
                    </AdaptableCard>
                </Container>
                </>
            )
               
            }
        </Loading>
   )
}

export default ProgramNew