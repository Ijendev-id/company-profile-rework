/** Pesan singkat yang bertahan satu kali pindah halaman (contoh: "Artikel berhasil disimpan"). */
export function useFlash() {
  const message = useState<string>('admin-flash', () => '')

  const set = (text: string) => {
    message.value = text
  }

  /** Mengambil pesan lalu menghapusnya, supaya tidak tampil dua kali. */
  const take = () => {
    const text = message.value
    message.value = ''
    return text
  }

  return { set, take }
}
