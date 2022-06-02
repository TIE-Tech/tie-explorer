<template>
  <header-view />

  <div class="page-wrapper">
    <div class="page-body">
      <div class="container-xl">
        <div class="page-header d-print-none m-0 p-0 pb-3">
          <div class="row align-items-center">
            <div class="col">
              <h2 class="page-title">
                Transaction Details
              </h2>
            </div>
          </div>
        </div>

        <div class="card">
          <ul class="nav nav-tabs" data-bs-toggle="tabs">
            <li class="nav-item">
              <a href="#overview" class="nav-link active" data-bs-toggle="tab">Overview</a>
            </li>
            <li class="nav-item">
              <a href="#logs" id="logs-tab" class="nav-link" data-bs-toggle="tab">Logs</a>
            </li>
          </ul>
          <div class="card-body">
            <div class="tab-content">
              <div class="tab-pane active" id="overview">
                <div class="row mb-4">
                  <div class="col-md-3">Transaction Hash:</div>
                  <div class="col-md-9">
                    <span class="fw-bold">{{ this.txInfo.transactionHash }}</span>
                  </div>
                </div>
                <div class="row mb-4">
                  <div class="col-md-3">Status:</div>
                  <div class="col-md-9">
                    <div v-if="this.txInfo.status === 1" class="badge bg-success-lt" style="line-height: 1rem;">
                      <svg xmlns="http://www.w3.org/2000/svg" class="icon text-green" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"></path><circle cx="12" cy="12" r="9"></circle><path d="M9 12l2 2l4 -4"></path></svg>
                      <span class="ms-1">Success</span>
                    </div>
                    <span v-else class="badge bg-danger">Failed</span>
                  </div>
                </div>
                <div class="row mb-4">
                  <div class="col-md-3">Block:</div>
                  <div class="col-md-9">
                    <a :href="'/block/' + this.txInfo.blockNumber" style="line-height: 1.8rem">
                      {{ this.txInfo.blockNumber }}
                    </a>
                    <span class="badge badge-arrow-in">{{ this.formatNumber(this.txInfo.blockConfirms) }} Block Confirms</span>
                  </div>
                </div>
                <div class="row mb-4">
                  <div class="col-md-3">Timestamp:</div>
                  <div class="col-md-9">
                    <span>{{ this.txInfo.dateTime }} ({{ this.formatTimestamp(this.txInfo.timestamp) }})</span>
                  </div>
                </div>
                <div class="hr mt-3 mb-4"></div>
                <div class="row mb-4">
                  <div class="col-md-3">From:</div>
                  <div class="col-md-9">
                    <span><a :href="'/address/'+this.txInfo.from">{{ this.txInfo.from }}</a></span>
                  </div>
                </div>
                <div class="row mb-4">
                  <div class="col-md-3">To:</div>
                  <div class="col-md-9">
                    <div v-html="this.txInfo.toStr"></div>
                  </div>
                </div>
                <div class="hr mt-3 mb-4"></div>
                <div id="token-transfer"></div>
                <div class="row mb-4">
                  <div class="col-md-3">Value:</div>
                  <div class="col-md-9">
                    <span>{{ this.formatAmount(this.txInfo.value) }} Tie</span>
                  </div>
                </div>
                <div class="row mb-4">
                  <div class="col-md-3">Transaction Fee:</div>
                  <div class="col-md-9">
                    <span>{{ this.formatAmount(this.txInfo.gas * this.txInfo.gasPrice) }} Tie</span>
                  </div>
                </div>
                <div class="hr mt-3 mb-4"></div>
                <div class="row mb-4">
                  <div class="col-md-3">Gas Price:</div>
                  <div class="col-md-9">
                    <span>{{ this.formatAmount(this.txInfo.gasPrice, 'gwei') }} Gwei</span>
                  </div>
                </div>
                <div class="row mb-4">
                  <div class="col-md-3">Gas Limit & Usage by Txn:</div>
                  <div class="col-md-9">
                    <span v-html="this.getGasLimitInfo(this.txInfo.gasUsed, this.txInfo.cumulativeGasUsed)"></span>
                  </div>
                </div>
                <div class="row mb-4">
                  <div class="col-md-3">Nonce:</div>
                  <div class="col-md-9">
                    <span>{{ this.formatNumber(this.txInfo.nonce) }}</span>
                  </div>
                </div>
                <div class="hr mt-3 mb-4"></div>
                <div class="row mb-4">
                  <div class="col-md-3">Input Data:</div>
                  <div class="col-md-9">
                    <span class="form-control input-data" style="background: transparent">{{ this.txInfo.input }}</span>
                  </div>
                </div>
              </div>
              <div class="tab-pane" id="logs"></div>
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
import {GetTransactionInfo} from "@/api/transaction";
import {ToThousands} from "@/utils/common";
import Web3 from 'web3'
import Moment from "moment";

