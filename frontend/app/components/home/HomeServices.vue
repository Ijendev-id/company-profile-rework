<script setup lang="ts">
import { services } from '~/data/home'
</script>

<template>
  <section id="features" class="services-area pt-120">
    <div class="container">
      <div class="row justify-content-center">
        <div class="col-lg-10">
          <SectionTitle :title="services.title" :highlight="services.highlight" center />
        </div>
      </div>
      <div class="row justify-content-center">
        <div v-for="item in services.items" :key="item.title" class="col-lg-4 col-md-7 col-sm-8">
          <div class="single-services text-center mt-30" v-bind="wow('fadeIn', 1, item.delay)">
            <div class="services-icon">
              <img class="shape" :src="asset('images/services/services-shape.svg')" alt="shape" />
              <img class="shape-1" :src="asset(item.shape)" alt="shape" />
              <i class="lni" :class="item.icon"></i>
            </div>
            <div class="services-content mt-30">
              <h4 class="services-title">
                <a href="#" @click.prevent>{{ item.title }}</a>
              </h4>
              <p class="text">{{ item.text }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
