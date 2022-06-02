import { request } from '@/utils/request'

export function GetDashboard() {
    return request({
        url: '/api/dashboard',
        method: 'GET',
        params: {}
    })
}

export function GetSearch(s, t) {
    return request({
        url: '/api/search',
        method: 'GET',
        params: {
            s: s,
            t: t,
        }
    })
}