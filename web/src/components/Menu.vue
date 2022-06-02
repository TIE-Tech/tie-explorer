<template>
  <ul class="navbar-nav">
    <li class="nav-item">
      <a class="nav-link" href="/">
              <span class="nav-link-icon d-md-none d-lg-inline-block">
                <svg width="18" height="16" viewBox="0 0 18 16" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <path d="M9 1C4.58172 1 1 4.58172 1 9C1 11.092 1.68538 13.1745 3 14.6H15C16.3146 13.1745 17 11.092 17 9C17 4.58172 13.4183 1 9 1Z" stroke="#CCCCCC" stroke-linecap="round" stroke-linejoin="round"/>
                  <path d="M8.9999 11.8C9.88356 11.8 10.5999 11.0836 10.5999 10.2C10.5999 9.31632 9.88356 8.59998 8.9999 8.59998C8.11625 8.59998 7.3999 9.31632 7.3999 10.2C7.3999 11.0836 8.11625 11.8 8.9999 11.8Z" stroke="#CCCCCC" stroke-linecap="round" stroke-linejoin="round"/>
                  <path d="M9 6.20001V8.60001" stroke="#CCCCCC" stroke-linecap="round" stroke-linejoin="round"/>
                  <path d="M9 3V3.8" stroke="#CCCCCC" stroke-linecap="round" stroke-linejoin="round"/>
                  <path d="M3 9.40002H3.8" stroke="#CCCCCC" stroke-linecap="round" stroke-linejoin="round"/>
                  <path d="M4.6001 5.40002L5.16578 5.9657" stroke="#CCCCCC" stroke-linecap="round" stroke-linejoin="round"/>
                  <path d="M14.2 9.40002H15" stroke="#CCCCCC" stroke-linecap="round" stroke-linejoin="round"/>
                  <path d="M13 5.9657L13.5657 5.40002" stroke="#CCCCCC" stroke-linecap="round" stroke-linejoin="round"/>
                </svg>
              </span>
        <span class="nav-link-title">
                DashBoard
              </span>
      </a>
    </li>
    <li class="nav-item">
      <a class="nav-link" href="/blocks">
              <span class="nav-link-icon d-md-none d-lg-inline-block">
                <svg width="16" height="18" viewBox="0 0 16 18" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <path d="M4.46408 3L1 5V9V13L4.46408 15L7.9282 17L11.3923 15L14.8564 13V9V5L11.3923 3L7.9282 1L4.46408 3Z" stroke="#CCCCCC" stroke-linecap="round" stroke-linejoin="round"/>
                  <path d="M7.92816 8.99999L3.77124 11.4M7.92816 8.99999V3.79999V8.99999ZM7.92816 8.99999L12.0851 11.4L7.92816 8.99999Z" stroke="#CCCCCC" stroke-linecap="round" stroke-linejoin="round"/>
                  <path d="M3.4248 7.7V6.4L4.55064 5.75M11.3056 5.75L12.4315 6.4V7.7L11.3056 5.75ZM9.05396 13.55L7.92812 14.2L6.80228 13.55H9.05396Z" stroke="#CCCCCC" stroke-linecap="round" stroke-linejoin="round"/>
                </svg>
              </span>
        <span class="nav-link-title">
                Blocks
              </span>
      </a>
    </li>
    <li class="nav-item">
      <a class="nav-link" href="/transactions">
              <span class="nav-link-icon d-md-none d-lg-inline-block">
                <svg width="18" height="14" viewBox="0 0 18 14" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <path d="M9 4.6H11V1L17 5.5L11 10V7H6.6V3.7L1 8.2L6.6 13V9.4H8.6" stroke="#CCCCCC" stroke-linecap="round" stroke-linejoin="round"/>
                </svg>
              </span>
        <span class="nav-link-title">
                Transactions
              </span>
      </a>
    </li>
    <li class="nav-item dropdown">
      <a class="nav-link dropdown-toggle" data-bs-toggle="dropdown" data-bs-auto-close="outside" role="button" aria-expanded="false">
              <span class="nav-link-icon d-md-none d-lg-inline-block">
                <svg width="18" height="18" viewBox="0 0 18 18" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <path d="M9 17C13.4183 17 17 13.4183 17 9C17 4.58172 13.4183 1 9 1C4.58172 1 1 4.58172 1 9C1 13.4183 4.58172 17 9 17Z" stroke="#CCCCCC" stroke-linecap="round" stroke-linejoin="round"/>
                  <path d="M9.40005 5.80001V4.20001M7.40005 5.80001V9.00001H10.2C11.0837 9.00001 11.8 8.28365 11.8 7.40001C11.8 6.51637 11.0837 5.80001 10.2 5.80001H9.40005H7.40005ZM7.40005 5.80001H5.80005H7.40005ZM7.40005 5.80001V4.20001V5.80001ZM7.40005 5.80001H9.40005H7.40005Z" stroke="#CCCCCC" stroke-linecap="round" stroke-linejoin="round"/>
                  <path d="M9.40005 12.2V13.8M7.40005 12.2V9H11C11.8837 9 12.6 9.71636 12.6 10.6C12.6 11.4836 11.8837 12.2 11 12.2H9.40005H7.40005ZM7.40005 12.2V13.8V12.2ZM7.40005 12.2H5.80005H9.40005H7.40005ZM7.40005 12.2H9.40005H7.40005Z" stroke="#CCCCCC" stroke-linecap="round" stroke-linejoin="round"/>
                </svg>
              </span>
        <span class="nav-link-title">
                Tokens
              </span>
      </a>
      <div class="dropdown-menu" style="background-color: #29292C;">
        <a class="dropdown-item" href="/tokens/erc-20">
          ERC-20
        </a>
        <a class="dropdown-item" href="/tokens/erc-721">
          ERC-721
        </a>
        <a class="dropdown-item" href="/tokens/tie">
          TIE
        </a>
      </div>
    </li>
    <li class="nav-item dropdown">
      <a class="nav-link dropdown-toggle" data-bs-toggle="dropdown" data-bs-auto-close="outside" role="button" aria-expanded="false">
              <span class="nav-link-icon d-md-none d-lg-inline-block">
                <svg width="18" height="18" viewBox="0 0 18 18" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <path d="M9.03572 0.999988L6.21729 3.81842L9.03565 6.63678L11.8541 3.81835L9.03572 0.999988Z" stroke="#CCCCCC" stroke-linecap="round" stroke-linejoin="round"/>
                  <path d="M14.2174 6.18157L11.3989 9L14.2173 11.8184L17.0357 8.99993L14.2174 6.18157Z" stroke="#CCCCCC" stroke-linecap="round" stroke-linejoin="round"/>
                  <path d="M3.85408 6.18157L1.03564 9L3.85401 11.8184L6.67244 8.99993L3.85408 6.18157Z" stroke="#CCCCCC" stroke-linecap="round" stroke-linejoin="round"/>
                  <path d="M9.03572 11.3632L6.21729 14.1816L9.03565 17L11.8541 14.1816L9.03572 11.3632Z" stroke="#CCCCCC" stroke-linecap="round" stroke-linejoin="round"/>
                </svg>
              </span>
        <span class="nav-link-title">
                Tools
              </span>
      </a>
      <div class="dropdown-menu" style="background-color: #29292C;">
        <a class="dropdown-item" href="/contract_verify">
          Verify Contract
        </a>
        <a class="dropdown-item" target="_blank" href="http://bridge.testnet.test.tie.tech">
          Tie Bridge
        </a>
      </div>
    </li>
  </ul>
  <div class="d-none d-md-flex">
    <span class="nav-link" style="color: #CCCCCC !important;"> | </span>
  </div>
