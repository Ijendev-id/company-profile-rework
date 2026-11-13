// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2025-07-15',
  devtools: { enabled: true },

  // Komponen di subfolder dipakai tanpa prefix nama folder:
  // components/home/HomeHero.vue -> <HomeHero />
  components: [{ path: '~/components', pathPrefix: false }],

  // CSS tambahan khusus halaman blog (CSS template tetap dimuat lewat app.head.link di bawah).
  css: ['~/assets/css/blog.css', '~/assets/css/admin.css'],

  // Halaman admin dirender di browser saja (tanpa SSR) dan dilarang diindeks mesin pencari.
  routeRules: {
    '/admin/**': { ssr: false, headers: { 'X-Robots-Tag': 'noindex, nofollow' } },
  },

  // Alamat API backend Go. Bisa ditimpa lewat environment variable atau file .env di folder frontend:
  //   NUXT_API_BASE         -> dipakai server Nuxt saat render (boleh alamat internal)
  //   NUXT_PUBLIC_API_BASE  -> dipakai browser dan untuk URL gambar (harus alamat yang bisa dibuka pengunjung)
  runtimeConfig: {
    apiBase: 'http://localhost:8080',
    public: {
      apiBase: 'http://localhost:8080',
    },
  },

  app: {
    head: {
      htmlAttrs: { lang: 'id' },
      charset: 'utf-8',
      viewport: 'width=device-width, initial-scale=1',
      // Semua aset template ada di public/assets_landing.
      // Meta SEO per halaman diatur lewat usePageSeo() (lihat app/composables).
      link: [
        { rel: 'icon', type: 'image/png', href: '/assets_landing/images/logo/ijendev.png' },
        { rel: 'stylesheet', href: '/assets_landing/css/animate.css' },
        { rel: 'stylesheet', href: '/assets_landing/css/glightbox.min.css' },
        { rel: 'stylesheet', href: '/assets_landing/css/lineIcons.css' },
        { rel: 'stylesheet', href: '/assets_landing/css/bootstrap.min.css' },
        { rel: 'stylesheet', href: '/assets_landing/css/style.css' },
      ],
    },
  },
})
