<script setup lang="ts">
import { admin } from '~/data/admin'
import type { AdminPost, PostPayload } from '~/types/admin'

// Form tulis/ubah artikel. Dipakai halaman "Tulis artikel" dan "Ubah artikel".
const props = defineProps<{
  initial?: AdminPost | null
  submitting: boolean
  errors: Record<string, string>
}>()

const emit = defineEmits<{ submit: [payload: PostPayload] }>()

const blogApi = useBlogApi()
const { data: categories } = useAsyncData('admin-form-categories', () => blogApi.categories())

const form = reactive({
  title: props.initial?.title ?? '',
  slug: props.initial?.slug ?? '',
  excerpt: props.initial?.excerpt ?? '',
  content: props.initial?.content ?? '',
  cover_image: props.initial?.cover_image ?? '',
  status: props.initial?.status ?? ('draft' as PostPayload['status']),
  meta_title: props.initial?.meta_title ?? '',
  meta_description: props.initial?.meta_description ?? '',
  category_id: props.initial?.category_id ?? ('' as number | ''),
})

function submit() {
  emit('submit', {
    title: form.title.trim(),
    slug: form.slug.trim(),
    excerpt: form.excerpt.trim(),
    content: form.content,
    cover_image: form.cover_image,
    status: form.status,
    meta_title: form.meta_title.trim(),
    meta_description: form.meta_description.trim(),
    category_id: form.category_id === '' ? null : Number(form.category_id),
  })
}
</script>

<template>
  <form novalidate @submit.prevent="submit">
    <div class="admin-card">
      <div class="admin-field">
        <label for="post-title">Judul</label>
        <input id="post-title" v-model="form.title" type="text" class="form-control" :class="{ 'is-invalid': errors.title }" maxlength="255" />
        <span v-if="errors.title" class="admin-error">{{ errors.title }}</span>
      </div>

      <div class="admin-field">
        <label for="post-slug">Slug (alamat URL)</label>
        <input id="post-slug" v-model="form.slug" type="text" class="form-control" :class="{ 'is-invalid': errors.slug }" maxlength="191" placeholder="otomatis dari judul" />
        <span class="admin-hint">Kosongkan untuk dibuat otomatis. Hanya huruf kecil, angka, dan tanda hubung.</span>
        <span v-if="errors.slug" class="admin-error">{{ errors.slug }}</span>
      </div>

      <div class="admin-field">
        <label for="post-excerpt">Ringkasan</label>
        <textarea id="post-excerpt" v-model="form.excerpt" class="form-control" :class="{ 'is-invalid': errors.excerpt }" rows="3" maxlength="1000"></textarea>
        <span class="admin-hint">Tampil di daftar artikel dan jadi deskripsi SEO bila deskripsi meta kosong.</span>
        <span v-if="errors.excerpt" class="admin-error">{{ errors.excerpt }}</span>
      </div>

      <div class="admin-field">
        <label>Isi artikel</label>
        <AdminMarkdownEditor v-model="form.content" :error="errors.content" />
      </div>
    </div>

    <div class="admin-card">
      <h2>Pengaturan</h2>

      <div class="row">
        <div class="col-md-6 admin-field">
          <label for="post-status">Status</label>
          <select id="post-status" v-model="form.status" class="form-select" :class="{ 'is-invalid': errors.status }">
            <option value="draft">{{ admin.statusLabels.draft }} (tidak tampil di situs)</option>
            <option value="published">{{ admin.statusLabels.published }} (tampil di situs)</option>
          </select>
          <span v-if="errors.status" class="admin-error">{{ errors.status }}</span>
        </div>

        <div class="col-md-6 admin-field">
          <label for="post-category">Kategori</label>
          <select id="post-category" v-model="form.category_id" class="form-select" :class="{ 'is-invalid': errors.category_id }">
            <option value="">Tanpa kategori</option>
            <option v-for="category in categories?.data ?? []" :key="category.id" :value="category.id">{{ category.name }}</option>
          </select>
          <span v-if="errors.category_id" class="admin-error">{{ errors.category_id }}</span>
        </div>
      </div>

      <div class="admin-field">
        <label>Gambar sampul</label>
        <AdminCoverUpload v-model="form.cover_image" :error="errors.cover_image" />
      </div>

      <details class="admin-field">
        <summary class="mb-3">Pengaturan SEO (opsional)</summary>

        <div class="admin-field">
          <label for="post-meta-title">Judul SEO</label>
          <input id="post-meta-title" v-model="form.meta_title" type="text" class="form-control" :class="{ 'is-invalid': errors.meta_title }" maxlength="255" />
          <span class="admin-hint">Kosong = memakai judul artikel.</span>
          <span v-if="errors.meta_title" class="admin-error">{{ errors.meta_title }}</span>
        </div>

        <div class="admin-field">
          <label for="post-meta-description">Deskripsi SEO</label>
          <textarea id="post-meta-description" v-model="form.meta_description" class="form-control" :class="{ 'is-invalid': errors.meta_description }" rows="2" maxlength="320"></textarea>
          <span class="admin-hint">Kosong = memakai ringkasan. Disarankan maksimal 160 karakter.</span>
          <span v-if="errors.meta_description" class="admin-error">{{ errors.meta_description }}</span>
        </div>
      </details>
    </div>

    <div class="d-flex gap-2">
      <button type="submit" class="btn btn-primary" :disabled="submitting">
        {{ submitting ? 'Menyimpan...' : 'Simpan' }}
      </button>
      <NuxtLink to="/admin/posts" class="btn btn-outline-secondary">Batal</NuxtLink>
    </div>
  </form>
</template>
