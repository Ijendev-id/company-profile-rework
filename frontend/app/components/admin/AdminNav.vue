<script setup lang="ts">
import { admin } from '~/data/admin'

const route = useRoute()
const { user, logout } = useAdminAuth()

const isActive = (to: string) => route.path === to || route.path.startsWith(`${to}/`)
</script>

<template>
  <header class="admin-nav">
    <div class="container admin-nav-inner">
      <NuxtLink to="/admin/posts" class="admin-brand">{{ admin.title }}</NuxtLink>

      <nav class="admin-links" aria-label="Menu admin">
        <NuxtLink v-for="item in admin.nav" :key="item.to" :to="item.to" :class="{ active: isActive(item.to) }">
          {{ item.label }}
        </NuxtLink>
        <a href="/blog" target="_blank" rel="noopener">Lihat situs</a>
      </nav>

      <div class="admin-user">
        <span v-if="user" class="admin-user-name">{{ user.name }}</span>
        <button type="button" class="btn btn-sm btn-outline-secondary" @click="logout()">Keluar</button>
      </div>
    </div>
  </header>
</template>
