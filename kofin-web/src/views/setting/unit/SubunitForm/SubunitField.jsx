import React from "react";
import { AdaptableCard } from "components/shared";
import { FormItem, Input } from "components/ui";
import { Field } from "formik";

const SubunitField=(props)=>{
    const {touched,errors} = props

    return (
        <AdaptableCard className="mb-4" divider>
            <h5>Input/Ubah Sub Unit</h5>
            <p className="mb-6">Form Sub Unit</p>
            <FormItem
                label="Kode Sub Unit"
                invalid={errors.unit_kode && touched.unit_kode}
                errorMessage={errors.unit_kode}
            >
                <Field
                    type="text"
                    autoComplete="off"
                    name="unit_kode"
                    placeholder="Kode"
                    component={Input}/>

            </FormItem>
            <FormItem 
                label="Nama Sub Unit"
                invalid={errors.unit_name && touched.unit_name}
                errorMessage={errors.unit_name}
            >
                <Field
                    type="text"
                    autoComplete="off"
                    name="unit_name"
                    placeholder="Name"
                    component={Input}/>

            </FormItem>
            <FormItem 
                label="Singkatan Sub Unit"
                invalid={errors.unit_abbr && touched.unit_abbr}
                errorMessage={errors.unit_abbr}
            >
                <Field
                    type="text"
                    autoComplete="off"
                    name="unit_abbr"
                    placeholder="Singkatan"
                    component={Input}/>

            </FormItem>
            <FormItem 
                label="Alamat Sub Unit"
                invalid={errors.unit_loc && touched.unit_loc}
                errorMessage={errors.unit_loc}
            >
                <Field
                    type="text"
                    autoComplete="off"
                    name="unit_loc"
                    placeholder="Lokasi"
                    component={Input}/>

            </FormItem>
            <FormItem 
                label="Kepala Sub Unit"
                invalid={errors.unit_head && touched.unit_head}
                errorMessage={errors.unit_head}
            >
                <Field
                    type="text"
                    autoComplete="off"
                    name="unit_head"
                    placeholder="Kepala"
                    component={Input}/>

            </FormItem>
        </AdaptableCard>
    )

}

export default SubunitField