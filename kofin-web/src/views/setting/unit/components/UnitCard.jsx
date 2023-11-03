import React, { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useNavigate } from "react-router-dom";
import reducer from "../store";
import { HiOutlinePencil } from "react-icons/hi";
import { injectReducer } from "store";
import { isEmpty } from 'lodash'
import { Button, Card, Table } from 'components/ui'
import { Loading } from 'components/shared'
import { getUnit } from "../store/dataSlice";

const { Tr, Th, Td, THead, TBody } = Table

injectReducer('unit',reducer)

const UnitCard=()=>{
    const dispatch = useDispatch()
    const navigate = useNavigate()

    const unitData = useSelector((state)=>state.unit.data.unitData)
    const loading = useSelector((state)=>state.unit.data.loading)

    const fetchData = ()=>{
        dispatch(getUnit())
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
    },[])

    return (
        <>
        <Loading loading={loading}>
            {!isEmpty(unitData) && (
                <>
                <div className="flex flex-col gap-4 h-full">
            <Card
                header={unitData[0].unit_name}
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
                    <Tr>
                    <Td>{unitData[0].unit_name}</Td>
                        <Td>{unitData[0].unit_loc}</Td>
                        <Td>{unitData[0].unit_head}</Td>
                    </Tr>
                </TBody>
            </Table>
            </Card>
            
        </div>
                </>
            )}
        </Loading>
        </>
    )



}

export default UnitCard