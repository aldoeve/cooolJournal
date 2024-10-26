<template>
  <div class="wrapper" v-if="isUserVerified">
    <div class="container">
      <img class="logo" src="../../public/LionProfilePic.jpg" />
      <div class="credentials">
        <label>Enter a Username</label>
        <input class="textBox" placeholder="Username" v-model="username"/>
        <label>Enter a Bio</label>
        <textarea rows="4" cols="50" placeholder="Bio" class="comment-box" v-model="bio"></textarea>
      </div>
      <div class="transistion-buttons-container" style="justify-content: flex-end">
        <button @click="gotoCreateAvatar" class="transistion-buttons">
          <slot>Next</slot>
        </button>
      </div>
    </div>
    <div class="enlargeMsg"><span>Please enlarge the window.</span></div>
  </div>
</template>

<script setup>
import { useRouter } from "vue-router";
import "../assets/welcome.css";
import { ref, onMounted } from "vue";
import axios from "axios";

const bio = ref('');
const username = ref('');
const usernameResponse = ref(null);
const bioResponse = ref(null);
const router = useRouter();
const isUserVerified = ref(false);
const email = ref(null);
const isUsernameValid = ref(true);
const isBioValid = ref(true);

onMounted(() => {
  verifyUser();
});

async function verifyUser() {
  try {
    const verifyResponse = await axios.post("/api/verifyUser", {});
    if (verifyResponse.data.verified[0] === "true") {
      isUserVerified.value = true;
    } else {
      router.push("/login");
    }
  } catch (error) {
    console.error("Verification failed:", error.message);
    router.push("/login");
  }
};

async function updateUsername() {
  try {
    const response = await axios.put("/api/updateUsername", {
      email: email.value,
      enteredUsername: username.value,
    });
    
  } catch(error) {
    console.error("Error:", error.message);
  }
}

async function updateBio() {
  try {
    const response = await axios.post("/api/updateBio", {
      email: email.value,
      enteredBio: bio.value,
    });
    response.data;
  } catch(error) {
    console.error("Error:", error.message);
  }
}


async function gotoCreateAvatar() {

  if (isBioValid && isUsernameValid) {
    try {
      const response = await axios.post("/api/retrieveUser", {});
      if(response.data.verified[0] === "true") {
        email.value = response.data.email[0];
        updateUsername();
        updateBio();
        router.push("/create/avatar");
      }
    } catch(error) {
      console.error("Error: ", error.message);

    }
    
  }

};

</script>

<style scoped>
@import "../assets/welcome.css";
</style>
