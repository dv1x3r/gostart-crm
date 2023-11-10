import { w2grid, w2alert, w2popup, query } from 'w2ui/w2ui-2.0.es6'

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
    { field: 'name', text: 'Name', size: '25%', sortable: true, editable: { type: 'text' } },
    { field: 'description', text: 'Description', size: '75%', sortable: true, editable: { type: 'text' } },
    {
      text: 'Summary', size: '120px',
      info: {
        fields: ['id', 'name', 'description'],
        showEmpty: true,
        showOn: 'mouseenter',
        options: { position: 'left' },
        // render: rec => `<b>${rec.name}</b>: ${rec.description}`,
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
        type: 'button', text: 'Preview', hint: 'Preview changes before saving',
        onClick: () => {
          w2popup.open({
            title: 'Preview Changes',
            with: 600,
            height: 550,
            body: `<pre>${JSON.stringify(todoGrid.getChanges(), null, 4)}</pre>`,
            actions: { Ok: w2popup.close }
          })
        },
      },
      // { type: 'break' },
    ],
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
