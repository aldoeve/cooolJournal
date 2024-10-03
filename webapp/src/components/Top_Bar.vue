<script setup>
  import { ref, onMounted, onBeforeUnmount } from 'vue';
  import { useRouter } from 'vue-router';
  import Side_Bar from './Side_Bar.vue';

  const isSmallScreen = ref(false);
  const sidebarTriggered = ref(false);
  const router = useRouter();

  const checkScreenSize = () => {
    isSmallScreen.value = window.innerWidth < 1200;

    if (!isSmallScreen.value) {
      sidebarTriggered.value = false;
    }
  };

  onMounted(() => {
      checkScreenSize(); 
    window.addEventListener('resize', checkScreenSize);
  });

  onBeforeUnmount(() => {
    window.removeEventListener('resize', checkScreenSize);
  });


  const sideToggle = () => {
    sidebarTriggered.value = !sidebarTriggered.value;
  };

  function gotoHome() {
    router.push('/home');
  }

  function gotoProfile() {
    router.push('/profile');
  }
</script>


<template>
      <div class="top-bar">

        <div v-if="isSmallScreen" class="left-button-container">
          <button @click="sideToggle" class="side-bar-button" style="margin-right: 16px;"><i class="fa fa-bars fa-lg"></i></button>
          <button class="logo-button" @click="gotoHome">
            <img class="logo" src="https://i.ibb.co/CW5Wvry/buttonpng.png" alt="buttonpng"/>
          </button>
        </div>

        <div v-else>
          <button class="logo-button" @click="gotoHome">
            <img class="logo" src="https://i.ibb.co/CW5Wvry/buttonpng.png" alt="buttonpng"/>
          </button>
        </div>
        

        <div class="search-box">
          <i class="fa fa-search fa-lg"></i>
          <input class="search-bar" type="text" placeholder="Search">
          
        </div>

        <div class="right-button-container">
          <button @click="gotoProfile" class="profile-button"  style="margin-right: 16px;"><i class="fa fa-user-o fa-lg"></i></button>
          <button class="profile-button"><i class="fa fa-cog fa-lg"></i></button>
        </div>
    </div>
    

    <Side_Bar v-if="sidebarTriggered" :sidebarTriggered="sidebarTriggered"></Side_Bar>
      
    
    
  </template>
  
<style scoped>
  @import '../assets/font-awesome-4.7.0/css/font-awesome.min.css';


  
.top-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100vw;
  height: 76px;
  background-color: #6d96a8;
  border-bottom: 1px solid black;
  padding: 16px;
}

.search-box {
  position: relative;
  display: flex;
  align-items: center;
  width: 640px;
  min-width: 180px;
  height: 48px;
  border-radius: 112px;
  margin-left: 16px;
  margin-right: 16px;
  border: 1px solid black;
  box-sizing: border-box;
  background-color: white;
}

.search-bar {
  padding-left: 50px;
  width: 100%;
  min-width: 100%;
  height: 100%;
  border: none;
  outline: none;
  border-radius: 112px;
}

.search-box .fa-search {
  position: absolute;
  left: 20px;
  color: black
}

.left-button-container {
  display: flex;
  align-items: center;
  width: fit-content;
  min-width: fit-content;
}

.right-button-container {
  display: flex;
  align-items: center;
  width: fit-content;
  min-width: fit-content;
}

.side-bar-button {
  width: 48px;
  min-width: 48px;
  height: 48px;
  border: none;
  cursor: pointer;
  border: none;
  border-radius: 100px;
  background-color: transparent;
  color:white;

}

.side-bar-button:hover {
  background-color: #4e6e7b;
}


.logo {
  width: 100%;
  height: 100%;
  object-fit: contain;
}

.logo-button {

  width: 48px;
  min-width: 48px;
  height: 48px;
  border: none;
  background-color: transparent;
  cursor: pointer;
  padding: 0;
}



.profile-button {
  width: 64px;
  min-width: 64px;
  height: 48px;
  border-radius: 112px;
  border: 1px solid black;
  font-size: 16px;
  padding: 8px;
  color: white;
  cursor: pointer;
  background-color: #6d96a8;
}

.profile-button:hover {
  background-color: #4e6e7b;
}
</style>
  
  
  