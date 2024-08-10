import { w2ui, w2layout, w2toolbar, w2utils } from 'w2ui/dist/w2ui.es6'
import { todoGrid } from './admin/todo'

window.w2ui = w2ui

w2utils.settings.dataType = 'JSON'
w2utils.formatters['safe'] = value => w2utils.encodeTags(value) ?? ""

const mainLayout = new w2layout({
  name: 'mainLayout',
  box: '#main-layout',
  panels: [
    {
      type: 'top', size: 40, toolbar: {
        items: [
          {
            type: 'button', id: 'todos', text: 'Todos', icon: 'w2ui-icon-colors',
            onClick: async () => {
              // await page.load('main', 'admin/todo')
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
      }
    },
    // { type: 'left', size: 260 },
    { type: 'main' },
  ]
})


