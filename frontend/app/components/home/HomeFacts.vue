<script setup lang="ts">
import { facts } from '~/data/home'

useCounterUp()
// Tombol play video (a.glightbox) membuka video YouTube di lightbox.
useLightbox({
  href: facts.video.url,
  type: 'video',
  source: 'youtube',
  width: 900,
  autoplayVideos: true,
})
</script>

<template>
  <section id="facts" class="video-counter pt-70">
    <div class="container">
      <div class="row">
        <div class="col-lg-6 order-lg-last">
          <div class="counter-wrapper mt-50" v-bind="wow('fadeIn', 1, 0.8)">
            <div class="counter-content">
              <SectionTitle :title="facts.title" :highlight="facts.highlight" />
              <p class="text">{{ facts.text }}</p>
            </div>
            <div class="row no-gutters">
              <div v-for="counter in facts.counters" :key="counter.label" class="col-4">
                <div
                  class="single-counter d-flex align-items-center justify-content-center"
                  :class="`counter-color-${counter.color}`"
                >
                  <div class="counter-items text-center">
                    <span
                      class="count countup text-uppercase"
                      :cup-end="counter.value"
                      :cup-append="counter.append"
                    ></span>
                    <p class="text">{{ counter.label }}</p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="col-lg-6">
          <div class="video-content mt-50" v-bind="wow('fadeIn', 1, 0.5)">
            <img class="dots" :src="asset(facts.video.dots)" alt="dots" />
            <div class="video-wrapper">
              <div class="video-image">
                <img :src="asset(facts.video.image)" alt="video" />
              </div>
              <div class="video-icon">
                <a :href="facts.video.url" class="video-popup glightbox">
                  <i class="lni lni-play"></i>
                </a>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
