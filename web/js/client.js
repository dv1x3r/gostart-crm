import Htmx from 'htmx.org'
window.htmx = Htmx
window.htmx.config.includeIndicatorStyles = false
window.htmx.config.scrollBehavior = 'smooth'

import Alpine from 'alpinejs'
import collapse from '@alpinejs/collapse'
window.Alpine = Alpine
Alpine.plugin(collapse)
Alpine.start()

