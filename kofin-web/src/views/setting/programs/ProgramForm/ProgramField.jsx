import { Field } from "formik";
import React from "react";
import { AdaptableCard } from "components/shared";
import { FormItem, Input, Select } from "components/ui";


const ProgramField=(props)=>{
    const {touched,errors,options,values} = props

    return(
        <AdaptableCard className="mb-4" divider>
              <h5>Input/Ubah Program</h5>
              <p className="mb-6">Form Program</p>
              <FormItem
                label="Kode Program"
                invalid={errors.program_kode && touched.program_kode}
                errorMessage={errors.program_kode}
              >
                    <Field
                        type="text"
                        autoComplete="off"
                        name="program_kode"
                        placeholder="Kode"
                        component={Input}/>

                </FormItem>
                <FormItem
                    label="Nama Program"
                    invalid={errors.program_name && touched.program_name}
                    errorMessage={errors.program_name}
                >
                    <Field
                        type="text"
                        autoComplete="off"
                        name="program_name"
                        placeholder="Nama"
                        component={Input}/>

                </FormItem>
                <FormItem
                    label="Sub Unit"
                    invalid={errors.unit_id && touched.unit_id}
                    errorMessage={errors.unit_id}
                >
                    <Field name="unit_id">
                        {({ field, form }) => (
                            <Select
                                field={field}
                                form={form}
                                options={options}
                                value={options.filter(
                                    (option) =>
                                        option.value ===
                                        values.unit_id
                                )}
                                onChange={(option) =>
                                    form.setFieldValue(
                                        field.name,
                                        option.value
                                    )
                                }
                            />
                        )}
                    </Field>
                    </FormItem>
        </AdaptableCard>
    )
}

export default ProgramField