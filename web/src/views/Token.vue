<template>
  <header-view />

  <div class="page-wrapper">
    <div class="page-body">
      <div class="container-xl">
        <div class="page-header d-print-none m-0 p-0 pb-3">
          <div class="row">
            <div class="col-md-1 fw-bold" style="font-size: 1.4rem">{{ this.tokenInfo.name }}</div>
          </div>
        </div>

        <div class="row">
          <div class="col-md-12">
            <div class="card">
              <div class="card-header">
                <h3 class="card-title">Overview [{{ this.upperCase(this.tokenInfo.type) }}]</h3>
              </div>
              <div class="card-body">
                <div class="row mb-3">
                  <div class="col-4">Contract: </div>
                  <div class="col-8">{{ this.tokenInfo.contractAddress }}</div>
                </div>
                <div class="row mb-3">
                  <div class="col-4">Holders: </div>
                  <div class="col-8">{{ this.holders }}</div>
                </div>
                <div class="row mb-3">
                  <div class="col-4">Transfers: </div>
                  <div class="col-8">{{ this.transfers }}</div>
                </div>
                <div class="row mb-3" v-if="this.tokenInfo.totalSupply > 0">
                  <div class="col-4">Total Supply: </div>
                  <div class="col-8">0</div>
                </div>
                <div class="row mb-3" v-if="this.tokenInfo.decimals > 0">
                  <div class="col-4">Decimals: </div>
                  <div class="col-8">0</div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="card mt-4">
          <ul class="nav nav-tabs" data-bs-toggle="tabs">
            <li class="nav-item">
              <a href="#pane-transfer" class="nav-link active" data-bs-toggle="tab">Transfers</a>
            </li>
            <li class="nav-item">
              <a href="#pane-holders" id="tab-balance-history" class="nav-link" data-bs-toggle="tab">Holders</a>
            </li>
          </ul>
          <div class="card-body p-0">
            <div class="tab-content">
              <div class="tab-pane active show" id="pane-transfer">
                <div class="table-responsive">
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
                  <tbody id="transfer-tbody">
                  <tr><td colspan="8" class="text-center"><span class="spinner-border" role="status"></span></td></tr>
                  </tbody>
                </table>
                </div>
              </div>
              <div class="tab-pane" id="pane-holders">
                <div class="table-responsive">
                  <table class="table table-vcenter card-table">
                  <thead>
                  <tr>
                    <th>Rank</th>
                    <th>Address</th>
                    <th>Quantity</th>
                  </tr>
                  </thead>
                  <tbody id="holder-tbody">
                  <tr><td colspan="8" class="text-center"><span class="spinner-border" role="status"></span></td></tr>
                  </tbody>
                </table>
                </div>
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
import {GetTokenInfo} from "@/api/token";
import Moment from "moment";
import Web3 from "web3";
export default {
  name: "TokenInfo",
  components: {FooterView, HeaderView},
  data() {
    return {
      tokenInfo: {},
      transfers: 0,
      holders: 0
    }
  },
  mounted() {
    if (this.$route.params.address) {
      this.getTokenInfo(this.$route.params.address)
    }
  },
  methods: {
    getTokenInfo(address){
      GetTokenInfo(address).then(res => {
        this.tokenInfo = res.data.info
        this.transfers = res.data.transfer.length
        this.holders = res.data.holder.length

        if (res.data.transfer.length > 0) {
          let data = res.data.transfer
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
          document.getElementById('transfer-tbody').innerHTML = html
        }

        if (res.data.holder.length > 0) {
          let data = res.data.holder
          let html = ''
          for (let i=0; i<data.length; i++) {
            html += '<tr>' +
                '<td>'+(i+1)+'</td>' +
                '<td><div>' +
                '<a href="/address/'+data[i].address+'">'+data[i].address+'</a>' +
                '</div></td>' +
                '<td>'+data[i].value+'</td></tr>'
          }
          document.getElementById('holder-tbody').innerHTML = html
        }
      }).catch(err => {
        console.log(err)
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
    upperCase(str) {
      if (str !== undefined) {
        return str.toLocaleUpperCase()
      }
      return ''
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
</style>
