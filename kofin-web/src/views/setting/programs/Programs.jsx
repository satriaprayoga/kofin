import React from "react";
import { Button, FormItem, FormContainer, Select, Input } from 'components/ui'
import { Field, Form, Formik } from 'formik'
import { HiArrowSmLeft } from 'react-icons/hi'
import * as Yup from 'yup'
import { Container } from "components/shared";


const validationSchema = Yup.object().shape({
    organizationSize: Yup.string().required(
        'Please select your organization size'
    ),
})

const colourOptions = [
    { value: 'ocean', label: 'Ocean', color: '#00B8D9' },
    { value: 'blue', label: 'Blue', color: '#0052CC' },
    { value: 'purple', label: 'Purple', color: '#5243AA' },
    { value: 'red', label: 'Red', color: '#FF5630' },
    { value: 'orange', label: 'Orange', color: '#FF8B00' },
    { value: 'yellow', label: 'Yellow', color: '#FFC400' },
    { value: 'green', label: 'Green', color: '#36B37E' },
    { value: 'forest', label: 'Forest', color: '#00875A' },
    { value: 'slate', label: 'Slate', color: '#253858' },
    { value: 'silver', label: 'Silver', color: '#666666' },
]

const Programs=()=>{
    return(
        <Container>
            <div className="text-center">
            <h3 className="mb-2">Daftar Program</h3>
            <div className="mt-8 max-w-[300px] lg:min-w-[300px] mx-auto">
            <Select
                size="sm"
                className="mb-4"
                placeholder="Please Select"
                options={colourOptions}
            ></Select>
            </div>
        </div>
        </Container>
       
    )
}

export default Programs