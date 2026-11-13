<script setup lang="ts">
import { blog } from '~/data/blog'

// Form pencarian. Tanpa JavaScript tetap jalan sebagai form GET biasa (/blog?q=...).
const props = defineProps<{
  search: string
  category: string
}>()

const term = ref(props.search)
watch(
  () => props.search,
  (value) => (term.value = value),
)

function submit() {
  navigateTo({
    path: '/blog',
    query: { q: term.value.trim() || undefined, category: props.category || undefined },
  })
}
</script>

<template>
  <form class="blog-search" action="/blog" method="get" role="search" @submit.prevent="submit">
    <input v-if="category" type="hidden" name="category" :value="category" />
    <input
      v-model="term"
      type="search"
      name="q"
      maxlength="100"
      :placeholder="blog.searchPlaceholder"
      aria-label="Cari artikel"
    />
    <button type="submit" aria-label="Cari"><i class="lni lni-search-alt"></i></button>
  </form>
</template>
