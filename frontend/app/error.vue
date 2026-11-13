<script setup lang="ts">
import type { NuxtError } from '#app'
import { blog } from '~/data/blog'

// Halaman error (404 artikel tidak ditemukan, atau backend tidak bisa dihubungi).
const props = defineProps<{ error: NuxtError }>()

const notFound = computed(() => props.error.statusCode === 404)
const title = computed(() => (notFound.value ? 'Halaman tidak ditemukan' : 'Terjadi kesalahan'))
const message = computed(() =>
  notFound.value
    ? 'Halaman yang kamu cari tidak ada atau sudah dipindahkan.'
    : 'Konten tidak dapat dimuat saat ini. Silakan coba lagi beberapa saat lagi.',
)

useHead({ title: () => title.value })

const goBlog = () => clearError({ redirect: '/blog' })
const goHome = () => clearError({ redirect: '/' })
</script>

<template>
  <NuxtLayout>
    <SiteHeader>
      <PageHero :title="title" :subtitle="message" />
    </SiteHeader>

    <section class="blog-section">
      <div class="container text-center">
        <button type="button" class="main-btn" @click="goBlog">{{ blog.backLabel }}</button>
        <p class="mt-4"><a href="/" @click.prevent="goHome">Ke beranda</a></p>
      </div>
    </section>
  </NuxtLayout>
</template>
