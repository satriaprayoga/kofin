import { Field } from "formik";
import React from "react";
import { AdaptableCard } from "components/shared";
import { FormItem, Input, Select } from "components/ui";

const statusOptions=[
    {value:0,label:'Draft'},
    {value:1,label:'Final'}
]


const SetupField=(props)=>{
    const {touched,errors,options,values} = props

    return(
       
              <>
              <h5>Setup Budget</h5>
              <p className="mb-6">Form Budget</p>
                <FormItem
                    label="Tahun"
                    invalid={errors.budget_year && touched.budget_year}
                    errorMessage={errors.budget_year}
                >
                    <Field name="budget_year">
                        {({ field, form }) => (
                            <Select
                                placeholder="Pilih Tahun"
                                field={field}
                                form={form}
                                options={options}
                                value={options.filter(
                                    (option) =>
                                        option.value ===
                                        values.budget_year
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
                <FormItem
                    label="Status"
                    invalid={errors.budget_year && touched.budget_status}
                    errorMessage={errors.budget_status}
                >
                    <Field name="budget_status">
                        {({ field, form }) => (
                            <Select
                                placeholder="Pilih Status"
                                field={field}
                                form={form}
                                options={statusOptions}
                                value={statusOptions.filter(
                                    (option) =>
                                        option.value ===
                                        values.budget_status
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

export default SetupField