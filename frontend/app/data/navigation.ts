export interface NavItem {
  label: string
  /** Anchor di halaman utama, contoh '#about' (scroll halus, aktif lewat scroll spy). */
  hash?: string
  /** Halaman lain, contoh '/blog'. */
  to?: string
}

export const navItems: NavItem[] = [
  { label: 'Home', hash: '#home' },
  { label: 'About', hash: '#about' },
  { label: 'Why', hash: '#facts' },
  { label: 'Blog', to: '/blog' },
]
