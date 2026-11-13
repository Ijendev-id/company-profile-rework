/** Menampung error dari API untuk ditampilkan di form: pesan umum + pesan per field. */
export function useFormErrors() {
  const errors = ref<Record<string, string>>({})
  const message = ref('')

  const clear = () => {
    errors.value = {}
    message.value = ''
  }

  const apply = (error: unknown) => {
    const apiError = toApiError(error)
    errors.value = { ...apiError.fields }

    // Slug yang sudah dipakai (409) ditandai di kolom slug.
    if (apiError.status === 409) errors.value.slug = apiError.message

    message.value = apiError.status === 422 ? 'Periksa kembali isian yang ditandai.' : apiError.message
  }

  return { errors, message, clear, apply }
}
