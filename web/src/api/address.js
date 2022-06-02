import { request } from '@/utils/request'

export function GetAddressList(page, pageSize = 20) {
    return request({
        url: '/api/addresses',
        method: 'GET',
        params: {
            p: page,
            ps: pageSize
        }
    })
}

export function GetAddressInfo(address) {
    return request({
        url: '/api/address/'+address,
        method: 'GET',
        params: {}
    })
}