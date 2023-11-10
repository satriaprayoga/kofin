import { Field } from "formik";
import React from "react";
import { AdaptableCard } from "components/shared";
import { FormItem, Input} from "components/ui";


const KegiatanField=(props)=>{
    const {touched,errors} = props

    return(
        <AdaptableCard className="mb-4" divider>
              <h5>Input/Ubah Kegiatan</h5>
              <p className="mb-6">Form Kegiatan</p>
              <FormItem
                label="Kode Kegiatan"
                invalid={errors.kegiatan_kode && touched.kegiatan_kode}
                errorMessage={errors.kegiatan_kode}
              >
                    <Field
                        type="text"
                        autoComplete="off"
                        name="kegiatan_kode"
                        placeholder="Kode"
                        component={Input}/>

                </FormItem>
                <FormItem
                    label="Nama Kegiatan"
                    invalid={errors.kegiatan_name && touched.kegiatan_name}
                    errorMessage={errors.kegiatan_name}
                >
                    <Field
                        type="text"
                        autoComplete="off"
                        name="kegiatan_name"
                        placeholder="Nama"
                        component={Input}/>

                </FormItem>
        </AdaptableCard>
    )
}

export default KegiatanField