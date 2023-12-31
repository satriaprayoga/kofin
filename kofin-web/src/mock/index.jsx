import { createServer } from 'miragejs'
import appConfig from 'configs/app.config'

import { signInUserData } from './data/authData'
import { unitData,subUnitData } from './data/unitData'
import { programsData } from './data/programData'
import { kegiatansData } from './data/kegiatanData'
import { budgetData, programBudgetData, kegiatanBudgetData } from './data/budgetData'

import { authFakeApi } from './fakeApi'
import unitFakeApi from './fakeApi/unitFakeApi'
import programFakeApi from './fakeApi/programFakeApi'
import kegiatanFakeApi from './fakeApi/kegiatanFakeApi'
import budgetFakeApi from './fakeApi/budgetFakeApi'

const { apiPrefix } = appConfig

export default function mockServer({ environment = 'test' }) {
    return createServer({
        environment,
        seeds(server) {
            server.db.loadData({
                signInUserData,
                unitData,
                subUnitData,
                programsData,
                kegiatansData,
                budgetData,
                programBudgetData,
                kegiatanBudgetData
            })
        },
        routes() {
            this.urlPrefix = ''
            this.namespace = ''
            this.passthrough((request) => {
                let isExternal = request.url.startsWith('http')
                return isExternal
            })
            this.passthrough()

            authFakeApi(this, apiPrefix)
            unitFakeApi(this,apiPrefix)
            programFakeApi(this,apiPrefix)
            kegiatanFakeApi(this,apiPrefix)
            budgetFakeApi(this,apiPrefix)
        },
    })
}
