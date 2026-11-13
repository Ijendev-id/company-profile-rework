<script setup lang="ts">
import { site } from '~/data/site'

const route = useRoute()
const api = useBlogApi()
const { mediaUrl } = useMedia()

const slug = String(route.params.slug)

const { data, error } = await useAsyncData(`blog-post-${slug}`, () => api.post(slug))

// Slug tidak ada atau artikel masih draft -> halaman 404. Backend mati/error -> 503.
if (error.value || !data.value) {
  const notFound = error.value?.statusCode === 404
  throw createError({
    statusCode: notFound ? 404 : 503,
    statusMessage: notFound ? 'Artikel tidak ditemukan' : 'Artikel tidak dapat dimuat saat ini',
    fatal: true,
  })
}

const post = computed(() => data.value!.data)

// Baris info di bawah judul: kategori · tanggal · penulis
const subtitle = computed(() =>
  [
    post.value.category?.name,
    formatDate(post.value.published_at),
    post.value.author_name ? `oleh ${post.value.author_name}` : '',
  ]
    .filter(Boolean)
    .join(' · '),
)

usePageSeo({
  title: post.value.meta_title || `${post.value.title} | ${site.name}`,
  description: post.value.meta_description || post.value.excerpt,
  image: post.value.cover_image ? mediaUrl(post.value.cover_image) : undefined,
  url: `${site.url}/blog/${post.value.slug}`,
  type: 'article',
})
</script>

<template>
  <SiteHeader>
    <PageHero :title="post.title" :subtitle="subtitle" />
  </SiteHeader>

  <section class="blog-section">
    <div class="container">
      <div class="row justify-content-center">
        <div class="col-lg-8">
          <BlogArticle :post="post" />
        </div>
      </div>
    </div>
  </section>
</template>
