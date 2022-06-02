<template>
  <header class="navbar navbar-expand-md navbar-light d-print-none" style="box-shadow: none;">
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

      <div style="height: 2.8rem">
        <search-view />
      </div>
    </div>
  </header>

  <div class="navbar-expand-md">
    <div class="collapse navbar-collapse" id="navbar-menu">
      <div class="navbar navbar-light">
        <div class="container-xl justify-content-end">
          <menu-view />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import SearchView from "@/components/Search";
import MenuView from "@/components/Menu";
export default {
  name: "HeaderView",
  components: {MenuView, SearchView}
}
</script>

<style scoped>

</style>