/** Identitas, kontak, dan logo situs. Ubah di sini, berlaku di semua halaman. */
export const site = {
  name: 'Ijendev.id',
  url: 'https://ijendev.id',
  author: 'Ijen Developer Indonesia',

  logo: {
    default: 'images/logo/iii.png',
    sticky: 'images/logo/a.png', // dipakai navbar saat halaman di-scroll
    share: 'images/logo/ijendev.png', // gambar untuk og:image / twitter:image
  },

  // Dua nomor WhatsApp ini sesuai kode lama: "primary" dipakai tombol navbar,
  // CTA footer, dan kontak footer; "secondary" dipakai tombol hero dan ikon sosial footer.
  whatsapp: {
    primary: '6281235026920',
    secondary: '6282139590210',
  },

  contact: {
    phone: '+62 812-3502-6920',
    address: ['Dsn Krajan, Ds Aliyan, Kec Rogojampi', 'Kab Banyuwangi, Jawa Timur, 68462.'],
  },

  social: [
    { label: 'WhatsApp', icon: 'lni-whatsapp', href: waMe('6282139590210') },
    { label: 'Instagram', icon: 'lni-instagram-filled', href: 'https://www.instagram.com/ijendev.id/' },
    {
      label: 'LinkedIn',
      icon: 'lni-linkedin-original',
      href: 'https://www.linkedin.com/company/ijendev-id-ijend-developer-indonesia',
    },
  ],
}
