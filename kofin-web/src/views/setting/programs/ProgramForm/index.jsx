import { Form, Formik } from "formik";
import { cloneDeep } from "lodash";
import React, { forwardRef } from "react";
import { FormContainer } from "src/components/ui";
import * as Yup from 'yup'
import ProgramField from "./ProgramField";

const validationSchema = Yup.object().shape({
    program_kode: Yup.string().required('Kode harus diisi'),
    program_name: Yup.string().required('Nama harus diisi'),
    unit_id: Yup.number().required('sub unit harus diisi'),
})

const ProgramForm=forwardRef((props,ref)=>{
    const { type, initialData, onFormSubmit, onDiscard, onDelete, options } = props

    return(
        <>
        <Formik
            innerRef={ref}
            initialValues={{
                ...initialData
            }}
            validationSchema={validationSchema}
            onSubmit={(values,{setSubmitting})=>{
                const formData = cloneDeep(values)
                onFormSubmit?.(formData,setSubmitting)
            }}
        >
            {({ values, touched, errors, isSubmitting }) => (
                <Form>
                    <FormContainer>
                    <div className="grid grid-cols-1 lg:grid-cols-3 gap-4">
                        <div className="lg:col-span-2">
                            <ProgramField
                                touched={touched}
                                errors={errors}
                                values={values}
                                options={options}
                                />
                        </div>
                    </div>
                    </FormContainer>
                </Form>
            )}
        </Formik>
        </>
    )
})

export default ProgramForm