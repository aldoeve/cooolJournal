<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import axios from "axios";
import "../assets/welcome.css";
import "../assets/util.css";

const router = useRouter();

const email = ref("");
const password = ref("");
const showPassword = ref(false);

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
    }
  }
  
  
}

function toggle() {
  showPassword.value = !showPassword.value;
}

function toggleButtonText() {
  return showPassword.value ? "Hide" : "Show";
}
</script>

<style scoped>
@import "../assets/welcome.css";
@import "../assets/util.css";
</style>

<template>
  <div class="wrapper">
    <div class="container">
      <img class="logo" src="../../public/LionProfilePic.jpg" alt="Logo" />

      <div class="credentials">
        <label>Email</label>
        <input
          class="textBox"
          v-model="email"
          @keydown.space.prevent
          placeholder="Email"
        />

        <label>Password</label>
        <div class="shiftPasswordBox">
          <input
            v-if="showPassword"
            class="textBox"
            v-model="password"
            @keydown.space.prevent
            placeholder="Password"
          />
          <input
            v-else
            class="textBox"
            type="password"
            v-model="password"
            @keydown.space.prevent
            placeholder="Password"
          />
          <button class="buttons toggleButton" @click="toggle" style="width: 15%">
            <span>{{ toggleButtonText() }}</span>
          </button>
        </div>

        <a href="#" style="color: #6d96a8"><u>Forgot Password?</u></a>
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
