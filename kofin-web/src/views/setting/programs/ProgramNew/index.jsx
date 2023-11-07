import React, { useEffect, useState } from "react";
import ProgramForm from "../ProgramForm";
import { useDispatch, useSelector } from "react-redux";
import { getSubunits } from "../../unit/store/dataSlice";
import { injectReducer } from "store";
import reducer from "../../unit/store";


injectReducer('unit',reducer)
const ProgramNew=()=>{

    const [options,setOptions]=useState([])
    const dispatch = useDispatch()
    const data = useSelector((state)=>state.unit.data.subunitsData)
    
    const fetchData=()=>{
     dispatch(getSubunits())
    }
 
    useEffect(()=>{
     fetchData()
     const optData=[]
     data[0].forEach((d)=>{
        optData.push({
            value:d.unit_id,
            label:d.unit_name
        })
     })
     setOptions(optData)
    },[])
 
   return(
        <ProgramForm type="new" options={options}/>
   )
}

export default ProgramNew