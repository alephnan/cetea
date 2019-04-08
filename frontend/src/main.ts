import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";

import GoogleSignin from "./plugins/google_signin";
import { client_id } from "@/../config";

Vue.config.productionTip = false;

const APP_ELEMENT_ID = "app";
Vue.use(GoogleSignin, {
  clientId: client_id
});
new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#" + APP_ELEMENT_ID);
