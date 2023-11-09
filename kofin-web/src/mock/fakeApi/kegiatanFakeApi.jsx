import paginate from "utils/paginate"
import sortBy from "utils/sortBy"
import wildCardSearch from "utils/wildCardSearch"

export default function kegiatanFakeApi(server,apiPrefix){
    server.post(`${apiPrefix}/kegiatans`,(schema,{requestBody})=>{
        const body = JSON.parse(requestBody)
        const {pageIndex, pageSize, sort, query, programId} = body
        const {order,key}=sort
        const kegiatans = schema.db.kegiatansData.where({program_id:programId})
        const sanitizeKegiatans=kegiatans.filter(
            (elm)=> typeof elm !=='function'
        )
        let data = sanitizeKegiatans
        let total = kegiatans.length

        if ((key == 'kegiatan_name') && order){
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
    

    server.get(`${apiPrefix}/kegiatans/get`,(schema,{queryParams})=>{
        const id=queryParams.id
        const kegiatan=schema.db.kegiatansData.find(id)
        return kegiatan
    })

    server.post(`${apiPrefix}/kegiatans/create`,(schema,{requestBody})=>{
        const data = JSON.parse(requestBody)
        schema.db.kegiatansData.insert(data)
        return true
    })

    server.put(`${apiPrefix}/kegiatans/update`,(schema,{requestBody})=>{
        const data = JSON.parse(requestBody)
        const {id} = data
        schema.db.kegiatansData.update({id},data)
        return true
    })

    server.del(`${apiPrefix}/kegiatans/delete`,(schema,{requestBody})=>{
        const {id} = JSON.parse(requestBody)
        schema.db.kegiatansData.remove({id})
        return true
    })

    server.get(`${apiPrefix}/kegiatans/program/`,(schema,{requestBody})=>{
        const {id} = JSON.parse(requestBody)
        return schema.db.kegiatansData.where({program_id:id})
    })
}