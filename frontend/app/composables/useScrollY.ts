/** Posisi scroll vertikal (reaktif). Aman dipakai di SSR: listener baru dipasang setelah mounted. */
export function useScrollY() {
  const y = ref(0)
  const update = () => {
    y.value = window.scrollY
  }

  onMounted(() => {
    update()
    window.addEventListener('scroll', update, { passive: true })
  })
  onBeforeUnmount(() => window.removeEventListener('scroll', update))

  return y
}
