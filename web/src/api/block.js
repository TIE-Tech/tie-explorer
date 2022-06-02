import { request } from '@/utils/request'

export function GetBlockList(page, pageSize = 20) {
    return request({
        url: '/api/blocks',
        method: 'GET',
        params: {
            p: page,
            ps: pageSize
        }
    })
}

export function GetBlockInfo(number) {
    return request({
        url: '/api/block/'+number,
        method: 'GET',
        params: {}
    })
}