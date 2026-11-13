/** Tanggal berbahasa Indonesia, contoh: "12 November 2026". Zona waktu dikunci agar server dan browser menampilkan hal yang sama. */
export function formatDate(value?: string | null): string {
  if (!value) return ''

  const date = new Date(value)
  if (Number.isNaN(date.getTime())) return ''

  return new Intl.DateTimeFormat('id-ID', { dateStyle: 'long', timeZone: 'Asia/Jakarta' }).format(date)
}
