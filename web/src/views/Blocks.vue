<template>
  <header-view />

  <div class="page-wrapper">
    <div class="page-body">
      <div class="container-xl">
        <div class="page-header d-print-none m-0 p-0 pb-3">
          <div class="row align-items-center">
            <div class="col">
              <h2 class="page-title">
                Blocks
              </h2>
            </div>
          </div>
        </div>

        <div class="card">
          <div class="m-0" id="header-pagination"></div>
          <div class="card-body p-0">
            <div class="table-responsive">
              <table class="table table-vcenter table-mobile-md card-table table-nowrap">
                <thead>
                  <tr>
                    <th>Block</th>
                    <th>Age</th>
                    <th>Txn</th>
                    <th>Uncles</th>
                    <th>Miner</th>
                    <th>Gas Limit</th>
                    <th>Gas Used</th>
                    <th>Size</th>
                  </tr>
                </thead>
                <tbody class="tbody-height" id="blocks-tbody">
                <tr><td colspan="8" class="text-center"><span class="spinner-border" role="status"></span></td></tr>
                </tbody>
              </table>
            </div>
          </div>
          <div class="m-0" style="background-color: unset;color: #CCCCCC;" id="footer-pagination"></div>
        </div>
      </div>
    </div>

    <footer-view />
  </div>
</template>

<script>
import FooterView from "@/components/Footer";
import HeaderView from "@/components/Header";
import { GetBlockList } from "@/api/block";
import { Pagination } from "@/utils/pagination";
import Moment from "moment";
export default {
  name: "BlockList",
  components: {HeaderView, FooterView},
  mounted() {
    let page = 1;
    if (this.$route.query.p > 0) {
      page = this.$route.query.p
    }
    this.getBlocks(page)
  },
  methods: {
    getBlocks(page) {
      GetBlockList(page).then(res => {
        if (res.code === 0){
          if (res.data.blocks.length > 0) {
            let html = ''
            let pagination = Pagination('blocks', res.data.pagination)
            const data = res.data.blocks
            for (let i=0; i<data.length; i++) {
              html += '<tr>' +
                  '<td data-label="Block"><a href="/block/'+data[i].number+'">'+data[i].number+'</a></td>' +
                  '<td data-label="Age">'+this.formatTimestamp(data[i].timestamp)+'</td>' +
                  '<td data-label="Txn">'+data[i].txs+'</td>' +
                  '<td data-label="Uncles">'+data[i].uncles+'</td>' +
                  '<td data-label="Miner" class="text-truncate"><a href="/address/'+data[i].miner+'">'+data[i].miner+'</a></td>' +
                  '<td data-label="Gas Limit">'+data[i].gasLimit+'</td>' +
                  '<td data-label="Gas Used">'+data[i].gasUsed+'</td>' +
                  '<td data-label="Size">'+data[i].size+' Bytes</td>' +
                  '</tr>'
            }
            document.getElementById("blocks-tbody").innerHTML = html
            document.getElementById("footer-pagination").innerHTML = pagination
            document.getElementById("header-pagination").innerHTML = pagination
            return
          }
        }
        document.getElementById("blocks-tbody").innerHTML = '<tr><td colspan="8" class="text-center fw-bold">暂无数据</td></tr>'
        document.getElementById("footer-pagination").style = 'display: none !important'
        document.getElementById("header-pagination").style = 'display: none !important'
      }).catch(() => {
        document.getElementById("blocks-tbody").innerHTML = '<tr><td colspan="8" class="text-center fw-bold">暂无数据</td></tr>'
        document.getElementById("footer-pagination").style = 'display: none !important'
        document.getElementById("header-pagination").style = 'display: none !important'
      })
    },
    formatTimestamp(timestamp) {
      if (isNaN(timestamp)) {
        return ''
      }
      return Moment.unix(timestamp).fromNow()
    }
  }
}
</script>