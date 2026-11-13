/**
 * Library JS bawaan template. Tambah entri di sini kalau ada library baru.
 * (bootstrap.bundle.min.js dan main.js sengaja tidak dimuat: menu mobile,
 * sticky navbar, scroll spy, dll. sekarang ditangani komponen Vue.)
 */
const vendorFiles = {
  wow: 'js/wow.min.js',
  glightbox: 'js/glightbox.min.js',
  countUp: 'js/count-up.min.js',
  particles: 'js/particles.min.js',
} as const

export type VendorName = keyof typeof vendorFiles

export const loadVendor = (name: VendorName) => loadScript(asset(vendorFiles[name]))
