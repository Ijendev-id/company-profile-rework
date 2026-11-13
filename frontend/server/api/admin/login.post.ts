/**
 * POST /api/admin/login
 * Meneruskan email & password ke backend Go. Kalau berhasil, token disimpan di cookie httpOnly
 * (tidak dikirim ke JavaScript browser); browser hanya menerima data user.
 */
export default defineEventHandler(async (event) => {
  assertSameOrigin(event)

  const body = await readBody<{ email?: unknown; password?: unknown }>(event)
  const { apiBase } = useRuntimeConfig(event)

  try {
    const result = await $fetch<{ token: string; expires_at: string; user: unknown }>(
      `${apiBase}/api/v1/auth/login`,
      { method: 'POST', body: { email: body?.email, password: body?.password } },
    )

    const secondsLeft = Math.floor((new Date(result.expires_at).getTime() - Date.now()) / 1000)
    setCookie(event, ADMIN_COOKIE, result.token, adminCookieOptions(event, Math.max(secondsLeft, 60)))

    return { user: result.user }
  } catch (error) {
    const fetchError = error as { statusCode?: number; data?: unknown }

    // Kegagalan dari backend (password salah, terlalu banyak percobaan, dll) diteruskan apa adanya.
    if (fetchError.statusCode) {
      setResponseStatus(event, fetchError.statusCode)
      return fetchError.data ?? { error: 'Login gagal' }
    }

    // Backend tidak bisa dihubungi sama sekali
    console.error('login: backend tidak dapat dihubungi:', error)
    setResponseStatus(event, 502)
    return { error: 'Server tidak dapat dihubungi. Coba lagi nanti.' }
  }
})
