<script>
import { ref } from 'vue';
import { useRouter } from 'vue-router';

export default {
  setup() {
    const router = useRouter();
    const username = ref('');
    const password = ref('');
    const showPassword = ref(false);

    const toggle = () => {
      showPassword.value = !showPassword.value;
    };

    const toggleButtonText = () => {
      return showPassword.value ? 'Hide' : 'Show';
    };
    

    const gotoLogin = () => {
      router.push('/');
    };

    const gotoCreateUser = () => {
      router.push('/create/username');
    };

    return { username, password, showPassword, toggle, toggleButtonText, gotoLogin, gotoCreateUser };
  }
};
</script>

<template>
    <body>
      <div class = "container" >
        <img class="logo" src="../assets/temp.png">
        <div class="credentials">

          <label>Email</label>
          <input class = "textBox" placeholder="Email"></input>
          <label>Password</label>
          <input class = "textBox" placeholder="Password"></input>
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
    </body>
  </template>