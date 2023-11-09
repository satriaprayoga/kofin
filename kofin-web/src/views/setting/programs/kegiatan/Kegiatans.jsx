import React from "react";
import { AdaptableCard, Container } from "components/shared";
import reducer from "./store";
import { injectReducer } from "store";
import KegiatanTableTools from "./components/KegiatanTableTools";
import KegiatanTable from "./components/KegiatanTable";

injectReducer('kegiatans', reducer)


const Kegiatans=()=>{
    return(
        <AdaptableCard className="h-full" bodyClass="h-full">
            <div className="lg:flex items-center justify-between mb-4">
                <h3 className="mb-4 lg:mb-0">Kegiatan</h3>
               <KegiatanTableTools/>
            </div>
            <KegiatanTable/>
        </AdaptableCard>

       
    )
}

export default Kegiatans