/**
 * Perantara ke API admin backend Go: /api/admin/<path>  ->  <apiBase>/api/v1/admin/<path>
 *
 * Browser tidak pernah memegang token. Di sini token dibaca dari cookie httpOnly lalu dipasang
 * sebagai header Authorization. Body (termasuk upload file) diteruskan apa adanya.
 */
export default defineEventHandler(async (event) => {
  assertSameOrigin(event)

  const token = getCookie(event, ADMIN_COOKIE)
  if (!token) {
    setResponseStatus(event, 401)
    return { error: 'silakan login terlebih dahulu' }
  }

  // Hanya path biasa yang diteruskan (menolak "..", "?", "#", dsb).
  const path = getRouterParam(event, 'path') ?? ''
  if (!/^[A-Za-z0-9_\-/]+$/.test(path) || path.includes('//')) {
    setResponseStatus(event, 400)
    return { error: 'path tidak valid' }
  }

  const { apiBase } = useRuntimeConfig(event)
  const target = `${apiBase}/api/v1/admin/${path}${getRequestURL(event).search}`

  try {
    return await proxyRequest(event, target, {
      // cookie dikosongkan: backend tidak perlu cookie, cukup header Authorization
      headers: { authorization: `Bearer ${token}`, cookie: '' },
    })
  } catch (error) {
    console.error('proxy admin: backend tidak dapat dihubungi:', error)
    setResponseStatus(event, 502)
    return { error: 'Server tidak dapat dihubungi. Coba lagi nanti.' }
  }
})
