<template>
  <div class="sidenav" v-if="showSidenav">
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
      <div v-if="showSpinner" class="sidenav-projectlist-spinner-container">
        <Spinner/>
      </div>
      <div v-if="!showSpinner" class="sidenav-projectlist-container">
        <ul class="sidenav-projectlist">
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
import Spinner from "@/components/Spinner.vue"; // @ is an alias to /src

@Component({
  components: {
    Spinner
  }
})
export default class Sidenav extends Vue {
  get projectNames() {
    return this.$store.state.projects || [];
  }

  get showSpinner() {
    return !this.$store.state.projects;
  }

  get showSidenav() {
    return this.$store.state.sidebar;
  }
}
</script>

<style scoped lang="less">
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