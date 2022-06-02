<template>
  <header-view />

  <div class="page-wrapper">
    <div class="page-body">
      <div class="container-xl">
        <div class="page-header d-print-none m-0 p-0 pb-3">
          <div class="row">
            <div class="col-md-1 fw-bold" style="font-size: 1.4rem" v-if="this.addressInfo.isContract">Contract</div>
            <div class="col-md-1 fw-bold" style="font-size: 1.4rem" v-else>Address</div>
            <div class="col-md-10 text-muted" style="line-height: 2rem">{{ this.$route.params.address }}</div>
          </div>
        </div>

        <div class="row">
          <div class="col-md-6">
            <div class="card">
              <div class="card-header">
                <h3 class="card-title">Overview</h3>
              </div>
              <div class="card-body">
                <div class="row mb-3">
                  <div class="col-4">Balance: </div>
                  <div class="col-8">{{ this.formatAmount(this.addressInfo.balance) }} Tie</div>
                </div>
                <div class="row">
                  <div class="col-4">Token: </div>
                  <div class="col-8" id="token-dropdown">0</div>
                </div>
              </div>
            </div>
          </div>
          <div class="col-md-6">
            <div class="card">
              <div class="card-header">
                <h3 class="card-title">More Info</h3>
              </div>
              <div class="card-body">
                <div class="row mb-3">
                  <div class="col-4">Transactions: </div>
                  <div class="col-8">{{ this.transactions }}</div>
                </div>
                <div class="row">
                  <div class="col-4">Transfers: </div>
                  <div class="col-8">{{ this.transfers }}</div>
                </div>
                <div class="row mt-3" v-if="this.addressInfo.isContract">
                  <div class="col-2">Creator: </div>
                  <div class="col-10">
                    <div class="d-flex">
                    <a class="text-truncate" :href="'/address/'+this.addressInfo.creator">{{ this.addressInfo.creator }}</a>
                    <span class="me-2 ms-2">at</span>
                    <a class="text-truncate" :href="'/transaction/'+this.addressInfo.createAt">{{ this.addressInfo.createAt }}</a>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="card mt-4">
          <ul class="nav nav-tabs" data-bs-toggle="tabs">
            <li class="nav-item">
              <a href="#pane-transaction" class="nav-link active" data-bs-toggle="tab">Transactions</a>
            </li>
            <li class="nav-item" v-if="this.display.erc20">
              <a href="#pane-erc-20" id="tab-erc-20" class="nav-link" data-bs-toggle="tab">ERC-20 Transfer</a>
            </li>
            <li class="nav-item" v-if="this.display.erc721">
              <a href="#pane-erc-721" id="tab-erc-721" class="nav-link" data-bs-toggle="tab">ERC-721 Transfer</a>
            </li>
            <li class="nav-item" v-if="this.display.log">
              <a href="#pane-log" id="tab-log" class="nav-link" data-bs-toggle="tab">Log</a>
            </li>
            <li class="nav-item" v-if="this.addressInfo.isContract">
              <a href="#pane-contract" id="tab-contract" class="nav-link" data-bs-toggle="tab">Contract</a>
            </li>
            <li class="nav-item">
              <a href="#pane-balance-history" @click="this.renderLineChart()" id="tab-balance-history" class="nav-link" data-bs-toggle="tab">Coin Balance History</a>
            </li>
          </ul>
          <div class="card-body p-0">
            <div class="tab-content">
              <div class="tab-pane active show" id="pane-transaction">
                <div id="txns-total-tips"></div>
                <div class="table-responsive">
                    <table class="table table-vcenter card-table">
                      <thead>
                      <tr>
                        <th>Tx Hash</th>
                        <th>Type</th>
                        <th>Block</th>
                        <th>Age</th>
                        <th>From</th>
                        <th></th>
                        <th>To</th>
                        <th>Value</th>
                        <th>Tx Fee</th>
                      </tr>
                      </thead>
                      <tbody id="transactions-tbody">
                      <tr><td colspan="9" class="text-center"><span class="spinner-border" role="status"></span></td></tr>
                      </tbody>
                    </table>
                  </div>
              </div>
              <div class="tab-pane" id="pane-erc-20">
                <table class="table table-vcenter card-table">
                  <thead>
                  <tr>
                    <th>Tx Hash</th>
                    <th>Block</th>
                    <th>Age</th>
                    <th>From</th>
                    <th></th>
                    <th>To</th>
                    <th>Value</th>
                    <th>Token</th>
                  </tr>
                  </thead>
                  <tbody id="erc-20-tbody">
                  <tr><td colspan="8" class="text-center"><span class="spinner-border" role="status"></span></td></tr>
                  </tbody>
                </table>
              </div>
              <div class="tab-pane" id="pane-erc-721">
                <table class="table table-vcenter card-table">
                  <thead>
                  <tr>
                    <th>Tx Hash</th>
                    <th>Block</th>
                    <th>Age</th>
                    <th>From</th>
                    <th></th>
                    <th>To</th>
                    <th>Value</th>
                    <th>Token</th>
                  </tr>
                  </thead>
                  <tbody id="erc-721-tbody">
                  <tr><td colspan="8" class="text-center"><span class="spinner-border" role="status"></span></td></tr>
                  </tbody>
                </table>
              </div>
              <div class="tab-pane p-3" id="pane-log"></div>
              <div class="tab-pane p-3" id="pane-contract">
                <div class="mb-3">
                  Are you the contract creator? <a class="" href="/verify_contract?a=">Verify and Publish</a> your contract source code today!
                </div>
                <div class="form-control input-data">{{ this.addressInfo.contractCode }}</div>
              </div>
              <div class="tab-pane" id="pane-balance-history">
                <div id="charts" style="height: 20rem"></div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <footer-view />
  </div>
