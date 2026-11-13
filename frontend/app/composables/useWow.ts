/** Aktifkan animasi scroll WOW.js (elemen ber-class "wow"). Panggil sekali di layout. */
export function useWow() {
  onMounted(async () => {
    try {
      await loadVendor('wow')
      new window.WOW().init()
    } catch (error) {
      console.error(error)
    }
  })
}
