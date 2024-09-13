<template>
    <div class="container">
        <img class="logo" src="../assets/temp.png">
        <div class = "credentials">
            <label>Enter a Username</label>
            <input class="textBox" placeholder="Username"/>
            <label>Enter a Bio</label>
            <textarea rows="4" cols="50" placeholder="Bio" class="comment-box"></textarea>
        </div>
            <div class="transistion-buttons-container" style="justify-content: flex-end;">
                <button @click="gotoCreateAvatar" class="transistion-buttons"><slot>Next</slot></button>
            </div>
    </div>
</template>

<script setup>
import {ref} from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();
function gotoCreateAvatar() {
    router.push('/create/avatar');
}
</script>