<template>
  <div class="landing-wrapper" v-if="isUserVerified">
    <Top_Bar></Top_Bar>
    <div class="main-content">
      <Side_Bar></Side_Bar>
      <Content></Content>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import axios from 'axios';
import Top_Bar from '@/components/Top_Bar.vue';
import Side_Bar from '@/components/Side_Bar.vue';
import Content from '@/components/Content.vue';

const isUserVerified = ref(false);
const router = useRouter();

async function verifyUser() {
  try {
    const verifyResponse = await axios.post("/api/verifyUser", {});
    if (verifyResponse.data.verified[0] === "true") {
      isUserVerified.value = true;
    } else {
      router.push("/login");
    }
  } catch (error) {
    console.error("Verification failed:", error.message);
    router.push("/login");
  }
}

onMounted(() => {
  verifyUser();
});
</script>

<style scoped>
@import '../assets/landing.css';

</style>
