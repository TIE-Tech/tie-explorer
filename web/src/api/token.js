import { request } from '@/utils/request'

export function GetTokenList(page, type, pageSize = 20) {
    return request({
        url: '/api/tokens',
        method: 'GET',
        params: {
            t: type,
            p: page,
            ps: pageSize
        }
    })
}

export function GetTokenInfo(address) {
    return request({
        url: '/api/token/'+address,
        method: 'GET',
        params: {}
    })
}