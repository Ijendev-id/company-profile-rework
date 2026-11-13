<script setup lang="ts">
import { admin } from '~/data/admin'

// Editor Markdown: kotak teks + tombol format + pratinjau. Nilainya teks Markdown.
const model = defineModel<string>({ required: true })
defineProps<{ error?: string }>()

const api = useAdminApi()
const { mediaBase } = useMedia()

const textarea = ref<HTMLTextAreaElement>()
const tab = ref<'write' | 'preview'>('write')
const uploading = ref(false)
const problem = ref('')

// Pratinjau memakai pengubah Markdown yang sama dengan halaman blog publik.
const preview = computed(() => renderMarkdown(model.value, { mediaBase }))

function replaceRange(start: number, end: number, text: string, selectFrom: number, selectTo: number) {
  model.value = model.value.slice(0, start) + text + model.value.slice(end)
  nextTick(() => {
    textarea.value?.focus()
    textarea.value?.setSelectionRange(start + selectFrom, start + selectTo)
  })
}

/** Mengapit teks terpilih, contoh: **tebal** */
function wrap(before: string, after = before, placeholder = 'teks') {
  const el = textarea.value
  if (!el) return
  const { selectionStart: start, selectionEnd: end } = el
  const selected = model.value.slice(start, end) || placeholder
  replaceRange(start, end, before + selected + after, before.length, before.length + selected.length)
}

/** Menambah awalan di setiap baris terpilih, contoh: "## " atau "- " */
function prefixLines(prefix: (index: number) => string) {
  const el = textarea.value
  if (!el) return
  const { selectionStart: start, selectionEnd: end } = el
  const lineStart = model.value.lastIndexOf('\n', start - 1) + 1
  const block = model.value.slice(lineStart, end) || 'teks'
  const result = block
    .split('\n')
    .map((line, index) => prefix(index) + line)
    .join('\n')
  replaceRange(lineStart, end, result, 0, result.length)
}

function insertAtCursor(text: string) {
  const el = textarea.value
  const position = el ? el.selectionEnd : model.value.length
  replaceRange(position, position, text, text.length, text.length)
}

async function onPickImage(event: Event) {
  const input = event.target as HTMLInputElement
  const file = input.files?.[0]
  input.value = ''
  if (!file) return

  problem.value = ''
  uploading.value = true
  try {
    const url = await api.uploadImage(file)
    insertAtCursor(`\n\n![deskripsi gambar](${url})\n\n`)
  } catch (error) {
    problem.value = toApiError(error).message
  } finally {
    uploading.value = false
  }
}
</script>

<template>
  <div>
    <div class="admin-editor" :class="{ 'is-invalid': error }">
      <div class="admin-editor-bar">
        <template v-if="tab === 'write'">
          <button type="button" class="btn btn-sm btn-outline-secondary" title="Tebal" @click="wrap('**')"><b>B</b></button>
          <button type="button" class="btn btn-sm btn-outline-secondary" title="Miring" @click="wrap('*')"><i>I</i></button>
          <button type="button" class="btn btn-sm btn-outline-secondary" title="Judul besar" @click="prefixLines(() => '## ')">H2</button>
          <button type="button" class="btn btn-sm btn-outline-secondary" title="Judul sedang" @click="prefixLines(() => '### ')">H3</button>
          <button type="button" class="btn btn-sm btn-outline-secondary" title="Daftar poin" @click="prefixLines(() => '- ')">Poin</button>
          <button type="button" class="btn btn-sm btn-outline-secondary" title="Daftar angka" @click="prefixLines((i) => `${i + 1}. `)">1.2.</button>
          <button type="button" class="btn btn-sm btn-outline-secondary" title="Kutipan" @click="prefixLines(() => '> ')">Kutip</button>
          <button type="button" class="btn btn-sm btn-outline-secondary" title="Kode" @click="wrap('`')">Kode</button>
          <button type="button" class="btn btn-sm btn-outline-secondary" title="Tautan" @click="wrap('[', '](https://)', 'teks tautan')">Link</button>
          <label class="btn btn-sm btn-outline-secondary mb-0" :class="{ disabled: uploading }" title="Sisipkan gambar">
            {{ uploading ? 'Mengunggah...' : 'Gambar' }}
            <input
              type="file"
              class="d-none"
              :accept="admin.allowedImageTypes.join(',')"
              :disabled="uploading"
              data-testid="editor-image-input"
              @change="onPickImage"
            />
          </label>
        </template>

        <div class="admin-editor-tabs">
          <button type="button" class="btn btn-sm" :class="tab === 'write' ? 'btn-primary' : 'btn-outline-secondary'" @click="tab = 'write'">Tulis</button>
          <button type="button" class="btn btn-sm" :class="tab === 'preview' ? 'btn-primary' : 'btn-outline-secondary'" @click="tab = 'preview'">Pratinjau</button>
        </div>
      </div>

      <textarea v-show="tab === 'write'" ref="textarea" v-model="model" placeholder="Tulis isi artikel dalam Markdown..." spellcheck="true"></textarea>

      <!-- Aman: renderMarkdown() menonaktifkan HTML mentah -->
      <!-- eslint-disable-next-line vue/no-v-html -->
      <div v-if="tab === 'preview'" class="admin-editor-preview blog-article-content" v-html="preview || '<p><em>Belum ada isi.</em></p>'"></div>
    </div>

    <span v-if="problem || error" class="admin-error">{{ problem || error }}</span>
  </div>
</template>
