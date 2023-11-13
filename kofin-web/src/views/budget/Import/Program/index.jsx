import React, { useRef, useEffect, useMemo, useState} from "react";
import { injectReducer } from "src/store";
import reducer from "./store";
import { useDispatch, useSelector } from "react-redux";
import { getPrograms, getSubunits, setSubunitId } from "./store/dataSlice";
import { AdaptableCard, Loading } from "src/components/shared";
import { Select,Table, Checkbox, Button } from "src/components/ui";
import {
    flexRender,
    getCoreRowModel,
    getFilteredRowModel,
    getPaginationRowModel,
    useReactTable,
} from '@tanstack/react-table'
import { isEmpty } from "lodash";

injectReducer("importProgram",reducer)

const { Tr, Th, Td, THead, TBody } = Table

function IndeterminateCheckbox({ indeterminate, onChange, ...rest }) {
    const ref = useRef(null)

    useEffect(() => {
        if (typeof indeterminate === 'boolean') {
            ref.current.indeterminate = !rest.checked && indeterminate
        }
    }, [ref, indeterminate])

    return <Checkbox ref={ref} onChange={(_, e) => onChange(e)} {...rest} />
}

const Program=()=>{
    const dispatch=useDispatch()
    const [rowSelection, setRowSelection] = useState({})

    const loading = useSelector((state)=>state.importProgram.data.loading)
    const programs= useSelector((state)=>state.importProgram.data.programsData)
    const subunits= useSelector((state)=>state.importProgram.data.subunitData)
    const subunitId = useSelector((state)=>state.importProgram.data.subunitId)

    const fetchData=async(data)=>{
      
        dispatch(setSubunitId(data))
        dispatch(getPrograms({id:data}))
    }

    useEffect(()=>{
        dispatch(getSubunits())
      
        console.log(subunitId)
        dispatch(getPrograms({id:subunitId}))
    },[])

    const columns = useMemo(() => {
        return [
            {
                id: 'select',
                header: ({ table }) => (
                    <IndeterminateCheckbox
                        {...{
                            checked: table.getIsAllRowsSelected(),
                            indeterminate: table.getIsSomeRowsSelected(),
                            onChange: table.getToggleAllRowsSelectedHandler(),
                        }}
                    />
                ),
                cell: ({ row }) => (
                    <div className="px-1">
                        <IndeterminateCheckbox
                            {...{
                                checked: row.getIsSelected(),
                                disabled: !row.getCanSelect(),
                                indeterminate: row.getIsSomeSelected(),
                                onChange: row.getToggleSelectedHandler(),
                            }}
                        />
                    </div>
                ),
            },
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
            }
        ]
    }, [])

    const table = useReactTable({
        data:programs,
        columns,
        state: {
            rowSelection,
        },
        enableRowSelection: true, //enable row selection for all rows
        // enableRowSelection: row => row.original.age > 18, // or enable row selection conditionally per row
        onRowSelectionChange: setRowSelection,
        getCoreRowModel: getCoreRowModel(),
        getFilteredRowModel: getFilteredRowModel(),
        getPaginationRowModel: getPaginationRowModel(),
    })

    const setButton=()=>{
        if (table.getIsSomeRowsSelected() || table.getIsAllRowsSelected()){
            return false
        }else{
            return true
        }
    }

    return (
        <>
        <AdaptableCard className="h-full" bodyClass="h-full">
        <div className="lg:flex items-center justify-between mb-4">
            <h3 className="mb-4 lg:mb-0">Import Program</h3>
            <div className="flex lg:flex-col lg:flex-row lg:items-center gap-4">
            <div
                        className="md:mb-0 mb-4"
                       
                    >
                    <Select
                         
                        options={subunits} 
                        onChange={(option)=>{fetchData(option.value);table.resetRowSelection(true)}} 
                        placeholder="Pilih Subunit" 
                        value={subunits.filter(
                                    (option) =>
                                        option.value ===
                                        subunitId
                                )}/>
                </div>
                <Button block variant="solid" disabled={setButton()}>
                    Import
                </Button>
            </div>
        </div>
        <Loading loading={loading}>
            {!isEmpty(programs) && (
                <Table>
                <THead>
                    {table.getHeaderGroups().map((headerGroup) => (
                        <Tr key={headerGroup.id}>
                            {headerGroup.headers.map((header) => {
                                return (
                                    <Th
                                        key={header.id}
                                        colSpan={header.colSpan}
                                    >
                                        {flexRender(
                                            header.column.columnDef.header,
                                            header.getContext()
                                        )}
                                    </Th>
                                )
                            })}
                        </Tr>
                    ))}
                </THead>
                <TBody>
                    {table.getRowModel().rows.map((row) => {
                        return (
                            <Tr key={row.id}>
                                {row.getVisibleCells().map((cell) => {
                                    return (
                                        <Td key={cell.id}>
                                            {flexRender(
                                                cell.column.columnDef.cell,
                                                cell.getContext()
                                            )}
                                        </Td>
                                    )
                                })}
                            </Tr>
                        )
                    })}
                </TBody>
            </Table>
            )}
        </Loading>

        </AdaptableCard>
        </>
    )
}

export default Program