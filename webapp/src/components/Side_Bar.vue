<template>
    <div class="side-bar-container" v-if="!isSmallScreen">
        <button class="side-button"><i class="fa fa-home fa-2x"></i><i>Home</i></button>
        <button class="side-button"><i class="fa fa-line-chart fa-2x"></i><i>Trending</i></button>
        <hr></hr>
        <button class="side-button" @click="togglePopup"><i>All Topics</i></button>
        <Topic_Popup v-if="popupTriggered" :togglePopup="togglePopup"></Topic_Popup>
        <button class="side-button"><i>A Buttton</i></button>
        <button class="side-button"><i>A Buttton</i></button>
        <button class="side-button"><i>A Buttton</i></button>
        <button class="side-button"><i>A Buttton</i></button>
        <button class="side-button"><i>A Buttton</i></button>
        <button class="side-button"><i>A Buttton</i></button>
        <button class="side-button"><i>A Buttton</i></button>
        <button class="side-button"><i>A Buttton</i></button>
        <button class="side-button"><i>A Buttton</i></button>
        <button class="side-button"><i>A Buttton</i></button>
        <button class="side-button"><i>A Buttton</i></button>
        <button class="side-button"><i>A Buttton</i></button>
        <button class="side-button"><i>A Buttton</i></button>
    </div>



    <div class="side-bar-container" v-if="isSmallScreen && props.sidebarTriggered" style="position: fixed; margin-top: 76px;">
        <button class="side-button"><i class="fa fa-home fa-2x"></i><i>Home</i></button>
        <button class="side-button"><i class="fa fa-line-chart fa-2x"></i><i>Trending</i></button>
        <hr></hr>
        <button class="side-button" @click="togglePopup"><i>All Topics</i></button>
        <Topic_Popup v-if="popupTriggered" :togglePopup="togglePopup"></Topic_Popup>
        <button class="side-button"><i>A Buttton</i></button>
        <button class="side-button"><i>A Buttton</i></button>
        <button class="side-button"><i>A Buttton</i></button>
        <button class="side-button"><i>A Buttton</i></button>
        <button class="side-button"><i>A Buttton</i></button>
        <button class="side-button"><i>A Buttton</i></button>
        <button class="side-button"><i>A Buttton</i></button>
        <button class="side-button"><i>A Buttton</i></button>
        <button class="side-button"><i>A Buttton</i></button>
        <button class="side-button"><i>A Buttton</i></button>
        <button class="side-button"><i>A Buttton</i></button>
        <button class="side-button"><i>A Buttton</i></button>
        <button class="side-button"><i>A Buttton</i></button>
        
        


    </div>
</template>

<script setup>
  import Topic_Popup from './Topic_Popup.vue';
  import { ref, onBeforeUnmount, onMounted } from 'vue';

  const popupTriggered = ref(false);
  const isSmallScreen = ref(true);


  const togglePopup = () => {
    popupTriggered.value = !popupTriggered.value;
  }

  const checkScreenSize = () => {
    isSmallScreen.value = window.innerWidth < 1200;
    };

    onMounted(() => {
        checkScreenSize();
      window.addEventListener('resize', checkScreenSize);
      
      
    });

    onBeforeUnmount(() => {
      window.removeEventListener('resize', checkScreenSize);
    });

    const props = defineProps({
      sidebarTriggered: Boolean
    });

</script>

<style scoped>
  @import '../assets/font-awesome-4.7.0/css/font-awesome.min.css';


  .side-bar-container {
    background-color: white;
    width: 320px;
    min-width: 320px;
    height: 100%;
    border-right: 1px solid black;
    display: flex;
    flex-direction: column;
    align-items: center;
    padding-top: 16px;
    overflow: hidden;
  }

  .side-bar-container:hover {
    overflow-y: scroll;
    scrollbar-width: thin;
  }

.side-button {
    flex-shrink: 0;
    display:flex;
    justify-content: flex-start;
    align-items: center;        
    box-sizing: border-box;
    border-radius: 5px;
    width: 90%;
    min-width: 90%;
    height: 48px;
    border:none;
    background-color: white;
    cursor: pointer;

} 
.side-button:hover {
    background-color: rgb(224, 224, 224);
}

hr {
    margin-top: 16px;
    margin-bottom: 16px;
    border: 0;
    clear:both;
    display:block;
    width: 90%;               
    background-color:black;
    height: 1px;
    flex-shrink: 0;
  }

</style>


