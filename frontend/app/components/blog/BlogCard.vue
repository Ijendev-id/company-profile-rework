<script setup lang="ts">
import type { PostSummary } from '~/types/blog'

defineProps<{ post: PostSummary }>()

const { mediaUrl } = useMedia()
</script>

<template>
  <NuxtLink :to="`/blog/${post.slug}`" class="blog-card">
    <img
      v-if="post.cover_image"
      class="blog-card-cover"
      :src="mediaUrl(post.cover_image)"
      :alt="post.title"
      loading="lazy"
    />
    <div v-else class="blog-card-cover blog-card-cover--empty" aria-hidden="true"></div>

    <div class="blog-card-body">
      <span v-if="post.category" class="blog-card-category">{{ post.category.name }}</span>
      <h3 class="blog-card-title">{{ post.title }}</h3>
      <p v-if="post.excerpt" class="blog-card-excerpt">{{ post.excerpt }}</p>
      <time v-if="post.published_at" class="blog-card-date" :datetime="post.published_at">
        {{ formatDate(post.published_at) }}
      </time>
    </div>
  </NuxtLink>
</template>
