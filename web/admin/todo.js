import { w2grid, w2alert, w2popup, w2utils, query } from 'w2ui/w2ui-2.0.es6'


export const todoGrid = new w2grid({
  name: 'todoGrid',
  url: '/admin/todo/data',
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
    // seachable for "all fields search"
    { field: 'id', text: 'ID', size: '100px', sortable: true, clipboardCopy: true },
    { field: 'name', text: 'Name', size: '25%', sortable: true, render: row =>  w2utils.encodeTags(row.name) },
    { field: 'description', text: 'Description', size: '75%', sortable: true, render: row => w2utils.encodeTags(row.description) },
    { field: 'qty', text: 'Quantity', size: '100px', sortable: true, editable: { type: 'float' } },
    {
      text: 'Summary', size: '120px',
      info: {
        fields: ['id', 'name', 'description'],
        showEmpty: true,
        showOn: 'mouseenter',
        options: { position: 'left' },
        render: rec => `<b>${rec.name}</b>: ${rec.description}`,
      },
      render: () => '<span class="text-slate-400">Mouse over</span>'
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
            title: 'Preview Changes',
            with: 600,
            height: 550,
            body: `<pre>${JSON.stringify(todoGrid.getChanges(), null, 4)}</pre>`,
            actions: { Close: w2popup.close }
          })
        },
      },
      // { type: 'break' },
    ],
  },
  contextMenu: [
    { id: 'edit', text: 'Edit', icon: 'w2ui-icon-pencil' },
    { id: 'delete', text: 'Delete', icon: 'w2ui-icon-cross' },
  ],
  onChange: (event) => { event.owner.toolbar.enable('preview') },
  onRestore: (event) => {
    event.onComplete = () => {
      if (event.owner.getChanges().length == 0) {
        event.owner.toolbar.disable('preview')
      }
    }
  },
  onAdd: function(event) {
    // this.add({ id: 55 });
    // this.scrollIntoView(recid);
    // this.editField(recid, 1)
    // w2popup.open({
    // })
    // w2alert('add');
  },
  onEdit: function(event) {
    // w2alert('edit');
  },
  // onDelete: function(event) {
  //   console.log('delete has default behavior');
  // },
  // onSave: function(event) {
  //   w2alert('save');
  // },
})
