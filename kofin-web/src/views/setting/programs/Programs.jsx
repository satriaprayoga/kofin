import React from "react";
import { AdaptableCard, Container } from "components/shared";
import reducer from "./store";
import { injectReducer } from "store";
import ProgramTableTools from "./components/ProgramTableTools";
import ProgramTable from "./components/ProgramTable";

injectReducer('programs', reducer)


const Programs=()=>{
    return(
        <AdaptableCard className="h-full" bodyClass="h-full">
            <div className="lg:flex items-center justify-between mb-4">
                <h3 className="mb-4 lg:mb-0">Program</h3>
               <ProgramTableTools/>
            </div>
            <ProgramTable/>
        </AdaptableCard>

       
    )
}

export default Programs