<template>
  <div>
    <div class="orgline">
      organization:
      <span class="organization">google.com</span>
    </div>
    <div class="project">
      <span class="projectlabel">project:</span>
      <span>my-first-project</span>
      <span class="projectid">( uplifted-scout-234505 )</span>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

@Component({
  components: {}
})
export default class Home extends Vue {}
</script>

<style>
.orgline {
  font-size: 9pt;
  color: #828282;
}
.project {
  font-size: 14pt;
}
.organization {
  color: black;
}
.projectlabel {
  color: #828282;
}
.projectid {
  font-size: 10pt;
  color: #828282;
}
</style>