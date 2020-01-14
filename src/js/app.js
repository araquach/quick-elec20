import Vue from 'vue'
import Vuelidate from 'vuelidate'
import Slider from "./components/Slider"
import ContactForm from "./components/ContactForm";
import Buefy from "buefy";

Vue.use(Vuelidate)
Vue.use(Buefy)

Vue.component('slider-component', Slider)
Vue.component('contact-form', ContactForm)

new Vue({
    el: '#app',
});
