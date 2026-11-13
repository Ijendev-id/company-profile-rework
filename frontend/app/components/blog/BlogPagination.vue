<script setup lang="ts">
import type { PageMeta } from '~/types/blog'

const props = defineProps<{ meta: PageMeta }>()

const route = useRoute()

// Link halaman tetap membawa filter yang aktif (category, q). Halaman 1 tidak perlu ?page=1.
function pageLink(page: number) {
  return { path: route.path, query: { ...route.query, page: page > 1 ? String(page) : undefined } }
}

// Tampilkan maksimal 5 nomor di sekitar halaman aktif.
const pages = computed(() => {
  const { page, total_pages: total } = props.meta
  const end = Math.min(total, Math.max(page + 2, 5))
  const start = Math.max(1, end - 4)
  return Array.from({ length: end - start + 1 }, (_, i) => start + i)
})
</script>

<template>
  <nav v-if="meta.total_pages > 1" aria-label="Halaman artikel">
    <ul class="pagination justify-content-center blog-pagination">
      <li class="page-item" :class="{ disabled: meta.page <= 1 }">
        <NuxtLink v-if="meta.page > 1" class="page-link" :to="pageLink(meta.page - 1)" rel="prev">&laquo;</NuxtLink>
        <span v-else class="page-link">&laquo;</span>
      </li>

      <li v-for="page in pages" :key="page" class="page-item" :class="{ active: page === meta.page }">
        <NuxtLink
          class="page-link"
          :to="pageLink(page)"
          :aria-current="page === meta.page ? 'page' : undefined"
        >
          {{ page }}
        </NuxtLink>
      </li>

      <li class="page-item" :class="{ disabled: meta.page >= meta.total_pages }">
        <NuxtLink v-if="meta.page < meta.total_pages" class="page-link" :to="pageLink(meta.page + 1)" rel="next">
          &raquo;
        </NuxtLink>
        <span v-else class="page-link">&raquo;</span>
      </li>
    </ul>
  </nav>
</template>
