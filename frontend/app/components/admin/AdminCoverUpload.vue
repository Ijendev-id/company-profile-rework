<script setup lang="ts">
import { admin } from '~/data/admin'

// Gambar sampul artikel. Nilainya path hasil upload (/uploads/xxx.jpg) atau kosong.
const model = defineModel<string>({ required: true })
defineProps<{ error?: string }>()

const api = useAdminApi()
const { mediaUrl } = useMedia()

const uploading = ref(false)
const problem = ref('')

async function onPick(event: Event) {
  const input = event.target as HTMLInputElement
  const file = input.files?.[0]
  input.value = '' // agar file yang sama bisa dipilih ulang
  if (!file) return

  problem.value = ''
  uploading.value = true
  try {
    model.value = await api.uploadImage(file)
  } catch (error) {
    problem.value = toApiError(error).message
  } finally {
    uploading.value = false
  }
}
</script>

<template>
  <div>
    <div class="admin-cover">
      <img v-if="model" :src="mediaUrl(model)" alt="Pratinjau gambar sampul" />

      <div>
        <label class="btn btn-sm btn-outline-primary mb-0" :class="{ disabled: uploading }">
          {{ uploading ? 'Mengunggah...' : model ? 'Ganti gambar' : 'Pilih gambar' }}
          <input
            type="file"
            class="d-none"
            :accept="admin.allowedImageTypes.join(',')"
            :disabled="uploading"
            data-testid="cover-input"
            @change="onPick"
          />
        </label>
        <button v-if="model" type="button" class="btn btn-sm btn-outline-danger ms-2" @click="model = ''">
          Hapus
        </button>
        <span class="admin-hint">JPG, PNG, GIF, atau WebP. Maksimal {{ admin.maxUploadMB }} MB.</span>
      </div>
    </div>

    <span v-if="problem || error" class="admin-error">{{ problem || error }}</span>
  </div>
</template>
