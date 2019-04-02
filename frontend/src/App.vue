<template>
  <div id="app">
    <Header/>
    <Main/>
    <BuildMetadata/>
    <!--
    <div id="nav">
      <router-link to="/">Home</router-link> |
      <router-link to="/about">About</router-link>
    </div>
    <router-view />
    -->
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import Header from "@/components/Header.vue"; // @ is an alias to /src
import BuildMetadata from "@/components/BuildMetadata.vue"; // @ is an alias to /src
import Main from "@/components/Main.vue"; // @ is an alias to /src

@Component({
  components: {
    BuildMetadata,
    Header,
    Main,
  }
})
export default class Home extends Vue {
  beforeCreate() {
     // Pings API to check that it's up.
    fetch("http://localhost:8080/api/health",  {
        method: "GET",
        cache: "no-cache",
        credentials: "same-origin",
      }).then(response => {
        // TODO: alert in the UI
        console.log(response.json())
      });
  }
}
</script>


<style lang="less">
#app {
  font-family:  Geneva, sans-serif;
  display: flex;
  min-height: 100vh;
  flex-direction: column;
}
// #nav {
//   padding: 30px;
//   a {
//     font-weight: bold;
//     color: #2c3e50;
//     &.router-link-exact-active {
//       color: #42b983;
//     }
//   }
// }
</style>
