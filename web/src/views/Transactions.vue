<template>
  <header-view />

  <div class="page-wrapper">
    <div class="page-body">
      <div class="container-xl">
        <div class="page-header d-print-none m-0 p-0 pb-3">
          <div class="row align-items-center">
            <div class="col">
              <h2 class="page-title">
                Transactions
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
                  <th>Tx Hash</th>
                  <th>Type</th>
                  <th>Block</th>
                  <th>Age</th>
                  <th>From</th>
                  <th>To</th>
                  <th>Value</th>
                  <th>Tx Fee</th>
                </tr>
                </thead>
                <tbody class="tbody-height" id="transactions-tbody">
                <tr><td colspan="8" class="text-center"><span class="spinner-border" role="status"></span></td></tr>
                </tbody>
              </table>
            </div>
          </div>
          <div class="m-0" style="background-color: unset;color: #CCCCCC" id="footer-pagination"></div>
        </div>
      </div>
    </div>

    <footer-view />
  </div>
</template>

<script>
import HeaderView from "@/components/Header";
import FooterView from "@/components/Footer";
import {GetTransactionList} from "@/api/transaction";
import {Pagination} from "@/utils/pagination";
import Moment from "moment";
import Web3 from "web3";
export default {
  name: "TransactionList",
  components: {FooterView, HeaderView},
  mounted() {
    let page = 1;
    let block = 0;
    if (this.$route.query.p > 0) {
      page = this.$route.query.p
    }
    if (this.$route.query.block > 0) {
      block = this.$route.query.block
    }
    this.getTransactions(page, block)
  },
  methods: {
    getTransactions(page, block) {
      GetTransactionList(page, block).then(res => {
        if (res.code === 0) {
          if (res.data.transactions.length > 0) {
            let html = ''
            let pagination = Pagination('transactions', res.data.pagination)
            const data = res.data.transactions
            for (let i=0; i<data.length; i++) {
              html += '<tr><td data-label="Tx Hash"><div class="text-truncate text-md" style="max-width: 15rem">' +
                  '<a href="/transaction/'+data[i].transactionHash+'" class="hash-tag text-truncate">'+data[i].transactionHash+'</a>' +
                  '</div></td><td data-label="Type">'+data[i].genre+'</td>' +
                  '<td data-label="Block"><a href="/block/'+data[i].blockNumber+'" class="hash-tag text-truncate">'+data[i].blockNumber+'</a></td>' +
                  '<td data-label="Age" style="color: #bbb !important">'+this.formatTimestamp(data[i].timestamp)+'</td>' +
                  '<td data-label="From"><div class="text-truncate text-md" style="max-width: 15rem"><a href="/address/'+data[i].from+'" class="hash-tag text-truncate">'+data[i].from+'</a></div></td>' +
                  '<td data-label="To"><div class="text-truncate text-md" style="max-width: 15rem"><a href="/address/'+data[i].to+'" class="hash-tag text-truncate">'+data[i].to+'</a></div></td>' +
                  '<td data-label="Value" style="color: #bbb !important">'+this.formatAmount(data[i].value)+' Tie</td>' +
                  '<td data-label="Tx Fee" style="color: #bbb !important">'+this.formatAmount(data[i].gasPrice, 'Gwei')+' Gwei</td></tr>'
            }
            document.getElementById("transactions-tbody").innerHTML = html
            document.getElementById("footer-pagination").innerHTML = pagination
            document.getElementById("header-pagination").innerHTML = pagination
            return
          }
        }
        document.getElementById("transactions-tbody").innerHTML = '<tr><td colspan="8" class="text-center fw-bold">暂无数据</td></tr>'
        document.getElementById("footer-pagination").style = 'display: none !important'
        document.getElementById("header-pagination").style = 'display: none !important'
      }).catch(() => {
        document.getElementById("transactions-tbody").innerHTML = '<tr><td colspan="8" class="text-center fw-bold">暂无数据</td></tr>'
        document.getElementById("footer-pagination").style = 'display: none !important'
        document.getElementById("header-pagination").style = 'display: none !important'
      })
    },
    formatTimestamp(timestamp) {
      if (isNaN(timestamp)) {
        return ''
      }
      return Moment.unix(timestamp).fromNow()
    },
    formatAmount(number, unit='ether') {
      if (isNaN(number)) {
        return 0
      }
      return Web3.utils.fromWei(number.toString(), unit)
    }
  }
}
</script>
