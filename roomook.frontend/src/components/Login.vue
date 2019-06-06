<template>
  <v-app class="app">
    <v-layout align-center>
      <v-flex xs8 sm6 offset-md4 offset-xs2 md4 align-self-center: true>
        <v-card>
          <v-flex>
            <v-layout row wrap>
              <h2 class="login-heading">Roomook Login</h2>
              <v-flex xs10 offset-xs1 class="login-form">
                <v-text-field label="Email or Username" outline :autofocus="false" v-model="username"></v-text-field>
              </v-flex>
              <v-flex xs10 offset-xs1 class="password-form">
                <v-text-field 
                 v-on:keyup.enter="login()"
                  v-model="password"
                  :append-icon="show1 ? 'visibility' : 'visibility_off'"
                  :rules="[rules.required, rules.min]"
                  :type="show1 ? 'text' : 'Password'"
                  label="Password"
                  hint="At least 8 characters"
                  counter
                  @click:append="show1 = !show1"
                  outline
                ></v-text-field>
              </v-flex>
            </v-layout>
          </v-flex>
          <v-flex>
            <v-btn color="green" class="signIn" @click="login()">Sign In</v-btn>
          </v-flex>
           <v-alert
              
               icon="warning"
              :value="alert"
               type="error"
            >The email or password you entered don't match!</v-alert>
        </v-card>
      </v-flex>
    </v-layout>
  </v-app>
</template>

<script>
import axios from "axios";
export default {
  name: "login",
  data() {
    return {
      username: "",
      show1: false,
      show2: true,
      show3: false,
      show4: false,
      password: "",
      alert: false,
      
      rules: {
        required: value => !!value || "Required.",
        min: v => v.length >= 8 || "Min 8 characters",
        emailMatch: () => "The email and password you entered don't match"
      }
    };
  },
  methods: {
    login() {
      axios({
        method: "post",
        url: this.$url + "/login",
        data: {
          user: this.username,
          password: this.password
        }
      })
        .then(response => {
         if (response.status === 200) {
            localStorage.setItem("Authentication", response.data.jwt);
            this.$router.push("/"); 
          }
          
        })
        .catch(err => {
          this.alert = true;
          console.log(err);
        });
        
    }
  }
};
</script> 

  <style scoped>
.login-heading {
  width: 100%;
  padding: 12px 20px;
  margin: 8px 0;
  display: inline-block;
  font-size: 35px;
}
.signIn {
  text-decoration: none;
  text-align: center;
  color: #fff;
  font-weight: bold;
  padding: 8px 20px;
  margin-top: 0px;
  margin-bottom: 15px;
  font-size: 16px;
  border-radius: 10px;
  background-color: #666666;
  box-shadow: 0 5px 5px #313131, 0 9px 0 #393939,
    0px 9px 10px rgba(0, 0, 0, 0.4), inset 0px 2px 9px rgba(255, 255, 255, 0.2),
    inset 0 -2px 9px rgba(0, 0, 0, 0.2);
  position: relative;
  border-bottom: 1px solid rgba(255, 255, 255, 0.2);
  display: inline-block;
  font-family: Arial, Helvetica, sans;
  text-shadow: 0px -1px 0px rgba(0, 0, 0, 0.2);
  margin-left: auto;
  margin-right: auto;
}

.app {
  background-image: url(https://images.pexels.com/photos/417074/pexels-photo-417074.jpeg?auto=compress&cs=tinysrgb&dpr=2&h=750&w=1260);
  background-repeat: no-repeat;
  background-position-x: 0px;
  background-position-y: 0px;
}
</style>
