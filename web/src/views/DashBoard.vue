<template>
  <header class="navbar navbar-expand-md navbar-light d-print-none">
    <div class="container-xl">
      <h1 class="navbar-brand navbar-brand-autodark d-none-navbar-horizontal pe-0 pe-md-3">
        <a href="/" style="color: #EEEEEE !important;">
          <span class="navbar-brand-image">
            <svg width="110" height="32" viewBox="0 0 234 107" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M163.494 24.436V35.67H148.242V82H134.22V35.67H118.968V24.436H163.494ZM184.553 24.436V82H170.531V24.436H184.553ZM208.737 35.67V47.314H227.515V58.138H208.737V70.766H229.975V82H194.715V24.436H229.975V35.67H208.737Z" fill="white"/>
              <path fill-rule="evenodd" clip-rule="evenodd" d="M63.1904 4.95503C60.9727 1.1139 56.0611 -0.202164 52.2199 2.01551C48.3788 4.23319 47.0627 9.14482 49.2804 12.9859L60.7532 32.8573C62.9708 36.6984 67.8825 38.0145 71.7236 35.7968C75.5647 33.5792 76.8808 28.6675 74.6631 24.8264L63.1904 4.95503ZM6.39623 25.5346C6.39623 21.0993 9.9918 17.5037 14.4271 17.5037H37.3726C41.808 17.5037 45.4035 21.0993 45.4035 25.5346C45.4035 29.97 41.808 33.5656 37.3726 33.5656H14.4271C9.99179 33.5656 6.39623 29.97 6.39623 25.5346ZM66.0545 53.0692C66.0545 61.6231 59.1202 68.5574 50.5663 68.5574C42.0124 68.5574 35.0781 61.6231 35.0781 53.0692C35.0781 44.5153 42.0124 37.581 50.5663 37.581C59.1202 37.581 66.0545 44.5153 66.0545 53.0692ZM1.07728 65.9512C-1.1404 69.7923 0.175672 74.704 4.0168 76.9216C7.85793 79.1393 12.7696 77.8232 14.9872 73.9821L26.46 54.1107C28.6777 50.2696 27.3616 45.358 23.5205 43.1403C19.6793 40.9226 14.7677 42.2387 12.55 46.0798L1.07728 65.9512ZM46.7236 103.797C42.8825 106.015 37.9708 104.698 35.7532 100.857L24.2804 80.9859C22.0627 77.1448 23.3788 72.2332 27.2199 70.0155C31.0611 67.7978 35.9727 69.1139 38.1904 72.955L49.6631 92.8264C51.8808 96.6675 50.5647 101.579 46.7236 103.797ZM85.4445 88.9371C89.8798 88.9371 93.4754 85.3415 93.4754 80.9062C93.4754 76.4708 89.8798 72.8753 85.4445 72.8753H62.499C58.0636 72.8753 54.4681 76.4708 54.4681 80.9062C54.4681 85.3415 58.0636 88.9371 62.499 88.9371H85.4445ZM95.5204 29.1403C99.3615 31.358 100.678 36.2696 98.4599 40.1107L86.9872 59.9821C84.7695 63.8232 79.8579 65.1393 76.0168 62.9216C72.1756 60.704 70.8596 55.7923 73.0772 51.9512L84.55 32.0798C86.7677 28.2387 91.6793 26.9226 95.5204 29.1403Z" fill="url(#paint0_linear_859_183)"/>
              <defs>
                <linearGradient id="paint0_linear_859_183" x1="-2.37334" y1="-218.956" x2="185.471" y2="-182.165" gradientUnits="userSpaceOnUse">
                  <stop offset="0.293673" stop-color="#CE52FF"/>
                  <stop offset="0.702837" stop-color="#789EFF"/>
                  <stop offset="1" stop-color="white"/>
                </linearGradient>
              </defs>
            </svg>
          </span>
        </a>
      </h1>
      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbar-menu">
        <span class="navbar-toggler-icon"></span>
      </button>

      <div class="collapse navbar-collapse justify-content-end" id="navbar-menu">
        <menu-view />
      </div>
    </div>
  </header>

  <div class="scan-info">
    <div class="container-xl" style="padding-top: 3.4rem">
      <div class="row justify-content-between align-items-center">
        <div class="col-md-12 col-lg-7">
          <h1 class="h2 mb-3" style="text-align: left; color: #EEEEEE;">The Tie Blockchain Explorer</h1>
          <div style="height: 2.8rem">
            <search-view />
          </div>
        </div>
      </div>
    </div>
  </div>

  <div class="page-wrapper">
    <div class="container-xl">
      <div class="col-12" style="margin-top: -3rem">
          <div class="card card-info mb-4">
            <div class="card-body">
              <div class="card-body">
                <div class="row">
                  <div class="col-md-6 col-lg-4">
                    <div class="subheader">Average block time</div>
                    <div class="h2 pb-4">{{ this.avg }} <span class="text-light">seconds</span></div>

                    <div class="subheader">Total transactions</div>
                    <div class="h2">{{ this.txns }}</div>
                  </div>
                  <div class="col-md-6 col-lg-4">
                    <div class="subheader">Total blocks</div>
                    <div class="h2 pb-4">{{ this.blocks }}</div>
                    <div class="subheader">Addresses</div>
                    <div class="h2 pt-1">{{ this.addresses }}</div>
                  </div>
                  <div class="col-lg-4 col-md-12">
                    <h4 class="subheader" style="font-size: .8rem;">TRANSACTION HISTORY IN 14 DAYS</h4>
                    <div class="radar" id="charts"></div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

      <div class="row">
        <div class="col-md-12 col-lg-6">
          <div class="card card-gradient">
            <div class="card-header">Latest blocks</div>
            <div class="card-body p-0 scroll-card">
              <div class="table-responsive">
                <table class="table table-vcenter table-mobile-md table-nowrap card-table">
                  <tbody class="tbody-height-double" id="blocks-tbody"><tr><td><div class="loading"><span class="spinner-border" role="status"></span></div></td></tr></tbody>
                </table>
              </div>
            </div>
            <div class="card-footer text-center p-2">
              <a class="btn btn-link" style="color: #EEEEEE" href="/blocks">view all blocks</a>
            </div>
          </div>
        </div>
        <div class="col-md-12 col-lg-6">
          <div class="card card-gradient">
            <div class="card-header">Latest Transactions</div>
            <div class="card-body p-0 scroll-card">
              <div class="table-responsive">
                <table class="table table-vcenter table-mobile-md card-table">
                  <tbody class="tbody-height-double" id="transactions-tbody"><tr><td><div class="loading"><span class="spinner-border" role="status"></span></div></td></tr></tbody>
                </table>
              </div>
            </div>
            <div class="card-footer text-center p-2">
              <a href="/transactions" style="color: #EEEEEE" class="btn btn-link">view all transactions</a>
            </div>
          </div>
        </div>
      </div>
    </div>

    <footer-view/>
  </div>
