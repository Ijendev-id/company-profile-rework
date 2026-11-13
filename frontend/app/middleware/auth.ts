// Penjaga halaman admin: belum login -> diarahkan ke /admin/login.
// Dipasang di tiap halaman admin lewat definePageMeta({ middleware: 'auth' }).
export default defineNuxtRouteMiddleware(async (to) => {
  const { user, fetchUser } = useAdminAuth()

  if (user.value) return
  if (await fetchUser()) return

  return navigateTo({ path: '/admin/login', query: { redirect: to.fullPath } })
})
