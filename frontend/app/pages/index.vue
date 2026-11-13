<script setup lang="ts">
import { seo } from '~/data/seo'

usePageSeo(seo.home!)
</script>

<template>
  <SiteHeader>
    <HomeHero />
  </SiteHeader>
  <HomeServices />
  <HomeAbout />
  <HomeFacts />
  <HomeClients />
</template>
