import React from "react";
import ProgramForm from "../ProgramForm";

const options = [
    { value: 1, label: 'Foo' },
    { value: 2, label: 'Bar' },
]

const ProgramNew=()=>{

   return(
        <ProgramForm type="new" options={options}/>
   )
}

export default ProgramNew