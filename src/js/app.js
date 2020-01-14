import Vue from 'vue'
import Vuelidate from 'vuelidate'
import Nav from './components/Nav'
import Slider from "./components/Slider"
import ContactForm from "./components/ContactForm"
import Buefy from "buefy"

Vue.use(Vuelidate)
Vue.use(Buefy)

Vue.component('nav-component', Nav)
Vue.component('slider-component', Slider)
Vue.component('contact-form', ContactForm)

window.axios = require('axios')

new Vue({
    el: '#app'
});
