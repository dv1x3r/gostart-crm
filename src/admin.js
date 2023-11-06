import 'w2ui/w2ui-2.0.css'
import { w2ui, w2layout, w2toolbar, w2sidebar, w2grid, query, w2utils } from 'w2ui/w2ui-2.0.es6'
import { grid } from './admin/gridTodo'

window.w2ui = w2ui

let config = {
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
  grid.load('api/todo')
}


let layout = new w2layout({
  name: 'layout',
  box: '#page',
  panels: [
    { type: 'top', size: 40, style: 'border: 1px solid #efefef;', html: new w2toolbar(config.toolbar) },
    { type: 'left', size: 260, style: 'border: 1px solid #efefef;', html: "test"},
    { type: 'main', style: 'border: 1px solid #efefef;' },
  ]
})

showTodos()

