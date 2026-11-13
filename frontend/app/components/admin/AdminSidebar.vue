<script setup lang="ts">
import { admin } from '~/data/admin'

defineProps<{ open: boolean }>()
const emit = defineEmits<{ close: [] }>()

const route = useRoute()
const { user, logout } = useAdminAuth()

const isActive = (to: string) => route.path === to || route.path.startsWith(`${to}/`)

// Pindah halaman -> tutup sidebar (di layar kecil).
watch(() => route.path, () => emit('close'))
</script>

<template>
  <aside id="admin-sidebar" class="admin-sidebar" :class="{ open }" aria-label="Menu admin">
    <NuxtLink to="/admin/posts" class="admin-brand">{{ admin.title }}</NuxtLink>

    <nav class="admin-side-nav">
      <div v-for="group in admin.nav" :key="group.title" class="admin-side-group">
        <div class="admin-side-title">{{ group.title }}</div>
        <NuxtLink v-for="item in group.items" :key="item.to" :to="item.to" :class="{ active: isActive(item.to) }">
          <i v-if="item.icon" :class="['lni', item.icon]" aria-hidden="true"></i>
          {{ item.label }}
        </NuxtLink>
      </div>
    </nav>

    <div class="admin-side-footer">
      <a href="/blog" target="_blank" rel="noopener">
        <i class="lni lni-link" aria-hidden="true"></i>
        Lihat situs
      </a>
      <div class="admin-user">
        <span v-if="user" class="admin-user-name">{{ user.name }}</span>
        <button type="button" class="btn btn-sm btn-outline-secondary" @click="logout()">Keluar</button>
      </div>
    </div>
  </aside>
</template>
