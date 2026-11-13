const cache = new Map<string, Promise<void>>()

/** Muat file <script> satu kali saja; panggilan berikutnya memakai hasil yang sama. */
export function loadScript(src: string): Promise<void> {
  const cached = cache.get(src)
  if (cached) return cached

  const promise = new Promise<void>((resolve, reject) => {
    const el = document.createElement('script')
    el.src = src
    el.async = true
    el.onload = () => resolve()
    el.onerror = () => {
      cache.delete(src)
      reject(new Error(`Gagal memuat script: ${src}`))
    }
    document.body.appendChild(el)
  })

  cache.set(src, promise)
  return promise
}
