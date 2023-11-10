import { THEME_ENUM } from 'constants/theme.constant'

/**
 * Since some configurations need to be match with specific themes,
 * we recommend to use the configuration that generated from demo.
 */

export const themeConfig = {
    themeColor: 'sky',
    direction: THEME_ENUM.DIR_LTR,
    mode: THEME_ENUM.MODE_LIGHT,
    primaryColorLevel: 500,
    cardBordered: false,
    panelExpand: false,
    controlSize: 'md',
    navMode: THEME_ENUM.NAV_MODE_THEMED,
    layout: {
        type: THEME_ENUM.LAYOUT_TYPE_STACKED_SIDE,
        sideNavCollapse: false,
    },
}
