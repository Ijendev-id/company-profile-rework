<script setup lang="ts">
// Kotak pesan di halaman admin. type: success | danger | info
withDefaults(defineProps<{ type?: 'success' | 'danger' | 'info' }>(), { type: 'info' })
</script>

<template>
  <div class="alert" :class="`alert-${type}`" role="alert">
    <slot />
  </div>
</template>
