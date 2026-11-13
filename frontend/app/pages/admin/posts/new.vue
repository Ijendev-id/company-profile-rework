<script setup lang="ts">
import type { PostPayload } from '~/types/admin'

definePageMeta({ layout: 'admin', middleware: 'auth' })
useHead({ title: 'Tulis artikel' })

const api = useAdminApi()
const flash = useFlash()
const { errors, message, clear, apply } = useFormErrors()
const submitting = ref(false)

async function save(payload: PostPayload) {
  clear()
  submitting.value = true
  try {
    await api.post('/posts', payload)
    flash.set('Artikel berhasil dibuat.')
    await navigateTo('/admin/posts')
  } catch (error) {
    apply(error)
    window.scrollTo({ top: 0, behavior: 'smooth' })
  } finally {
    submitting.value = false
  }
}
</script>

<template>
  <div>
    <div class="admin-page-head">
      <h1>Tulis artikel</h1>
    </div>

    <AdminAlert v-if="message" type="danger">{{ message }}</AdminAlert>
    <AdminPostForm :submitting="submitting" :errors="errors" @submit="save" />
  </div>
</template>
