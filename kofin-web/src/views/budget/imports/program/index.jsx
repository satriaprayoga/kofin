import React, { useRef, useEffect, useMemo, useState} from "react";
import { injectReducer } from "src/store";
import reducer from "./store";
import { useDispatch, useSelector } from "react-redux";
import { AdaptableCard, Loading } from "src/components/shared";
import { Select,Table, Checkbox, Button, toast, Notification } from "src/components/ui";
import {
    flexRender,
    getCoreRowModel,
    getFilteredRowModel,
    getPaginationRowModel,
    useReactTable,
} from '@tanstack/react-table'
import { isEmpty } from "lodash";

import { getBudgets, getPrograms, importProgram, setBudgetId } from "./store/dataSlice";
import { useNavigate } from "react-router-dom";


injectReducer("importBudgetProgram",reducer)

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
    const navigate = useNavigate()

    const [rowSelection, setRowSelection] = useState({})

    const programs = useSelector((state)=>state.importBudgetProgram.data.programsData)
    const budgets = useSelector((state)=>state.importBudgetProgram.data.budgetsData)
    const budgetId = useSelector((state)=>state.importBudgetProgram.data.budgetId)
    const loading = useSelector((state)=>state.importBudgetProgram.data.loading)
    
    const fetchData=async(data)=>{
      dispatch(setBudgetId(data))
      dispatch(getPrograms(data))
     
     // console.log(programs)
     
    } 

    useEffect(()=>{
        dispatch(getBudgets())
       // console.log(budgetId)
        dispatch(getPrograms(budgetId))
        //console.log(budgets)
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

    const handleClick=async(e)=>{
        e.preventDefault()
        const rows = table.getSelectedRowModel().rows
        const success = await importProgram(rows)
        if (success){
             popNotification('Import')
        }
    }

    const popNotification = (keyword) => {
        toast.push(
            <Notification
                title={`${keyword} Berhasil!`}
                type="success"
                duration={2500}
            >
                {keyword} Program Berhasil
            </Notification>,
            {
                placement: 'top-center',
            }
        )
        dispatch(getBudgets)
        dispatch(getPrograms(budgetId))
        navigate('/budget/import/program')
    }



    return (
        <>
        <AdaptableCard className="h-full" bodyClass="h-full">
        <div className="lg:flex items-center justify-between mb-4">
            <h3 className="mb-4 lg:mb-0">Import Program</h3>
            <Select
                         
                         options={budgets} 
                         onChange={(option)=>{fetchData(option.value);table.resetRowSelection(true)}} 
                         placeholder="Subunit" 
                         value={budgets.filter(
                                     (option) =>
                                         option.value ===
                                         budgetId
                                 )}/>  
            <div className="flex lg:flex-col lg:flex-row lg:items-center gap-4">
            <div className="md:mb-0 mb-4">
            
            </div>
                <Button block variant="solid" disabled={setButton()} onClick={handleClick}>
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