<template>
  <div class="pt-4 mt-auto">
    <footer class="footer foot-info d-print-none">
      <div class="container-xl">
        <div class="row text-center align-items-center flex-row-reverse">
          <div class="col-lg-auto ms-lg-auto">
            <ul class="list-inline list-inline-dots mb-0">
              <li
                class="list-inline-item add-metamask-btn"
                @click="addMetamask"
              >
                <img src="/images/fox.png" class="icon" alt="" />
                <span class="link-secondary text-white"
                  >Add TIE Chain Testnet</span
                >
              </li>
              <li class="list-inline-item">
                <a href="javascript:" class="link-secondary text-white"
                  >Documentation</a
                >
              </li>
              <li class="list-inline-item">
                <a
                  href="javascript:"
                  class="link-secondary text-white"
                  rel="noopener"
                  >Source code</a
                >
              </li>
            </ul>
          </div>
          <div class="col-12 col-lg-auto mt-3 mt-lg-0">
            <ul class="list-inline list-inline-dots mb-0">
              <li class="list-inline-item">
                &copy; 2022 TieScan. All rights reserved
              </li>
            </ul>
          </div>
        </div>
      </div>
    </footer>
  </div>
</template>

<script>
  import Web3 from "web3";
  export default {
    name: "FooterView",
    methods: {
      async addMetamask() {
        if (window.ethereum) {
          try {
            await window.ethereum.request({
              method: "wallet_switchEthereumChain",
              params: [
                {
                  chainId: Web3.utils.numberToHex(188), // 目标链ID
                },
              ],
            });
          } catch (e) {
            if (e.code === 4902) {
              try {
                await window.ethereum.request({
                  method: "wallet_addEthereumChain",
                  params: [
                    {
                      chainId: Web3.utils.numberToHex(188), // 目标链ID
                      chainName: "TIE Chain Testnet",
                      nativeCurrency: {
                        name: "TIE",
                        symbol: "TIE",
                        decimals: 18,
                      },
                      rpcUrls: ["https://rpc.testnet.tie.tech"], // 节点
                      blockExplorerUrls: ["https://testnet.tiescan.tech"],
                    },
                  ],
                });
              } catch (ee) {
                //
              }
            }
          }
        }
      },
    },
  };
</script>

<style lang="scss" scoped>
  .foot-info {
    border-top: none;
    padding: 1.8rem 0;
    color: #cccccc;
    background: linear-gradient(
      90.11deg,
      #30333f 27.51%,
      rgba(32, 35, 44, 0.84) 105.24%
    );
  }
  .add-metamask-btn {
    line-height: 16px;
    background: rgba(238, 238, 238, 0.05);
    border-radius: 4px;
    padding: 13px 15px;
    user-select: none;
    cursor: pointer;
    &:hover {
      background: #eeeeee;
      .text-white {
        color: #111111 !important;
      }
    }
    .icon {
      width: 16px;
      height: 16px;
      margin-right: 5px;
    }
  }
  @media (max-width: 991.98px) {
    .add-metamask-btn {
      display: block;
      margin-left: auto;
      width: fit-content;
      margin-right: auto;
      margin-bottom: 10px;
      font-size: 15px;
      margin-inline-end: auto !important;
    }
  }
</style>
