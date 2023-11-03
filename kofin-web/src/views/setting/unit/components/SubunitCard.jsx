import React, { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useNavigate } from "react-router-dom";
import reducer from "../store";
import { HiOutlinePencil } from "react-icons/hi";
import { injectReducer } from "store";
import { isEmpty, isUndefined } from 'lodash'
import { Button, Card, Table } from 'components/ui'
import { Loading } from 'components/shared'
import { getSubunits} from "../store/dataSlice";

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

    const onEdit = () => {
        navigate(`/setting/unit/edit`)
    }
    const headerExtraContent = (
        <span className="flex items-center">
            <Button className="mr-2" icon={<HiOutlinePencil />} variant="solid" onClick={onEdit}>
            </Button>
        </span>
    )

    useEffect(()=>{
        fetchData()
        console.log(subunitsData)
    },[])

    return (
        <>
      <Loading loading={loading}>
            {!isEmpty(subunitsData) && (
                <>
                <Card
                bodyClass="h-full"
                header="Daftar Subunit"
                headerExtra={headerExtraContent}
            >
               <Table>
                <THead>
                    <Tr>
                    <Th>Nama Unit</Th>
                        <Th>Lokasi</Th>
                        <Th>Kepala</Th>
                    </Tr>
                </THead>
                <TBody>
                    {subunitsData[0].map((s,index)=>{
                        return (
                            <Tr key={index}>
                    <Td>{s.unit_name}</Td>
                        <Td>{s.unit_loc}</Td>
                        <Td>{s.unit_head}</Td>
                    </Tr>
                        )
                    })}
                </TBody>
            </Table>
                
            </Card>
                </>
            )}
        </Loading>
        </>
    )



}

export default SubunitCard