export default {
  name: "TransactionInfo",
  components: {FooterView, HeaderView},
  data() {
    return {
      txInfo: {}
    }
  },
  mounted() {
    if (this.$route.params.hash) {
      this.getTransaction(this.$route.params.hash)
    }
  },
  methods: {
    getTransaction(hash) {
      GetTransactionInfo(hash).then(res => {
        if (res.code === 0) {
          this.txInfo = res.data.transaction
          if (res.data.tokenTransfer.length > 0) {
            let transfer = res.data.tokenTransfer
            let html = '<div class="row"><div class="col-md-3">Tokens Transferred:</div><div class="col-md-9"><div class="mb-4">'
            for (let i=0; i<transfer.length; i++) {
              let tokenId = ''
              if (transfer[i].tokenId !== "") {
                tokenId = 'TokenId [<a href="/token/'+transfer[i].tokenContractAddress+'">' + transfer[i].tokenId + '</a>]'
              }
              html += '<div class="d-flex"><div style="width: 5rem;">From</div><div class="text-truncate" style="max-width: 20rem;"><a href="/address/'+transfer[i].from+'">'+transfer[i].from+'</a></div></div>' +
                  '<div class="d-flex"><div style="width: 5rem;">To</div><div class="text-truncate" style="max-width: 20rem;"><a href="/address/'+transfer[i].to+'">'+transfer[i].to+'</a></div></div>' +
                  '<div class="d-flex"><div style="width: 5rem;">For</div><div class="text-truncate" style="max-width: 20rem;">'+tokenId+' <a href="/token/'+transfer[i].tokenContractAddress+'">'+transfer[i].tokenName+'</a></div></div>'
            }
            html += '</div></div></div><div class="hr mt-0"></div>'
            document.getElementById("token-transfer").innerHTML = html
          }

          if (res.data.log.length > 0) {
            document.getElementById("logs-tab").innerText = 'Logs ('+ res.data.log.length +')'
            let html = ''
            let log = res.data.log
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
            }
            document.getElementById("logs").innerHTML = html
          } else {
            document.getElementById("logs-tab").style = 'display:none'
          }
        }
      })
    },
    formatNumber(number) {
      return ToThousands(number)
    },
    formatAmount(number, unit='ether') {
      if (isNaN(number)) {
        return 0
      }
      return Web3.utils.fromWei(number.toString(), unit)
    },
    formatTimestamp(timestamp) {
      if (isNaN(timestamp)) {
        return ''
      }
      return Moment.unix(timestamp).fromNow()
    },
    getGasLimitInfo(gasUsed, gasCumulative) {
      let html = ''
      if (gasCumulative > 0) {
        html += '<span class="position-relative">' + this.formatNumber(gasUsed)
        html += '&nbsp;&nbsp;|&nbsp;&nbsp;' + this.formatNumber(gasCumulative)
        const r = parseFloat((gasUsed / gasCumulative * 100).toString()).toFixed(0)
        let color = 'bg-green'
        if (r >= 80) {
          color = 'bg-red'
        } else if (r >= 60) {
          color = 'bg-yellow'
        } else if (r >= 40) {
          color = 'bg-purple'
        } else if (r >= 20) {
          color = 'bg-blue'
        }
        html += '&nbsp;&nbsp;(' + r + '%)'
        html += '<div class="mt-2 progress progress-sm position-absolute" style="min-width: 8rem;">' +
            '<div class="progress-bar '+color+'" style="width: '+r+'%"></div></div></span>'
      }
      return html
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
    height: 7.8rem;
  }

  .badge-arrow-in {
    position: relative;
    padding: 0.5rem 0.5rem 0.5rem 1.4rem;
    letter-spacing: 0.8px;
    margin-left: 0.4rem;
    border: none;
  }
  .badge-arrow-in::after {
    position: absolute;
    left: 0;
    top: 0;
    content: "";
    border-top: 0.9rem solid transparent;
    border-bottom: 0.8rem solid transparent;
    border-left: 0.8rem solid #29292C;
  }
</style>
