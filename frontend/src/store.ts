import Vue from "vue";
import Vuex from "vuex";
import {AuthState} from "./enum";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    auth: {
      state: AuthState.LoggedOut,
      email: null,
    },
    sidebar: false,
    projects: null
  },
  mutations: {
    auth(state, payload) {
      if(!payload.email) {
        if(payload.state == AuthState.Verifying || payload.state == AuthState.Verified) {
          // TODO: throw invalid state transition error.
        }
      }
      state.auth = payload;
    },
    showSidebar(state) {
      state.sidebar = true;
    },
    projects(state, payload) {
      state.projects = payload;
    }
  },
  actions: {
    signin: ({commit, dispatch}, payload) => {
      commit("auth", {state: AuthState.LoggingIn});
      // https://developers.google.com/identity/sign-in/web/reference#gapiauth2offlineaccessoptions
      // const prompt = "select_account";
      // const prompt = "consent;"
      (window as any).auth2.grantOfflineAccess().then((response: any) => {
        // Change state
        if(!response.code) {
          commit("auth", {state: AuthState.Error});
          console.log("Error.");
          return;
        }
        const googleUser = (window as any).auth2.currentUser.get();
        const profile = googleUser.getBasicProfile();
        const email = profile.getEmail();
        // BasicProfile.getId()
        // BasicProfile.getName()
        // BasicProfile.getGivenName()
        // BasicProfile.getFamilyName()
        // BasicProfile.getImageUrl()
        // BasicProfile.getEmail();
        commit("auth", {
          state: AuthState.Verifying,
          email
        });
        commit("showSidebar");
        dispatch("verify", response);
      });
    },
    verify: ({commit, dispatch, state}, payload) => {
      console.log("Verifying with backend.");

      const response = payload;
      // https://developers.google.com/identity/sign-in/web/reference#googleusergetid
      const googleUser = (window as any).auth2.currentUser.get();
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
        // TODO: Handle error response
        commit("auth", {
          state: AuthState.Verified,
          email: state.auth.email
        });
        dispatch('handleVerificationResponse', response);
      });
    },
    handleVerificationResponse: ({commit}, payload) => {
      payload.json().then((json: any) => {
        const projectNames = json.projects;
        commit("projects", projectNames);
      });
    }
  }
});
