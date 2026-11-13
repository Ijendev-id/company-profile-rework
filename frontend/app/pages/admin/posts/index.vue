<script setup lang="ts">
import { admin } from '~/data/admin'
import type { ListResponse, PostSummary } from '~/types/blog'

definePageMeta({ layout: 'admin', middleware: 'auth' })
useHead({ title: 'Artikel' })

const route = useRoute()
const api = useAdminApi()
const flash = useFlash()

const success = ref(flash.take())
const problem = ref('')

// Filter diambil dari URL: /admin/posts?page=2&status=draft&q=kata
const page = computed(() => {
  const value = Number(route.query.page)
  return Number.isInteger(value) && value > 0 ? value : 1
})
const statusFilter = computed(() => (typeof route.query.status === 'string' ? route.query.status : ''))
const search = computed(() => (typeof route.query.q === 'string' ? route.query.q.trim() : ''))

const { data, error, refresh } = useAsyncData(
  'admin-posts',
  () =>
    api.get<ListResponse<PostSummary>>('/posts', {
      page: page.value,
      per_page: admin.perPage,
      status: statusFilter.value || undefined,
      q: search.value || undefined,
    }),
  { watch: [page, statusFilter, search] },
)

const posts = computed(() => data.value?.data ?? [])
const loadError = computed(() => (error.value ? toApiError(error.value).message : ''))

const searchInput = ref(search.value)
const statusInput = ref(statusFilter.value)
watch(search, (value) => (searchInput.value = value))
watch(statusFilter, (value) => (statusInput.value = value))

function applyFilters() {
  navigateTo({ path: '/admin/posts', query: { q: searchInput.value.trim() || undefined, status: statusInput.value || undefined } })
}

async function remove(post: PostSummary) {
  if (!window.confirm(`Hapus artikel "${post.title}"?\n\nArtikel yang dihapus tidak dapat dikembalikan.`)) return

  success.value = ''
  problem.value = ''
  try {
    await api.del(`/posts/${post.id}`)
    success.value = 'Artikel berhasil dihapus.'

    // Menghapus artikel terakhir di halaman ini -> mundur satu halaman.
    if (posts.value.length === 1 && page.value > 1) {
      await navigateTo({ path: '/admin/posts', query: { ...route.query, page: page.value > 2 ? String(page.value - 1) : undefined } })
    } else {
      await refresh()
    }
  } catch (err) {
    problem.value = toApiError(err).message
  }
}
</script>

<template>
  <div>
    <div class="admin-page-head">
      <h1>Artikel</h1>
      <NuxtLink to="/admin/posts/new" class="btn btn-primary">+ Tulis artikel</NuxtLink>
    </div>

    <AdminAlert v-if="success" type="success">{{ success }}</AdminAlert>
    <AdminAlert v-if="problem || loadError" type="danger">{{ problem || loadError }}</AdminAlert>

    <form class="admin-filters" role="search" @submit.prevent="applyFilters">
      <input v-model="searchInput" type="search" class="form-control" placeholder="Cari judul atau ringkasan..." aria-label="Cari artikel" />
      <select v-model="statusInput" class="form-select" aria-label="Filter status" @change="applyFilters">
        <option value="">Semua status</option>
        <option value="published">{{ admin.statusLabels.published }}</option>
        <option value="draft">{{ admin.statusLabels.draft }}</option>
      </select>
      <button type="submit" class="btn btn-outline-primary">Cari</button>
    </form>

    <div class="table-responsive">
      <table class="admin-table">
        <thead>
          <tr>
            <th>Judul</th>
            <th>Kategori</th>
            <th>Status</th>
            <th>Tanggal</th>
            <th></th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="post in posts" :key="post.id">
            <td>
              <NuxtLink :to="`/admin/posts/${post.id}`" class="admin-title">{{ post.title }}</NuxtLink>
              <span class="admin-sub">/blog/{{ post.slug }}</span>
            </td>
            <td>{{ post.category?.name ?? '-' }}</td>
            <td>
              <span class="admin-badge" :class="`admin-badge--${post.status}`">{{ admin.statusLabels[post.status] }}</span>
            </td>
            <td>{{ formatDate(post.published_at ?? post.created_at) }}</td>
            <td>
              <div class="admin-actions">
                <NuxtLink :to="`/admin/posts/${post.id}`" class="btn btn-sm btn-outline-primary">Ubah</NuxtLink>
                <a v-if="post.status === 'published'" :href="`/blog/${post.slug}`" target="_blank" rel="noopener" class="btn btn-sm btn-outline-secondary">Lihat</a>
                <button type="button" class="btn btn-sm btn-outline-danger" @click="remove(post)">Hapus</button>
              </div>
            </td>
          </tr>
          <tr v-if="!posts.length && data">
            <td colspan="5" class="text-center text-muted py-4">
              {{ search || statusFilter ? 'Tidak ada artikel yang cocok.' : 'Belum ada artikel. Klik "Tulis artikel" untuk membuat yang pertama.' }}
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div class="mt-4">
      <BlogPagination v-if="data" :meta="data.meta" />
    </div>
  </div>
</template>
