<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import axios from "axios";
import InputText from 'primevue/inputtext';
import Password from 'primevue/password';
import "../assets/welcome.css";
import "../assets/util.css";

const router = useRouter();
const email = ref("");
const password = ref("");
const errorMessage = ref("");

function gotoSignUp() {
  router.push("/create");
}

async function gotoHome() {

  try {
    const response = await axios.post("/api/loginUser", {
      enteredEmail: email.value,
      enteredPass: password.value,
    });
    if(response.data.userExists[0] === "true") {
      router.push("/");
    }
  } catch (error) {
    if (error.response && error.response.status === 409) {
      console.error("Error:", error.response.data.error[0]);
      errorMessage.value = "Invalid username or password."
      return;
    }
  }
  
  
}

</script>

<style scoped>
@import "../assets/welcome.css";
@import "../assets/util.css";
</style>

<template>
  <div class="wrapper">
    <div class="container">
      <img class="logo" src="../../public/logo.png" alt="Logo" />

      <div class="credentials">
        <label>Email</label>
        <InputText id="email" v-model="email" @keydown.space.prevent required style="margin-bottom: 2vh;" autocomplete="off"/>

        <label>Password</label>
        <Password v-model="password" @keydown.space.prevent :feedback="false" style="margin-bottom: 2vh;" />

        <p style="font-size: small;"><a href="#" class="no-highlight">Forgot Password?</a></p>

        <div class="error-container">
            <div v-if="errorMessage" class="error">{{ errorMessage }}</div>
          </div>

      </div>
      <div class="account-buttons-container">
        <button class="buttons" @click="gotoHome">Log In</button>
        <button class="buttons" @click="gotoSignUp">Sign Up</button>
      </div>
    </div>
    <div class="enlargeMsg">
      <span>Please expand the window.</span>
    </div>
  </div>
</template>
