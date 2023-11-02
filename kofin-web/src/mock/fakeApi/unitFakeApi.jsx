export default function unitFakeApi(server,apiPrefix){
    server.get(`${apiPrefix}/units`,(schema)=>{
        return schema.db.unitData[0]
    })

    server.get(`${apiPrefix}/units/subunits`,(schema)=>{
        return schema.db.subUnitData
    })
}