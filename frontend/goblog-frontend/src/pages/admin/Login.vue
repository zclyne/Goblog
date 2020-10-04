<template>
    <q-page class="flex flex-center">
        <q-card class="login-card">

            <q-tabs
                    v-model="tab"
                    dense
                    class="text-grey"
                    active-color="primary"
                    indicator-color="primary"
                    align="justify"
                    narrow-indicator>
                <q-tab name="login" label="Login" />
                <q-tab name="register" label="Register" />
            </q-tabs>
            <q-separator />
            <q-tab-panels v-model="tab" animated>
                <q-tab-panel name="login">
                    <q-input label="Username" v-model="username" class="login-input-field"/>
                    <q-input label="Password" v-model="password" type="password" class="login-input-field"/>
                    <q-btn class="q-mt-md" color="primary" @click="login">Submit</q-btn>
                </q-tab-panel>
                <q-tab-panel name="register">
                    <q-input label="Username" v-model="username" class="login-input-field"/>
                    <q-input label="Password" v-model="password" class="login-input-field"/>
                    <q-input label="Email" v-model="email" class="login-input-field"/>
                    <q-btn class="q-mt-md" color="primary" @click="login">Submit</q-btn>
                </q-tab-panel>
            </q-tab-panels>
        </q-card>
    </q-page>
</template>

<script>
    import axios from 'axios'
    export default {
        name: 'Login',
        data () {
            return {
                tab: 'login',
                username: '',
                password: '',
                email: '',
                userType: 2
            }
        },
        methods: {
            login () {
                if (this.username === '' || this.password === '') {
                    // TODO: prompt to input username and password
                    return
                } else {
                    axios.post('/api/user/login', {
                        username: this.username,
                        password: this.password
                    }).then(this.loginSuccess)
                }
            },
            loginSuccess (res) {
                // TODO: prompt login successful
                console.log(res)
                // save user info in Vuex
                this.$store.commit('user/saveUserToken', res.data.token)
                // jump to the admin page
                this.$router.push('/admin/blogs')
            },
            register () {
                if (this.username === '' || this.password === '' || this.email === '') {
                    // TODO: prompt to input username, password and email
                    return
                } else {
                    axios.post('/api/user/register', {
                        username: this.username,
                        password: this.password,
                        email: this.email,
                        type: this.type
                    }).then(this.loginSuccess)
                }
            }
        }
    }
</script>

<style scoped>
    .login-card {
        width: 400px;
    }
    .login-input-field {
        width: 100%;
    }
</style>
