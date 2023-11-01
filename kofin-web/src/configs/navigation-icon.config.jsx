import React from 'react'
import {
    HiOutlineColorSwatch,
    HiOutlineDesktopComputer,
    HiOutlineTemplate,
    HiOutlineViewGridAdd,
    HiOutlineHome,
    HiOutlineAdjustments,
    HiOutlineChartPie,
    HiOutlineOfficeBuilding
} from 'react-icons/hi'

const navigationIcon = {
    home: <HiOutlineHome />,
    setting:<HiOutlineAdjustments/>,
    budget:<HiOutlineChartPie/>,
    unit:<HiOutlineOfficeBuilding/>,
    singleMenu: <HiOutlineViewGridAdd />,
    collapseMenu: <HiOutlineTemplate />,
    groupSingleMenu: <HiOutlineDesktopComputer />,
    groupCollapseMenu: <HiOutlineColorSwatch />,
}

export default navigationIcon
