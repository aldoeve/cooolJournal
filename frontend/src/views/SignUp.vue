<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import axios from "axios";
import "../assets/welcome.css";
import "../assets/util.css";

const router = useRouter();
const email = ref('');  
const password = ref(''); 

function gotoLogin(){
  router.push('/login');
};

async function gotoCreateUser(){
  try {
    const response = await axios.post("/api/createUser", {
      enteredEmail: email.value,
      enteredPass: password.value,
    });

    if (response.data.userExists[0] === "false") {
      router.push("create/username");
      
    }
  } catch (error) {
    if (error.response && error.response.status === 409) {
      console.error("Error:", error.response.data.error[0]);
    }
    
  }
  
};

</script>


<template>
    <div class="wrapper">
      <div class="container">
        <img class="logo" src="../../public/LionProfilePic.jpg">
        <div class="credentials">
          <label>Email</label>
          <input class="textBox" placeholder="Email" v-model="email" @keydown.space.prevent />
          <label>Password</label>
          <input class="textBox" placeholder="Password" v-model="password" @keydown.space.prevent />
          <input id="checkbox" type="checkbox" />
          <label for="checkbox"> I agree to these <a href="#" style="color: #6d96a8"><u>Terms and Conditions</u></a>.</label>
        </div>
        <div class="account-buttons-container">
          <button class="buttons" @click="gotoCreateUser">Create Account</button>
          <button class="buttons" @click="gotoLogin">Back to Login</button>
        </div>
      </div>
      <div class="enlargeMsg"><span>Please enlarge the window.</span></div>
    </div>
</template>

<style scoped>
  @import '../assets/welcome.css';
</style>
