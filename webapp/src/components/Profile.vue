<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
</script>

<template>
  <div class="profileRow">
    <div class="profileCol">
      <div class="circleBase circleImage">
        <img src="../../LionProfilePic.jpg" />
      </div>
      <button class="profileCol saveButton"><span>Save</span></button>
    </div>
  </div>
  <div class="profileRow"></div>
  <div class="profileRow"></div>
</template>

<style>
.circleBase {
  background-color: white;
  border-radius: 50%;
  width: 40vh;
  height: 40vh;
}
.userName {
  background-color: white;
  width: 40vh;
  height: 40vh;
}
.saveButton {
  background-color: white;
  size: 1vh;
}
.profileRow {
  display: flex;
}
.profileCol {
  flex: 50%;
}
.saveButtonPosition {
  position: inherit;
  bottom: 0;
  left: 0;
}
</style>
