<script setup lang="ts">
// Sama seperti Route::fallback di Laravel: URL yang tidak dikenal (login, dashboard, dll)
// diarahkan ke landing page. Route yang lebih spesifik (mis. /blog/...) selalu diprioritaskan.
definePageMeta({
  middleware: () => navigateTo('/', { redirectCode: 302 }),
})
</script>

<template>
  <div></div>
</template>
