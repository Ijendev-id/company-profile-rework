/** Mengubah path gambar dari API (/uploads/xxx.jpg) menjadi URL lengkap ke server backend. */
export function useMedia() {
  const mediaBase = useRuntimeConfig().public.apiBase.replace(/\/+$/, '')

  const mediaUrl = (path?: string | null): string => {
    if (!path) return ''
    return path.startsWith('/uploads/') ? `${mediaBase}${path}` : path
  }

  return { mediaBase, mediaUrl }
}
