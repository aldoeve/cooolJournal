<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const router   = useRouter();

const email = ref('');
const password = ref('');
const showPassword = ref(false);

function gotoSignUp() {
  router.push('/signup');
}
function toggle() {
  showPassword.value = !showPassword.value;
}
function toggleButtonText() {
  return showPassword.value ? 'Hide':'Show';
}
</script>

<template>
  <body>
  <div class="container">
    <img class="logo" src="../assets/temp.png">

    <div class="credentials">
      <label>Email</label>
      <input class="textBox" v-model="email" @keydown.space.prevent placeholder="Email"/>
      <label>Password</label>
      
      <div class="shiftPasswordBox">
        <input v-if="showPassword" class="textBox" v-model="password" @keydown.space.prevent placeholder="Password"/>
        <input v-else class="textBox" type="password" v-model="password" @keydown.space.prevent placeholder="Password" />
        <button class="buttons toggleButton" @click="toggle" style="width:15%">
          <span> {{ toggleButtonText() }} </span>
        </button>
      </div>
      <a href="#" style="color: #6d96a8;"><u>Forgot Password?</u></a>
    </div>
      <div class="account-buttons-container">
        <button class="buttons"><slot>Log In</slot></button>
        <button class="buttons" @click="gotoSignUp"><slot>Sign Up</slot></button>
      </div>
    
  </div>
</body>
</template>
