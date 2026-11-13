// Library bawaan template (assets_landing/js) yang dipasang sebagai global di window.
export {}

declare global {
  interface Window {
    WOW: new (options?: Record<string, unknown>) => { init: () => void }
    GLightbox: (options?: Record<string, unknown>) => { destroy?: () => void }
    counterUp: new (options?: Record<string, unknown>) => { start: () => void }
    particlesJS: (id: string, config: Record<string, unknown>) => void
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    pJSDom?: any[]
  }
}
