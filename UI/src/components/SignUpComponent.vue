<template>
    <v-container fill-height fluid>
        <v-row align="center" justify="center" dense>
            <v-col cols="3">
                <div class="login-form">
                    <v-col align="center"><h2>Sign Up</h2></v-col>
                    <v-col align="center">
                            <h3 :class="{'success': accountCreated, 'error': accountCreatedError}" >{{message}}</h3>
                    </v-col>
                    <v-form fluid>
                        <v-text-field
                            class="username"
                            label="Username"
                            :rules="usernameRules" 
                            hide-details="auto"
                            v-model="username">
                        </v-text-field>
                        <v-text-field
                            class="email"
                            label="Email"
                            :rules="emailRules"
                            v-model="email">
                        </v-text-field>
                        <v-text-field
                            class="password"
                            label="Password"
                            type="password"
                            :rules="passwordRules"
                            v-model="password">
                        </v-text-field>
                        <v-col align="center">
                            <v-btn
                                medium
                                class="register-button"
                                color="success"
                                justify="center"
                                @click="sendLoginRequest">
                                Sign Up
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
            password: "",
            email: "",
            usernameRules: [
                value => !!value || 'Required.',
                value => (value && value.length > 3) || 'Min of 3 chars.'
            ],
            passwordRules: [
                value => value.length >= 8 || 'Password should be 8 or more chars.'
            ],
            emailRules: [
                (value) => {
                    var validRegex = /^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/;
                    if (value.match(validRegex)){
                        return true
                    }
                    return 'Please enter a valid email.'
                }
            ]
        }
    },
    methods: {
        sendLoginRequest: function () {

            // Build the form data for the login
            const formData = new FormData();
            formData.append('username', this.username);
            formData.append('password', this.password);
            formData.append('email', this.email);

            let vm = this;

            this.axios.post("https://node404.com/auth/register", formData)
                      .then((res) => {
                          console.log(res);
                          if (res.status == 200) {
                              vm.message = "Account Created."
                              vm.accountCreated = true;
                          } else {
                              vm.message = "Failed to create account."
                              vm.accountCreated = false;
                              vm.accountCreatedError = true;
                          }
                      })
                      .catch((error) => {
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