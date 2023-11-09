import { Form, Formik } from "formik";
import cloneDeep from "lodash/cloneDeep";
import React, { forwardRef, useState } from "react";
import { HiOutlineTrash } from "react-icons/hi";
import { ConfirmDialog, StickyFooter } from "components/shared";
import { Button, FormContainer } from "components/ui";
import * as Yup from 'yup'
import ProgramField from "./ProgramField";
import { AiOutlineSave } from "react-icons/ai";


const validationSchema = Yup.object().shape({
    program_kode: Yup.string().required('Kode harus diisi'),
    program_name: Yup.string().required('Nama harus diisi'),
    unit_id: Yup.number().required('sub unit harus diisi'),
})

const DeleteSubunitButton = ({onDelete})=>{
    const [dialogOpen,setDialogOpen]=useState(false)

    const onConfirmDialogOpen=()=>{
        setDialogOpen(true)
    }

    const onConfirmDialogClose=()=>{
        setDialogOpen(false)
    }

    const handleConfirm=()=>{
        onDelete?.(setDialogOpen)
    }

    return (
        <>
            <Button
                className="text-red-600"
                variant="plain"
                size="sm"
                icon={<HiOutlineTrash/>}
                type="button"
                onClick={onConfirmDialogOpen}
            >
                Delete
            </Button>
            <ConfirmDialog
                isOpen={dialogOpen}
                onClose={onConfirmDialogClose}
                onRequestClose={onConfirmDialogClose}
                type="danger"
                title="Hapus Sub Unit"
                onCancel={onConfirmDialogClose}
                onConfirm={handleConfirm}
                confirmButtonColor="red-600"
            >
                <p>
                    Anda yakin ingin menghapus sub unit ini?
                </p>

            </ConfirmDialog>
        </>
    )
}

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
                    <StickyFooter
                            className="-mx-8 px-8 flex items-center justify-between py-4"
                            stickyClass="border-t bg-white dark:bg-gray-800 border-gray-200 dark:border-gray-700"
                    >
                        <div>
                                {type === 'edit' && (
                                    <DeleteSubunitButton
                                        onDelete={onDelete}
                                    />
                                )}
                            </div>
                            <div className="md:flex items-center">
                                    <Button
                                        size="sm"
                                        className="ltr:mr-2 rtl:ml-2"
                                        onClick={() => onDiscard?.()}
                                        type="button"
                                    >
                                        Discard
                                    </Button>
                                    <Button
                                        size="sm"
                                        variant="solid"
                                        loading={isSubmitting}
                                        icon={<AiOutlineSave />}
                                        type="submit"
                                    >
                                        Save
                                    </Button>
                                </div>
                    </StickyFooter>
                    </FormContainer>
                </Form>
            )}
        </Formik>
        </>
    )
})

export default ProgramForm