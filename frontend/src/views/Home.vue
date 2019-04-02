<template>
  <div>
    <div id="orgline">organization: <span id="organization">google.com</span></div>
    <div id="project">
      <span id="projectlabel">project:</span>
      <span id="projectname">my-first-project</span>
      <span id="projectid">( uplifted-scout-234505 )</span>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

@Component({
  components: {
  }
})
export default class Home extends Vue {}
</script>


