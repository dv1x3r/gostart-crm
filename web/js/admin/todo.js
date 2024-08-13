import { w2grid } from 'w2ui/dist/w2ui.es6'
import * as utils from './utils'

export function createTodoGrid() {
  return new w2grid({
    name: 'todoGrid',
    url: {
      get: '/todo',
      save: '/todo/save',
      remove: '/todo/delete',
    },
    httpHeaders: { 'X-CSRF-Token': utils.getCsrfToken() },
    recid: 'id',
    limit: 1000,
    recordHeight: 30,
    multiSearch: true,
    reorderRows: false,
    show: {
      footer: true,
      toolbar: true,
      toolbarAdd: true,
      toolbarEdit: false,
      toolbarDelete: true,
      toolbarSave: true,
      toolbarSearch: true,
      toolbarReload: true,
      searchSave: false,
    },
    columns: [
      { field: 'id', text: 'ID', size: '60px', sortable: true },
      { field: 'name', text: 'Name', size: '120px', render: 'safe', editable: { type: 'text' }, sortable: true },
      { field: 'description', text: 'Description', size: '300px', render: 'safe', editable: { type: 'text' }, sortable: true },
      { field: 'created_at', text: 'Created at', size: '135px', render: 'datetime', sortable: true },
      { field: 'updated_at', text: 'Updated at', size: '135px', render: 'datetime', sortable: true },
    ],
    defaultOperator: {
      'text': 'contains',
    },
    onAdd: async event => await utils.gridNewRowAdd(event),
    onChange: async event => await utils.gridNewRowChange(event),
    onSave: async event => await utils.gridShowSaved(event),
    onDelete: async event => await utils.gridShowDeleted(event),
  })
}

