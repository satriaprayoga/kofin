import React, { useEffect, useMemo, useRef } from "react";
import { HiOutlinePencil } from "react-icons/hi";
import { useDispatch, useSelector } from "react-redux";
import { useNavigate } from "react-router-dom";
import useThemeClass from "utils/hooks/useThemeClass";
import { getSubunits } from "../store/dataSlice";
import { DataTable } from "src/components/shared";
import reducer from "../store";
import { injectReducer } from "src/store";



const ActionColumn = ({row})=>{
    const {textTheme} = useThemeClass()
    const navigate = useNavigate()

    const onEdit = () =>{
        navigate(`/setting/unit/subunit/edit/${row.unit_id}`)
    }

    return (
        <div className="flex justify-end text-lg">
            <span
                className={`cursor-pointer p-2 hover:${textTheme}`}
                onClick={onEdit}
            >
                <HiOutlinePencil />
            </span>
        </div>
    )
}

const SubunitTable = () =>{
    const tableRef = useRef(null)
    const dispatch = useDispatch()

    const data = useSelector((state)=>state.unit.data.subunitsData)
    const loading = useSelector((state)=>state.unit.data.loading)

    const fetchData = ()=>{
        dispatch(getSubunits())
    }

    useEffect(()=>{
        fetchData()
        console.log(data.data)
    },[])

    const columns = useMemo(
        ()=>[
            {
                header:'Kode Unit',
                cell:(props)=>{
                    const row = props.row.original
                    return <span className="capitalize">{row.unit_kode}</span>
                }
            },
            {
                header:'Nama Unit',
                cell:(props)=>{
                    const row = props.row.original
                    return <span className="capitalize">{row.unit_name}</span>
                }
            },
            {
                header:'ABBR',
                cell:(props)=>{
                    const row = props.row.original
                    return <span className="capitalize">{row.unit_abbr}</span>
                }
            },
            {
                header:'Lokasi',
                cell:(props)=>{
                    const row = props.row.original
                    return <span className="capitalize">{row.unit_loc}</span>
                }
            },
            {
                header:'Kepala',
                cell:(props)=>{
                    const row = props.row.original
                    return <span className="capitalize">{row.unit_head}</span>
                }
            },
            {
                header:'',
                id:'action',
                cell:(props)=><ActionColumn row={props.row.original}/>
            }
        ]
    )

    return (
        <>
        <DataTable
            ref={tableRef}
            columns={columns}
            data={data.data}
            loading={loading}
        />
        </>
    )
}

export default SubunitTable