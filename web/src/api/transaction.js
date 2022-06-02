import { request } from '@/utils/request'

export function GetTransactionList(page, block = 0, pageSize = 20) {
    return request({
        url: '/api/transactions',
        method: 'GET',
        params: {
            p: page,
            ps: pageSize,
            block: block
        }
    })
}

export function GetTransactionInfo(hash) {
    return request({
        url: '/api/transaction/' + hash,
        method: 'GET',
        params: {}
    })
}