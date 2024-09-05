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

    const gotoSignUp = () => {
      router.push('/signup');
    };

    return { username, password, showPassword, toggle, toggleButtonText, gotoSignUp };
  }
};
</script>

<template>
  <div class="container">
    <input
      class="textBox"
      v-model="username"
      @keydown.space.prevent
      placeholder="Username"
    />
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
      <button class="buttons toggleButton" @click="toggle">
        <span> {{ toggleButtonText() }} </span>
      </button>
    </div>
    <button class="buttons"><slot>Log In</slot></button>
    <button class="buttons" @click="gotoSignUp"><slot>Sign Up</slot></button>
  </div>
</template>
