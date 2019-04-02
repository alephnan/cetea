<template>
  <div id="sidenav" class="sidenav hidden">
    <div class="sidenav-header">
      <div class="sidenav-header-title">Projects</div>
      <div class="sidenav-header-button">
        <div class="button-container">
          <button href="#" class="button">
            New
          </button>
        </div>
      </div>
    </div>
    <div>
      <div id="sidenav-projectlist-spinner-container" class="sidenav-projectlist-spinner-container">
        <div class="sidenav-projectlist-spinner lds-roller"><div></div><div></div><div></div><div></div><div></div><div></div><div></div><div></div></div>
      </div>
      <div id="sidenav-projectlist-container" class="sidenav-projectlist-container hidden">
        <ul id="sidenav-projectlist" class="sidenav-projectlist">
          <!-- TODO: handle too long name. truncate, use ellipsis. -->
          <li v-for="item in projectNames">
            <a href="">{{ item }}</a>
          </li>
        </ul>
        <a href="" class="show-more">
          Show more
        </a>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";

@Component
export default class Sidenav extends Vue {
  get projectNames() {
    return this.$store.state.projects || [];
  }
}
</script>

<style scoped lang="less">
.sidenav.hidden {
  display: none;
}
.sidenav {
  order: 1;
  flex: 0 0 200px;
  overflow: hidden;
  padding-left: 25px;
  padding-right: 25px;
  padding-top: 35px;
  background: url("http://thecodeplayer.com/uploads/media/geometry.png");
}
.sidenav-header {
  display: flex;
  justify-content: center;
}
.sidenav-header-title {
  flex-grow: 99;
  align-items: center;
  display: flex;
  font-weight: bold;
}
.sidenav-header-button {
  flex: 1;
}
.sidenav-projectlist-spinner-container {
  display: flex;
  align-content: center;
  justify-content: center;
  padding-top: 40px;
}
.sidenav-projectlist-spinner-container.hidden {
  display: none;
}
.sidenav-projectlist-container.hidden {
  display: none;
}
.sidenav-projectlist {
  list-style-type: none;
  padding: 0px;
}
.sidenav-projectlist li {
  margin-top: 8px;
}
.sidenav-projectlist a {
  text-decoration: none;
  font-weight: bold;
}
.sidenav-projectlist a:hover {
  text-decoration: underline;
}
.sidenav-projectlist a:visited {
  color: rgb(43, 125, 233);
  text-decoration: underline;
}
.show-more {
  font-size: 14px;
  color: grey;
  text-decoration: none;
}
.show-more:hover {
  color: rgb(43, 125, 233);
}
</style>