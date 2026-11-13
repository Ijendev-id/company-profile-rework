import { site } from '~/data/site'

export interface PageSeo {
  title: string
  description: string
  keywords?: string
  /** URL absolut gambar share; bawaan: logo situs */
  image?: string
  /** URL absolut halaman; bawaan: URL situs */
  url?: string
  type?: 'website' | 'article'
  /** Isi khusus Open Graph / Twitter jika berbeda dari title & description */
  og?: { title?: string; description?: string }
  twitter?: { title?: string; description?: string }
}

/** Atur <title> dan meta SEO (description, Open Graph, Twitter) untuk satu halaman. */
export function usePageSeo(seo: PageSeo) {
  const image = seo.image ?? `${site.url}${asset(site.logo.share)}`
  const ogTitle = seo.og?.title ?? seo.title
  const ogDescription = seo.og?.description ?? seo.description

  useSeoMeta({
    title: seo.title,
    description: seo.description,
    ogTitle,
    ogDescription,
    ogImage: image,
    ogUrl: seo.url ?? site.url,
    ogType: seo.type ?? 'website',
    ogSiteName: site.name,
    twitterCard: 'summary_large_image',
    twitterTitle: seo.twitter?.title ?? ogTitle,
    twitterDescription: seo.twitter?.description ?? ogDescription,
    twitterImage: image,
  })

  useHead({
    meta: [
      { name: 'author', content: site.author },
      ...(seo.keywords ? [{ name: 'keywords', content: seo.keywords }] : []),
    ],
  })
}
