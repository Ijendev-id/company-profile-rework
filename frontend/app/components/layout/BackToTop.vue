<script setup lang="ts">
const scrollY = useScrollY()
const visible = computed(() => scrollY.value > 50)

function toTop(event: MouseEvent) {
  event.preventDefault()
  window.scrollTo({ top: 0, behavior: 'smooth' })
}
</script>

<template>
  <a href="#" class="back-to-top" :style="{ display: visible ? 'flex' : 'none' }" @click="toTop">
    <i class="lni lni-chevron-up"></i>
  </a>
</template>
