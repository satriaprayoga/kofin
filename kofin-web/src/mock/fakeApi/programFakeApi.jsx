export default function programFakeApi(server,apiPrefix){
    server.get(`${apiPrefix}/programs`,(schema)=>{
        return schema.db.programData
    })

    server.get(`${apiPrefix}/programs/get`,(schema,{queryParams})=>{
        const id=queryParams.id
        const program=schema.db.programData.find(id)
        return program
    })

    server.put(`${apiPrefix}/programs/update`,(schema,{requestBody})=>{
        const data = JSON.parse(requestBody)
        const {id} = data
        schema.db.programData.update({id},data)
        return true
    })

    server.del(`${apiPrefix}/programs/delete`,(schema,{requestBody})=>{
        const {id} = JSON.parse(requestBody)
        schema.db.programData.remove({id})
        return true
    })

    server.get(`${apiPrefix}/programs/unit/`,(schema,{requestBody})=>{
        const {id} = JSON.parse(requestBody)
        return schema.db.programData.where({unit_id:id})
    })
}