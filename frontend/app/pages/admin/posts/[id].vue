<script setup lang="ts">
import type { AdminPost, PostPayload } from '~/types/admin'

definePageMeta({ layout: 'admin', middleware: 'auth' })
useHead({ title: 'Ubah artikel' })

const route = useRoute()
const api = useAdminApi()
const flash = useFlash()
const { errors, message, clear, apply } = useFormErrors()
const submitting = ref(false)

const id = Number(route.params.id)
const validId = Number.isInteger(id) && id > 0

const { data, error } = useAsyncData(`admin-post-${id}`, () =>
  validId ? api.get<{ data: AdminPost }>(`/posts/${id}`) : Promise.reject(new ApiError(404, 'Artikel tidak ditemukan.')),
)

const loadError = computed(() => (error.value ? toApiError(error.value) : null))

async function save(payload: PostPayload) {
  clear()
  submitting.value = true
  try {
    await api.put(`/posts/${id}`, payload)
    flash.set('Perubahan berhasil disimpan.')
    await navigateTo('/admin/posts')
  } catch (err) {
    apply(err)
    window.scrollTo({ top: 0, behavior: 'smooth' })
  } finally {
    submitting.value = false
  }
}
</script>

<template>
  <div>
    <div class="admin-page-head">
      <h1>Ubah artikel</h1>
      <a v-if="data?.data.status === 'published'" :href="`/blog/${data.data.slug}`" target="_blank" rel="noopener" class="btn btn-outline-secondary">
        Lihat di situs
      </a>
    </div>

    <template v-if="loadError">
      <AdminAlert type="danger">{{ loadError.status === 404 ? 'Artikel tidak ditemukan.' : loadError.message }}</AdminAlert>
      <NuxtLink to="/admin/posts" class="btn btn-outline-secondary">Kembali ke daftar</NuxtLink>
    </template>

    <template v-else-if="data">
      <AdminAlert v-if="message" type="danger">{{ message }}</AdminAlert>
      <AdminPostForm :initial="data.data" :submitting="submitting" :errors="errors" @submit="save" />
    </template>
  </div>
</template>
