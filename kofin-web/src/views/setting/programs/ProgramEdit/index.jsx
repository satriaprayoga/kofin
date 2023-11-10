import React, { useEffect, useState } from "react";
import { Notification, toast } from "components/ui";
import ProgramForm from "../ProgramForm";
import { useDispatch, useSelector } from "react-redux";
import { deleteProgram, getProgram, getSubunits } from "./store/dataSlice";
import { injectReducer } from "store";
import reducer from "./store";
import { useLocation, useNavigate } from "react-router-dom";
import { apiPutProgram } from "src/services/ProgramService";
import { AdaptableCard, Container, Loading } from "components/shared";
import { isEmpty } from "lodash";


injectReducer('programEdit',reducer)
const ProgramEdit=()=>{
    const navigate = useNavigate()
    const location = useLocation()
    const dispatch = useDispatch()

    const programData=useSelector((state)=>state.programEdit.data.program)
    const data = useSelector((state)=>state.programEdit.data.subunitData)
    const loading = useSelector((state)=>state.programEdit.data.loading)
    
    const fetchData= async (data)=>{
     dispatch(getSubunits())
     dispatch(getProgram(data))
    }

    useEffect(()=>{
        const path = location.pathname.substring(
            location.pathname.lastIndexOf('/')+1
        )
        const requestParam = {id:path}
        fetchData(requestParam)
    },[location.pathname])

    

    const handleDiscard=()=>{
        navigate('/setting/programs')
    }

    const updateProgram=async(data)=>{
        const response = await apiPutProgram(data)
        return response.data
    }

    const handleFormSubmit = async (values,setSubmitting)=>{
        setSubmitting(true)
        const succes = await updateProgram(values)
        setSubmitting(false)
        if (succes){
            popNotification('Perubahan')
        }
    }

    const handleDelete=async (setDialogOpen)=>{
        setDialogOpen(false)
        const success = await deleteProgram({id:programData.id})
        if (success){
            popNotification('Penghapusan')
        }
    }


    const popNotification = (keyword) => {
        toast.push(
            <Notification
                title={`${keyword} Berhasil!`}
                type="success"
                duration={2500}
            >
                {keyword} Program Berhasil
            </Notification>,
            {
                placement: 'top-center',
            }
        )
        navigate('/setting/programs')
    }


 
    useEffect(()=>{
        const path = location.pathname.substring(
            location.pathname.lastIndexOf('/')+1
        )
        const requestParam = {id:path}
        fetchData(requestParam)
     
    },[location.pathname])
 
   return(
        <Loading loading={loading}>
            {!isEmpty(programData) && (
                <>
                <Container>
                    <AdaptableCard>
                    <ProgramForm type="edit" 
                options={data}
                initialData={programData}
                onFormSubmit={handleFormSubmit}
                onDiscard={handleDiscard}
                onDelete={handleDelete}/>
                    </AdaptableCard>
                </Container>
                </>
            )
               
            }
        </Loading>
   )
}

export default ProgramEdit