</template>

<script>
import FooterView from "@/components/Footer";
import SearchView from "@/components/Search";
import MenuView from "@/components/Menu";
import {GetBlockList} from "@/api/block";
import {GetTransactionList} from "@/api/transaction";
import {GetDashboard} from "@/api/dashboard";
import {ToThousands} from "@/utils/common";
import Web3 from 'web3';
import Moment from "moment";
import * as echarts from 'echarts'

export default {
  name: "DashBoard",
  components: {
    MenuView,
    SearchView,
    FooterView,
  },
  data() {
    return {
      avg: 0,
      txns: 0,
      blocks: 0,
      addresses: 0,
      nameList: [],
      uList: []
    }
  },
  mounted() {
    this.getDashboard()
    this.getBlocks(1, 10)
    this.getTransactions(1, 0, 10)
  },
  methods: {
    getDashboard(){
      GetDashboard().then(res => {
        if (res.code === 0) {
          this.avg = res.data.avg
          this.txns = res.data.txns
          this.blocks = res.data.blocks
          this.addresses = res.data.addresses
          let arr = res.data.txnCountArr
          this.uList = arr.map(n => {
            return n.count
          })
          this.nameList = arr.map(n => {
            return n.date
          })
          this.radarBtn()
        }
      }).catch(err => {
        console.log(err)
      })
    },
    getBlocks(page, pageSize) {
      GetBlockList(page, pageSize).then(res => {
        if (res.code === 0) {
          if (res.data.blocks.length > 0) {
            let html = ''
            const data = res.data.blocks
            for (let i=0; i<data.length; i++) {
              html += '<tr><td>' +
                  '<div class="d-flex py-1 align-items-center">' +
                  '<span class="avatar avatar-sm p-2">Bk</span>' +
                  '<div class="flex-fill ms-2">' +
                  '<div><a href="/block/'+data[i].number+'">'+data[i].number+'</a></div>' +
                  '<div>'+this.formatTimestamp(data[i].timestamp)+'</div>' +
                  '</div></div></td>' +
                  '<td>' +
                  '<div class="text-truncate text-md" style="max-width: 18rem"><span class="me-2 text-muted">Miner</span><a class="" href="/address/'+data[i].miner+'">'+data[i].miner+'</a></div>' +
                  '<div class="text-muted"><a href="/transactions?block='+data[i].number+'">'+data[i].txs+' txns </a></div>' +
                  '</td>' +
                  '<td>' +
                  '<div><span class="badge p-2"><b>'+ToThousands(data[i].gasUsed)+'</b> Gas Used</span></div>' +
                  '</td></tr>'
            }
            document.getElementById("blocks-tbody").innerHTML = html
            return
          }
        }
        document.getElementById("blocks-tbody").innerHTML = '<tr><td class="empty-content">暂无数据</td></tr>'
      }).catch(() => {
        document.getElementById("blocks-tbody").innerHTML = '<tr><td class="empty-content">暂无数据</td></tr>'
      })
    },
    getTransactions(page, block, pageSize){
      GetTransactionList(page, block, pageSize).then(res => {
        if (res.code === 0) {
          if (res.data.transactions.length > 0) {
            let html = '';
            const data = res.data.transactions;
            for (let i=0; i<data.length; i++) {
            html += '<tr><td>' +
                '<div class="d-flex py-1 align-items-center">' +
                '<span class="avatar avatar-sm p-2">Tx</span>' +
                '<div class="flex-fill ms-2 text-truncate text-md" style="max-width: 10rem">' +
                '<div class="text-truncate"><a href="/transaction/'+data[i].transactionHash+'">'+data[i].transactionHash+'</a></div>' +
                '<div class="text-muted">'+this.formatTimestamp(data[i].timestamp)+'</div>' +
                '</div></div></td>' +
                '<td><div class="text-truncate text-md" style="max-width: 18rem"><span class="me-2 text-muted">From</span><a class="" href="/address/'+data[i].from+'">'+data[i].from+'</a></div>' +
                '<div class="text-truncate text-md" style="max-width: 18rem"><span class="me-2 text-muted">To</span><a class="" href="/address/'+data[i].to+'">'+data[i].to+'</a></div></td>' +
                '<td><span class="badge p-2"><b>'+Web3.utils.fromWei(data[i].value)+'</b> Tie</span></td></tr>'
            }
            document.getElementById("transactions-tbody").innerHTML = html
            return
          }
        }
        document.getElementById("transactions-tbody").innerHTML = '<tr><td><div class="empty-content">暂无数据</div></td></tr>'
      }).catch(() => {
        document.getElementById("transactions-tbody").innerHTML = '<tr><td class="empty-content">暂无数据</td></tr>'
      })
    },
    formatTimestamp(timestamp) {
      if (isNaN(timestamp)) {
        return ''
      }
      return Moment.unix(timestamp).fromNow()
    },
    radarBtn() {
      let myChart = echarts.init(document.getElementById("charts"))
      let option = {
        grid: {
          top: 5,
          right: 0,
          bottom: 25,
          left: 10
        },
        xAxis: {
          type: 'category',
          data: this.nameList,
          offset: 6,
          axisLine: {
            show: false,
            lineStyle: {
              color:'#CCC'
            }
          },
          axisTick: {
            show: false
          }
        },
        yAxis: {
          show:false
        },
        series: [
          {
            type: 'line',
            data: this.uList,
            color: '#6E6CC3',
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
    }
  }
}
</script>

<style lang="scss">
.loading {
  height: 26.3rem;
  line-height: 26.3rem;
  text-align: center;
}

.empty-content {
  height: 26.3rem;
  line-height: 26.3rem;
  text-align: center;
  font-weight: bold;
  color: #EEEEEE !important;
}

.card-info {
  background: #30333F !important;

  .h2 {
    color: #2CC6A0;
  }

  .subheader {
    color: #EEEEEE;
  }
}

.card-header {
  color: #EEEEEE !important;
}

.scan-info {
  display: block;
  height: 16rem;
}

.scroll-card {
  height: 27.3rem;
  overflow: auto;
}

.radar{
  width: 100%;
  height: 135px;
  margin: 0 auto;
}

.card-gradient {
  border: 1px solid transparent !important;
  background: linear-gradient(#29292C,#29292C) padding-box,
  repeating-linear-gradient(-45deg,#30333F 0, #30333F 0.3em,#29292C 0,#29292C 0.6em) !important;
}

</style>