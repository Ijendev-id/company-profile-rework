<script setup lang="ts">
import { clients } from '~/data/home'
</script>

<template>
  <section id="clients" class="services-area pt-120">
    <div class="container">
      <div class="row justify-content-center">
        <div class="col-lg-10">
          <SectionTitle :title="clients.title" :highlight="clients.highlight" center />
        </div>

        <!-- Sama seperti HTML lama: brand-area berada langsung di dalam row -->
        <div class="brand-area pt-10">
          <div class="container">
            <div class="row">
              <div class="col-lg-12">
                <div
                  class="brand-logo d-flex align-items-center justify-content-center justify-content-md-between"
                >
                  <div
                    v-for="logo in clients.logos"
                    :key="logo.name"
                    class="single-logo mt-30"
                    v-bind="wow('fadeIn', logo.duration, logo.delay)"
                  >
                    <img :src="asset(logo.image)" :alt="logo.name" />
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
