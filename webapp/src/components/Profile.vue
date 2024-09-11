<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";

const changePicture = ref(false);

function requestProfileChange() {
  changePicture != changePicture;
}
</script>

<template>
  <div class="profileLayout">
    <div class="circleBase circleImage imgPosition">
      <img src="../../LionProfilePic.jpg" />
    </div>
    <button class="saveButton imgButtonPosition" @click="requestProfileChange">
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
    </button>
  </div>
</template>

<style>
.profileLayout {
  display: grid;
  grid-template-columns: 50% 20%;
  grid-template-rows: 50% 50% 50%;
  gap: 5vh;
  place-self: center;
  justify-content: center;
  position: relative;
  row-gap: 5vh;
}
.circleBase {
  background-color: white;
  border-radius: 50%;
  width: 40vh;
  height: 40vh;
}
.username {
  background-color: white;
  color: black;
  width: 40vh;
  height: 4vh;
  text-align: center;
  font-size: large;
  font-weight: bold;
  overflow: hidden;
  grid-area: 3/1;
}
.bio {
  width: 40vh;
  height: 20vh;
  text-align: left;
  padding-left: 1vh;
}
.saveButton {
  background-color: white;
  size: 1vh;
  width: 11vh;
  height: 4vh;
  text-align: center;
  font-size: large;
  font-weight: bold;
}
.imgPosition {
  grid-area: 1/1;
}
.imgButtonPosition {
  margin-top: 36vh;
  grid-area: 1/2;
}
.usernamePosition {
  grid-area: 2/1;
}
.usernameButtonPosition {
  grid-area: 2/2;
  margin-left: 8vh;
}
.bioPosition {
  grid-area: 3/1;
}
.bioButtonPosition {
  grid-area: 3/2;
  margin-left: 8vh;
  margin-top: 6vh;
}
</style>
