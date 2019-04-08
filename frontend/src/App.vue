<template>
  <div id="app">
    <Header/>
    <div class="main">
      <Sidenav/>
      <div class="content">
        <router-view/>
      </div>
    </div>
    <BuildMetadata/>
    <div>
      <router-link to="/">Home</router-link>|
      <router-link to="/about">About</router-link>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
// @ is an alias to /src
import BuildMetadata from "@/components/BuildMetadata.vue";
import Header from "@/components/Header.vue";
import Sidenav from "@/components/Sidenav.vue";

@Component({
  components: {
    BuildMetadata,
    Header,
    Sidenav
  }
})
export default class Home extends Vue {
  beforeCreate() {
    // Pings API to check that it's up.
    fetch("http://localhost:8080/api/health", {
      method: "GET",
      cache: "no-cache",
      credentials: "same-origin"
    }).then(response => {
      // TODO: alert in the UI
      console.log(response.json());
    });
    // TODO: check cookies for logged in state.
  }
  mounted() {
    this.$store.dispatch("loadAuthClient");
  }
}
</script>

<style lang="less">
#app {
  font-family: Geneva, sans-serif;
  display: flex;
  min-height: 100vh;
  flex-direction: column;
}
.main {
  display: -webkit-flex;
  display: flex;
  flex: 1;
}
.content {
  flex: 1;
  order: 2;
  padding-left: 25px;
  padding-right: 25px;
  padding-top: 35px;
  border-left: 1px solid rgb(235, 235, 235);
}
</style>
