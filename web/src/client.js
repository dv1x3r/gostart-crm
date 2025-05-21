import Htmx from 'htmx.org'
window.htmx = Htmx

// https://www.youtube.com/watch?v=9AtijVV11SA
Htmx.config.disableInheritance = true

// styles are part of the main.css
Htmx.config.includeIndicatorStyles = false

// convenient scroll experience
Htmx.config.scrollIntoViewOnBoost = false
Htmx.config.scrollBehavior = 'smooth'

// https://github.com/bigskysoftware/htmx/issues/2910
Htmx.config.historyCacheSize = 0
Htmx.config.refreshOnHistoryMiss = true

document.body.addEventListener('htmx:beforeSwap', function(evt) {
  // responses with HX-Force-Swap header should swap
  if (evt.detail.xhr.getResponseHeader('HX-Force-Swap')) {
    evt.detail.shouldSwap = true
    evt.detail.isError = false
  }

  // trigger alpine destruction for the current DOM
  if (evt.target && evt.detail.shouldSwap) {
    Alpine.destroyTree(evt.target);
  }
})

document.body.addEventListener('htmx:afterSwap', function(evt) {
  // re-initialize alpine after htmx swaps in new content
  if (evt.target) {
    Alpine.initTree(evt.target);
  }
})

// grammar utilities
window.makePlural = (word, count) => (count == -1 || count == 1) ? word : word + 's'
window.getHaveOrHas = count => count == -1 || count == 1 ? 'has' : 'have'

import Alpine from 'alpinejs'
import anchor from '@alpinejs/anchor'
import collapse from '@alpinejs/collapse'
import focus from '@alpinejs/focus'

window.Alpine = Alpine
Alpine.plugin(anchor)
Alpine.plugin(collapse)
Alpine.plugin(focus)

import htmxRequestObserver from './alpine/htmx'
import tableComponent from './alpine/table_component'
import tableModel from './alpine/table_model'
import sse from './alpine/sse'

Alpine.data('htmxRequestObserver', htmxRequestObserver)
Alpine.data('tableComponent', tableComponent)
Alpine.data('tableModel', tableModel)
Alpine.data('sse', sse)

Alpine.start()

