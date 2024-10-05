<template>
  <div class="landing-wrapper">
    <Top_Bar></Top_Bar>
    <div class="main-content">
      <Side_Bar></Side_Bar>
      <Content></Content>
    </div>
  </div>
</template>

<script setup>
import Top_Bar from '@/components/Top_Bar.vue';
import Side_Bar from '@/components/Side_Bar.vue';
import Content from '@/components/Content.vue';
</script>

<style scoped>
@import '../assets/landing.css';

</style>
