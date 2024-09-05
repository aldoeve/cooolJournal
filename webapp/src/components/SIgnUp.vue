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

    return { username, password, showPassword, toggle, toggleButtonText, gotoLogin };
  }
};
</script>

<template>
    <body>
      <div class = "container" >
        
        <input class = "textBox" placeholder="Username">
        </input>

        <input class = "textBox" placeholder="Password">
        </input>


        <button class = "buttons" @click="gotoLogin">
            <slot>
                Sign Up
            </slot>
        </button>
      </div>
    </body>
  </template>