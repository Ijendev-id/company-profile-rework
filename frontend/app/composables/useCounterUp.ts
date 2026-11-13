/**
 * Animasi angka (elemen ber-class "countup" dengan atribut cup-end / cup-append).
 * Opsi bawaan sama dengan kode lama; bisa ditimpa lewat parameter.
 */
export function useCounterUp(options: Record<string, unknown> = {}) {
  onMounted(async () => {
    try {
      await loadVendor('countUp')
      const counter = new window.counterUp({
        start: 0,
        duration: 2000,
        intvalues: true,
        interval: 100,
        append: 'k',
        ...options,
      })
      counter.start()
    } catch (error) {
      console.error(error)
    }
  })
}
