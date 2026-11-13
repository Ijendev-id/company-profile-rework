<script setup lang="ts">
const visible = ref(true)

// Sama seperti kode lama: tunggu halaman selesai dimuat, lalu sembunyikan setelah 0,5 detik.
onMounted(() => {
  const hide = () => setTimeout(() => (visible.value = false), 500)
  if (document.readyState === 'complete') hide()
  else window.addEventListener('load', hide, { once: true })
})
</script>

<template>
  <div v-if="visible" class="preloader">
    <div class="loader">
      <div class="spinner">
        <div class="spinner-container">
          <div class="spinner-rotator">
            <div class="spinner-left">
              <div class="spinner-circle"></div>
            </div>
            <div class="spinner-right">
              <div class="spinner-circle"></div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
