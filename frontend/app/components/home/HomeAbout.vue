<script setup lang="ts">
import { about } from '~/data/home'
</script>

<template>
  <section id="about">
    <div class="about-area pt-70">
      <div class="container">
        <div class="row">
          <div class="col-lg-6">
            <div class="about-content mt-50" v-bind="wow('fadeInLeftBig', 1, 0.5)">
              <SectionTitle :title="about.title" :highlight="about.highlight" />
              <p class="text">{{ about.paragraphs.join(' ') }}</p>
              <a href="#" class="main-btn" @click.prevent>{{ about.buttonLabel }}</a>
            </div>
          </div>
          <div class="col-lg-6">
            <div class="about-image text-center mt-50" v-bind="wow('fadeInRightBig', 1, 0.5)">
              <img :src="asset(about.image)" alt="about" />
            </div>
          </div>
        </div>
      </div>
      <div class="about-shape-1">
        <img :src="asset(about.shape)" alt="shape" />
      </div>
    </div>
  </section>
</template>
