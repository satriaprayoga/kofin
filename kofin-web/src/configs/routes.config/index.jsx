import React from 'react'
import authRoute from './authRoute'

export const publicRoutes = [...authRoute]

export const protectedRoutes = [
    {
        key: 'home',
        path: '/home',
        component: React.lazy(() => import('views/Home')),
        authority: [],
    },
    {
        key:'unit',
        path:'/setting/unit',
        component:React.lazy(()=>import('views/setting/unit/Unit')),
        authority:[],
        meta: {
            header: 'Pengaturan Unit Kerja',
        },
    },
    {
        key:'addSubunit',
        path:'/setting/unit/subunit/new',
        component:React.lazy(()=>import('views/setting/unit/SubunitNew')),
        authority:[],
        meta: {
            header: 'Tambah Sub Unit',
            headerContainer: true,
        },
    },
    {
        key:'editSubunit',
        path:'/setting/unit/subunit/edit/:subunitId',
        component:React.lazy(()=>import('views/setting/unit/SubunitEdit')),
        authority:[],
        meta: {
            header: 'Ubah Sub Unit',
            headerContainer: true,
        },
    },
    {
        key:'editUnit',
        path:'/setting/unit/edit',
        component:React.lazy(()=>import('views/setting/unit/UnitEdit')),
        authority:[],
        meta: {
            header: 'Edit Unit',
            headerContainer: true,
        },
    },
    {
        key:'program',
        path:'/setting/programs',
        component:React.lazy(()=>import('views/setting/programs/Programs')),
        authority:[],
       
    },
    {
        key:'addProgram',
        path:'/setting/programs/new',
        component:React.lazy(()=>import('views/setting/programs/ProgramNew')),
        authority:[],
        meta: {
            header: 'Tambah Program',
            headerContainer: true,
        },
    },
    {
        key:'editProgram',
        path:'/setting/programs/edit/:programId',
        component:React.lazy(()=>import('views/setting/programs/ProgramEdit')),
        authority:[],
        meta: {
            header: 'Edit Program',
            headerContainer: true,
        },
    },
    {
        key:'kegiatan',
        path:'/setting/programs/kegiatan/:programId',
        component:React.lazy(()=>import('views/setting/programs/kegiatan/Kegiatans')),
        authority:[],
        meta: {
            header: 'Daftar Kegiatan',
            headerContainer: true,
        },
    },
    {
        key:'addKegiatan',
        path:'/setting/programs/kegiatan/new/:programId',
        component:React.lazy(()=>import('views/setting/programs/kegiatan/KegiatanNew')),
        authority:[],
        meta: {
            header: 'Tambah Kegiatan',
            headerContainer: true,
        },
    },
    {
        key:'editKegiatan',
        path:'/setting/programs/kegiatan/edit/:kegiatanId',
        component:React.lazy(()=>import('views/setting/programs/kegiatan/KegiatanEdit')),
        authority:[],
        meta: {
            header: 'Tambah Kegiatan',
            headerContainer: true,
        },
    },
    /** Example purpose only, please remove */
    {
        key: 'singleMenuItem',
        path: '/single-menu-view',
        component: React.lazy(() => import('views/demo/SingleMenuView')),
        authority: [],
    },
    {
        key: 'collapseMenu.item1',
        path: '/collapse-menu-item-view-1',
        component: React.lazy(() => import('views/demo/CollapseMenuItemView1')),
        authority: [],
    },
    {
        key: 'collapseMenu.item2',
        path: '/collapse-menu-item-view-2',
        component: React.lazy(() => import('views/demo/CollapseMenuItemView2')),
        authority: [],
    },
    {
        key: 'groupMenu.single',
        path: '/group-single-menu-item-view',
        component: React.lazy(() =>
            import('views/demo/GroupSingleMenuItemView')
        ),
        authority: [],
    },
    {
        key: 'groupMenu.collapse.item1',
        path: '/group-collapse-menu-item-view-1',
        component: React.lazy(() =>
            import('views/demo/GroupCollapseMenuItemView1')
        ),
        authority: [],
    },
    {
        key: 'groupMenu.collapse.item2',
        path: '/group-collapse-menu-item-view-2',
        component: React.lazy(() =>
            import('views/demo/GroupCollapseMenuItemView2')
        ),
        authority: [],
    },
]
