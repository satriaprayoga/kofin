import React, { useEffect, useMemo, useRef } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { getProgramKegiatans, setTableData } from '../store/dataSlice'
import { DataTable } from 'src/components/shared'
import { cloneDeep } from 'lodash'
import { useLocation, useNavigate } from 'react-router-dom'
import { HiOutlinePencil } from 'react-icons/hi'
import useThemeClass from 'src/utils/hooks/useThemeClass'

const ActionColumn = ({row})=>{
    const {textTheme} = useThemeClass()
    const navigate = useNavigate()

    const onEdit = () =>{
        navigate(`/setting/programs/kegiatan/edit/${row.id}`)
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


const KegiatanTable=()=>{
    const tableRef = useRef(null)
    const dispatch = useDispatch()
    const location = useLocation()

    const {pageIndex, pageSize, sort, query, total}=useSelector(
        (state)=>state.kegiatans.data.tableData
    )

    const loading = useSelector((state)=>state.kegiatans.data.loading)
    const data = useSelector((state)=>state.kegiatans.data.kegiatansData)

    const fetchData=()=>{
       // const newTableData = cloneDeep(tableData)
        const path = location.pathname.substring(
            location.pathname.lastIndexOf('/')+1
        )
      //  dispatch(setTableData(newTableData))
        dispatch(getProgramKegiatans(path,{pageIndex,pageSize,sort,query}))
    }

    useEffect(()=>{
        fetchData()
    },[pageIndex,pageSize,sort,location.pathname])

    const tableData = useMemo(
        () => ({ pageIndex, pageSize, sort, query, total }),
        [pageIndex, pageSize, sort, query, total]
    )

    const columns = useMemo(
        () => [
            {
                header: 'Kode',
                accessorKey: 'kegiatan_kode',
                cell: (props) => {
                    const row = props.row.original
                    return <span className="capitalize">{row.kegiatan_kode}</span>
                },
            },
            {
                header: 'Nama',
                accessorKey: 'kegiatan_name',
                cell: (props) => {
                    const row = props.row.original
                    return <span className="capitalize">{row.kegiatan_name}</span>
                },
            },
            {
                header: 'Kode Program',
                accessorKey: 'program_kode',
                cell: (props) => {
                    const row = props.row.original
                    return <span className="capitalize">{row.program_kode}</span>
                },
            },
            {
                header: 'Program',
                accessorKey: 'program_name',
                cell: (props) => {
                    const row = props.row.original
                    return <span className="capitalize">{row.program_name}</span>
                },
            },
            {
                header:'',
                id:'action',
                cell:(props)=><ActionColumn row={props.row.original}/>
            }
            
        ],
        []
    )

    const onPaginationChange = (page) => {
        const newTableData = cloneDeep(tableData)
        newTableData.pageIndex = page
        dispatch(setTableData(newTableData))
    }

    const onSelectChange = (value) => {
        const newTableData = cloneDeep(tableData)
        newTableData.pageSize = Number(value)
        newTableData.pageIndex = 1
        dispatch(setTableData(newTableData))
    }

    const onSort = (sort, sortingColumn) => {
        const newTableData = cloneDeep(tableData)
        newTableData.sort = sort
        dispatch(setTableData(newTableData))
    }

    return (
        <>
        <DataTable
                ref={tableRef}
                columns={columns}
                data={data}
                skeletonAvatarColumns={[0]}
                skeletonAvatarProps={{ className: 'rounded-md' }}
                loading={loading}
                pagingData={tableData}
                onPaginationChange={onPaginationChange}
                onSelectChange={onSelectChange}
                onSort={onSort}
            />
        </>
    )

}

export default KegiatanTable