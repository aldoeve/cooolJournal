<script setup>
import { ref, onBeforeMount } from "vue";
import Top_Bar from "@/components/Top_Bar.vue";
import { useRouter } from "vue-router";
import axios from "axios";
import "../assets/welcome.css";
import "../assets/util.css";

const router = useRouter();
const bio = ref('');
const username = ref('');

onBeforeMount(() => {
  
}); 

function gotoHome() {
  router.push("/");
}
</script>

<style scoped>
.profile-wrapper {
  position: fixed;
    inset: 0;
    background-color: white;
    display: flex;
    flex-direction: column;
}

.profile-info-container {
    background-color: white;
    width: 320px;
    min-width: 320px;
    height: 100%;
    border-right: 1px solid black;
    display: flex;
    flex-direction: column;
    align-items: center;
    padding-top: 16px;
    overflow: hidden;
  }

</style>

<template>
  <div class="profile-wrapper">
    <Top_Bar></Top_Bar>
    <div class="profile-info-container">
      <p>{{ bio.toString() }}</p>
    </div>
  </div>
</template>
