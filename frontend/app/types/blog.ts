// Bentuk data dari API backend Go (lihat backend/internal/handlers/posts.go dan categories.go).

export interface Category {
  id: number
  name: string
  slug: string
}

/** Artikel versi ringkas (untuk daftar). */
export interface PostSummary {
  id: number
  title: string
  slug: string
  excerpt: string
  cover_image: string
  status: 'draft' | 'published'
  published_at: string | null
  created_at: string
  updated_at: string
  category: Category | null
}

/** Artikel lengkap (untuk halaman baca). */
export interface PostDetail extends PostSummary {
  content: string
  meta_title: string
  meta_description: string
  author_name: string
}

export interface PageMeta {
  page: number
  per_page: number
  total: number
  total_pages: number
}

export interface ListResponse<T> {
  data: T[]
  meta: PageMeta
}
