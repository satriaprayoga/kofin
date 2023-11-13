import paginate from "utils/paginate"
import sortBy from "utils/sortBy"
import wildCardSearch from "utils/wildCardSearch"

export default function programFakeApi(server,apiPrefix){
    server.post(`${apiPrefix}/programs`,(schema,{requestBody})=>{
        const body = JSON.parse(requestBody)
        const {pageIndex, pageSize, sort, query} = body
        const {order,key}=sort
        const programs = schema.db.programsData
        const sanitizePrograms=programs.filter(
            (elm)=> typeof elm !=='function'
        )
        let data = sanitizePrograms
        let total = programs.length

        if ((key == 'program_name' || key == 'unit_name') && order){
            data.sort(sortBy(key, order === 'desc', (a) => a.toUpperCase()))
        }else {
            data.sort(sortBy(key, order === 'desc', parseInt))
        }
        if (query) {
            data = wildCardSearch(data, query)
            total = data.length
        }

        data = paginate(data, pageSize, pageIndex)

        const responseData = {
            data: data,
            total: total,
        }
        return responseData
    })
    

    server.get(`${apiPrefix}/programs/get`,(schema,{queryParams})=>{
        const id=queryParams.id
        const program=schema.db.programsData.find(id)
        return program
    })

    server.post(`${apiPrefix}/programs/create`,(schema,{requestBody})=>{
        const data = JSON.parse(requestBody)
        schema.db.programsData.insert(data)
        return true
    })

    server.put(`${apiPrefix}/programs/update`,(schema,{requestBody})=>{
        const data = JSON.parse(requestBody)
        const {id} = data
        schema.db.programsData.update({id},data)
        return true
    })

    server.del(`${apiPrefix}/programs/delete`,(schema,{requestBody})=>{
        const {id} = JSON.parse(requestBody)
        schema.db.programsData.remove({id})
        return true
    })

    server.get(`${apiPrefix}/programs/unit/`,(schema,{queryParams})=>{
        const id=queryParams.id
        if(parseInt(id)===0){
            const program=schema.db.programsData
            return program
           
        }else{
            const program=schema.db.programsData.where({unit_id:parseInt(id)})
            return program
        }
      
    })
}