<!--  <div class="d-none d-md-flex  ms-2">-->
<!--    <a href="?theme=dark" class="nav-link px-0 hide-theme-dark" title="" data-bs-toggle="tooltip" data-bs-placement="bottom" data-bs-original-title="Enable dark mode">-->
<!--      <svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">-->
<!--        <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>-->
<!--        <path d="M12 3c.132 0 .263 0 .393 0a7.5 7.5 0 0 0 7.92 12.446a9 9 0 1 1 -8.313 -12.454z"></path>-->
<!--      </svg>-->
<!--    </a>-->
<!--    <a href="?theme=light" class="nav-link px-0 hide-theme-light" title="" data-bs-toggle="tooltip" data-bs-placement="bottom" data-bs-original-title="Enable light mode">-->
<!--      <svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">-->
<!--        <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>-->
<!--        <circle cx="12" cy="12" r="4"></circle>-->
<!--        <path d="M3 12h1m8 -9v1m8 8h1m-9 8v1m-6.4 -15.4l.7 .7m12.1 -.7l-.7 .7m0 11.4l.7 .7m-12.1 -.7l-.7 .7"></path>-->
<!--      </svg>-->
<!--    </a>-->
<!--  </div>-->
  <ul class="navbar-nav">
    <li class="nav-item dropdown">
      <a href="javascript:" class="nav-link" data-bs-toggle="dropdown">
        <div class="nav-link-title">
          TestNet
        </div>
      </a>
      <div class="dropdown-menu" style="background-color: #29292C;">
        <a href="javascript:" class="dropdown-item">MainNet</a>
        <a href="javascript:" class="dropdown-item">TestNet</a>
      </div>
    </li>
    <li class="nav-item dropdown">
      <a href="javascript:" class="nav-link dropdown-toggle" data-bs-toggle="dropdown">
        <div class="nav-link-title">
          En
        </div>
      </a>
      <div class="dropdown-menu" style="background-color: #29292C;">
        <a href="javascript:" class="dropdown-item">中文</a>
        <a href="javascript:" class="dropdown-item">English</a>
      </div>
    </li>
  </ul>
</template>

<script>
export default {
  name: "MenuView"
}
</script>

<style scoped>

</style>