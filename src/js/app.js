import Vue from 'vue'
import Vuelidate from 'vuelidate'
import Slider from "./components/Slider"
import Buefy from "buefy";

Vue.use(Vuelidate)
Vue.use(Buefy)

Vue.component('slider-component', Slider)

new Vue({
    el: '#app',
});
