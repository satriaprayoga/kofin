import React from "react";
import KegiatanTableSearch from "./KegiatanTableSearch";
import { Link } from "react-router-dom";
import { Button } from "components/ui";
import { HiPlusCircle } from "react-icons/hi";

const KegiatanTableTools=()=>{
    return(
        <div className="flex lg:flex-col lg:flex-row lg:items-center gap-4">
            <KegiatanTableSearch/>
            <Link
                className="md:mb-0 mb-4"
                to="/setting/programs/new"
            >
                <Button block variant="solid" size="sm" icon={<HiPlusCircle />}>
                    Tambah Kegiatan
                </Button>
            </Link>
        </div>
    )
}

export default KegiatanTableTools