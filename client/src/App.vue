<script setup lang="ts">
import { ref } from "vue"
const socket = new WebSocket("ws://localhost:8080")

socket.addEventListener("message", (event) => {
  console.log("Message from server!", event.data)
})

const message = ref("")

const sendMsg = () => {
  socket.send(message.value)
}
</script>

<template>
  <main>
    <input type="text" v-model="message">
    <button @click="sendMsg" />
  </main>
</template>

<style scoped></style>
