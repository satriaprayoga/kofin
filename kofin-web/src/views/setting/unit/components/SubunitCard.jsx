import React, { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useNavigate } from "react-router-dom";
import reducer from "../store";
import { HiOutlinePencil, HiOutlinePlus } from "react-icons/hi";
import { injectReducer } from "store";
import { isEmpty, isUndefined } from 'lodash'
import { Button, Card, Table } from 'components/ui'
import { Loading } from 'components/shared'
import { getSubunits} from "../store/dataSlice";
import SubunitTable from "./SubunitTable";

const { Tr, Th, Td, THead, TBody } = Table

injectReducer('unit',reducer)

const SubunitCard=()=>{
    const dispatch = useDispatch()
    const navigate = useNavigate()

    const subunitsData = useSelector((state)=>state.unit.data.subunitsData)
    const loading = useSelector((state)=>state.unit.data.loading)

    const fetchData = ()=>{
        dispatch(getSubunits())
    }

    const onCreate = () => {
        navigate(`/setting/unit/subunit/new`)
    }
    const headerExtraContent = (
        <span className="flex items-center">
            <Button className="mr-2" icon={<HiOutlinePlus />} variant="solid" onClick={onCreate}>
            </Button>
        </span>
    )

    useEffect(()=>{
        fetchData()
       // console.log(subunitsData)
    },[])

    return (
        <>
      <Card
                bodyClass="h-full"
                header="Daftar Subunit"
                headerExtra={headerExtraContent}
            >
              <SubunitTable/>
                
            </Card>
        </>
    )



}

export default SubunitCard

