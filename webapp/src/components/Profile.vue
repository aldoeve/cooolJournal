<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";

const changePicture = ref(false);

function requestProfileChange() {
  changePicture != changePicture;
}
</script>

<template>
  <div class="profileRow">
    <div class="profileCol">
      <div class="circleBase circleImage">
        <img src="../../LionProfilePic.jpg" />
      </div>
      <button
        class="profileCol saveButton saveButtonPosition"
        @click="requestProfileChange"
      >
        <span>Change</span>
      </button>
    </div>
  </div>
  <div class="profileRow">
    <div class="profileCol">
      <div class="username"><span>User</span></div>
    </div>
    <button class="profileCol saveButton">
      <span>Change</span>
    </button>
  </div>
  <div class="profileRow">
    <div class="profileCol">
      <div class="username bio"><span>Hello, I am here to have fun.</span></div>
    </div>
    <button class="profileCol saveButton">
      <span>Change</span>
    </button>
  </div>
</template>

<style>
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
}
.bio {
  width: 40vh;
  height: 20vh;
}
.saveButton {
  background-color: white;
  size: 1vh;
  height: 4vh;
  text-align: center;
  font-size: large;
  font-weight: bold;
}
.profileRow {
  display: flex;
  margin-bottom: 10vh;
}
.profileCol {
  flex: 50%;
  margin-left: 5vh;
}
</style>
