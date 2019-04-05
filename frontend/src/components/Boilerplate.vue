<template>
  <div></div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import Spinner from "@/components/Spinner.vue"; // @ is an alias to /src

@Component({
  components: {
    Spinner
  }
})
export default class Boilerplate extends Vue {}
</script>

<style scoped lang="less"></style>
