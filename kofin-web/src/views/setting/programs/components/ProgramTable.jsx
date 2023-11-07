import React, { useEffect, useMemo, useRef } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { getPrograms, setTableData } from '../store/dataSlice'
import { DataTable } from 'src/components/shared'
import { cloneDeep } from 'lodash'

const ProgramTable=()=>{
    const tableRef = useRef(null)
    const dispatch = useDispatch()

    const {pageIndex, pageSize, sort, query, total}=useSelector(
        (state)=>state.programs.data.tableData
    )

    const loading = useSelector((state)=>state.programs.data.loading)
    const data = useSelector((state)=>state.programs.data.programsData)

    const fetchData=()=>{
        dispatch(getPrograms({pageIndex,pageSize,sort,query}))
    }

    useEffect(()=>{
        fetchData()
    },[pageIndex,pageSize,sort])

    const tableData = useMemo(
        () => ({ pageIndex, pageSize, sort, query, total }),
        [pageIndex, pageSize, sort, query, total]
    )

    const columns = useMemo(
        () => [
            {
                header: 'Kode',
                accessorKey: 'program_kode',
                cell: (props) => {
                    const row = props.row.original
                    return <span className="capitalize">{row.program_kode}</span>
                },
            },
            {
                header: 'Nama',
                accessorKey: 'program_name',
                cell: (props) => {
                    const row = props.row.original
                    return <span className="capitalize">{row.program_name}</span>
                },
            },
            {
                header: 'Unit',
                accessorKey: 'unit_name',
                cell: (props) => {
                    const row = props.row.original
                    return <span className="capitalize">{row.unit_name}</span>
                },
            },
            
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

export default ProgramTable