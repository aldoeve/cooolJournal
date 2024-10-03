<template>
    <div class="popup-wrapper">
        <div class="popup-window">
            <div class="exit-button-container">
                <button @click="togglePopup" class="exit-button"><i class="fa fa-times fa-2x"></i></button>
            </div>
            
        </div>    
    </div>
    
   
</template>

<script setup>
    const props = defineProps({
        togglePopup: Function
    });
</script>

<style scoped>
@import '../assets/font-awesome-4.7.0/css/font-awesome.min.css';

.popup-wrapper {
    position: fixed;
    inset: 0;
    background: rgba(0, 0, 0, 0.3);
    margin-top: 76px;
}

.popup-window {
    position: fixed; 
    top: 50%;        
    left: 50%;       
    transform: translate(-50%, -50%);
    background-color: white; 
    margin-top: 76px;
    width: 50vw;
    height: 100vh;
    padding: 20px;
    
}

.exit-button-container {
    display: flex;
    justify-content: flex-end;
}

.exit-button {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 50px;
    height: 50px;
    background-color: white;
    border:none;
    color: black;
    border-radius: 100px;
}

.exit-button:hover {
    background-color: rgb(224, 224, 224);
}
</style>