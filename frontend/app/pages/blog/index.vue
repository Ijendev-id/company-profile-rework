<script setup lang="ts">
import { blog } from '~/data/blog'
import { seo } from '~/data/seo'

usePageSeo(seo.blog!)

const route = useRoute()
const api = useBlogApi()

// Filter diambil dari URL: /blog?page=2&category=berita&q=kata
const page = computed(() => {
  const value = Number(route.query.page)
  return Number.isInteger(value) && value > 0 ? value : 1
})
const category = computed(() => (typeof route.query.category === 'string' ? route.query.category : ''))
const search = computed(() => (typeof route.query.q === 'string' ? route.query.q.trim() : ''))

// Data dimuat di server (bagus untuk SEO) dan dimuat ulang otomatis kalau filter berubah.
const {
  data: list,
  error,
  status,
  refresh,
} = await useAsyncData(
  'blog-posts',
  () =>
    api.posts({
      page: page.value,
      per_page: blog.perPage,
      category: category.value || undefined,
      q: search.value || undefined,
    }),
  { watch: [page, category, search] },
)

const { data: categories } = await useAsyncData('blog-categories', () => api.categories())

const posts = computed(() => list.value?.data ?? [])
const isFiltered = computed(() => !!category.value || !!search.value)
</script>

<template>
  <SiteHeader>
    <PageHero :title="blog.title" :subtitle="blog.subtitle" />
  </SiteHeader>

  <section class="blog-section">
    <div class="container">
      <div class="blog-toolbar">
        <BlogCategoryFilter :categories="categories?.data ?? []" :active="category" :search="search" />
        <BlogSearch :search="search" :category="category" />
      </div>

      <div v-if="error" class="blog-state">
        <p>{{ blog.loadErrorText }}</p>
        <button type="button" class="main-btn" @click="refresh()">{{ blog.retryLabel }}</button>
      </div>

      <template v-else>
        <div :class="{ 'blog-loading': status === 'pending' }">
          <BlogList :posts="posts" :filtered="isFiltered" />
        </div>
        <BlogPagination v-if="list" :meta="list.meta" />
      </template>
    </div>
  </section>
</template>
