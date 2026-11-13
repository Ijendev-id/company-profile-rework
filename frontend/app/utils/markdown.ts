import MarkdownIt from 'markdown-it'

// html: false -> tag HTML mentah di dalam artikel tidak diproses (ditampilkan sebagai teks),
// sehingga script tidak bisa disisipkan lewat isi artikel. Link berskema berbahaya
// (javascript:, vbscript:, dll) juga ditolak otomatis oleh markdown-it.
const md = new MarkdownIt({ html: false, linkify: true, typographer: false, breaks: false })

type RenderEnv = {
  /** Alamat backend untuk gambar berpath /uploads/... */
  mediaBase?: string
}

// Link ke situs lain dibuka di tab baru.
const defaultLinkOpen =
  md.renderer.rules.link_open ?? ((tokens, idx, options, _env, self) => self.renderToken(tokens, idx, options))

md.renderer.rules.link_open = (tokens, idx, options, env, self) => {
  const token = tokens[idx]!
  if (/^https?:\/\//i.test(String(token.attrGet('href') ?? ''))) {
    token.attrSet('target', '_blank')
    token.attrSet('rel', 'noopener noreferrer')
  }
  return defaultLinkOpen(tokens, idx, options, env, self)
}

// Gambar: dimuat saat mendekati layar, dan path /uploads/... diarahkan ke server backend.
const defaultImage =
  md.renderer.rules.image ?? ((tokens, idx, options, _env, self) => self.renderToken(tokens, idx, options))

md.renderer.rules.image = (tokens, idx, options, env, self) => {
  const token = tokens[idx]!
  const src = String(token.attrGet('src') ?? '')
  const mediaBase = typeof env?.mediaBase === 'string' ? env.mediaBase : ''
  if (mediaBase && src.startsWith('/uploads/')) {
    token.attrSet('src', mediaBase + src)
  }
  token.attrSet('loading', 'lazy')
  return defaultImage(tokens, idx, options, env, self)
}

/** Mengubah Markdown menjadi HTML aman untuk ditampilkan dengan v-html. */
export function renderMarkdown(source: string, env: RenderEnv = {}): string {
  return md.render(source, env)
}
