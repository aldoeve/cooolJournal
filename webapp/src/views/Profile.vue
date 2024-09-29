<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import "../assets/welcome.css";
import "../assets/profile.css";
import "../assets/util.css";

const router = useRouter();

const changePicture = ref(false);

function requestProfileChange() {
  changePicture != changePicture;
}

function gotoHome() {
  router.push("/Home");
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
        <img src="../../LionProfilePic.jpg" />
      </div>
      <button class="profileButton imgButtonLocation" @click="requestProfileChange">
        <span>Change</span>
      </button>
      <div class="username"><span>John John</span></div>
      <button class="profileButton usernameButtonLocation"><span>Change</span></button>
      <div class="bio bioLocation"><span>Hello, I am just here to be silly.</span></div>
      <button class="profileButton bioButtonLocation">
        <span>Change</span>
      </button>
    </div>
    <button class="profileButton profileToHome homeButton" @click="gotoHome">
      <span>Home</span>
    </button>
  </div>
</template>
