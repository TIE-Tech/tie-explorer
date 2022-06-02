<template>
  <header-view />

  <div class="page-wrapper">
    <div class="page-body">
      <div class="container-xl">
        <div class="text-center">
          <h2>Verify & Publish Contract Source Code</h2>
        </div>

        <div class="card">
          <div class="card-header">
            <h4 class="card-title">Contract Source Code</h4>
          </div>
          <div class="card-body">
            <div class="row">
              <div class="col-md-6">
                <div class="mb-3">
                  <label class="form-label">Contract Address</label>
                  <input type="text" class="form-control" style="background: transparent" name="example-text-input" placeholder="Input placeholder">
                </div>
              </div>
              <div class="col-md-6">
                <div class="mb-3">
                  <label class="form-label">Contract Name</label>
                  <input type="text" class="form-control" style="background: transparent" name="example-text-input" placeholder="Contract Name">
                </div>
              </div>
              <div class="col-md-6">
                <div class="mb-3">
                  <label class="form-label">Compiler</label>
                  <select class="form-select" style="background: transparent">
                    <option value="no" selected>Latest</option>
                    <option value="yes">Yes</option>
                  </select>
                </div>
              </div>
              <div class="col-md-6">
                <div class="mb-3">
                  <label class="form-label">Include nightly builds</label>
                  <div>
                    <select class="form-select" style="background: transparent">
                      <option value="No">No</option>
                      <option value="Yes">yes</option>
                    </select>
                  </div>
                </div>
              </div>
            </div>
            <div class="row mb-3">
              <div class="col-md-12">
                <label class="form-label">Please select the Standard-Input-Json (*.json) file to upload<span class="ms-1 text-danger">*</span></label>
                <input type="file" class="form-control" style="background: transparent">
              </div>
            </div>

            <div class="accordion mb-3">
              <div class="accordion-item">
                <h2 class="accordion-header">
                  <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#abi-encode">
                    <label>
                      ABI-encoded Constructor Arguments <span class="text-secondary">(if required by the contract)</span>
                      <!--                      <a href="https://solidity.readthedocs.io/en/develop/abi-spec.html" target="_blank" class="text-cyan">ABI-encoded</a>-->
                    </label>
                  </button>
                </h2>
                <div id="abi-encode" class="accordion-collapse collapse" data-bs-parent="#accordion-example">
                  <div class="accordion-body pt-0">
                    <div class="form-group mb-3 row">
                      <div class="col-md-4">
                        <label class="col-form-label">Try to fetch constructor arguments automatically</label>
                      </div>
                      <div class="col-md-8" style="padding: 0.5rem;">
                        <label class="form-check form-check-inline">
                          <input class="form-check-input" name="autodetect_constructor" type="radio" checked="">
                          <span class="form-check-label">No</span>
                        </label>
                        <label class="form-check form-check-inline">
                          <input class="form-check-input" name="autodetect_constructor" type="radio">
                          <span class="form-check-label">Yes</span>
                        </label>
                      </div>
                    </div>
                    <textarea class="form-control" rows="5" style="background: transparent"></textarea>
                  </div>
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
export default {
  name: "InputJson",
  components: {FooterView, HeaderView}
}
</script>

<style scoped>

</style>