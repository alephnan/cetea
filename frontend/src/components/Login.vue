<template>
  <a class="login-button" @click="login">{{msg}}</a>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import {AuthState} from "../enum";

@Component
export default class Login extends Vue {
  login () {
    this.$store.dispatch("signin")
  }

  get msg () {
    const auth = this.$store.state.auth;
    switch(auth.state) {
      case AuthState.LoggedOut:
        return "Sign in";
      case AuthState.LoggingIn:
        return "Logging in";
      case AuthState.Verifying:
        return auth.email + " ( Verifying ) ";
      case AuthState.Verified: 
        return auth.email;
      case AuthState.Error:
        return "Error";
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="less">
.login-button {
  padding-top: 9px;
  padding-bottom: 9px;
  padding-left: 23px;
  padding-right: 23px;
  font-size: 14px;
  background-color: rgb(43, 125, 233);
  color: white;
  border-bottom-left-radius: 4px;
  border-bottom-right-radius: 4px;
  border-top-left-radius: 4px;
  border-top-right-radius: 4px;
  font-family: "Google Sans",Roboto,RobotoDraft,Helvetica,Arial,sans-serif;
  margin-left: 10px;
}
.login-button:hover {
  cursor: pointer;
  box-shadow: 1px 4px 5px 1px rgba(0,0,0,0.1);
}
</style>
