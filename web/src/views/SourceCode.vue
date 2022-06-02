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
                  <label class="form-label">Optimization</label>
                  <div>
                    <select class="form-select" style="background: transparent">
                      <option value="No">No</option>
                      <option value="Yes">yes</option>
                    </select>
                  </div>
                </div>
              </div>
            </div>
            <div class="row">
              <div class="col-md-12">
                <div class="mb-3">
                  <label class="form-label">Enter the Solidity Contract Code below <span class="ms-1 text-danger">*</span></label>
                  <textarea class="form-control" name="example-textarea-input" rows="10" placeholder="Content.." style="background: transparent"></textarea>
                </div>
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

            <div class="accordion mb-3">
              <div class="accordion-item">
                <h2 class="accordion-header">
                  <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#contract-libraries">
                    <label>
                      Contract Libraries
                    </label>
                  </button>
                </h2>
                <div id="contract-libraries" class="accordion-collapse collapse" data-bs-parent="#accordion-example">
                  <div class="accordion-body pt-0">
                    <div class="text-secondary pb-3"><b>Note: Library names are case sensitive and affects the keccak library hash</b></div>
                    <div class="d-md-flex mb-3">
                      <div class="form-inline w-100">
                        <label class="text-monospace text-secondary mr-sm-3 mb-1">Library_1 Name:</label>
                        <input name="library_name[]" type="text" maxlength="42" style="background: transparent" class="form-control">
                      </div>
                      <div class="d-none d-md-block align-self-end me-3 ms-3">
                        <span class="btn btn-xs btn-icon btn-soft-secondary rounded-circle">
                          <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-arrow-narrow-right" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
                             <desc>Download more icon variants from https://tabler-icons.io/i/arrow-narrow-right</desc>
                             <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
                             <line x1="5" y1="12" x2="19" y2="12"></line>
                             <line x1="15" y1="16" x2="19" y2="12"></line>
                             <line x1="15" y1="8" x2="19" y2="12"></line>
                          </svg>
                        </span>
                      </div>
                      <div class="form-inline w-100">
                        <label class="text-monospace text-secondary mr-sm-2 mb-1">Library_1 Contract Address:</label>
                        <input name="library_address[]" maxlength="42" style="background: transparent" class="form-control">
                      </div>
                    </div>
                    <div class="d-md-flex mb-3">
                      <div class="form-inline w-100">
                        <label class="text-monospace text-secondary mr-sm-3 mb-1">Library_2 Name:</label>
                        <input name="library_name[]" type="text" maxlength="42" style="background: transparent" class="form-control">
                      </div>
                      <div class="d-none d-md-block align-self-end me-3 ms-3">
                        <span class="btn btn-xs btn-icon btn-soft-secondary rounded-circle">
                          <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-arrow-narrow-right" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
                             <desc>Download more icon variants from https://tabler-icons.io/i/arrow-narrow-right</desc>
                             <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
                             <line x1="5" y1="12" x2="19" y2="12"></line>
                             <line x1="15" y1="16" x2="19" y2="12"></line>
                             <line x1="15" y1="8" x2="19" y2="12"></line>
                          </svg>
                        </span>
                      </div>
                      <div class="form-inline w-100">
                        <label class="text-monospace text-secondary mr-sm-2 mb-1">Library_2 Contract Address:</label>
                        <input name="library_address[]" maxlength="42" style="background: transparent" class="form-control">
                      </div>
                    </div>
                    <div class="d-md-flex mb-3">
                      <div class="form-inline w-100">
                        <label class="text-monospace text-secondary mr-sm-3 mb-1">Library_3 Name:</label>
                        <input name="library_name[]" type="text" maxlength="42" style="background: transparent" class="form-control">
                      </div>
                      <div class="d-none d-md-block align-self-end me-3 ms-3">
                        <span class="btn btn-xs btn-icon btn-soft-secondary rounded-circle">
                          <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-arrow-narrow-right" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
                             <desc>Download more icon variants from https://tabler-icons.io/i/arrow-narrow-right</desc>
                             <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
                             <line x1="5" y1="12" x2="19" y2="12"></line>
                             <line x1="15" y1="16" x2="19" y2="12"></line>
                             <line x1="15" y1="8" x2="19" y2="12"></line>
                          </svg>
                        </span>
                      </div>
                      <div class="form-inline w-100">
                        <label class="text-monospace text-secondary mr-sm-2 mb-1">Library_3 Contract Address:</label>
                        <input name="library_address[]" maxlength="42" style="background: transparent" class="form-control">
                      </div>
                    </div>
                    <div class="d-md-flex mb-3">
                      <div class="form-inline w-100">
                        <label class="text-monospace text-secondary mr-sm-3 mb-1">Library_4 Name:</label>
                        <input name="library_name[]" type="text" maxlength="42" style="background: transparent" class="form-control">
                      </div>
                      <div class="d-none d-md-block align-self-end me-3 ms-3">
                        <span class="btn btn-xs btn-icon btn-soft-secondary rounded-circle">
                          <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-arrow-narrow-right" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
                             <desc>Download more icon variants from https://tabler-icons.io/i/arrow-narrow-right</desc>
                             <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
                             <line x1="5" y1="12" x2="19" y2="12"></line>
                             <line x1="15" y1="16" x2="19" y2="12"></line>
                             <line x1="15" y1="8" x2="19" y2="12"></line>
                          </svg>
                        </span>
                      </div>
                      <div class="form-inline w-100">
                        <label class="text-monospace text-secondary mr-sm-2 mb-1">Library_4 Contract Address:</label>
                        <input name="library_address[]" maxlength="42" style="background: transparent" class="form-control">
                      </div>
                    </div>
                    <div class="d-md-flex mb-3">
                      <div class="form-inline w-100">
                        <label class="text-monospace text-secondary mr-sm-3 mb-1">Library_5 Name:</label>
                        <input name="library_name[]" type="text" maxlength="42" style="background: transparent" class="form-control">
                      </div>
                      <div class="d-none d-md-block align-self-end me-3 ms-3">
                        <span class="btn btn-xs btn-icon btn-soft-secondary rounded-circle">
                          <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-arrow-narrow-right" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
                             <desc>Download more icon variants from https://tabler-icons.io/i/arrow-narrow-right</desc>
                             <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
                             <line x1="5" y1="12" x2="19" y2="12"></line>
                             <line x1="15" y1="16" x2="19" y2="12"></line>
                             <line x1="15" y1="8" x2="19" y2="12"></line>
                          </svg>
                        </span>
                      </div>
                      <div class="form-inline w-100">
                        <label class="text-monospace text-secondary mr-sm-2 mb-1">Library_5 Contract Address:</label>
                        <input name="library_address[]" maxlength="42" style="background: transparent" class="form-control">
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <div class="accordion mb-3">
              <div class="accordion-item">
                <h2 class="accordion-header">
                  <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#misc-settings">
                    <label>Misc Settings <span class="text-secondary">(Runs, EvmVersion &amp; Include nightly builds)</span></label>
                  </button>
                </h2>
                <div id="misc-settings" class="accordion-collapse collapse" data-bs-parent="#accordion-example">
                  <div class="accordion-body pt-0">
                    <div class="row">
                      <div class="col-md-4">
                        <div class="form-group mb-3 ">
                          <label class="form-label">Runs</label>
                          <div>
                            <input type="email" class="form-control" style="background: transparent" aria-describedby="emailHelp" placeholder="Runs" value="200">
                          </div>
                        </div>
                      </div>
                      <div class="col-md-4">
                        <div class="form-group mb-3 ">
                          <label class="form-label">EVM Version to target</label>
                          <div>
                            <select class="form-select" style="background: transparent"></select>
                          </div>
                        </div>
                      </div>
                      <div class="col-md-4">
                        <div class="form-group mb-3 ">
                          <label class="form-label">Include nightly builds</label>
                          <select class="form-select" style="background: transparent">
                            <option value="no" selected>No</option>
                            <option value="yes">Yes</option>
                          </select>
                        </div>
                      </div>
                    </div>
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
  name: "SourceCode",
  components: {FooterView, HeaderView}
}
</script>

<style scoped>

</style>