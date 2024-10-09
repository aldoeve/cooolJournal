<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import axios from "axios";
import "../assets/welcome.css";
import "../assets/profile.css";
import "../assets/util.css";

const router = useRouter();

const data = ref(null);

function gotoHome() {
  router.push("/home");
}

async function updateProfilePic() {
  try {
    const response = await axios.post("/api/updateProfilePic", {
      picture: "yes",
    });
    if (response.status !== 200) {
      throw new Error(response.status);
    }
    data.value = response.data;
  } catch (error) {
    console.error("Error:", error.message);
    return;
  }
}

async function updateUsername() {
  try {
    const response = await axios.post("/api/updateUsername", {
      name: "yey",
    });
    if (response.status !== 200) {
      throw new Error(response.status);
    }
    data.value = response.data;
  } catch (error) {
    console.error("Error:", error.message);
    return;
  }
}

async function updateBio() {
  try {
    const response = await axios.post("/api/updateBio", {
      bio: "i",
    });
    if (response.status !== 200) {
      throw new Error(response.status);
    }
    data.value = response.data;
  } catch (error) {
    console.error("Error:", error.message);
    return;
  }
}
</script>

<style scoped>
@import "../assets/top_bar.css";
@import "../assets/font-awesome-4.7.0/css/font-awesome.min.css";
@import "../assets/welcome.css";
</style>

<template>
  <div class="wrapper">
    <div class="profileContainer">
      <div class="circleBase circleImage imgPosition">
        <img src="../../public/LionProfilePic.jpg" />
      </div>
      <button class="profileButton imgButtonLocation" @click="updateProfilePic">
        <span>Change</span>
      </button>
      <div class="username"><span>John John</span></div>
      <button class="profileButton usernameButtonLocation" @click="updateUsername">
        <span>Change</span>
      </button>
      <div class="bio bioLocation"><span>Hello, I am just here to be silly.</span></div>
      <button class="profileButton bioButtonLocation" @click="updateBio">
        <span>Change</span>
      </button>
    </div>
    <button class="profileButton profileToHome homeButton" @click="gotoHome">
      <span>Home</span>
    </button>
    <div class="enlargeMsg"><span>Please expand the window.</span></div>
  </div>
</template>
