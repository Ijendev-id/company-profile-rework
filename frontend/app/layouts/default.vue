<script setup lang="ts">
// WOW.js cukup dijalankan sekali di sini. Elemen ber-class "wow" di halaman
// mana pun (termasuk setelah pindah halaman) ikut teranimasi otomatis.
useWow()
</script>

<template>
  <div>
    <AppPreloader />
    <slot />
    <TheFooter />
    <BackToTop />
  </div>
</template>
