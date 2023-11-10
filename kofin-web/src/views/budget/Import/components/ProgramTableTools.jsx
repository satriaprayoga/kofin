import React from "react";
import ProgramTableSearch from "./ProgramTableSearch";
import { Link } from "react-router-dom";
import { Button } from "components/ui";
import { HiPlusCircle } from "react-icons/hi";

const ProgramTableTools=()=>{
    return(
        <div className="flex lg:flex-col lg:flex-row lg:items-center gap-4">
            <ProgramTableSearch/>
        </div>
    )
}

export default ProgramTableTools