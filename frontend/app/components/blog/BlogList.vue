<script setup lang="ts">
import { blog } from '~/data/blog'
import type { PostSummary } from '~/types/blog'

// Kisi artikel. `filtered` = daftar sedang dipersempit pencarian/kategori (hanya mengubah teks saat kosong).
defineProps<{
  posts: PostSummary[]
  filtered?: boolean
}>()
</script>

<template>
  <div v-if="posts.length" class="row">
    <div v-for="post in posts" :key="post.id" class="col-lg-4 col-md-6 mb-4">
      <BlogCard :post="post" />
    </div>
  </div>
  <p v-else class="blog-state">{{ filtered ? blog.emptyFilteredText : blog.emptyText }}</p>
</template>
