export default function budgetFakeApi(server,apiPrefix){
    server.put(`${apiPrefix}/budget/setup`,(schema,{requestBody})=>{
        const body = JSON.parse(requestBody)
        schema.db.budgetData.update({id:1},body)
        return true
    })
    server.get(`${apiPrefix}/budget`,(schema)=>{
        const budget=schema.db.budgetData.find(1)
        return budget
    })
  
}
