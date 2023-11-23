import React, { useEffect, useState } from "react";
import { Notification, toast } from "components/ui";
import KegiatanForm from "../KegiatanForm";
import { useDispatch, useSelector } from "react-redux";
import { getProgram} from "./store/dataSlice";
import { injectReducer } from "store";
import reducer from "./store";
import { useLocation, useNavigate } from "react-router-dom";
import { apiCreateKegiatanData } from "src/services/KegiatanService";
import { AdaptableCard, Loading } from "components/shared";
import { isEmpty } from "lodash";


injectReducer('kegiatansNew',reducer)
const KegiatanNew=()=>{
    const navigate = useNavigate()
    const location = useLocation()

    const dispatch = useDispatch()
    const data = useSelector((state)=>state.kegiatansNew.data.programData)
    const loading = useSelector((state)=>state.kegiatansNew.data.loading)
    
    const fetchData= async (data)=>{
     dispatch(getProgram(data))
    }

    const handleDiscard=()=>{
        navigate('/setting/programs')
    }

    const addKegiatan=async(data)=>{
        const response = await apiCreateKegiatanData(data)
        return response.data
    }

    const handleFormSubmit = async (values,setSubmitting)=>{
        setSubmitting(true)
        values.program_kode=data.program_kode
        values.program_id=data.program_id
        values.program_name=data.program_name
        console.log(values)
        const success = await addKegiatan(values)
       
        setSubmitting(false)
        if (success){
            toast.push(
                <Notification
                    title={'Berhasil ditambah!'}
                    type="success"
                    duration={2500}>
                       Kegiatan Berhasil Ditambahkan! 
                </Notification>,
                {
                    placement:'top-center'
                }
            )
            navigate('/setting/programs')
        }
    }


 
    useEffect(()=>{
        const path = location.pathname.substring(
            location.pathname.lastIndexOf('/')+1
        )
        const requestParam = path
        fetchData(requestParam)
       
    },[location.pathname])
 
   return(
        <Loading loading={loading}>
            {!isEmpty(data) && (
                <AdaptableCard>
                    <h3>{data.program_name}</h3>
                    <>
                 <KegiatanForm type="new" 
                onDiscard={handleDiscard}
                onFormSubmit={handleFormSubmit}/>
                </>
                </AdaptableCard>
            )
               
            }
        </Loading>
   )
}

export default KegiatanNew