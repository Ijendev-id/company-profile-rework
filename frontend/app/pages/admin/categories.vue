<script setup lang="ts">
import type { Category } from '~/types/blog'

definePageMeta({ layout: 'admin', middleware: 'auth' })
useHead({ title: 'Kategori' })

const api = useAdminApi()
const blogApi = useBlogApi()

const { data, refresh } = useAsyncData('admin-categories', () => blogApi.categories())
const categories = computed(() => data.value?.data ?? [])

const success = ref('')
const problem = ref('')
const create = useFormErrors()
const edit = useFormErrors()

// --- Tambah kategori ---
const newName = ref('')
const newSlug = ref('')
const creating = ref(false)

async function add() {
  create.clear()
  success.value = problem.value = ''
  creating.value = true
  try {
    await api.post('/categories', { name: newName.value.trim(), slug: newSlug.value.trim() })
    newName.value = newSlug.value = ''
    success.value = 'Kategori berhasil ditambahkan.'
    await refresh()
  } catch (error) {
    create.apply(error)
  } finally {
    creating.value = false
  }
}

// --- Ubah kategori (satu baris dalam mode ubah pada satu waktu) ---
const editingId = ref<number | null>(null)
const editName = ref('')
const editSlug = ref('')
const saving = ref(false)

function startEdit(category: Category) {
  edit.clear()
  editingId.value = category.id
  editName.value = category.name
  editSlug.value = category.slug
}

async function saveEdit() {
  if (editingId.value === null) return
  edit.clear()
  success.value = problem.value = ''
  saving.value = true
  try {
    await api.put(`/categories/${editingId.value}`, { name: editName.value.trim(), slug: editSlug.value.trim() })
    editingId.value = null
    success.value = 'Kategori berhasil diubah.'
    await refresh()
  } catch (error) {
    edit.apply(error)
  } finally {
    saving.value = false
  }
}

// --- Hapus kategori ---
async function remove(category: Category) {
  if (!window.confirm(`Hapus kategori "${category.name}"?\n\nArtikel di kategori ini tidak ikut terhapus, hanya menjadi tanpa kategori.`)) return

  success.value = problem.value = ''
  try {
    await api.del(`/categories/${category.id}`)
    success.value = 'Kategori berhasil dihapus.'
    await refresh()
  } catch (error) {
    problem.value = toApiError(error).message
  }
}
</script>

<template>
  <div>
    <div class="admin-page-head">
      <h1>Kategori</h1>
    </div>

    <AdminAlert v-if="success" type="success">{{ success }}</AdminAlert>
    <AdminAlert v-if="problem" type="danger">{{ problem }}</AdminAlert>

    <form class="admin-card" novalidate @submit.prevent="add">
      <h2>Tambah kategori</h2>
      <AdminAlert v-if="create.message.value" type="danger">{{ create.message.value }}</AdminAlert>

      <div class="row">
        <div class="col-md-5 admin-field">
          <label for="cat-name">Nama</label>
          <input id="cat-name" v-model="newName" type="text" class="form-control" :class="{ 'is-invalid': create.errors.value.name }" maxlength="100" />
          <span v-if="create.errors.value.name" class="admin-error">{{ create.errors.value.name }}</span>
        </div>
        <div class="col-md-5 admin-field">
          <label for="cat-slug">Slug (opsional)</label>
          <input id="cat-slug" v-model="newSlug" type="text" class="form-control" :class="{ 'is-invalid': create.errors.value.slug }" maxlength="120" placeholder="otomatis dari nama" />
          <span v-if="create.errors.value.slug" class="admin-error">{{ create.errors.value.slug }}</span>
        </div>
        <div class="col-md-2 admin-field d-flex align-items-end">
          <button type="submit" class="btn btn-primary w-100" :disabled="creating">{{ creating ? 'Menyimpan...' : 'Tambah' }}</button>
        </div>
      </div>
    </form>

    <div class="table-responsive">
      <table class="admin-table">
        <thead>
          <tr>
            <th>Nama</th>
            <th>Slug</th>
            <th></th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="category in categories" :key="category.id">
            <template v-if="editingId === category.id">
              <td>
                <input v-model="editName" type="text" class="form-control form-control-sm" :class="{ 'is-invalid': edit.errors.value.name }" maxlength="100" aria-label="Nama kategori" />
                <span v-if="edit.errors.value.name" class="admin-error">{{ edit.errors.value.name }}</span>
              </td>
              <td>
                <input v-model="editSlug" type="text" class="form-control form-control-sm" :class="{ 'is-invalid': edit.errors.value.slug }" maxlength="120" aria-label="Slug kategori" />
                <span v-if="edit.errors.value.slug" class="admin-error">{{ edit.errors.value.slug }}</span>
                <span v-if="edit.message.value && !edit.errors.value.slug && !edit.errors.value.name" class="admin-error">{{ edit.message.value }}</span>
              </td>
              <td>
                <div class="admin-actions">
                  <button type="button" class="btn btn-sm btn-primary" :disabled="saving" @click="saveEdit">Simpan</button>
                  <button type="button" class="btn btn-sm btn-outline-secondary" @click="editingId = null">Batal</button>
                </div>
              </td>
            </template>
            <template v-else>
              <td class="admin-title">{{ category.name }}</td>
              <td>{{ category.slug }}</td>
              <td>
                <div class="admin-actions">
                  <button type="button" class="btn btn-sm btn-outline-primary" @click="startEdit(category)">Ubah</button>
                  <button type="button" class="btn btn-sm btn-outline-danger" @click="remove(category)">Hapus</button>
                </div>
              </td>
            </template>
          </tr>
          <tr v-if="!categories.length && data">
            <td colspan="3" class="text-center text-muted py-4">Belum ada kategori.</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
