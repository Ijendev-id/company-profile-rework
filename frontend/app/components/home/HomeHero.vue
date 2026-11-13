<script setup lang="ts">
import { hero } from '~/data/home'
import { site } from '~/data/site'

useParticles('particles-1')
</script>

<template>
  <div id="home" class="header-hero bg_cover" :style="{ backgroundImage: `url(${asset(hero.background)})` }">
    <div class="container">
      <div class="row justify-content-center">
        <div class="col-lg-8">
          <div class="header-hero-content text-center">
            <h3 class="header-sub-title" v-bind="wow('fadeInUp', 1.3, 0.2)">{{ hero.subtitle }}</h3>
            <h2 class="header-title" v-bind="wow('fadeInUp', 1.3, 0.5)">{{ hero.title }}</h2>
            <p class="text" v-bind="wow('fadeInUp', 1.3, 0.8)">{{ hero.text }}</p>
            <a class="main-btn" :href="waApi(site.whatsapp.secondary)" v-bind="wow('fadeInUp', 1.3, 1.1)">
              {{ hero.ctaLabel }}
            </a>
          </div>
        </div>
      </div>
      <div class="row">
        <div class="col-lg-12">
          <div class="header-hero-image text-center" v-bind="wow('fadeIn', 1.3, 1.4)">
            <img :src="asset(hero.image)" alt="hero" />
          </div>
        </div>
      </div>
    </div>
    <div id="particles-1" class="particles"></div>
  </div>
</template>
