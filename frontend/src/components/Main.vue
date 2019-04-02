<template>
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
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import Sidenav from "@/components/Sidenav.vue"; // @ is an alias to /src

@Component({
  components: {
    Sidenav,
  }
})
export default class Main extends Vue {
}
</script>

<style scoped lang="less">
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
</style>