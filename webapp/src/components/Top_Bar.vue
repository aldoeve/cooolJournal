<script setup>
    import { ref } from 'vue';
    import { useRouter } from 'vue-router';
    
    const router = useRouter();
    
    function gotoHome() {
      router.push('/home');
    }

    function gotoProfile() {
      router.push('/profile');
    }
    
</script>

<template>
      <div class="top-bar">
        
        <button class="logo-button" @click="gotoHome">
          <img class="logo" src="https://i.ibb.co/CW5Wvry/buttonpng.png" alt="buttonpng"/>
        </button>

        <div class="search-box">
          <i class="fa fa-search fa-lg"></i>
          <input class="search-bar" type="text" placeholder="Search">
          
        </div>

        <div class="top-button-container">
          <button @click="gotoProfile" class="profile-button"><i class="fa fa-user-o fa-lg"></i></button>
          <button class="profile-button"><i class="fa fa-cog fa-lg"></i></button>
        </div>
    </div>
    
  </template>
  
<style scoped>
  @import '../assets/top_bar.css';
  @import '../assets/font-awesome-4.7.0/css/font-awesome.min.css';
</style>
  
  
  