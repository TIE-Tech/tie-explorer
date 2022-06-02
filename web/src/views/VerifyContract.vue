<template>
  <header-view />

  <div class="page-wrapper">
    <div class="page-body">
      <div class="container-xl">
        <div class="text-center mb-3 mt-2">
          <div class="w-md-75 w-lg-75 mx-md-auto ">
            <h1 class="h2">Verify &amp; Publish Contract Source Code </h1>
            <small class="text-muted text-muted fw-bold">COMPILER TYPE AND VERSION SELECTION</small>
          </div>
        </div>

        <div class="hr"></div>

        <div class="row">
          <div class="col-md-10 offset-md-1 text-secondary mb-4">
            Source code verification provides <b>transparency</b> for users interacting with smart contracts.
            By uploading the source code, Etherscan will match the compiled code with that on the blockchain.
            Just like contracts, a "smart contract" should provide end users with more information on what they are "digitally signing" for and give users an opportunity to audit the code to independently verify that it actually does what it is supposed to do.
          </div>
        </div>

        <div class="w-md-75 w-lg-50 mx-md-auto mt-2">
          <div class="card-body">
              <form action="/">
                <div class="form-group mb-3">
                  <label class="form-label">Please enter the Contract Address you would like to verify</label>
                  <div>
                    <input type="text" class="form-control" style="background: transparent" id="address-input" v-model="address" name="address" aria-describedby="emailHelp" placeholder="0x...">
                    <div class="invalid-feedback">This field is required.</div>
                  </div>
                </div>
                <div class="form-group mb-3 ">
                  <label class="form-label">Please select Compiler Type</label>
                  <div class="form-selectgroup form-selectgroup-boxes d-flex flex-column">
                    <label class="form-selectgroup-item flex-fill">
                      <input type="radio" v-model="verify" name="verify-mode" value="source-code" class="form-selectgroup-input" checked="">
                      <div class="form-selectgroup-label d-flex align-items-center p-3" style="background: transparent;">
                        <div class="me-3">
                          <span class="form-selectgroup-check"></span>
                        </div>
                        <div>
                          Via flattened source code
                        </div>
                      </div>
                    </label>
                    <label class="form-selectgroup-item flex-fill">
                      <input type="radio" v-model="verify" name="verify-mode" value="input-json" class="form-selectgroup-input">
                      <div class="form-selectgroup-label d-flex align-items-center p-3" style="background: transparent;">
                        <div class="me-3">
                          <span class="form-selectgroup-check"></span>
                        </div>
                        <div>
                          Via Standard Input JSON
                        </div>
                      </div>
                    </label>
                    <label class="form-selectgroup-item flex-fill">
                      <input type="radio" v-model="verify" name="verify-mode" value="vyper-contract" class="form-selectgroup-input">
                      <div class="form-selectgroup-label d-flex align-items-center p-3" style="background: transparent;">
                        <div class="me-3">
                          <span class="form-selectgroup-check"></span>
                        </div>
                        <div>
                          Vyper contract
                        </div>
                      </div>
                    </label>
                  </div>
                </div>
                <div class="form-footer d-flex justify-content-center">
                  <button type="button" class="btn btn-primary" @click="continueVerify">Continue</button>
                  <button type="button" class="btn btn-secondary ms-3" @click="resetForm">Reset</button>
                </div>
              </form>
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
export default {
  name: "VerifyContract",
  components: {FooterView, HeaderView},
  data() {
    return {
      address: this.$route.query.a,
      verify: "source-code",
    }
  },
  methods: {
    continueVerify() {
      if (!this.address) {
        document.getElementById("address-input").classList.add("is-invalid")
        return
      }
      document.getElementById("address-input").classList.remove("is-invalid")
      location.replace("/contract_verify/"+this.verify+"?a="+this.address)
    },
    resetForm() {
      let address = ""
      if (this.$route.query.a) {
        address = this.$route.query.a
      }
      this.address = address
      this.verify = "source-code"
      document.getElementById("address-input").classList.remove("is-invalid")
    }
  }
}
</script>

<style scoped>

@media (min-width: 768px) {
  .w-md-75 {
    width: 75%!important;
  }

  .mx-md-auto {
    margin-left: auto!important;
    margin-right: auto!important;
  }
}

@media (min-width: 992px) {
  .w-lg-50 {
    width: 50%!important;
  }
}
</style>