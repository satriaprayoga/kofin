import React from 'react'
import {
    HiOutlineColorSwatch,
    HiOutlineDesktopComputer,
    HiOutlineTemplate,
    HiOutlineViewGridAdd,
    HiOutlineHome,
    HiOutlineAdjustments,
    HiOutlineChartPie,
    HiOutlineOfficeBuilding,
    HiOutlinePresentationChartLine,
    HiOutlineCog,
    HiOutlineSwitchVertical,
    HiOutlineX
} from 'react-icons/hi'

const navigationIcon = {
    home: <HiOutlineHome />,
    setting:<HiOutlineAdjustments/>,
    budget:<HiOutlineChartPie/>,
    unit:<HiOutlineOfficeBuilding/>,
    program:<HiOutlinePresentationChartLine/>,
    setup:<HiOutlineCog/>,
    inout:<HiOutlineSwitchVertical/>,
    hapus:<HiOutlineX/>,
    singleMenu: <HiOutlineViewGridAdd />,
    collapseMenu: <HiOutlineTemplate />,
    groupSingleMenu: <HiOutlineDesktopComputer />,
    groupCollapseMenu: <HiOutlineColorSwatch />,
}

export default navigationIcon
