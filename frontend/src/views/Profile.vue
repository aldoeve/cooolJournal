<script setup>
import { ref, onBeforeMount } from "vue";
import Top_Bar from "@/components/Top_Bar.vue";
import { useRouter, useRoute } from "vue-router";
import axios from "axios";
import "../assets/welcome.css";
import "../assets/util.css";

const router = useRouter();
const route = useRoute();
const bio = ref('');
const profileUsername = ref('');
const userUsername = ref('');

onBeforeMount(() => {
  getProfile();
  getUserUsername();
}); 

async function getProfile() {
  try {
    const response = await axios.get(`/api/getUserProfile/${route.params.username}`);
    profileUsername.value = response.data.username[0];
    bio.value = response.data.bio[0];
    
  } catch(error) {
    console.error("Error:", error.message)
    router.push('/');
  }
}

async function getUserUsername() {
  try {
    const response = await axios.get('/api/getUsernameFromJWT');
    userUsername.value = response.data.username[0];
    
  } catch(error) {
    console.error("Error:", error.message)
  }
}

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
      <p v-if="profileUsername">{{ profileUsername }}</p>
      <p v-if="bio">{{ bio }}</p>
      <button v-if="userUsername === profileUsername"><i>Edit Profile</i></button>
    </div>
  </div>
</template>
