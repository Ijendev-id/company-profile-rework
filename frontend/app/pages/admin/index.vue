<script setup lang="ts">
// /admin langsung membuka daftar artikel.
definePageMeta({ middleware: () => navigateTo('/admin/posts', { replace: true }) })
</script>

<template>
  <div></div>
</template>
