import './htmx-1.9.6.min.js'
import './w2ui-2.0.min.css'

import { w2layout, w2toolbar, w2sidebar, w2grid, query } from './w2ui-2.0.es6.min.js'

let config = {
  layout: {
    name: 'layout',
    box: '#page',
    panels: [
      { type: 'top', size: 40, style: 'border: 1px solid #efefef;' },
      { type: 'left', size: 260, style: 'border: 1px solid #efefef;' },
      { type: 'main', style: 'border: 1px solid #efefef;' },
    ]
  },
  toolbar: {
    name: 'toolbar',
    items: [
      { type: 'button', id: 'todos', text: 'Todos', icon: 'w2ui-icon-colors', onClick: () => { action('load', 'main', 'todos'); grid.render('#grid'); } },
      { type: 'break' },
      { type: 'button', id: 'info', text: 'Info', icon: 'w2ui-icon-info', onClick: () => action('html', 'left', 'Sidebar!') },
      { type: 'break' },
      { type: 'button', id: 'settings', text: 'Settings', icon: 'w2ui-icon-settings', onClick: () => action('html', 'left', 'Settings!') },
      { type: 'spacer' },
      { type: 'button', id: 'logout', text: 'Logout', icon: 'w2ui-icon-cross', onClick: () => action('html', 'left', 'Logout!') },
    ],
  },
}

let layout = new w2layout(config.layout)
layout.html('top', new w2toolbar(config.toolbar))

window.action = function(method, param1, param2) {
  layout[method](param1, param2)
}


let grid = new w2grid({
  name: 'grid',
  // box: '#grid',
  header: 'List of Names',
  reorderRows: false,
  show: {
    header: true,
    footer: true,
    toolbar: true,
    lineNumbers: true
  },
  columns: [
    { field: 'recid', text: 'ID', size: '30px' },
    { field: 'fname', text: 'First Name', size: '30%' },
    { field: 'lname', text: 'Last Name', size: '30%' },
    { field: 'email', text: 'Email', size: '40%' },
    { field: 'sdate', text: 'Start Date', size: '120px' }
  ],
  searches: [
    { type: 'int', field: 'recid', label: 'ID' },
    { type: 'text', field: 'fname', label: 'First Name' },
    { type: 'text', field: 'lname', label: 'Last Name' },
    { type: 'date', field: 'sdate', label: 'Start Date' }
  ],
  onExpand(event) {
    query('#' + event.detail.box_id).html('<div style="padding: 10px; height: 100px">Expanded content</div>')
  }
})
// grid.load('data/list.json')


// window.action = function (param1) {
//     if (param1 == 'reorderRows') {
//         grid.reorderRows = !grid.reorderRows
//     } else {
//         grid.show[param1] = !grid.show[param1]
//     }
//     grid.refresh()
// }
