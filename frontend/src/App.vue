<template>
  <div id="app">
    <Header/>
    <div id="main">
      <Sidenav />
      <div id="content">
        <div id="orgline">organization: <span id="organization">google.com</span></div>
        <div id="project">
          <span id="projectlabel">project:</span>
          <span id="projectname">my-first-project</span>
          <span id="projectid">( uplifted-scout-234505 )</span>
        </div>
      </div>
    </div>
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
 // @ is an alias to /src
import BuildMetadata from "@/components/BuildMetadata.vue";
import Header from "@/components/Header.vue";
import Sidenav from "@/components/Sidenav.vue";

@Component({
  components: {
    BuildMetadata,
    Header,
    Sidenav,
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
#main {
  display: -webkit-flex;
  display: flex;
  flex: 1;
}
#content {
  flex: 1;
  order: 2;
  padding-left: 25px;
  padding-right: 25px;
  padding-top: 35px;
  border-left: 1px solid rgb(235, 235, 235);
}
#project {
  font-size: 14pt;
}
#projectlabel {
  color: #828282;
}
#projectid {
  font-size: 10pt;
  color: #828282;
}
#orgline {
  font-size: 9pt;
  color: #828282;
}
#organization {
  color: black;
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
