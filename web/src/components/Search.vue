<template>
  <form class="h-100" action="/search" method="GET">
    <div style="float: left; background:#30333F">
      <div class="input-group h-100" style="flex-wrap: nowrap">
        <div>
          <select v-model="selected" class="form-select h-100" style="background-color: transparent;border: none;color: #ccc;min-width: 7.2rem;">
            <option value="">All filters</option>
            <option value="address">Addresses</option>
            <option value="token">Tokens</option>
          </select>
        </div>
        <span style="font-size: 1.4rem;line-height: 2.6rem;">|</span>
        <div class="position-relative">
          <input type="text" style="background-color: transparent;border: none;margin-right: -0.6rem;color: #ccc;" v-model="inputVal" class="form-control h-100 w-md-30" @input="inputSearch" @focus="inputSearch" @blur="inputBlur" placeholder="Search by Address / Txn Hash / Block / Token" autocomplete="off">
          <div class="dropdown-menu" id="search-result" style="position: absolute;transform: translate3d(0px, 0px, 0px);width: 100%;background: rgb(48, 51, 63);color: #ccc;" data-popper-placement="bottom-start"></div>
        </div>
        <span style="font-size: 1.4rem;line-height: 2.6rem;">|</span>
        <a href="javascript:" @click="inputSearch" class="btn btn-icon" style="border: none; background: transparent;min-width: 3rem;" aria-label="Button">
          <svg width="17" height="17" viewBox="0 0 17 17" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M7.76265 14.5253C11.4975 14.5253 14.5253 11.4975 14.5253 7.76265C14.5253 4.02776 11.4975 1 7.76265 1C4.02776 1 1 4.02776 1 7.76265C1 11.4975 4.02776 14.5253 7.76265 14.5253Z" stroke="#6E6CC3" stroke-linejoin="round"/>
            <path d="M10.0128 5.11453C9.43693 4.53863 8.64133 4.18243 7.76254 4.18243C6.88375 4.18243 6.08815 4.53863 5.51221 5.11453" stroke="#6E6CC3" stroke-linecap="round" stroke-linejoin="round"/>
            <path d="M12.6245 12.6245L16 16" stroke="#6E6CC3" stroke-linecap="round" stroke-linejoin="round"/>
          </svg>
        </a>
      </div>
    </div>
  </form>
</template>

<script>
import {GetSearch} from "@/api/dashboard";

let t = null
export default {
  name: "SearchView",
  data() {
    return {
      selected: "",
      inputVal: "",
    }
  },
  methods: {
    getSearch(s, t) {
      GetSearch(s, t).then(res => {
        if (res.code === 0) {
          let html = ''
          if (res.data.exists) {
            html = '<div class="row"><div class="col-md-12"><div class="mb-1 p-1">Found <span class="text-danger fw-bold">1</span> matching results</div></div></div>' +
                '<a href="/'+res.data.type+'/'+this.inputVal+'" class="dropdown-item"><div class="text-truncate text-warning">'+ this.inputVal +'</div><div class="ms-auto">'+res.data.type.toUpperCase()+'</div></a>'
          } else {
            html = '<div class="row"><div class="col-md-12"><div class="mb-1 p-1">Found <span class="text-danger fw-bold">0</span> matching results</div></div></div>'
          }
          document.getElementById('search-result').classList.add("show")
          document.getElementById('search-result').innerHTML = html
        }
      })
    },
    inputSearch() {
      if (t !== null) {
        clearTimeout(t)
      }

      if (this.inputVal === "") {
        document.getElementById('search-result').classList.remove("show")
        return
      }
      t = setTimeout(() => {
        this.getSearch(this.inputVal, this.selected)
      }, 1000)
    },
    inputBlur() {
      setTimeout(function () {
        document.getElementById('search-result').classList.remove("show")
      }, 100)
    }
  }
}
</script>

<style scoped>
  .form-select {
    border-top-right-radius: unset;
    border-bottom-right-radius: unset;
  }

  @media (min-width: 768px) {
    .w-md-30 {
      width: 30rem;
    }
  }
</style>