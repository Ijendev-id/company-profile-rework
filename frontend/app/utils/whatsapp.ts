/** Link WhatsApp gaya api.whatsapp.com (dipakai tombol navbar & hero). */
export const waApi = (phone: string) =>
  `https://api.whatsapp.com/send/?phone=${phone}&text&type=phone_number&app_absent=0`

/** Link WhatsApp gaya wa.me, opsional dengan teks pesan awal. */
export const waMe = (phone: string, text?: string) =>
  `https://wa.me/${phone}${text ? `?text=${encodeURIComponent(text)}` : ''}`
