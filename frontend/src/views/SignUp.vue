<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import '../assets/welcome.css';

const router = useRouter();
const email = ref('');
const password = ref('');

function gotoLogin(){
  router.push('/');
};

function gotoCreateUser(){
  router.push('/create/username');
};
</script>

<style scoped>
  @import '../assets/welcome.css';
</style>

<template>
    <div class="wrapper">
      <div class = "container" >
        <img class="logo" src="../../public/LionProfilePic.jpg">
        <div class="credentials">

          <label>Email</label>
          <input class = "textBox" placeholder="Email" v-model="email"
          @keydown.space.prevent></input>
          <label>Password</label>
          <input class = "textBox" placeholder="Password" v-model="password"
          @keydown.space.prevent></input>
          <input id="checkbox" type="checkbox" />
          <label for="checkbox"> I agree to these <a href="#" style="color: #6d96a8"><u>Terms and Conditions</u></a>.</label>
        </div>
          <div class="account-buttons-container">
            <button class = "buttons" @click="gotoCreateUser">
                <slot>Create Account</slot>
            </button>
            <button class = "buttons" @click="gotoLogin">
                <slot>Back to Login</slot>
            </button>
          </div>
      </div>
      <div class="enlargeMsg"><span>Please enlarge the window.</span></div>
    </div>
  </template>