/**
 * GLightbox untuk elemen ber-class "glightbox" (contoh: tombol play video).
 * Opsi diteruskan apa adanya ke GLightbox.
 */
export function useLightbox(options: Record<string, unknown> = {}) {
  let instance: { destroy?: () => void } | undefined

  onMounted(async () => {
    try {
      await loadVendor('glightbox')
      instance = window.GLightbox(options)
    } catch (error) {
      console.error(error)
    }
  })

  onBeforeUnmount(() => instance?.destroy?.())
}
