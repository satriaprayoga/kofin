import React, { useEffect, useRef } from "react";
import { useDispatch, useSelector } from "react-redux";
import { getPrograms, getSubunits, setTableData } from "../store/dataSlice";
import { Input, Select } from "components/ui";
import { HiOutlineSearch } from "react-icons/hi";
import cloneDeep from "lodash/cloneDeep";
import debounce from "lodash/debounce";

const options=[
    {
        value:"test1",
        label:"test"
    },
    {
        value:"test2",
        label:"test"
    },
]

const ProgramTableSearch = () =>{
    const dispatch = useDispatch()
    const searchInput = useRef()

    const tableData = useSelector(
        (state)=>state.programs.data.tableData
    )

    const debounceFn = debounce(handleDebounceFn,500)

    function handleDebounceFn(val){
        const newTableData=cloneDeep(tableData)
        newTableData.query=val
        newTableData.pageIndex=1
        if (typeof val === 'string' && val.length > 1){
            fetchData(newTableData)
        }
        if (typeof val === 'string' && val.length === 0) {
            fetchData(newTableData)
        }

    }

    const fetchData=(data)=>{
        dispatch(setTableData(data))
        dispatch(getPrograms(data))
    }

    const onEdit=(e)=>{
        debounceFn(e.target.value)
    }

    return (
        <Input
        ref={searchInput}
        className="max-w-md md:w-52 md:mb-0 mb-4"
        size="sm"
        placeholder="Cari Program"
        prefix={<HiOutlineSearch className="text-lg" />}
        onChange={onEdit}
    />
    )
}

export default ProgramTableSearch