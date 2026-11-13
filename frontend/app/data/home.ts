/** Seluruh isi teks & gambar halaman utama. Ubah konten cukup di file ini. */

export const hero = {
  subtitle: 'Bangun Solusi Digital Tanpa Ribet',
  title: 'Website, Aplikasi Mobile, dan IoT untuk percepat pertumbuhan bisnismu.',
  text: 'Desain modern & cepat diakses | Pengembangan aman dan scalable | Dukungan after-sales yang responsif',
  ctaLabel: 'Konsultasi Gratis',
  background: 'images/header/banner-bg.svg',
  image: 'images/header/header.png',
}

export const services = {
  title: 'Saatnya melangkah ke level baru,',
  highlight: 'dalam transformasi digital bisnis Anda',
  items: [
    {
      icon: 'lni-baloon',
      shape: 'images/services/services-shape-1.svg',
      title: 'Dampak Nyata',
      text: 'Teknologi tanpa arah hanyalah alat kosong. Di IjenDev.id, setiap baris kode kami tulis dengan tujuan: membantu bisnis Anda tumbuh, memperluas pengaruh, dan menciptakan nilai berkelanjutan bagi masa depan.',
      delay: 0.2,
    },
    {
      icon: 'lni-cog',
      shape: 'images/services/services-shape-2.svg',
      title: 'Profesional Tim',
      text: 'Inovasi terbaik lahir dari kerja sama. Bersama kami, Anda tidak hanya bekerja dengan developer — tetapi dengan tim kreatif yang terdiri dari perancang, pemikir, dan engineer yang berkomitmen pada keberhasilan visi Anda.',
      delay: 0.5,
    },
    {
      icon: 'lni-bolt-alt',
      shape: 'images/services/services-shape-3.svg',
      title: 'Proses Jelas & Terbuka',
      text: 'Kami percaya hasil terbaik dimulai dari transparansi. Setiap tahap pengembangan disusun dengan rencana yang disepakati bersama, sehingga Anda tahu arah, waktu, dan tujuan di setiap langkahnya.',
      delay: 0.8,
    },
  ],
}

export const about = {
  title: 'Masa depan tidak menunggu',
  highlight: '— ia dibangun dari inovasi hari ini',
  paragraphs: [
    'Itulah semangat yang mendorong Ijen Developer Indonesia (IjenDev.id) untuk terus berkreasi menghadirkan solusi digital yang berdampak nyata.',
    'Kami adalah software house teknologi informasi yang berfokus pada pengembangan aplikasi berbasis website, mobile, dan desktop, serta integrasi sistem berbasis IoT dan digital marketing. Berdiri sejak 2025, kami hadir sebagai mitra bagi bisnis, institusi, dan individu yang ingin tumbuh dan beradaptasi di era digital.',
    'Dengan tim yang solid, profesional, dan penuh semangat, kami tidak hanya membangun sistem — kami merancang pengalaman digital yang membantu Anda mencapai tujuan dengan cara yang lebih cepat, cerdas, dan efisien.',
  ],
  buttonLabel: 'Kisah Kami Selengkapnya',
  image: 'images/about/about1.svg',
  shape: 'images/about/about-shape-1.svg',
}

export const facts = {
  title: 'Fakta Menarik',
  highlight: 'Tentang Ijendev.id',
  text: 'Dibangun oleh tim muda kreatif dari Indonesia, Ijendev.id menghadirkan layanan pembuatan website, aplikasi, dan sistem digital yang cepat, aman, dan sesuai kebutuhan bisnis Anda.',
  // append: teks di belakang angka. Kosongkan ('') untuk menghilangkan sufiks bawaan "k";
  // hapus properti append untuk memakai sufiks bawaan.
  counters: [
    { value: 15, label: 'Clients', color: 1, append: '' as string | undefined },
    { value: 10, label: 'Active Users', color: 2, append: undefined as string | undefined },
    { value: 25, label: 'Project', color: 3, append: '' as string | undefined },
  ],
  video: {
    url: 'https://www.youtube.com/watch?v=r44RKWyfcFw',
    image: 'images/video/video.png',
    dots: 'images/video/dots.svg',
  },
}

export const clients = {
  title: 'Mereka sudah bertransformasi bersama IjenDev.id,',
  highlight: 'Kini giliran bisnis Anda untuk naik level ke dunia digital',
  logos: [
    { name: 'Adzdzikra', image: 'images/brands/adzdzikra.png', duration: 1, delay: 0.2 },
    { name: 'PLN', image: 'images/brands/pln.png', duration: 1.5, delay: 0.2 },
    { name: 'Polinema', image: 'images/brands/polinema.png', duration: 1.5, delay: 0.3 },
    { name: 'Malang Kab', image: 'images/brands/malangkab.png', duration: 1.5, delay: 0.4 },
    { name: 'E-commerce', image: 'images/brands/ecommerce-html.svg', duration: 1.5, delay: 0.5 },
  ],
}
