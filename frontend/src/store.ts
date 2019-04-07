import Vue from "vue";
import Vuex from "vuex";
import { AuthState } from "./enum";

Vue.use(Vuex);
export default new Vuex.Store({
  state: {
    auth2: null,
    auth: {
      state: AuthState.Preparing,
      email: null
    },
    sidebar: false,
    projects: null
  },
  mutations: {
    loggingout(state, payload) {
      state.auth = {
        state: AuthState.LoggingOut,
        email: null
      }
    },
    logout(state, payload) {
      state.auth = {
        state: AuthState.LoggedOut,
        email: null
      }
    },
    loadedAuth2(state, payload) {
      state.auth2 = payload;
      state.auth.state = AuthState.LoggedOut;
    },
    auth(state, payload) {
      if (!payload.email) {
        if (
          payload.state == AuthState.Verifying ||
          payload.state == AuthState.Verified
        ) {
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
    loadAuthClient: ({ commit }, payload) => {
      Vue.prototype.$getGapi().then((auth2: any) => {
        commit("loadedAuth2", auth2);
      });
    },
    signin: ({ commit, dispatch, state }, payload) => {
      commit("auth", { state: AuthState.LoggingIn });
      // https://developers.google.com/identity/sign-in/web/reference#gapiauth2offlineaccessoptions
      // const prompt = "select_account";
      // const prompt = "consent;"
      (state.auth2 as any).grantOfflineAccess().then((response: any) => {
        // Change state
        if (!response.code) {
          commit("auth", { state: AuthState.Error });
          console.log("Error.");
          return;
        }
        const googleUser = (state.auth2 as any).currentUser.get();
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
    logout: ({ commit  }, payload) => {
      commit("loggingout")
      fetch("http://localhost:8080/api/auth/logout", {
        method: "POST",
        cache: "no-cache",
        credentials: "same-origin",
        headers: {
          "Content-Type": "application/json",
          "X-Requested-With": "XMLHttpRequest"
        },
      }).then(_ => {
        // TODO: handle error
        commit("logout")
      });
    },
    verify: ({ commit, dispatch, state }, payload) => {
      console.log("Verifying with backend.");

      const response = payload;
      // https://developers.google.com/identity/sign-in/web/reference#googleusergetid
      const googleUser = (state.auth2 as any).currentUser.get();
      const { id_token } = googleUser.getAuthResponse();
      fetch("http://localhost:8080/api/authorization", {
        method: "POST",
        cache: "no-cache",
        credentials: "same-origin",
        headers: {
          "Content-Type": "application/json",
          "X-Requested-With": "XMLHttpRequest"
        },
        body: JSON.stringify({
          code: response.code,
          // TODO: verify id_token on server
          id_token: id_token
        })
      }).then(response => {
        // TODO: Handle error response
        commit("auth", {
          state: AuthState.Verified,
          email: state.auth.email
        });
        dispatch("handleVerificationResponse", response);
      });
    },
    handleVerificationResponse: ({ commit }, payload) => {
      payload.json().then((json: any) => {
        const projectNames = json.projects;
        commit("projects", projectNames);
      });
    }
  }
});
