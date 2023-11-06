import { w2grid } from 'w2ui/w2ui-2.0.es6'

export const grid = new w2grid({
  name: 'grid',
  header: 'List of Names',
  reorderRows: false,
  show: {
    header: true,
    footer: true,
    toolbar: true,
    lineNumbers: true
  },
  columns: [
    { field: 'id', text: 'ID', size: '30px' },
    { field: 'name', text: 'First Name', size: '30%' },
    { field: 'description', text: 'Last Name', size: '30%' },
    { field: 'email', text: 'Email', size: '40%' },
    { field: 'sdate', text: 'Start Date', size: '120px' }
  ],
  searches: [
    { type: 'int', field: 'id', label: 'ID' },
    { type: 'text', field: 'name', label: 'Name' },
  ],
})
