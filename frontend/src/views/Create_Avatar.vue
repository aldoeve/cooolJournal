<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import "../assets/welcome.css";
import axios from "axios";

const router = useRouter();
const image = ref(null);
const data = ref(null);

function gotoCreateUser() {
  router.push("/create/username");
}

function onFileChange(e) {
  const files = e.target.files || e.dataTransfer.files;
  if (!files.length) return;

  const file = files[0];
  image.value = URL.createObjectURL(file);
}

async function createUser() {
  try {
    const response = await axios.post("/api/createuser", {
      userstats: image.value,
    });
    if (response.status !== 200) {
      throw new Error(response.status);
    }
    data.value = response.data;
  } catch (error) {
    console.error("Error:", error.message);
    return;
  }
  router.push("/home");
}
</script>

<style scoped>
@import "../assets/welcome.css";
</style>

<template>
  <div class="wrapper">
    <div class="container">
      <img class="logo" src="../../public/LionProfilePic.jpg" />

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
        <button class="transistion-buttons" @click="createUser">
          <slot>Finish</slot>
        </button>
      </div>
    </div>
    <div class="enlargeMsg"><span>Please enlarge the window.</span></div>
  </div>
</template>
