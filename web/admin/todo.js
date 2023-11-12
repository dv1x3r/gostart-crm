import { w2grid, w2alert, w2popup, w2form } from 'w2ui/w2ui-2.0.es6'
import { safeRender, enablePreview, disablePreview, getCookie } from './utils'


export const todoGrid = new w2grid({
  name: 'todoGrid',
  url: '/admin/todo/data',
  httpHeaders: { 'X-CSRF-Token': getCookie('_csrf') },
  recid: 'id',
  liveSearch: true,
  show: {
    footer: true,
    toolbar: true,
    toolbarAdd: true,
    toolbarEdit: true,
    toolbarDelete: true,
    toolbarSave: true,
    searchLogic: false,
  },
  columns: [
    {
      field: 'id', text: 'ID', size: '100px',
      sortable: true, clipboardCopy: true,
    },
    {
      field: 'name', text: 'Name', size: '25%',
      sortable: true, render: row => safeRender(row.name),
    },
    {
      field: 'description', text: 'Description', size: '75%',
      sortable: true, render: row => safeRender(row.description),
    },
    {
      field: 'qty', text: 'Quantity', size: '100px',
      sortable: true, editable: { type: 'float' },
    },
    {
      text: 'Summary', size: '120px',
      info: {
        showEmpty: true,
        showOn: 'mouseenter',
        options: { position: 'left' },
        render: rec => `<b>${safeRender(rec.name)}</b>: ${safeRender(rec.description)}`,
      },
      render: () => '<span class="text-slate-400">Mouse over</span>',
    },
  ],
  searches: [
    { type: 'int', field: 'id', label: 'ID' },
    { type: 'text', field: 'name', label: 'Name' },
    { type: 'text', field: 'description', label: 'Description' },
  ],
  toolbar: {
    items: [
      {
        id: 'preview',
        text: 'Preview Changes',
        tooltip: 'Preview changes before saving',
        type: 'button',
        disabled: true,
        onClick: () => {
          w2popup.open({
            title: 'Preview Changes', with: 600, height: 550,
            body: `<pre>${JSON.stringify(todoGrid.getChanges(), null, 4)}</pre>`,
            actions: { Close: w2popup.close }
          })
        },
      },
      // { type: 'break' },
    ],
  },
  onChange: enablePreview,
  onRestore: disablePreview,
  onAdd: () => {
    w2popup.open({
      title: 'New Todo', width: 800, height: 600, showMax: true,
      body: '<div id="todoForm" class="w-full h-full"></div>',
    })
    new w2form({
      box: '#todoForm',
      url: '/admin/todo/data',
      httpHeaders: { 'X-CSRF-Token': getCookie('_csrf') },
      style: 'border: 0px; background-color: transparent;',
      fields: [
        { field: 'id', type: 'int', required: true, html: { label: 'ID' } },
        { field: 'name', type: 'text', required: true, html: { label: 'Name' } },
        { field: 'description', type: 'text', required: true, html: { label: 'Description' } },
        { field: 'qty', type: 'float', required: true, html: { label: 'Quantity' } },
        // {
        //   field: 'type', type: 'list',
        //   html: { label: 'Person Type' },
        //   options: { items: ['Employee', 'Contractor', 'Other'] }
        // }
      ],
      // record: {
      //   name: 'todo',
      //   description: 'important stuff',
      // },
      actions: {
        Close() { w2popup.close() },
        Save() { if (this.validate().length == 0) { this.save() } },
      }
    })
  },
  onEdit: () => {
    w2alert('edit form');
  },
})
