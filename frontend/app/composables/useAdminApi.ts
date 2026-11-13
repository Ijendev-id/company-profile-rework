import { admin } from '~/data/admin'

/** Error dari API dalam bentuk seragam, siap ditampilkan di form. */
export class ApiError extends Error {
  constructor(
    public status: number,
    message: string,
    /** Pesan per field, contoh { title: "wajib diisi" } (status 422) */
    public fields: Record<string, string> = {},
  ) {
    super(message)
  }
}

export function toApiError(error: unknown): ApiError {
  if (error instanceof ApiError) return error

  const fetchError = error as { status?: number; statusCode?: number; data?: { error?: string; fields?: Record<string, string> } }
  const status = fetchError.status ?? fetchError.statusCode ?? 0

  if (!status) return new ApiError(0, 'Tidak dapat menghubungi server. Periksa koneksi lalu coba lagi.')

  return new ApiError(status, fetchError.data?.error ?? 'Terjadi kesalahan pada server.', fetchError.data?.fields ?? {})
}

/**
 * Pemanggil API admin lewat perantara Nuxt (/api/admin/...). Token tidak pernah terlihat di sini.
 * Kalau sesi habis (401), otomatis kembali ke halaman login.
 */
export function useAdminApi() {
  const { user } = useAdminAuth()

  async function request<T>(path: string, options: Parameters<typeof $fetch>[1] = {}): Promise<T> {
    try {
      return (await $fetch(`/api/admin${path}`, options)) as T
    } catch (error) {
      const apiError = toApiError(error)
      if (apiError.status === 401) {
        user.value = null
        await navigateTo({ path: '/admin/login', query: { expired: '1' } })
      }
      throw apiError
    }
  }

  /** Upload gambar. Hasilnya path seperti "/uploads/xxxx.jpg". */
  async function uploadImage(file: File): Promise<string> {
    if (!admin.allowedImageTypes.includes(file.type)) {
      throw new ApiError(415, 'File harus berupa gambar JPG, PNG, GIF, atau WebP.')
    }
    if (file.size > admin.maxUploadMB * 1024 * 1024) {
      throw new ApiError(413, `Ukuran file maksimal ${admin.maxUploadMB} MB.`)
    }

    const form = new FormData()
    form.append('file', file)
    const result = await request<{ url: string }>('/uploads', { method: 'POST', body: form })
    return result.url
  }

  return {
    get: <T>(path: string, query?: object) => request<T>(path, { query }),
    post: <T>(path: string, body: object) => request<T>(path, { method: 'POST', body }),
    put: <T>(path: string, body: object) => request<T>(path, { method: 'PUT', body }),
    del: (path: string) => request<void>(path, { method: 'DELETE' }),
    uploadImage,
  }
}
