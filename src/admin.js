import 'w2ui/w2ui-2.0.min.css'
import { w2ui, w2layout, w2toolbar, w2sidebar, w2grid, query, w2utils } from 'w2ui/w2ui-2.0.es6'

window.w2ui = w2ui

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
      { type: 'button', id: 'todos', text: 'Todos', icon: 'w2ui-icon-colors', onClick: showTodos },
      { type: 'break' },
      { type: 'button', id: 'info', text: 'Info', icon: 'w2ui-icon-info', onClick: () => { action('html', 'left', 'Sidebar!'); grid.render('#grid'); } },
      { type: 'break' },
      { type: 'button', id: 'settings', text: 'Settings', icon: 'w2ui-icon-settings', onClick: () => action('html', 'left', 'Settings!') },
      { type: 'spacer' },
      { type: 'button', id: 'logout', text: 'Logout', icon: 'w2ui-icon-cross', onClick: () => action('html', 'left', 'Logout!') },
    ],
  },
}

async function showTodos() {
  await layout.load('main', 'todos')
  grid.render('#grid')
  grid.load('data')
}

let grid = new w2grid({
  name: 'grid',
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

let layout = new w2layout(config.layout)
layout.html('top', new w2toolbar(config.toolbar))
showTodos()

