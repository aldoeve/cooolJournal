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
const data = ref(null);

function gotoSignUp() {
  router.push("/signup");
}

async function gotoHome() {
  try {
    const response = await axios.post("/api/loginUser", {
      enteredUser: email.value,
      enteredPass: password.value,
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
