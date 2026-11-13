<script setup lang="ts">
// Judul section dengan garis kecil di atasnya. `highlight` tampil dengan warna aksen.
defineProps<{
  title: string
  highlight?: string
  center?: boolean
}>()
</script>

<template>
  <div class="section-title" :class="{ 'text-center pb-40': center }">
    <div class="line" :class="{ 'm-auto': center }"></div>
    <h3 class="title">{{ title }} <span v-if="highlight">{{ highlight }}</span></h3>
  </div>
</template>
