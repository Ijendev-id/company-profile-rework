export interface AdminNavItem {
  label: string
  to: string
  icon?: string
}
export interface AdminNavGroup {
  title: string
  items: AdminNavItem[]
}

/** Teks dan pengaturan halaman admin. */
export const admin = {
  title: 'Admin Ijendev',
  perPage: 10,

  // Harus sama dengan batas di backend (MAX_UPLOAD_MB di .env) dan jenis file yang diizinkan.
  maxUploadMB: 5,
  allowedImageTypes: ['image/jpeg', 'image/png', 'image/gif', 'image/webp'],

  // Menu sidebar, dikelompokkan. Fitur baru: tambah satu item di grup yang ada, atau satu grup baru.
  // Contoh grup baru: { title: 'Pengaturan', items: [{ label: 'Profil', to: '/admin/profile', icon: 'lni-user' }] }
  // `icon` memakai class LineIcons (sudah ada di template).
  nav: [
    {
      title: 'Blog',
      items: [
        { label: 'Artikel', to: '/admin/posts', icon: 'lni-pencil-alt' },
        { label: 'Kategori', to: '/admin/categories', icon: 'lni-tag' },
      ],
    },
  ] as AdminNavGroup[],

  statusLabels: { draft: 'Draft', published: 'Terbit' } as Record<string, string>,
}
