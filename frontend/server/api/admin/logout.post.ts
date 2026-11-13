/** POST /api/admin/logout: menghapus cookie login. */
export default defineEventHandler((event) => {
  assertSameOrigin(event)
  deleteCookie(event, ADMIN_COOKIE, { path: '/' })
  return { ok: true }
})
