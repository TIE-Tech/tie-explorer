export function Pagination(act, p) {
    let pagination = '<div class="d-md-flex justify-content-between m-3"><div class="pb-2" style="line-height: 2.2rem;">Showing <span>'+(parseInt(p.offset) + 1)+'</span> to <span>'+p.limit+'</span> of <span>'+p.total+'</span> entries</div>'

    let disableFirst = ''
    if (p.page === 1) {
        disableFirst = 'disabled'
    }

    let disablePrev = ''
    let pagePrev = p.page - 1
    if (pagePrev <= 0) {
        pagePrev = p.page
        disablePrev = 'disabled'
    }

    let disableNext = ''
    let pageNext = p.page + 1
    if (pageNext > p.totalPage) {
        pageNext = p.totalPage
        disableNext = 'disabled'
    }

    let disableLast = ''
    if (p.page === p.totalPage) {
        disableLast = 'disabled'
    }

    pagination += '<div>' +
        '<a href="/'+act+'?p=1" class="btn btn-sm me-2 p-2 '+disableFirst+'">First</a>' +
        '<a href="/'+act+'?p='+pagePrev+'" class="btn btn-icon me-2 '+disablePrev+'">' +
        '<svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"></path><polyline points="15 6 9 12 15 18"></polyline></svg>' +
        '</a>' +
        '<span class="btn btn-sm me-2 p-2 disabled">Page '+p.page+' of '+p.totalPage+'</span>' +
        '<a href="/'+act+'?p='+pageNext+'" class="btn btn-icon me-2 '+disableNext+'">' +
        '<svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"></path><polyline points="9 6 15 12 9 18"></polyline></svg>' +
        '</a>' +
        '<a href="/'+act+'?p='+p.totalPage+'" class="btn btn-sm p-2 me-2 '+disableLast+'">Last</a>' +
        '</div></div>'

    return pagination
}