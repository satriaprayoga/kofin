import React, { useEffect } from "react";
import { injectReducer } from "store";
import reducer from "./store";
import { useDispatch, useSelector } from "react-redux";
import { useLocation, useNavigate } from "react-router-dom";
import { deleteSubunit, getSubunit, updateSubunit } from "./store/dataSlice";
import { AdaptableCard, Container, Loading } from "components/shared";
import { isEmpty } from "lodash";
import SubunitForm from "../SubunitForm";
import { Notification, toast } from "components/ui";

injectReducer('subunitEdit',reducer)

const SubunitEdit=()=>{
    const dispatch = useDispatch()

    const location = useLocation()
    const navigate = useNavigate()

    const subunitData = useSelector(
        (state)=>state.subunitEdit.data.subunitData
    )

    const loading = useSelector(
        (state)=>state.subunitEdit.data.loading
    )

    const fetchData = (data) =>{
        dispatch(getSubunit(data))
       // console.log(subunitData.data)
    }

    const handleFormSubmit = async (values,setSubmitting)=>{
        const path = location.pathname.substring(
            location.pathname.lastIndexOf('/')+1
        )
        const requestParam = path
        setSubmitting(true)
        const succes = await updateSubunit(requestParam,values)
        setSubmitting(false)
        if (succes){
            popNotification('Perubahan')
        }
    }

    const handleDiscard = () => {
        navigate('/setting/unit')
    }

    const handleDelete=async (setDialogOpen)=>{
        const path = location.pathname.substring(
            location.pathname.lastIndexOf('/')+1
        )
        const requestParam = path
        setDialogOpen(false)
        const success = await deleteSubunit(requestParam)
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
                {keyword} Subunit Berhasil
            </Notification>,
            {
                placement: 'top-center',
            }
        )
        navigate('/setting/unit')
    }

    useEffect(()=>{
        const path = location.pathname.substring(
            location.pathname.lastIndexOf('/')+1
        )
        const requestParam = path
        fetchData(requestParam)
    },[location.pathname])

    return(
        <>
        <Loading loading={loading}>
            {!isEmpty(subunitData.data) && (
                <>
                    <Container>
                        <AdaptableCard>
                        <SubunitForm
                        type="edit"
                        initialData={subunitData.data}
                        onFormSubmit={handleFormSubmit}
                        onDiscard={handleDiscard}
                        onDelete={handleDelete}/>
                        </AdaptableCard>
                    </Container>
                </>
            )}
        </Loading>
        </>
    )
}

export default SubunitEdit