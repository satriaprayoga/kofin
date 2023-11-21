import React, { useEffect } from "react";
import { AdaptableCard, Container, Loading } from "components/shared";
import reducer from "./store";
import { injectReducer } from "store";
import KegiatanTableTools from "./components/KegiatanTableTools";
import KegiatanTable from "./components/KegiatanTable";
import { useDispatch, useSelector } from "react-redux";
import { useLocation } from "react-router-dom";
import { getProgram } from "./store/dataSlice";
import { isEmpty } from "lodash";

injectReducer('kegiatans', reducer)


const Kegiatans=()=>{
    const dispatch = useDispatch()
    const location = useLocation()

    const loading =useSelector((state)=>state.kegiatans.data.loading)
    const data = useSelector((state)=>state.kegiatans.data.programData)

    const fetchData=async(data)=>{
        dispatch(getProgram(data))
    }

    useEffect(()=>{
        const path = location.pathname.substring(
            location.pathname.lastIndexOf('/')+1
        )
        const requestParam = path
        fetchData(requestParam)
        console.log(data)
    },[location.pathname])


    return(
        <>
       
        {!isEmpty(data) && (
            <AdaptableCard className="h-full" bodyClass="h-full">
            <div className="lg:flex items-center justify-between mb-4">
                <h4 className="mb-4 lg:mb-0">{data.program_name}</h4>
               <KegiatanTableTools program={data}/>
            </div>
            <KegiatanTable/>
        </AdaptableCard>
        )}
    
        </>

       
    )
}

export default Kegiatans