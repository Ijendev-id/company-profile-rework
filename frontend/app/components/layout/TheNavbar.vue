<script setup lang="ts">
import { navItems, type NavItem } from '~/data/navigation'
import { site } from '~/data/site'

const route = useRoute()

// Sticky: navbar berubah gaya + logo saat halaman di-scroll
const scrollY = useScrollY()
const isSticky = computed(() => scrollY.value > 0)
const logo = computed(() => asset(isSticky.value ? site.logo.sticky : site.logo.default))

// Scroll spy: menandai menu sesuai section yang sedang tampil
const hashes = navItems.flatMap((item) => (item.hash ? [item.hash] : []))
const activeHash = useScrollSpy(hashes)

// Menu mobile (dikelola Vue, tidak butuh bootstrap.bundle.js)
const menuOpen = ref(false)

function isActive(item: NavItem) {
  if (item.hash) return route.path === '/' && activeHash.value === item.hash
  return !!item.to && route.path.startsWith(item.to)
}

function onHashClick(event: MouseEvent, hash: string) {
  event.preventDefault()
  menuOpen.value = false
  if (route.path === '/') document.querySelector(hash)?.scrollIntoView({ behavior: 'smooth' })
  else navigateTo({ path: '/', hash })
}
</script>

<template>
  <div class="navbar-area" :class="{ sticky: isSticky }">
    <div class="container">
      <div class="row">
        <div class="col-lg-12">
          <nav class="navbar navbar-expand-lg">
            <NuxtLink class="navbar-brand" to="/">
              <img :src="logo" alt="Logo" />
            </NuxtLink>

            <button
              class="navbar-toggler"
              :class="{ active: menuOpen }"
              type="button"
              aria-controls="navbarSupportedContent"
              :aria-expanded="menuOpen"
              aria-label="Toggle navigation"
              @click="menuOpen = !menuOpen"
            >
              <span class="toggler-icon"></span>
              <span class="toggler-icon"></span>
              <span class="toggler-icon"></span>
            </button>

            <div
              id="navbarSupportedContent"
              class="collapse navbar-collapse sub-menu-bar"
              :class="{ show: menuOpen }"
            >
              <ul id="nav" class="navbar-nav ms-auto">
                <li v-for="item in navItems" :key="item.label" class="nav-item">
                  <a
                    v-if="item.hash"
                    class="page-scroll"
                    :class="{ active: isActive(item) }"
                    :href="`/${item.hash}`"
                    @click="onHashClick($event, item.hash)"
                  >
                    {{ item.label }}
                  </a>
                  <NuxtLink
                    v-else-if="item.to"
                    :to="item.to"
                    :class="{ active: isActive(item) }"
                    @click="menuOpen = false"
                  >
                    {{ item.label }}
                  </NuxtLink>
                </li>
              </ul>
            </div>

            <div class="navbar-btn d-none d-sm-inline-block">
              <a class="main-btn" :href="waApi(site.whatsapp.primary)" rel="nofollow">Hubungi Sekarang</a>
            </div>
          </nav>
        </div>
      </div>
    </div>
  </div>
</template>
