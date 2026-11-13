<script setup lang="ts">
import { admin } from '~/data/admin'

// Halaman admin tidak boleh masuk hasil pencarian Google.
useHead({
  titleTemplate: (title) => (title ? `${title} | ${admin.title}` : admin.title),
  meta: [{ name: 'robots', content: 'noindex, nofollow' }],
})

// Sidebar di layar kecil: tersembunyi, dibuka dengan tombol menu.
const sidebarOpen = ref(false)
</script>

<template>
  <div class="admin-shell">
    <header class="admin-topbar">
      <button type="button" class="admin-menu-btn" aria-label="Buka menu" aria-controls="admin-sidebar" :aria-expanded="sidebarOpen" @click="sidebarOpen = !sidebarOpen">
        <span></span><span></span><span></span>
      </button>
      <span class="admin-brand">{{ admin.title }}</span>
    </header>

    <AdminSidebar :open="sidebarOpen" @close="sidebarOpen = false" />
    <div v-if="sidebarOpen" class="admin-backdrop" @click="sidebarOpen = false"></div>

    <main class="admin-main">
      <div class="admin-content">
        <slot />
      </div>
    </main>
  </div>
</template>
