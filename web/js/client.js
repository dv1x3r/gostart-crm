import Htmx from 'htmx.org'
window.htmx = Htmx

window.filtersSet = trigger => {
  const checkboxes = document.querySelectorAll('[id^="filter-cb-"]:checked')
  const filters = Array.from(checkboxes).map(c => c.value)
  document.getElementById('hx-filters').value = filters.join('_')
  if (trigger) {
    htmx.trigger('#hx-filters', 'updated')
  }
}

import Alpine from 'alpinejs'
import morph from '@alpinejs/morph'
import collapse from '@alpinejs/collapse'
window.Alpine = Alpine
Alpine.plugin(morph)
Alpine.plugin(collapse)
Alpine.start()

