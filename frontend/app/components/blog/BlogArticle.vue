<script setup lang="ts">
import { blog } from '~/data/blog'
import type { PostDetail } from '~/types/blog'

const props = defineProps<{ post: PostDetail }>()

const { mediaBase, mediaUrl } = useMedia()

// Isi artikel ditulis dalam Markdown; HTML mentah dinonaktifkan di renderMarkdown().
const html = computed(() => renderMarkdown(props.post.content, { mediaBase }))
</script>

<template>
  <article class="blog-article">
    <img v-if="post.cover_image" class="blog-article-cover" :src="mediaUrl(post.cover_image)" :alt="post.title" />

    <!-- eslint-disable-next-line vue/no-v-html -->
    <div class="blog-article-content" v-html="html"></div>

    <div class="blog-article-footer">
      <NuxtLink to="/blog" class="main-btn">&larr; {{ blog.backLabel }}</NuxtLink>
    </div>
  </article>
</template>
