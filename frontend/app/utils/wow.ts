/**
 * Atribut animasi WOW.js, dipakai dengan v-bind supaya tidak menulis ulang
 * class + data-wow-* di setiap elemen.
 *
 *   <div class="mt-30" v-bind="wow('fadeInUp', 1.3, 0.2)">
 *   -> class="mt-30 wow fadeInUp" data-wow-duration="1.3s" data-wow-delay="0.2s"
 */
export function wow(animation = 'fadeIn', duration = 1, delay = 0) {
  return {
    class: `wow ${animation}`,
    'data-wow-duration': `${duration}s`,
    'data-wow-delay': `${delay}s`,
  }
}
