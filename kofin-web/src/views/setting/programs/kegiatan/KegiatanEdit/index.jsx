import React, { useEffect, useState } from "react";
import { Notification, toast } from "components/ui";
import KegiatanForm from "../KegiatanForm";
import { useDispatch, useSelector } from "react-redux";
import { deleteKegiatan, getKegiatan} from "./store/dataSlice";
import { injectReducer } from "store";
import reducer from "./store";
import { useLocation, useNavigate } from "react-router-dom";
import { apiPutKegiatan } from "src/services/KegiatanService";
import { AdaptableCard, Loading } from "components/shared";
import { isEmpty } from "lodash";


injectReducer('kegiatanEdit',reducer)
const KegiatanEdit=()=>{
    const navigate = useNavigate()
    const location = useLocation()
    const dispatch = useDispatch()

    const kegiatanData=useSelector((state)=>state.kegiatanEdit.data.kegiatan)
    const loading = useSelector((state)=>state.kegiatanEdit.data.loading)
    
    const fetchData= async (data)=>{
     dispatch(getKegiatan(data))
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

    const updateKegiatan=async(data)=>{
        const response = await apiPutKegiatan(data)
        return response.data
    }

    const handleFormSubmit = async (values,setSubmitting)=>{
        setSubmitting(true)
        const succes = await updateKegiatan(values)
        setSubmitting(false)
        if (succes){
            popNotification('Perubahan')
        }
    }

    const handleDelete=async (setDialogOpen)=>{
        setDialogOpen(false)
        const success = await deleteKegiatan({id:kegiatanData.id})
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
                {keyword} Kegiatan Berhasil
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
            {!isEmpty(kegiatanData) && (
                <>
                <AdaptableCard>
                    <h3>{kegiatanData.program_name}</h3>
                </AdaptableCard>
                 <KegiatanForm type="edit" 
                initialData={kegiatanData}
                onFormSubmit={handleFormSubmit}
                onDiscard={handleDiscard}
                onDelete={handleDelete}/>
                </>
            )
               
            }
        </Loading>
   )
}

export default KegiatanEdit