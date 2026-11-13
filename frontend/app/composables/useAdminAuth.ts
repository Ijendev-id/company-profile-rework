import type { AdminUser } from '~/types/admin'

/** Status login admin. Token sendiri disimpan di cookie httpOnly oleh server Nuxt; di sini hanya data user. */
export function useAdminAuth() {
  const user = useState<AdminUser | null>('admin-user', () => null)

  /** Cek ke server apakah cookie login masih berlaku. */
  async function fetchUser(): Promise<boolean> {
    try {
      const result = await $fetch<{ data: AdminUser }>('/api/admin/me')
      user.value = result.data
      return true
    } catch {
      user.value = null
      return false
    }
  }

  async function login(email: string, password: string) {
    const result = await $fetch<{ user: AdminUser }>('/api/admin/login', {
      method: 'POST',
      body: { email, password },
    })
    user.value = result.user
  }

  async function logout() {
    await $fetch('/api/admin/logout', { method: 'POST' }).catch(() => undefined)
    user.value = null
    await navigateTo('/admin/login')
  }

  return { user, fetchUser, login, logout }
}
