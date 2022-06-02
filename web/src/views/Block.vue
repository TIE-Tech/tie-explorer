<template>
  <header-view/>

  <div class="page-wrapper">
    <div class="page-body">
      <div class="container-xl">
        <div class="page-header d-print-none m-0 p-0 pb-3">
          <div class="row align-items-center">
            <div class="col">
              <h2 class="page-title">
                Block <span class="text-muted ps-2 small">#{{ this.block.number }}</span>
              </h2>
            </div>
          </div>
        </div>

        <div class="card">
          <ul class="nav nav-tabs" data-bs-toggle="tabs">
            <li class="nav-item">
              <a href="#overview" class="nav-link active" data-bs-toggle="tab">Overview</a>
            </li>
          </ul>
          <div class="card-body">
            <div class="tab-content">
              <div class="tab-pane active show" id="overview">
                <div class="row">
                  <div class="col-md-3">Block Height:</div>
                  <div class="col-md-9">
                    <span class="fw-bold">{{ this.block.number }}</span>
                  </div>
                  <div class="hr mt-3 mb-3"></div>
                  <div class="col-md-3">Timestamp:</div>
                  <div class="col-md-9">
                    <span>{{ this.block.dateTime }} ({{ this.formatTimestamp(this.block.timestamp) }})</span>
                  </div>
                  <div class="hr mt-3 mb-3"></div>
                  <div class="col-md-3">Transactions:</div>
                  <div class="col-md-9">
                    <div>
                      <a v-if="block.txs > 0" :href="'/transactions?block='+this.block.number" class="badge bg-azure">{{ this.block.txs }} transactions</a>
                      <span v-else class="badge bg-azure">{{ this.block.txs }} transactions</span>
                      in this block</div>
                  </div>
                  <div class="hr mt-3 mb-3"></div>
                  <div class="col-md-3">Miner:</div>
                  <div class="col-md-9">
                    <span>{{ this.block.miner }}</span>
                  </div>
                  <div class="hr mt-3 mb-3"></div>
                  <div class="col-md-3">Uncles:</div>
                  <div class="col-md-9">
                    <span>{{ this.block.uncles }}</span>
                  </div>
                  <div class="hr mt-3 mb-3"></div>
                  <div class="col-md-3">Difficulty:</div>
                  <div class="col-md-9">
                    <span>{{ this.formatNumber(this.block.difficulty) }}</span>
                  </div>
                  <div class="hr mt-3 mb-3"></div>
                  <div class="col-md-3">Total Difficulty:</div>
                  <div class="col-md-9">
                    <span>{{ this.formatNumber(this.block.totalDifficulty) }}</span>
                  </div>
                  <div class="hr mt-3 mb-3"></div>
                  <div class="col-md-3">Size:</div>
                  <div class="col-md-9">
                    <span>{{ this.block.sizeDecode }}</span>
                  </div>
                  <div class="hr mt-3 mb-3"></div>
                  <div class="col-md-3">Block Reward:</div>
                  <div class="col-md-9">
                    <span>{{ this.formatAmount(this.block.reward) }} TIE</span>
                  </div>
                  <div class="hr mt-3 mb-3"></div>
                  <div class="col-md-3">Gas Used:</div>
                  <div class="col-md-9">
                    <div class="pb-2" v-html="this.getGasLimitInfo(this.block.gasUsed, this.block.gasLimit)"></div>
                  </div>
                  <div class="hr mt-3 mb-3"></div>
                  <div class="col-md-3">Gas Limit:</div>
                  <div class="col-md-9">
                    <span>{{ this.formatNumber(this.block.gasLimit) }}</span>
                  </div>
                  <div class="hr mt-3 mb-3"></div>
                  <div class="col-md-3">Extra Data:</div>
                  <div class="col-md-9">
                    <span class="form-control extra-data" style="background: transparent">{{ this.block.extraData }}</span>
                  </div>
                  <div class="hr mt-3 mb-3"></div>
                  <div class="col-md-3">Block Hash:</div>
                  <div class="col-md-9">
                    <span>{{ this.block.hash }}</span>
                  </div>
                  <div class="hr mt-3 mb-3"></div>
                  <div class="col-md-3">Parent Hash:</div>
                  <div class="col-md-9">
                    <span>{{ this.block.parentHash }}</span>
                  </div>
                  <div class="hr mt-3 mb-3"></div>
                  <div class="col-md-3">Sha3Uncles:</div>
                  <div class="col-md-9">
                    <span>{{ this.block.sha3Uncles }}</span>
                  </div>
                  <div class="hr mt-3 mb-3"></div>
                  <div class="col-md-3">StateRoot:</div>
                  <div class="col-md-9">
                    <span>{{ this.block.stateRoot }}</span>
                  </div>
                  <div class="hr mt-3 mb-3"></div>
                  <div class="col-md-3">Nonce:</div>
                  <div class="col-md-9">
                    <span>{{ this.block.nonce }}</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <footer-view/>
  </div>
</template>

<script>
import HeaderView from "@/components/Header";
import FooterView from "@/components/Footer";
import {GetBlockInfo} from "@/api/block";
import {ToThousands} from "@/utils/common";
import Moment from 'moment'
import Web3 from "web3";

export default {
  name: "BlockInfo",
  components: {FooterView, HeaderView},
  data() {
    return {block: {}}
  },
  mounted() {
    if (this.$route.params.number > 0) {
      this.getBlock(this.$route.params.number)
    }
  },
  methods: {
    getBlock(number) {
      GetBlockInfo(number).then(res => {
        if (res.code === 0) {
          this.block = res.data
        }
      })
    },
    formatTimestamp(timestamp) {
      if (isNaN(timestamp)) {
        return ''
      }
      return Moment.unix(timestamp).fromNow()
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
    getGasLimitInfo(gasUsed, gasCumulative) {
      let html = ''
      if (gasCumulative > 0) {
        html += '<span class="position-relative">' + this.formatNumber(gasUsed)
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
        html += ' (' + r + '%)'
        html += '<div class="mt-2 mb-2 progress progress-sm position-absolute" style="min-width: 8rem;">' +
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

.extra-data {
  line-height: 1.4rem;
  overflow: auto;
  height: 7.8rem;
}
</style>
