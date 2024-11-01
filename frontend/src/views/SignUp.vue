<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import axios from "axios";
import InputText from 'primevue/inputtext';
import Password from 'primevue/password';
import "../assets/welcome.css";
import "../assets/util.css";

const router = useRouter();
const email = ref('');  
const password = ref(''); 
const repeatPassword = ref('');
const passwordError = ref('');


async function gotoCreateUser(){

  if(!/^[^@]+@\w+(\.\w+)+\w$/.test(email.value)) {
    passwordError.value = 'Invalid Email format.';
    return;
  }
  
  if(password.value.length < 8) {
    passwordError.value = 'Password must be at least 8 characters long.';
    return;
  }
  if(password.value === password.value.toLowerCase()) {
    passwordError.value = 'Password must include at least one uppercase letter.';
    return;
  }
  if(password.value === password.value.toLowerCase()) {
    passwordError.value = 'Password must include at least one lowercase letter.';
    return;
  }
  if(!/\d/.test(password.value)) {
    passwordError.value = 'Password must include at least one number (0-9).';
    return;
  }
  if (!/[!@#$%^&*(),.?":{}|<>]/.test(password.value)) {
    passwordError.value = 'Password must include at least one special character (!, @, #, $, %, ?).';
    return;
  }
  if (password.value !== repeatPassword.value) {
    passwordError.value = 'Passwords do not match.'; 
    return;
  }

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
      passwordError.value = 'Email has already been taken.'
      
    }
    
  }
  
};

</script>


<template>
    <div class="wrapper">
      <div class="container">
        <img class="logo" src="../../public/logo.png">
        <div class="credentials">
          <label>Email</label>
          <InputText id="email" v-model="email" @keydown.space.prevent required style="margin-bottom: 2vh;" autocomplete="off"/>
          
          <label>Password</label>
          <Password v-model="password" @keydown.space.prevent promptLabel="Choose a password" weakLabel="Too simple" mediumLabel="Average complexity" strongLabel="Complex" style="margin-bottom: 2vh;" />
          <label>Confirm Password</label>
          <Password v-model="repeatPassword" @keydown.space.prevent :feedback="false" style="margin-bottom: 2vh;" />

          <p style="font-size: small;">By signing up, you agree to the<a href="#" class="no-highlight">Terms of Service</a>and<a href="#"class="no-highlight">Privacy Policy</a>, including<a href="#"class="no-highlight">Cookie Use.</a></p>
          
          <div class="error-container">
            <div v-if="passwordError" class="error">{{ passwordError }}</div>
          </div>

        </div>
        <div class="account-buttons-container">
          <button class="buttons" @click="gotoCreateUser">Create Account</button>
        </div>
      </div>
      <div class="enlargeMsg"><span>Please enlarge the window.</span></div>
    </div>
</template>

<style scoped>
  @import '../assets/welcome.css';
</style>
