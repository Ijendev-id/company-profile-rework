import type { PageSeo } from '~/composables/usePageSeo'
import { site } from '~/data/site'

/** Meta SEO per halaman. Halaman blog nanti cukup menambah entri/pemanggilan usePageSeo sendiri. */
export const seo: Record<string, PageSeo> = {
  home: {
    title: 'Ijendev.id | Jasa Pembuatan Web dan Aplikasi',
    url: `${site.url}/`,
    description:
      'Ijendev.id adalah software house profesional yang menyediakan jasa pembuatan website, aplikasi mobile, dan solusi digital terintegrasi untuk bisnis, instansi, dan UMKM di seluruh Indonesia.',
    keywords:
      'ijendev,id, software house, jasa pembuatan website, jasa pembuatan aplikasi, jasa web malang, jasa app android, digital agency indonesia',
    og: {
      title: 'Ijendev.id | Jasa Pembuatan Website & Aplikasi Profesional',
      description:
        'Bangun website dan aplikasi impian Anda bersama Ijendev.id — software house terpercaya dari Indonesia. Solusi digital, cepat, transparan, dan hasil berkualitas.',
    },
    twitter: {
      title: 'Ijendev.id | Jasa Pembuatan Website & Aplikasi Profesional',
      description:
        'Software house terpercaya untuk solusi digital, website, aplikasi, dan IoT. Bekerja cepat, transparan, dan profesional.',
    },
  },

  blog: {
    title: 'Blog | Ijendev.id',
    description: 'Artikel dan tulisan terbaru dari Ijendev.id.',
  },
}
