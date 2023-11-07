import React from "react";
import { Button, FormItem, FormContainer, Select, Input } from 'components/ui'
import { Field, Form, Formik } from 'formik'
import { HiArrowSmLeft } from 'react-icons/hi'
import * as Yup from 'yup'
import { AdaptableCard, Container } from "components/shared";
import reducer from "./store";
import ProgramTableSearch from "./components/ProgramTableSearch";
import { injectReducer } from "store";
import ProgramTableTools from "./components/ProgramTableTools";
import ProgramTable from "./components/ProgramTable";

injectReducer('programs', reducer)


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
        <AdaptableCard className="h-full" bodyClass="h-full">
            <div className="lg:flex items-center justify-between mb-4">
                <h3 className="mb-4 lg:mb-0">Program</h3>
               <ProgramTableTools/>
            </div>
            <ProgramTable/>
        </AdaptableCard>

       
    )
}

export default Programs