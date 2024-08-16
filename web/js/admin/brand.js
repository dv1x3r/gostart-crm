import { w2grid } from 'w2ui/dist/w2ui.es6'
import * as utils from './utils'

export function createBrandGrid() {
  return new w2grid({
    name: 'brandGrid',
    url: {
      get: '/brand',
      save: '/brand/save',
      remove: '/brand/delete',
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
      { field: 'id', text: 'ID', size: '60px', hidden: true },
      { field: 'name', text: 'Name', size: '340px', render: 'safe', editable: { type: 'text' }, sortable: true },
      { field: 'related_products', text: '# Products', size: '90px', render: 'int', tooltip: 'Number of products related to this brand', sortable: true },
    ],
    searches: [
      { field: 'code', label: 'Code', type: 'text', _all: true },
      { field: 'name', label: 'Name', type: 'text', _all: true },
      { field: 'related_products', label: '# Products', type: 'float' },
    ],
    defaultOperator: {
      'text': 'contains',
    },
    sortData: [
      { field: 'name', direction: 'asc' },
    ],
    onSearch: event => utils.gridSearchAllowedAll(event),
    onAdd: async event => await utils.gridNewRowAdd(event),
    onChange: async event => await utils.gridNewRowChange(event),
    onSave: async event => await utils.gridShowSaved(event),
    onDelete: async event => await utils.gridShowDeleted(event),
  })
}

