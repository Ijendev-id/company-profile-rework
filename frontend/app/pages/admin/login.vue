<script setup lang="ts">
import { admin } from '~/data/admin'

// Halaman login berdiri sendiri (tanpa menu admin).
definePageMeta({ layout: false })
useHead({
  title: `Masuk | ${admin.title}`,
  meta: [{ name: 'robots', content: 'noindex, nofollow' }],
})

const route = useRoute()
const { login, fetchUser } = useAdminAuth()

const email = ref('')
const password = ref('')
const submitting = ref(false)
const problem = ref('')

// Tujuan setelah login. Hanya alamat di dalam /admin yang diterima (mencegah dialihkan ke situs lain).
const destination = computed(() => {
  const target = route.query.redirect
  return typeof target === 'string' && (target === '/admin' || target.startsWith('/admin/')) && !target.includes('//') && !target.includes('\\') ? target : '/admin/posts'
})

// Sudah punya sesi yang masih berlaku -> langsung masuk.
onMounted(async () => {
  if (await fetchUser()) await navigateTo(destination.value, { replace: true })
})

async function submit() {
  problem.value = ''
  submitting.value = true
  try {
    await login(email.value.trim(), password.value)
    await navigateTo(destination.value, { replace: true })
  } catch (error) {
    const apiError = toApiError(error)
    problem.value = apiError.status === 422 ? 'Isi email dan password dengan benar.' : apiError.message
    password.value = ''
  } finally {
    submitting.value = false
  }
}
</script>

<template>
  <div class="admin-login">
    <form class="admin-login-card" novalidate @submit.prevent="submit">
      <h1>{{ admin.title }}</h1>

      <AdminAlert v-if="route.query.expired && !problem" type="info">Sesi kamu telah berakhir. Silakan masuk lagi.</AdminAlert>
      <AdminAlert v-if="problem" type="danger">{{ problem }}</AdminAlert>

      <div class="admin-field">
        <label for="login-email">Email</label>
        <input id="login-email" v-model="email" type="email" class="form-control" autocomplete="username" required autofocus />
      </div>

      <div class="admin-field">
        <label for="login-password">Password</label>
        <input id="login-password" v-model="password" type="password" class="form-control" autocomplete="current-password" required />
      </div>

      <button type="submit" class="btn btn-primary w-100" :disabled="submitting">
        {{ submitting ? 'Memeriksa...' : 'Masuk' }}
      </button>
    </form>
  </div>
</template>
