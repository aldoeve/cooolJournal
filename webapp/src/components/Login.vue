<script>
import { ref } from "vue";

export default {
  setup() {
    const username = ref("");
    const password = ref("");
    const showPassword = ref(false);
    return { username, password, showPassword };
  },
  methods: {
    toggle() {
      this.showPassword = !this.showPassword;
    },
    toggleButtonText() {
      return this.showPassword ? "Hide" : "Show";
    },
  },
};
</script>

<template>
  <body>
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
      <button class="buttons"><slot>Sign Up</slot></button>
    </div>
  </body>
</template>