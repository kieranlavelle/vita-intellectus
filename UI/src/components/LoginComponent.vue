<template>
    <v-container fill-height fluid>
        <v-row align="center" justify="center" dense>
            <v-col cols="3">
                <div class="login-form">
                    <v-col align="center"><h2>Login</h2></v-col>
                    <v-col align="center">
                            <h3 :class="{'success': accountCreated, 'error': accountCreatedError}">{{message}}</h3>
                    </v-col>
                    <v-form fluid>
                        <v-text-field
                            class="username"
                            label="Username"
                            hide-details="auto"
                            v-model="username">
                        </v-text-field>
                        <v-text-field
                            class="password"
                            label="Password"
                            type="password"
                            v-model="password">
                        </v-text-field>
                        <v-col align="center">
                            <v-btn
                                medium
                                class="register-button"
                                color="success"
                                justify="center"
                                @click="sendLoginRequest">
                                Login
                            </v-btn>
                        </v-col>
                    </v-form>
                </div>
            </v-col>
        </v-row>
    </v-container>
</template>

<script>
  export default {
    name: 'SignUpComponent',
    data() {
        return {
            accountCreated: false,
            accountCreatedError: false,
            message: "",
            username: "",
            password: ""
        }
    },
    methods: {
        sendLoginRequest: function () {

            // Build the form data for the login
            const formData = new FormData();
            formData.append('username', this.username);
            formData.append('password', this.password);

            let vm = this;

            this.axios.post("https://node404.com/auth/login", formData)
                      .then((res) => {
                          console.log(res);
                          if (res.status == 200) {
                              vm.message = "Logged In."
                              vm.accountCreated = true;


                              //set the token
                              window.localStorage.setItem('token', res.data.access_token);
                              window.localStorage.setItem('username', res.data.username);
                              vm.$router.push('home');
                          } else {
                              vm.message = "Invalid Credentials."
                              vm.accountCreated = false;
                              vm.accountCreatedError = true;
                          }
                      })
                      .catch((error) => {
                            vm.message = "Invalid Credentials."
                            vm.accountCreated = false;
                            vm.accountCreatedError = true;
                          console.log(error);
                      })
        }
    },
    }
</script>

<style scoped>

.login-form {
    width: 100%;
    padding: 55px 25px 25px 25px;

    border: 1px solid rgba(0, 0, 0, 0.3);
    box-shadow: rgba(0, 0, 0, 0.55) 0px 5px 15px;

    background-color: rgb(255, 255, 255);
    opacity: 0.9;
}

.error {
    color: red;
}

.success {
    color: green;
}

.register-button {
    margin: 0 auto;
}

.username {
    padding-bottom: 25px;
}

.email {
    padding-bottom: 25px;
}

.password {
    padding-bottom: 25px;
}


.container {
    background-image: url("../assets/login-bg.jpg");
    background-size: cover;
    background-repeat: no-repeat;
}

</style>