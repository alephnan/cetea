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

/* Begin animated button */
.button-container {
  text-align: center;
}
.button {
  display: inline-block;
  padding: 6px 24px;
  border-radius: 4px;
  transition: all 0.2s ease-in;
  position: relative;
  overflow: hidden;
  z-index: 100;
  background-color: #39bda7;
  color: white;
  cursor: pointer;
}
.button:before {
  content: "";
  position: absolute;
  left: 50%;
  transform: translateX(-50%) scaleY(1) scaleX(1.25);
  top: 100%;
  width: 140%;
  height: 180%;
  background-color: rgba(0, 0, 0, 0.05);
  border-radius: 50%;
  display: block;
  transition: all 0.5s 0.1s cubic-bezier(0.55, 0, 0.1, 1);
  z-index: -1;
}
.button:after {
  content: "";
  position: absolute;
  left: 55%;
  transform: translateX(-50%) scaleY(1) scaleX(1.45);
  top: 180%;
  width: 160%;
  height: 190%;
  background-color: white;
  border-radius: 50%;
  display: block;
  transition: all 0.5s 0.1s cubic-bezier(0.55, 0, 0.1, 1);
  z-index: -1;
  border: 1px solid #39bda7;
}
.button:hover {
  color: #39bda7;
  border: 1px solid #39bda7;
}
.button:hover:before {
  top: -35%;
  background-color: white;
  transform: translateX(-50%) scaleY(1.3) scaleX(0.8);
}
.button:hover:after {
  top: -45%;
  background-color: white;
  transform: translateX(-50%) scaleY(1.3) scaleX(0.8);
}
/* End animated button */
/* Tooltip */
.popover {
  position: absolute;
  display: none;
  width: 250px;
  background-color: white;
  border: 1px solid rgb(235, 235, 235);
  border-radius: 3px;
  box-shadow: 0 5px 10px rgba(0,0,0,.2);
  -webkit-box-shadow: 0 5px 10px rgba(0,0,0,.2);
}
.popover .popover-content {
  padding: 10px;
}
.popover .arrow {
  border-style: solid;
  border-color: transparent;
  border-top-color: rgb(235, 235, 235);
  border-width: 11px;
  border-bottom-width: 0;
  bottom: -11px;
  color: #333;
  height: 0;
  margin-left: -11px;
  position: absolute;
  width: 0;
}
.popover .arrow:after {
  content: " ";
  bottom: 1px;
  margin-left: -10px;
  border-color: transparent;
  border-style: solid;
  border-width: 10px;
  border-bottom-width: 0;
  border-top-color: #ffe;
  position: absolute;
  display: block;
  width: 0;
  height: 0;
}
.tip {
  text-decoration: none;
  color: rgb(43, 125, 233);
  cursor: pointer;
}
.tip:hover {
  text-decoration: underline;
}
/* End tooltip */
</style>