<template>
  <div class="header">
    <div class="top-bar">
      <div class="top-bar-spacer"></div>
      <div class="top-bar-right">
        <i class="material-icons topnav-button" title="Search">search</i>
        <i class="material-icons topnav-button" title="Dark mode">brightness_4</i>
        <i class="material-icons topnav-button" title="Products">apps</i>
        <Login/>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import Login from "@/components/Login.vue";

@Component({
  components: {
    Login
  }
})
export default class Header extends Vue {}
</script>

<style scoped lang="less">
.header {
  padding-left: 150px;
  padding-top: 20px;
  font-size: 14pt;
  border-bottom: 1px solid rgb(235, 235, 235);
}
.top-bar {
  display: flex;
  align-items: center;
  margin-bottom: 10px;
}
.top-bar-spacer {
  flex-grow: 99;
}
.top-bar-right {
  flex-grow: 1;
  padding-right: 20px;
  display: flex;
  align-items: center;
}
.material-icons.topnav-button {
  color: rgb(95, 99, 104);
  padding: 8px;
}
.material-icons.topnav-button:hover {
  background-color: #ebecec;
  border-radius: 50%;
  cursor: pointer;
}
</style>
