import React from 'react'
import { Button, Card } from 'components/ui'
import { HiCheckCircle,HiOutlinePencil } from 'react-icons/hi'
import Table from 'components/ui/Table'
import { useNavigate } from 'react-router-dom'

const { Tr, Th, Td, THead, TBody } = Table

const Unit=()=>{
    const navigate = useNavigate()

    const onEdit = () => {
        navigate(`/setting/unit/edit`)
    }
    const headerExtraContent = (
        <span className="flex items-center">
            <Button className="mr-2" icon={<HiOutlinePencil />} variant="solid" onClick={onEdit}>
            </Button>
        </span>
    )

    const cardFooter = (
        <div className="flex">
            <Button className="mr-2" icon={<HiOutlinePencil />} variant="solid">
                <span>Edit</span>
            </Button>
        </div>
    )

    return (
        <div className="flex flex-col gap-4 h-full">
            <Card
                header="KONI Kabupaten Bogor"
                headerExtra={headerExtraContent}
            >
                <Table>
                <THead>
                    <Tr>
                        <Th>Kode</Th>
                        <Th>Kepala</Th>
                        <Th>Alamat</Th>
                    </Tr>
                </THead>
                <TBody>
                    <Tr>
                        <Td>KONI Kab.Bogor</Td>
                        <Td>Dedi Ade Bachtiar, SE. MM. MBA</Td>
                        <Td>Deket DISPORA Kab.Bogor</Td>
                    </Tr>
                </TBody>
            </Table>
            </Card>
            <Card
                 header="Daftar Sub Unit"
                 footer={cardFooter}
            >
                <div className="flex flex-col xl:flex-row gap-4">
                <Table>
                <THead>
                    <Tr>
                        <Th>Company</Th>
                        <Th>Contact</Th>
                        <Th>Country</Th>
                    </Tr>
                </THead>
                <TBody>
                    <Tr>
                        <Td>Alfreds Futterkiste</Td>
                        <Td>Maria Anders</Td>
                        <Td>Germany</Td>
                    </Tr>
                    <Tr>
                        <Td>Centro comercial Moctezuma</Td>
                        <Td>Francisco Chang</Td>
                        <Td>Mexico</Td>
                    </Tr>
                    <Tr>
                        <Td>Ernst Handel</Td>
                        <Td>Roland Mendel</Td>
                        <Td>Austria</Td>
                    </Tr>
                </TBody>
            </Table>
                <div>
            
        </div>
            </div>

            </Card>
            
        </div>
    )
}

export default Unit