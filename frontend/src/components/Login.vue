<template>
  <a class="login-button" @click="login">{{msg}}</a>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";

@Component
export default class Login extends Vue {
  msg = "Sign in "

  login () {
    // https://developers.google.com/identity/sign-in/web/reference#gapiauth2offlineaccessoptions
    // const prompt = "select_account";
    // const prompt = "consent;"
    (window as any).auth2.grantOfflineAccess().then((response: any) => {
      if(!response.code) {
        console.log("Error.")
        return;
      }
      (document as any).getElementById('sidenav').classList.remove("hidden");

      // https://developers.google.com/identity/sign-in/web/reference#googleusergetid
      const googleUser = (window as any).auth2.currentUser.get();

      console.log("Verifying with backend.");
      const {id_token} = googleUser.getAuthResponse();
      fetch("http://localhost:8080/api/authorization", {
          method: "POST",
          cache: "no-cache",
          credentials: "same-origin",
          headers: {
            "Content-Type": "application/json",
            "X-Requested-With": "XMLHttpRequest"
          },
          body: JSON.stringify({
            "code": response.code,
            // TODO: verify id_token on server
            "id_token": id_token
          })
      }).then(response => {
        response.json().then(json => {
          const projectNames = json.projects;
          const newUl = document.createElement('ul');
          newUl.id = "sidenav-projectlist";
          for(let i = 0 ; i < projectNames.length; i++) {
            const item = document.createElement('li');
            const a = document.createElement('a');
            a.setAttribute("href", "")
            a.appendChild(document.createTextNode(projectNames[i]));
            item.appendChild(a);
            // https://coderwall.com/p/o9ws2g/why-you-should-always-append-dom-elements-using-documentfragments
            newUl.appendChild(item);
          }
          const frag = document.createDocumentFragment();
          frag.appendChild(newUl);
          const ul: any = document.getElementById("sidenav-projectlist");
          ul.parentNode.replaceChild(frag, ul as any);

          (document as any).getElementById('sidenav-projectlist-spinner-container').classList.add("hidden");
          (document as any).getElementById('sidenav-projectlist-container').classList.remove("hidden");
        });
        // TODO: update login "verifying" state.
      });
      const profile = googleUser.getBasicProfile();
      const email = profile.getEmail();
      // BasicProfile.getId()
      // BasicProfile.getName()
      // BasicProfile.getGivenName()
      // BasicProfile.getFamilyName()
      // BasicProfile.getImageUrl()
      // BasicProfile.getEmail();
      this.msg = email + " ( verifying )";
    });

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
