import React from 'react'
import UnitCard from './components/UnitCard'
import SubunitCard from './components/SubunitCard'



const Unit=()=>{
    

    return (
        <>
         <div className="grid grid-flow-row auto-rows-max gap-4">
         <UnitCard/>
        <SubunitCard/>
         </div>
       
        </>
    )
}

export default Unit