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
  router.push("/home");
}
</script>

<style scoped>
@import "../assets/top_bar.css";
@import "../assets/font-awesome-4.7.0/css/font-awesome.min.css";
@import "../assets/welcome.css";
</style>

<template>
  <div class="wrapper">
    <div class="profileLayout">
      <div class="circleBase circleImage imgPosition">
        <img src="../../LionProfilePic.jpg" />
      </div>
      <button @click="requestProfileChange"><span>Change</span></button>
      <!--<button class="saveButton imgButtonPosition" @click="requestProfileChange">
        <span>Change</span>
      </button>
      <div class="username usernamePosition"><span>User</span></div>
      <button class="saveButton usernameButtonPosition">
        <span>Change</span>
      </button>
      <div class="username bio bioPosition">
        <span>Hello, I am here to have fun ah ha ha.</span>
      </div>
      <button class="saveButton bioButtonPosition">
        <span>Change</span>
      </button>-->
    </div>
  </div>
</template>
