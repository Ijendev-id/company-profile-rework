<script setup lang="ts">
import { hero } from '~/data/home'

// Banner judul halaman (versi ringkas dari hero landing page). Dipakai di dalam <SiteHeader>.
defineProps<{
  title: string
  subtitle?: string
}>()
</script>

<template>
  <div class="header-hero bg_cover blog-hero" :style="{ backgroundImage: `url(${asset(hero.background)})` }">
    <div class="container">
      <div class="row justify-content-center">
        <div class="col-lg-8">
          <div class="header-hero-content text-center">
            <h1 class="header-title">{{ title }}</h1>
            <p v-if="subtitle" class="text">{{ subtitle }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
