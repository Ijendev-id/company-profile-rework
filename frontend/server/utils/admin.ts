import type { H3Event } from 'h3'

/** Nama cookie penyimpan token login admin. Cookie ini httpOnly: tidak bisa dibaca JavaScript di browser. */
export const ADMIN_COOKIE = 'ijendev_admin'

/**
 * Pengaman CSRF: request yang mengubah data hanya diterima dari halaman situs ini sendiri.
 * Browser selalu menyertakan header Origin pada request POST/PUT/DELETE, jadi situs lain
 * tidak bisa memicu aksi admin lewat cookie milik admin yang sedang login.
 */
export function assertSameOrigin(event: H3Event) {
  if (['GET', 'HEAD', 'OPTIONS'].includes(event.method)) return

  const origin = getHeader(event, 'origin')
  if (!origin) return

  let originHost = ''
  try {
    originHost = new URL(origin).host
  } catch {
    // Origin tidak valid -> ditolak di bawah
  }

  if (originHost !== getRequestHost(event, { xForwardedHost: true })) {
    throw createError({ statusCode: 403, statusMessage: 'Asal request tidak diizinkan' })
  }
}

/** Pengaturan cookie login. `secure` aktif otomatis kalau situs diakses lewat HTTPS. */
export function adminCookieOptions(event: H3Event, maxAgeSeconds: number) {
  return {
    httpOnly: true,
    sameSite: 'lax' as const,
    secure: getRequestProtocol(event, { xForwardedProto: true }) === 'https',
    path: '/',
    maxAge: maxAgeSeconds,
  }
}
