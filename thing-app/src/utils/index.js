export function localGet(key) {
    const value = window.localStorage.getItem(key)
    console.log(value)
    try {
        return JSON.parse(window.localStorage.getItem(key))
    } catch (error) {
        return value
    }
}

export function localSet(key, value) {
    window.localStorage.setItem(key, JSON.stringify(value))
}

export function localRemove(key) {
    window.localStorage.removeItem(key)
}

// 判断内容是否含有表情字符，现有数据库不支持。
export function hasEmoji(str = '') {
    const reg = /[^\u0020-\u007E\u00A0-\u00BE\u2E80-\uA4CF\uF900-\uFAFF\uFE30-\uFE4F\uFF00-\uFFEF\u0080-\u009F\u2000-\u201f\u2026\u2022\u20ac\r\n]/g;
    return str.match(reg) && str.match(reg).length
}

// 单张图片上传
export const uploadImgServer = 'http://backend-api-02.newbee.ltd/manage-api/v1/upload/file'
    // 多张图片上传
export const uploadImgsServer = 'http://backend-api-02.newbee.ltd/manage-api/v1/upload/files'

export const pathMap = {
    login: '登录',
    introduce: '系统介绍',
    dashboard: '邮件统计',
    add: '添加商品',
    swiper: '轮播图配置',
    hot: '热销商品配置',
    new: '新品上线配置',
    recommend: '为你推荐配置',
    category: '分类管理',
    level2: '分类二级管理',
    level3: '分类三级管理',
    good: '商品管理',
    guest: '会员管理',
    order: '订单管理',
    order_detail: '订单详情',
    addProduct: '添加物资',
    mine: '我的物资',
    product: '易物广场',
    account: '修改账户'
}

import { MAPPING } from './constants'

/**
 * @description: 获取Api BaseUrl
 * @param {*} key
 * @return {*}
 * @author: gumingchen
 */
export function getApiBaseUrl() {
    const baseUrl = process.env.VUE_APP_PROXY === 'true' ?
        `/proxy${ MAPPING }` :
        process.env.VUE_APP_BASE_API + MAPPING
    return baseUrl
}

/**
 * @description: 置空json数据
 * @param {*} json
 * @return {*}
 * @author: gumingchen
 */
export function clearJson(data) {
    const json = data
    let key
    for (key in json) {
        if (json[key] instanceof Array) {
            json[key] = []
        } else if (typeof json[key] === 'object' && Object.prototype.toString.call(json[key]).toLowerCase() === '[object object]' && !json[key].length) {
            json[key] = {}
        } else {
            json[key] = ''
        }
    }
}

/**
 * @description: 日期转字符串
 * @param {*} time 日期 默认当前日期
 * @param {*} format 格式
 * @return {*}
 * @author: gumingchen
 */
export function parseDate2Str(time = new Date(), format = '{y}-{M}-{d} {h}:{m}:{s}') {
    let result = ''
    let date = new Date()
    const type = typeof time
    if (type === 'object') {
        date = time
    } else if (type === 'number') {
        date = new Date(time)
    }
    const formatObj = {
        y: date.getFullYear(),
        M: date.getMonth() + 1,
        d: date.getDate(),
        h: date.getHours(),
        m: date.getMinutes(),
        s: date.getSeconds(),
        a: date.getDay()
    }
    result = format.replace(/\{[yMdhmsa]+\}/g, (val) => {
        const key = val.replace(/\{|\}/g, '')
        const value = formatObj[key]
        if (key === 'a') {
            return ['日', '一', '二', '三', '四', '五', '六'][value]
        }
        return value.toString().padStart(2, '0')
    })
    return result
}

/**
 * @description: 字符串转日期
 * @param {*} time 日期字符串
 * @param {*} separator 分隔符
 * @return {*}
 * @author: gumingchen
 */
export function parseStr2Date(time = '', separator = ['-', ' ', ':']) {
    let result = new Date()
    const regexp = `/[${ separator.join('') }]/g`
    const data = time.split(eval(regexp))
    switch (data.length) {
        case 3:
            result = new Date(+data[0], +data[1] - 1, +data[2])
            break
        case 6:
            result = new Date(+data[0], +data[1] - 1, +data[2], +data[3], +data[4], +data[5])
            break
    }
    return result
}

/**
 * @description: 毫秒转日期字符串
 * @param {*} time 毫秒 默认当前日期毫秒数
 * @param {*} format 格式
 * @return {*}
 * @author: gumingchen
 */
export function parseTime2Str(time = new Date().getTime(), format = '{y}-{M}-{d} {h}:{m}:{s}') {
    return parseDate2Str(new Date(time), format)
}

/**
 * @description: 日期转字符串
 * @param {*} time 毫秒 默认当前日期毫秒数
 * @param {*} format 格式
 * @return {*}
 * @author: gumingchen
 */
export function parseStringDate(time) {
    let result = parseDate2Str(new Date(time), '{y}-{M}-{d} {h}:{m}:{s}')
    if (time) {
        const now = parseDate2Str(new Date(), '{y}-{M}-{d}')
        const formats = result.split(' ')
        if (now === formats[0]) {
            result = formats[1]
        }
    }
    return result
}

/**
 * @description: 生成UUID
 * @param {*}
 * @return {*}
 * @author: gumingchen
 */
export function getUUID() {
    let result = ''
    const str = 'xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx'
    result = str.replace(/[xy]/gu, c => {
        const r = (Math.random() * 16) | 0
        const v = c === 'x' ? r : (r & 0x3) | 0x8
        return v.toString(16)
    })
    return result
}