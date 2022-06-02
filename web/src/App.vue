<template>
  <router-view/>
</template>

<style lang="scss">
  body {
    overflow: auto !important;
    background-image: url("./assets/images/bg.png");
    background-repeat: no-repeat;
    -moz-background-size: 100% 100%;
    background-size: 100%;
  }

  .tbody-height td {
    height: 3rem;
    color: #CCCCCC !important;
  }

  .tbody-height-double td {
    height: 4.49rem;
  }

  .hash-tag {
    max-width: 10rem;
    display: inline-block;
    vertical-align: bottom;
  }

  .dropdown-item {
    color: #bbb !important;
  }

  @media (max-width: 767.98px){
    .text-md {
      max-width: none !important;
      width: 100%;
    }
    .text-md a.hash-tag {
      max-width: none !important;
      display: unset !important;
    }
    .tx-hash-tag {
      max-width: none !important;
      width: 18rem;
    }
  }
</style>
