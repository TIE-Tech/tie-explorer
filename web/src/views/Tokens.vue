<template>
  <header-view />

  <div class="page-wrapper">
    <div class="page-body">
      <div class="container-xl">
        <div class="page-header d-print-none m-0 p-0 pb-3">
          <div class="row align-items-center">
            <div class="col">
              <h2 class="page-title">
                {{ this.$route.params.type }} Tokens
              </h2>
            </div>
          </div>
        </div>

        <div class="card">
          <div class="m-0" id="header-pagination"></div>
          <div class="card-body p-0">
            <div class="table-responsive">
              <table v-if="this.token" class="table table-vcenter table-mobile-md table-nowrap card-table">
                <thead>
                <tr>
                  <th>Name</th>
                  <th>Symbol</th>
                  <th>Decimal</th>
                  <th>Total Supply</th>
                  <th>Holders</th>
                </tr>
                </thead>
                <tbody id="tokens-tbody">
                <tr><td colspan="5" class="text-center"><span class="spinner-border" role="status"></span></td></tr>
                </tbody>
              </table>
              <table v-else class="table table-vcenter table-mobile-md table-nowrap card-table">
                <thead>
                <tr>
                  <th>Address</th>
                  <th>Balance</th>
                  <th>Txn Sends</th>
                </tr>
                </thead>
                <tbody class="tbody-height" id="ether-tbody">
                <tr><td colspan="3" class="text-center"><span class="spinner-border" role="status"></span></td></tr>
                </tbody>
              </table>
            </div>
          </div>
          <div style="background-color: unset;color: #CCCCCC" id="footer-pagination"></div>
        </div>
      </div>
    </div>

    <footer-view />
  </div>
</template>



<script>
import HeaderView from "@/components/Header";
import FooterView from "@/components/Footer";
import {GetTokenList} from "@/api/token";
import {Pagination} from "@/utils/pagination";
import Web3 from "web3";
export default {
  name: "TokenList",
  components: {FooterView, HeaderView},
  data() {
    return {
      token: true
    }
  },
  mounted() {
    if (this.$route.params.type) {
      let page = 1;
      if (this.$route.query.p > 0) {
        page = this.$route.query.p
      }
      if (this.$route.params.type === 'tie') {
        this.token = false
      }
      this.getTokens(page, this.$route.params.type)
    }
  },
  methods: {
    getTokens(page, type) {
      GetTokenList(page, type).then(res => {
        if (res.code === 0) {
          let html = ''
          const data = res.data.tokens
          let pagination = Pagination('tokens', res.data.pagination)
          if (type === 'tie') {
            if (data.length > 0) {
              for (let i=0; i<data.length; i++) {
                html += '<tr><td data-label="Address" class="text-truncate"><a href="/address/'+data[i].address+'">'+data[i].address+'</a></td>' +
                    '<td data-label="Balance">'+this.formatAmount(data[i].balance)+' Tie</td>' +
                    '<td data-label="Txn Sends">'+data[i].senders+'</td></tr>'
              }
              document.getElementById("ether-tbody").innerHTML = html
              document.getElementById("footer-pagination").innerHTML = pagination
              document.getElementById("header-pagination").innerHTML = pagination
            } else {
              document.getElementById("ether-tbody").innerHTML = '<tr><td colspan="3" class="text-center fw-bold">暂无数据</td></tr>'
              document.getElementById("footer-pagination").style = 'display: none !important'
              document.getElementById("header-pagination").style = 'display: none !important'
            }
          } else {
            if (data.length > 0) {
              for (let i=0; i<data.length; i++) {
                let decimals = 'unknown'
                if (data[i].decimals > 0) {
                  decimals = data[i].decimals
                }
                let totalSupply = 'unknown'
                if (data[i].totalSupply > 0) {
                  totalSupply = data[i].totalSupply
                }
                html += '<tr><td data-label="Name"><a href="/token/'+data[i].contractAddress+'">'+data[i].name+'</a></td>' +
                    '<td data-label="Symbol">'+data[i].symbol+'</td>' +
                    '<td data-label="Decimal">'+decimals+'</td>' +
                    '<td data-label="Total Supply">'+totalSupply+'</td>' +
                    '<td data-label="Holders">'+data[i].holders+'</td></tr>'
              }
              document.getElementById("tokens-tbody").innerHTML = html
              document.getElementById("footer-pagination").innerHTML = pagination
              document.getElementById("header-pagination").innerHTML = pagination
            } else {
              document.getElementById("tokens-tbody").innerHTML = '<tr><td colspan="5" class="text-center fw-bold">暂无数据</td></tr>'
              document.getElementById("footer-pagination").style = 'display: none !important'
              document.getElementById("header-pagination").style = 'display: none !important'
            }
          }
        }
      }).catch(() => {
        if (this.token) {
          document.getElementById("tokens-tbody").innerHTML = '<tr><td colspan="5" class="text-center fw-bold">暂无数据</td></tr>'
        } else {
          document.getElementById("ether-tbody").innerHTML = '<tr><td colspan="3" class="text-center fw-bold">暂无数据</td></tr>'
        }
        document.getElementById("footer-pagination").style = 'display: none !important'
        document.getElementById("header-pagination").style = 'display: none !important'
      })
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
