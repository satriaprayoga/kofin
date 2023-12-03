import { Field } from "formik";
import React from "react";
import { FormItem, Input, Select } from "components/ui";




const MurniField=(props)=>{
    const {touched,errors,options,values} = props

    return(
       
              <>
              <h5>Setup Anggaran</h5>
              <p className="mb-6">Form Tahapan</p>
                <FormItem
                    label="Tahun"
                    invalid={errors.year && touched.year}
                    errorMessage={errors.year}
                >
                    <Field name="year">
                        {({ field, form }) => (
                            <Select
                                placeholder="Pilih Tahun"
                                field={field}
                                form={form}
                                options={options}
                                value={options.filter(
                                    (option) =>
                                        option.value ===
                                        values.year
                                )}
                                onChange={(option) =>
                                    {
                                        form.setFieldValue(
                                            field.name,
                                            option.value
                                        )
                                        
                                    }
                                }
                            />
                        )}
                    </Field>
                </FormItem>
              </>
       
    )
}

export default MurniField