<script setup lang="ts">
import { blog } from '~/data/blog'
import type { Category } from '~/types/blog'

// Daftar kategori sebagai tombol. Kata pencarian (q) ikut dibawa saat pindah kategori.
const props = defineProps<{
  categories: Category[]
  active: string
  search: string
}>()

function categoryLink(slug: string) {
  return {
    path: '/blog',
    query: { q: props.search || undefined, category: slug || undefined },
  }
}
</script>

<template>
  <ul v-if="categories.length" class="blog-categories">
    <li>
      <NuxtLink :to="categoryLink('')" :class="{ active: !active }">{{ blog.allCategoriesLabel }}</NuxtLink>
    </li>
    <li v-for="category in categories" :key="category.id">
      <NuxtLink :to="categoryLink(category.slug)" :class="{ active: active === category.slug }">
        {{ category.name }}
      </NuxtLink>
    </li>
  </ul>
</template>
