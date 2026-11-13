import type { PostSummary } from '~/types/blog'

export interface AdminUser {
  id: number
  name: string
  email: string
}

/** Artikel lengkap dari API admin (untuk form edit). */
export interface AdminPost extends PostSummary {
  content: string
  meta_title: string
  meta_description: string
  category_id: number | null
}

/** Isi form artikel yang dikirim ke API (POST/PUT). */
export interface PostPayload {
  title: string
  slug: string
  excerpt: string
  content: string
  cover_image: string
  status: 'draft' | 'published'
  meta_title: string
  meta_description: string
  category_id: number | null
}
