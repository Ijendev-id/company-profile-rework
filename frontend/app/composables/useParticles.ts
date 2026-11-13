import { particlesConfig } from '~/data/particles'

/** Pasang efek particles.js pada elemen dengan id tertentu, dan bersihkan saat komponen dilepas. */
export function useParticles(id: string, config: Record<string, unknown> = particlesConfig) {
  onMounted(async () => {
    try {
      await loadVendor('particles')
      if (document.getElementById(id)) window.particlesJS(id, structuredClone(config))
    } catch (error) {
      console.error(error)
    }
  })

  onBeforeUnmount(() => destroyParticles(id))
}

function destroyParticles(id: string) {
  const list = window.pJSDom
  if (!Array.isArray(list)) return

  const entry = list.find((p) => p?.pJS?.canvas?.el?.parentNode?.id === id)
  if (!entry) return

  try {
    entry.pJS.fn.vendors.destroypJS()
  } catch {
    // abaikan: elemen akan dilepas Vue juga
  }
  // destroypJS() mengosongkan window.pJSDom; pulihkan agar instance lain tetap jalan.
  window.pJSDom = list.filter((p) => p !== entry)
}
