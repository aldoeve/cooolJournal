<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import "../assets/welcome.css";

const router = useRouter();
const image = ref(null);

function gotoCreateUser() {
  router.push("/create/username");
}

function onFileChange(e) {
  const files = e.target.files || e.dataTransfer.files;
  if (!files.length) return;

  const file = files[0];
  image.value = URL.createObjectURL(file);
}
</script>

<style scoped>
@import "../assets/welcome.css";
</style>

<template>
  <div class="wrapper">
    <div class="container">
      <img class="logo" src="../assets/temp.png" />

      <div class="credentials">
        <div v-if="image">
          <img :src="image" class="preview" />
        </div>

        <div class="file-upload-area">
          <input type="file" accept="image/*" @change="onFileChange" />
        </div>
      </div>

      <div class="transistion-buttons-container">
        <button @click="gotoCreateUser" class="transistion-buttons">
          <slot>Back</slot>
        </button>
        <button class="transistion-buttons"><slot>Finish</slot></button>
      </div>
    </div>
    <div class="enlargeMsg"><span>Please enlarge the window.</span></div>
  </div>
</template>
