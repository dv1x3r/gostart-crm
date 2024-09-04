import { w2grid } from 'w2ui/dist/w2ui.es6'
import * as utils from './utils'

export function createSupplierGrid() {
  return new w2grid({
    name: 'supplierGrid',
    url: {
      get: '/supplier',
      save: '/supplier/save',
      remove: '/supplier/delete',
    },
    httpHeaders: { 'X-CSRF-Token': utils.getCsrfToken() },
    recid: 'id',
    multiSearch: true,
    reorderRows: true,
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
      { field: 'id', text: 'ID', size: '60px', hidden: true },
      { field: 'code', text: 'Code', size: '88px', render: 'safe', editable: { type: 'text' } },
      { field: 'name', text: 'Name', size: '120px', render: 'safe', editable: { type: 'text' } },
      { field: 'description', text: 'Description', size: '300px', render: 'safe', editable: { type: 'text' } },
      { field: 'published_products', text: '# Pub', size: '60px', render: 'int', tooltip: 'Number of published products', sortable: true },
      { field: 'related_products', text: '# All', size: '60px', render: 'int', tooltip: 'Total number of products in the local catalog', sortable: true },
      { field: 'is_published', text: 'Is Pub', size: '60px', editable: { type: 'checkbox' }, tooltip: 'Show this supplier' },
    ],
    searches: [
      { field: 'code', label: 'Code', type: 'text', _all: true },
      { field: 'name', label: 'Name', type: 'text', _all: true },
      { field: 'description', label: 'Description', type: 'text', _all: true },
      { field: 'is_published', label: 'Is Published', type: 'enum', options: utils.getSelectOptionsBool() },
      { field: 'related_products', label: '# Products', type: 'float' },
      { field: 'published_products', label: '# Published', type: 'float' },
    ],
    defaultOperator: {
      'text': 'contains',
    },
    onSearch: event => utils.gridSearchAllowedAll(event),
    onAdd: async event => {
      const initial = { is_published: true }
      await utils.gridNewRowAdd(event, initial)
    },
    onChange: async event => await utils.gridNewRowChange(event),
    onSave: async event => await utils.gridShowSaved(event),
    onDelete: async event => await utils.gridShowDeleted(event),
    onReorderRow: async event => await utils.gridPostReorderRow(event, '/supplier/reorder'),
    onLoad: async event => await utils.gridMarkRows(event, x => !x.is_published ? 'darkgrey' : null),
  })
}

