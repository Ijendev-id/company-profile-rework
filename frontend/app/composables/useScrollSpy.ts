/**
 * Menentukan section mana yang sedang tampil (untuk menandai menu aktif).
 * @param hashes  daftar selector section, contoh ['#home', '#about']
 * @param offset  jarak dari atas layar (px), sama dengan kode lama (73)
 */
export function useScrollSpy(hashes: string[], offset = 73) {
  const route = useRoute()
  const scrollY = useScrollY()
  // Awalnya menu pertama (Home) aktif, sama seperti HTML lama.
  const active = ref<string | null>(hashes[0] ?? null)

  function update() {
    const pos = window.scrollY + offset
    active.value =
      hashes.find((hash) => {
        const el = document.querySelector<HTMLElement>(hash)
        if (!el) return false
        const top = el.getBoundingClientRect().top + window.scrollY
        return top <= pos && top + el.offsetHeight > pos
      }) ?? null
  }

  function reset() {
    if (window.scrollY > 0) update()
    else active.value = hashes[0] ?? null
  }

  watch(scrollY, update)
  watch(
    () => route.path,
    () => nextTick(reset),
  )
  onMounted(() => nextTick(reset))

  return active
}
