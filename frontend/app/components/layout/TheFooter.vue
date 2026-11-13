<script setup lang="ts">
import { footer } from '~/data/footer'
import { site } from '~/data/site'

useParticles('particles-2')
</script>

<template>
  <footer id="footer" class="footer-area pt-120">
    <div class="container">
      <!-- CTA: tombol WhatsApp -->
      <div class="row mb-5 justify-content-center">
        <div class="col-lg-12">
          <div
            class="cta-box p-4 p-md-5 bg-white shadow-lg rounded-4 d-flex flex-column flex-lg-row align-items-center justify-content-between"
            v-bind="wow('fadeIn', 1, 0.3)"
          >
            <div class="text-center text-lg-start mb-3 mb-lg-0">
              <h3 class="fw-bold mb-2" style="color: #111">
                {{ footer.cta.title }} <span style="color: #0061f2">{{ footer.cta.highlight }}</span>
              </h3>
              <p class="mb-0 text-muted">{{ footer.cta.text }}</p>
            </div>

            <div class="d-flex align-items-center gap-3">
              <a
                :href="waMe(site.whatsapp.primary, footer.cta.message)"
                target="_blank"
                rel="noopener noreferrer"
                class="btn btn-success d-flex align-items-center gap-2 px-4 py-2 rounded-pill shadow-sm"
                style="font-weight: 600"
              >
                <i class="lni lni-whatsapp fs-5"></i> {{ footer.cta.buttonLabel }}
              </a>
            </div>
          </div>
        </div>
      </div>

      <div class="footer-widget pb-100">
        <div class="row">
          <!-- Tentang + sosial media -->
          <div class="col-lg-4 col-md-6 col-sm-8">
            <div class="footer-about mt-50" v-bind="wow('fadeIn', 1, 0.2)">
              <NuxtLink class="logo" to="/">
                <img :src="asset(site.logo.default)" alt="logo" />
              </NuxtLink>
              <p class="text">{{ footer.about }}</p>
              <ul class="social">
                <li v-for="item in site.social" :key="item.label">
                  <a :href="item.href" :aria-label="item.label">
                    <i class="lni" :class="item.icon"></i>
                  </a>
                </li>
              </ul>
            </div>
          </div>

          <!-- Kolom kosong (tempat link footer, belum dipakai) -->
          <div class="col-lg-5 col-md-7 col-sm-12"></div>

          <!-- Kontak -->
          <div class="col-lg-3 col-md-5 col-sm-12">
            <div class="footer-contact mt-50" v-bind="wow('fadeIn', 1, 0.8)">
              <div class="footer-title">
                <h4 class="title">{{ footer.contactTitle }}</h4>
              </div>
              <ul class="contact">
                <li>{{ site.contact.phone }}</li>
                <li>
                  <template v-for="(line, index) in site.contact.address" :key="line">
                    <br v-if="index > 0" />
                    {{ line }}
                  </template>
                </li>
              </ul>
            </div>
          </div>
        </div>
      </div>

      <div class="footer-copyright">
        <div class="row">
          <div class="col-lg-12">
            <div class="copyright d-sm-flex justify-content-between">
              <div class="copyright-content">
                <p class="text">
                  Designed and Developed by
                  <a :href="site.url" rel="nofollow">{{ site.name }}</a>
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div id="particles-2"></div>
  </footer>
</template>
