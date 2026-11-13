/**
 * Path aset template. Semua gambar/css/js ada di public/assets_landing.
 * Contoh: asset('images/logo/iii.png') -> '/assets_landing/images/logo/iii.png'
 */
export const ASSET_BASE = '/assets_landing'

export const asset = (path: string) => `${ASSET_BASE}/${path.replace(/^\/+/, '')}`
