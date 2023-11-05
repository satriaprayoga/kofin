export default function unitFakeApi(server,apiPrefix){
    server.get(`${apiPrefix}/units`,(schema)=>{
        return schema.db.unitData[0]
    })

    server.get(`${apiPrefix}/units/subunits`,(schema)=>{
        return schema.db.subUnitData
    })

    server.post(`${apiPrefix}/units/subunits/create`,
        (schema,{requestBody})=>{
            const data = JSON.parse(requestBody)
            schema.db.subUnitData.insert(data)
            return true
        })

    server.get(`${apiPrefix}/units/subunit`,(schema,{queryParams})=>{
        const id=queryParams.id
        const subunit=schema.db.subUnitData.find(id)
        return subunit
    })

    server.put(`${apiPrefix}/units/subunit/update`,(schema,{requestBody})=>{
        const data = JSON.parse(requestBody)
        const {id} = data
        schema.db.subUnitData.update({id},data)
        return true
    })

    server.del(`${apiPrefix}/units/subunit/delete`,(schema,{requestBody})=>{
        const {id} = JSON.parse(requestBody)
        schema.db.subUnitData.remove({id})
        return true
    })
}