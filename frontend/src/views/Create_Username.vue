<template>
  <div class="wrapper" v-if="isUserVerified">
    <div class="container">
      <img class="logo" src="../../public/logo.png" />
      <div class="credentials">
        <label>Enter a Username</label>
        <InputText id="username" v-model="username" @keydown.space.prevent required style="margin-bottom: 2vh;" autocomplete="off"/>

        <label>Enter a Bio</label>
        <Textarea rows="4" cols="50" v-model="bio" style="resize: none;" :maxlength="maxBioLength" ></Textarea>
        <p>{{ maxBioLength - bio.length }} characters remaining.</p>
      </div>

      <div class="error-container">
        <div v-if="errorMessage" class="error">{{ errorMessage }}</div>
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
import InputText from 'primevue/inputtext';
import Textarea from "primevue/textarea";

const bio = ref('');
const username = ref('');
const router = useRouter();
const isUserVerified = ref(false);
const email = ref(null);
const isUsernameValid = ref(true);
const isBioValid = ref(true);
const errorMessage = ref("");
const maxBioLength = ref(300);
const maxUsernameLength = ref(20);

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

  if(username.value.length < 5) {
    errorMessage.value = "Username must be at least 5 characters long."
    return false;
  }

  if(username.value.length > 20) {
    errorMessage.value = "Username must be at most 20 characters long."
    return false;
  }

  try {
    const response = await axios.put("/api/updateUsername", {
      email: email.value,
      enteredUsername: username.value,
    });
    return true;
  } catch(error) {
    console.error("Error:", error.message);
    errorMessage.value = "Could not update Username."
    return false;
  }
}

async function updateBio() {
  try {
    const response = await axios.post("/api/updateBio", {
      email: email.value,
      enteredBio: bio.value,
    });
    return true;
  } catch(error) {
    console.error("Error:", error.message);
    errorMessage.value = "Could not update Bio."
    return false;
  }
}


//Need to fix*********
//Make each routes authenticate the user
async function gotoCreateAvatar() {

  if (isBioValid && isUsernameValid) {
    try {
      const response = await axios.post("/api/retrieveUser", {});
      if(response.data.verified[0] === "true") {
        email.value = response.data.email[0];
        
        const usernameUpdated = await updateUsername();
        const bioUpdated = await updateBio();

        

        if (usernameUpdated && bioUpdated) {
          router.push("/create/avatar");
        }
      }
    } catch(error) {
      console.error("Error: ", error.message);
      errorMessage.value = "Could not find User."
      return;
    }
    
  }

};

</script>

<style scoped>
@import "../assets/welcome.css";
</style>
