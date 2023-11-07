import React from "react";
import ProgramTableSearch from "./ProgramTableSearch";
import { Link } from "react-router-dom";
import { Button } from "components/ui";
import { HiPlusCircle } from "react-icons/hi";

const ProgramTableTools=()=>{
    return(
        <div className="flex lg:flex-col lg:flex-row lg:items-center gap-4">
            <ProgramTableSearch/>
            <Link
                className="md:mb-0 mb-4"
                to="/setting/programs/new"
            >
                <Button block variant="solid" size="sm" icon={<HiPlusCircle />}>
                    Tambah Program
                </Button>
            </Link>
        </div>
    )
}

export default ProgramTableTools