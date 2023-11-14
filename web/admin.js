import 'w2ui/w2ui-2.0.css'
import { w2ui, w2layout, w2toolbar, w2utils } from 'w2ui/w2ui-2.0.es6'
import { todoGrid } from './admin/todo'

window.w2ui = w2ui

w2utils.settings.dataType = 'JSON'

const toolbar = new w2toolbar({
  name: 'toolbar',
  items: [
    {
      type: 'button', id: 'todos', text: 'Todos', icon: 'w2ui-icon-colors',
      onClick: async () => {
        await page.load('main', 'admin/todo')
        todoGrid.render('#grid')
      }
    },
    { type: 'break' },
    {
      type: 'button', id: 'info', text: 'Info', icon: 'w2ui-icon-info',
      onClick: () => {
        action('html', 'left', 'Sidebar!')
        grid.render('#grid')
      }
    },
    { type: 'break' },
    {
      type: 'button', id: 'settings', text: 'Settings', icon: 'w2ui-icon-settings',
      onClick: () => page.html('left', 'Settings!')
    },
    { type: 'spacer' },
    {
      type: 'button', id: 'logout', text: 'Log out', icon: 'w2ui-icon-cross',
      onClick: () => page.html('left', 'Logout!')
    },
  ],
})

const page = new w2layout({
  name: 'page',
  box: '#page',
  panels: [
    { type: 'top', size: 40, html: toolbar },
    // { type: 'left', size: 260 },
    { type: 'main' },
  ]
})

toolbar.items[0].onClick()