</template>

<script>
import HeaderView from "@/components/Header";
import FooterView from "@/components/Footer";
import {GetAddressInfo} from "@/api/address";
import Moment from "moment";
import Web3 from "web3";
import * as echarts from 'echarts'

export default {
  name: "AddressInfo",
  components: {FooterView, HeaderView},
  data() {
    return {
      addressInfo: {},
      transfers: 0,
      transactions: 0,
      display: {
        erc20: false,
        erc721: false,
        log: false,
      },
      render: false,
      LineList: []
    }
  },
  mounted() {
    if (this.$route.params.address) {
      this.getAddressInfo(this.$route.params.address)
    }
  },
  methods: {
    getAddressInfo(address) {
      GetAddressInfo(address).then(res => {
        this.addressInfo = res.data.info
        this.transactions = res.data.txns_total
        this.transfers = res.data.erc20Total + res.data.erc721Total
        if (res.data.txns.length > 0) {
          let data = res.data.txns
          let html = ''
          for (let i=0; i<data.length; i++) {
            html += '<tr><td><div class="text-truncate" style="max-width: 15rem">' +
                '<a href="/transaction/'+data[i].transactionHash+'" class="hash-tag text-truncate">'+data[i].transactionHash+'</a>' +
                '</div></td><td>'+data[i].genre+'</td>' +
                '<td><a href="/block/'+data[i].blockNumber+'" class="hash-tag text-truncate">'+data[i].blockNumber+'</a></td>' +
                '<td>'+this.formatTimestamp(data[i].timestamp)+'</td>' +
                '<td><div class="text-truncate" style="max-width: 15rem"><a href="/address/'+data[i].from+'" class="hash-tag text-truncate">'+data[i].from+'</a></div></td>' +
                '<td><span class="badge bg-purple-lt">'+data[i].direction+'</span></td>' +
                '<td><div class="text-truncate" style="max-width: 15rem"><a href="/address/'+data[i].to+'" class="hash-tag text-truncate">'+data[i].to+'</a></div></td>' +
                '<td>'+this.formatAmount(data[i].value)+' Tie</td>' +
                '<td>'+this.formatAmount(data[i].gasPrice, 'Gwei')+' Gwei</td></tr>'
          }
          //document.getElementById("txns-total-tips").innerHTML = '<p class="m-2"></p>'
          document.getElementById("transactions-tbody").innerHTML = html
        } else {
          document.getElementById("transactions-tbody").innerHTML = '<tr><td colspan="9" class="text-center fw-bold">暂无数据</td></tr>'
        }

        if (res.data.erc20.length > 0) {
          this.display.erc20 = true
          let data = res.data.erc20
          let html = ''
          for (let i=0; i<data.length; i++) {
            let tokenId = ''
            if (data[i].tokenId != '') {
              tokenId = '[' + data[i].tokenId + ']'
            }
            html += '<tr><td><div class="text-truncate" style="max-width: 15rem">' +
                '<a href="/transaction/'+data[i].transactionHash+'" class="hash-tag text-truncate">'+data[i].transactionHash+'</a>' +
                '</div></td>' +
                '<td><a href="/block/'+data[i].blockNumber+'" class="hash-tag text-truncate">'+data[i].blockNumber+'</a></td>' +
                '<td>'+this.formatTimestamp(data[i].timestamp)+'</td>' +
                '<td><div class="text-truncate" style="max-width: 15rem"><a href="/address/'+data[i].from+'" class="hash-tag text-truncate">'+data[i].from+'</a></div></td>' +
                '<td><span class="badge bg-purple-lt">'+data[i].direction+'</span></td>' +
                '<td><div class="text-truncate" style="max-width: 15rem"><a href="/address/'+data[i].to+'" class="hash-tag text-truncate">'+data[i].to+'</a></div></td>' +
                '<td>'+data[i].value+'</td>' +
                '<td>'+tokenId+' '+data[i].tokenName+'</td></tr>'
          }
          document.getElementById('erc-20-tbody').innerHTML = html
        }

        if (res.data.erc721.length > 0) {
          this.display.erc721 = true
          let data = res.data.erc721
          let html = ''
          for (let i=0; i<data.length; i++) {
            let tokenId = ''
            if (data[i].tokenId != '') {
              tokenId = '[' + data[i].tokenId + ']'
            }
            html += '<tr><td><div class="text-truncate" style="max-width: 15rem">' +
                '<a href="/transaction/'+data[i].transactionHash+'" class="hash-tag text-truncate">'+data[i].transactionHash+'</a>' +
                '</div></td>' +
                '<td><a href="/block/'+data[i].blockNumber+'" class="hash-tag text-truncate">'+data[i].blockNumber+'</a></td>' +
                '<td>'+this.formatTimestamp(data[i].timestamp)+'</td>' +
                '<td><div class="text-truncate" style="max-width: 15rem"><a href="/address/'+data[i].from+'" class="hash-tag text-truncate">'+data[i].from+'</a></div></td>' +
                '<td><span class="badge bg-purple-lt">'+data[i].direction+'</span></td>' +
                '<td><div class="text-truncate" style="max-width: 15rem"><a href="/address/'+data[i].to+'" class="hash-tag text-truncate">'+data[i].to+'</a></div></td>' +
                '<td>'+data[i].value+'</td>' +
                '<td><a href="/token/'+data[i].tokenContractAddress+'">'+tokenId+' '+data[i].tokenName+'</a></td></tr>'
          }
          document.getElementById('erc-721-tbody').innerHTML = html
        }

        if (res.data.tokens != null && res.data.tokens.length) {
          const tokens = res.data.tokens
          let tokenArr = {};
          for (let i=0; i<tokens.length; i++) {
            if (tokenArr[tokens[i].tokenType] != null) {
              tokenArr[tokens[i].tokenType].push(tokens[i])
            } else {
              tokenArr[tokens[i].tokenType] = [
                tokens[i]
              ]
            }
          }
          let html = '<button data-bs-toggle="dropdown" type="button" class="btn dropdown-toggle w-100 justify-content-between">'+res.data.tokens.length+' tokens</button>' +
              '<div class="dropdown-menu" style="max-height: 10rem; overflow: auto;">'
          for (let type in tokenArr) {
            html += '<div class="d-flex p-2 fw-bold text-muted">'+type.toLocaleUpperCase()+'<span class="badge bg-primary ms-auto">'+tokenArr[type].length+'</span></div>'
            for (let i=0; i<tokenArr[type].length; i++) {
              let tokenId = ''
              if (tokenArr[type][i].tokenId !== '') {
                tokenId = '['+tokenArr[type][i].tokenId+']'
              }
              html += '<a class="dropdown-item" href="/token/'+tokenArr[type][i].tokenContract+'"><span class="me-3">'+tokenId+tokenArr[type][i].tokenName+'</span><span class="ms-auto"> quantity: '+tokenArr[type][i].tokenValue+'</span></a>'
            }
          }
          html += '</div>'
          document.getElementById("token-dropdown").innerHTML = html
        }

        if (res.data.logs != null && res.data.logs.length > 0) {
          this.display.log = true
          let log = res.data.logs
          let html = ''
          for (let i=0; i<log.length; i++) {
            html += '<div class="row"><div class="col-auto">' +
                '<span class="avatar avatar-rounded">'+log[i].logIndex+'</span></div>' +
                '<div class="col"><div class="d-flex mb-3"><div style="width: 6rem;">Address:</div>' +
                '<div><a href="/address/'+log[i].address+'">'+log[i].address+'</a></div></div>' +
                '<div class="d-flex mb-3"><div style="width: 6rem;">topics:</div>' +
                '<div><div class="mb-2"><span class="badge bg-secondary me-2" style="width: 1.45rem">0</span>'+log[i].topics0 + '</div>'

            if (log[i].topics1 !== '') {
              html += '<div class="mb-2"><span class="badge bg-secondary me-2" style="width: 1.45rem">1</span>'+log[i].topics1 +'</div>'
            }
            if (log[i].topics2 !== '') {
              html += '<div class="mb-2"><span class="badge bg-secondary me-2" style="width: 1.45rem">2</span>'+log[i].topics2 +'</div>'
            }
            if (log[i].topics3 !== '') {
              html += '<div class="mb-2"><span class="badge bg-secondary me-2" style="width: 1.45rem">3</span>'+log[i].topics3+'</div>'
            }

            html += '</div></div><div class="d-flex"><div style="width: 6rem;">data:</div>' +
                '<div class="form-control bg-azure-lt">'+log[i].data+'</div></div>' +
                '</div><div class="hr mt-4 mb-4"></div></div>'

            document.getElementById("pane-log").innerHTML = html
          }
        }

        if (res.data.daily.length > 0) {
          let data = res.data.daily
          let obj = {}
          for (let i=0; i< data.length; i++) {
            obj[data[i].date] = data[i].value
          }

          let arr = []
          for (let i=30; i>=0; i--) {
            let now = Moment().subtract(i, 'days').format('YYYY-MM-DD')
            let val = '0.00'
            if (obj[now]) {
              val = parseFloat(this.formatAmount(parseFloat(obj[now]))).toFixed(4)
            }
            arr.push({ value: val, date: now})
          }
          this.LineList = arr
        }
      }).catch((err) => {
        console.log(err)
        document.getElementById("transactions-tbody").innerHTML = '<tr><td colspan="9" class="text-center fw-bold">暂无数据</td></tr>'
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
    },
    renderLineChart() {
      if (!this.render) {
        let arr = this.LineList
        this.uList = arr.map(n => {
          return n.value
        })
        this.nameList = arr.map(n => {
          return n.date
        })

        let myChart = echarts.init(document.getElementById("charts"))
        let option = {
          grid: {
            top: 10,
            right: 25,
            bottom: 35,
            left: 50
          },
          xAxis: {
            type: 'category',
            data: this.nameList,
            axisLine: {
              show: false
            },
            axisTick: {
              show: false
            }
          },
          yAxis: {
            show: true
          },
          series: [
            {
              type: 'line',
              data: this.uList,
              color: '#326092',
              smooth: true
            }
          ],
          tooltip: {
            trigger: "axis",
            axisPointer: {
              type: "shadow"
            }
          }
        }
        myChart.setOption(option)

        window.addEventListener('resize', function () {
          myChart.resize()
        })
        this.render = true
      }
    }
  }
}
</script>

<style lang="scss" scoped>
  .nav-tabs .nav-link {
    padding: .6rem;
  }

  .nav-tabs .nav-link.active {
    background-color: unset;
    border-color: transparent;
    color: #4299E1;
    border-bottom: 2px solid #4299E1;
  }

  .nav-tabs .nav-link.active:focus, .nav-tabs .nav-link.active:hover {
    border-color: transparent;
    color: #4299E1;
    border-bottom: 2px solid #4299E1;
  }

  .nav-tabs .nav-link:focus, .nav-tabs .nav-link:hover {
    border-color: transparent;
    color: #4299E1;
  }

  .input-data {
    line-height: 1.4rem;
    overflow: auto;
    height: 16rem;
  }
</style>
