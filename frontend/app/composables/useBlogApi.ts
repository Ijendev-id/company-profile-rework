import type { Category, ListResponse, PostDetail, PostSummary } from '~/types/blog'

export interface PostListQuery {
  page?: number
  per_page?: number
  category?: string
  q?: string
}

/**
 * Pemanggil API blog (backend Go). Panggil di dalam <script setup> atau composable lain.
 * Saat render di server dipakai `apiBase` (boleh alamat internal), di browser dipakai `public.apiBase`.
 */
export function useBlogApi() {
  const config = useRuntimeConfig()
  const baseURL = import.meta.server ? config.apiBase : config.public.apiBase

  const request = <T>(path: string, query?: object) => $fetch<T>(`/api/v1${path}`, { baseURL, query })

  return {
    posts: (query: PostListQuery = {}) => request<ListResponse<PostSummary>>('/posts', query),

    post: (slug: string) => request<{ data: PostDetail }>(`/posts/${encodeURIComponent(slug)}`),

    // Kategori hanya pelengkap filter: kalau gagal dimuat, daftar artikel tetap tampil tanpa filter.
    categories: () =>
      request<{ data: Category[] }>('/categories').catch(() => ({ data: [] as Category[] })),
  }
}
