import { w2grid, w2alert, w2popup, query } from 'w2ui/w2ui-2.0.es6'

export const todoGrid = new w2grid({
  name: 'todoGrid',
  url: '/admin/todo/data',
  header: 'List of Names',
  reorderRows: true,
  liveSearch: true,
  show: {
    header: false,
    footer: true,
    toolbar: true,
    toolbarAdd: true,
    toolbarDelete: true,
    toolbarSave: true,
    toolbarEdit: true,
    lineNumbers: false,
    selectColumn: true,
  },
  columns: [
    { field: 'id', text: 'ID', size: '30px', sortable: true, searchable: true, clipboardCopy: true }, // seachable for "all fields search"
    { field: 'name', text: 'First Name', size: '30%', editable: { type: 'text' } },
    { field: 'description', text: 'Last Name', size: '30%', editable: { type: 'text' } },
    { field: 'email', text: 'Email', size: '40%' },
    {
      field: 'sdate', text: 'Start Date', size: '120px',
      info: {
        render: (rec, ind, col_ind) => { return '<i>' + rec.name + '</i> <b>' + rec.description + '</b>' }
      }
    },
    {
      field: 'empty', text: 'Empty', size: '120px',
      info: {
        fields: ['id', 'name'],
        showOn: 'mouseenter',
        options: { position: 'top' },
        style: 'color: green'
      },
      render: () => '<span style="color: silver">Mouse over</span>'
    },
  ],
  searches: [
    { type: 'int', field: 'id', label: 'ID' },
    { type: 'text', field: 'name', label: 'Name' },
  ],
  toolbar: {
    items: [
      { id: 'add', type: 'button', text: 'Add Record', icon: 'w2ui-icon-plus' },
      { type: 'break' },
      { type: 'button', id: 'showChanges', text: 'Preview Changes' }
    ],
    onClick(event) {
      if (event.target == 'add') {
        let recid = grid.records.length + 1
        this.owner.add({ recid });
        this.owner.scrollIntoView(recid);
        this.owner.editField(recid, 1)
      }
      if (event.target == 'showChanges') {
        w2popup.open({
          title: 'Preview Changes',
          with: 600,
          height: 550,
          body: `<pre>${JSON.stringify(todoGrid.getChanges(), null, 4)}</pre>`,
          actions: { Ok: w2popup.close }
        })
      }
    }
  },
  contextMenu: [
    { id: 'view', text: 'View', icon: 'w2ui-icon-info' },
    { id: 'edit', text: 'Edit', icon: 'w2ui-icon-pencil' },
    { text: '--' },
    { id: 'delete', text: 'Delete', icon: 'w2ui-icon-cross' },
  ],
  onAdd: function(event) {
    w2alert('add');
  },
  onEdit: function(event) {
    w2alert('edit');
  },
  onDelete: function(event) {
    console.log('delete has default behavior');
  },
  onSave: function(event) {
    w2alert('save');
  },
  onReorderRow(event) {
    query('#log').html(`Record "${event.detail.recid}" moved before "${event.detail.moveBefore || event.detail.moveAfter}"`)
    console.log('move event', event)
  },
  onContextMenuClick(event) {
    console.log(event)
    query('#grid-log').html(event.detail.menuItem.text)
  }

})
