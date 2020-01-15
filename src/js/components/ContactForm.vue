<template>
    <form @submit.prevent="submit">
        <div v-if="submitStatus !== 'OK'">
            <div class="field">
                <label class="label has-text-white">Full Name</label>
                <div class="control">
                    <input class="input" v-model.trim="$v.name.$model" :class="{ 'is-danger': $v.name.$error }" placeholder="Your Full Name">
                </div>
                <div class="help is-danger" v-if="submitStatus === 'ERROR' && !$v.name.required">
                    Name is required
                </div>
            </div>
            <div class="field">
                <label class="label has-text-white">Email Address</label>
                <div class="control">
                    <input class="input" :class="{ 'is-danger': $v.email.$error }" v-model.trim="$v.email.$model" placeholder="Your Email Address">
                    <div class="help is-danger" v-if="submitStatus === 'ERROR' && !$v.email.required">
                        Email Address is required
                    </div>
                    <div class="help is-danger" v-if="submitStatus === 'ERROR' && !$v.email.email">
                        Valid Email Address is required
                    </div>
                </div>
            </div>
            <div class="field">
                <label class="label has-text-white">Message</label>
                <div class="control">
                    <input class="textarea" :class="{ 'is-danger': $v.message.$error }" v-model.trim="$v.message.$model" placeholder="Your Message">
                    <div class="help is-danger" v-if="submitStatus === 'ERROR' && !$v.message.required">
                        Message is required
                    </div>
                </div>
            </div>
            <br>
            <div class="field">
                <div class="control">
                    <button class="button is-primary" type="submit" :disabled="submitStatus === 'PENDING'">Send Message</button>
                </div>
            </div>
        </div>
        <div v-if="submitStatus === 'OK'">
            <p class="is-size-4 has-text-primary">Thanks for messaging us - We'll get back to you soon. If it's an emergency please call the number above.</p>
        </div>
    </form>
</template>

<script>
    import {required, email} from 'vuelidate/lib/validators'
    export default {
        data() {
            return {
                name: '',
                email: '',
                message: '',
                submitStatus: null
            }
        },

        validations: {
            name: { required },
            email: { required, email },
            message: { required }
        },

        methods:{
            fullMessage() {
                return `From: ${this.name}
                Email Address: ${this.email}
                Message: ${this.message}
                `
            },

            submit() {
                console.log('submit!')
                this.$v.$touch()
                if (this.$v.$invalid) {
                    this.submitStatus = 'ERROR'
                } else {
                    axios.post('/api/sendMessage', {
                        name: this.name,
                        email: this.email,
                        message: this.fullMessage()
                    })
                        .then(response => {
                            this.submitStatus = 'OK'
                        })
                        .catch((e) => {
                            console.error(e)
                        })
                }
            }
        }
    }
</script>