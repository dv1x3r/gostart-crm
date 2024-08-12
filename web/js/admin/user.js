import { w2form, w2grid, w2layout, w2popup } from 'w2ui/dist/w2ui.es6'
import * as utils from './utils'

export function createUserLayout() {
  const userGrid = new w2grid({
    name: 'userGrid',
    url: {
      get: '/user',
      remove: '/user/delete',
    },
    httpHeaders: { 'X-CSRF-Token': utils.getCsrfToken() },
    recid: 'id',
    limit: 1000,
    multiSearch: true,
    reorderRows: false,
    show: {
      footer: true,
      toolbar: true,
      toolbarAdd: false,
      toolbarEdit: false,
      toolbarDelete: true,
      toolbarSave: false,
      toolbarSearch: true,
      toolbarReload: true,
      searchSave: false,
    },
    columns: [
      { field: 'id', text: 'ID', size: '60px', hidden: true },
      { field: 'full_name', text: 'Full name', size: '170px', render: 'safe', sortable: true, frozen: true },
      { field: 'email', text: 'E-mail', size: '170px', render: 'safe', sortable: true },
      { field: 'created_at', text: 'Created at', size: '90px', render: 'date', tooltip: 'Date of sign-up', sortable: true },
      { field: 'visited_at', text: 'Visited at', size: '90px', render: 'date', tooltip: 'Most recent visit to a website', sortable: true },
      { field: 'is_active', text: 'Is Act', size: '60px', render: 'toggle', tooltip: 'User is active', sortable: true },
      { field: 'is_admin', text: 'Is Adm', size: '60px', render: 'toggle', tooltip: 'User has admin rights', sortable: true },
      { field: 'is_readonly', text: 'Is R/O', size: '60px', render: 'toggle', tooltip: 'User has read only rights', sortable: true },
    ],
    searches: [
      { field: 'full_name', label: 'Full name', type: 'text', _all: true },
      { field: 'email', label: 'Email', type: 'text', _all: true },
      { field: 'is_active', label: 'Is Active', type: 'enum', options: utils.getSelectOptionsBool() },
      { field: 'is_admin', label: 'Is Admin', type: 'enum', options: utils.getSelectOptionsBool() },
      { field: 'is_readonly', label: 'Is Read Only', type: 'enum', options: utils.getSelectOptionsBool() },
    ],
    defaultOperator: {
      'text': 'contains',
    },
    sortData: [
      { field: 'created_at', direction: 'desc' },
    ],
    toolbar: {
      items: [
        { id: 'visits', type: 'button', text: 'Visits', icon: 'fa fa-chart-simple', onClick: () => openVisitPopup() },
      ],
    },
    onDelete: async event => {
      await utils.gridShowDeleted(event)
      userForm.recid = null
      userForm.clear()
    },
    onSelect: async event => await utils.gridSetFormRecord(event, userForm),
    onSearch: event => utils.gridSearchAllowedAll(event),
    onLoad: async event => await utils.gridMarkRows(event, x => !x.is_active ? 'darkgrey' : null),
  })

  const userForm = new w2form({
    name: `userForm`,
    url: '/user',
    httpHeaders: { 'X-CSRF-Token': utils.getCsrfToken() },
    style: 'min-height: calc(100vh - 68px); max-height: calc(100vh - 68px);',
    fields: [
      { field: 'email', html: { label: 'E-mail', span: 4, group: 'Contact Details', groupCollapsible: true, attr: 'style="width: 100%;' }, type: 'text', required: true },
      { field: 'first_name', html: { label: 'First name', span: 4, attr: 'style="width: 100%;' }, type: 'text' },
      { field: 'last_name', html: { label: 'Last name', span: 4, attr: 'style="width: 100%;' }, type: 'text' },
      { field: 'is_active', html: { label: 'Is Active', span: 4, group: 'Access Control', groupCollapsible: true, attr: 'style="height: 21px"' }, type: 'checkbox' },
      { field: 'is_admin', html: { label: 'Is Admin', span: 4, attr: 'style="height: 21px"' }, type: 'checkbox' },
      { field: 'is_readonly', html: { label: 'Is Read Only', span: 4, attr: 'style="height: 21px"' }, type: 'checkbox' },
    ],
    actions: {
      async Save() {
        if (userForm.recid == null || userForm.recid == 0) {
          return
        }

        try {
          await userForm.save()
        } catch { return }

        await userGrid.reload()
        utils.fetchShowSuccess('User has been successfully saved!')
      },
    },
  })

  return new w2layout({
    name: 'userLayout',
    panels: [
      { type: 'left', size: '70%', style: 'margin-right: 5px;', html: userGrid, resizable: true },
      { type: 'main', size: '30%', style: 'margin-left: 5px;', html: userForm },
    ],
    onDestroy: () => {
      userGrid.destroy()
      userForm.destroy()
    }
  })
}

function openVisitPopup() {
  const visitGrid = new w2grid({
    name: 'visitGrid',
    url: '/user/visit',
    httpHeaders: { 'X-CSRF-Token': utils.getCsrfToken() },
    recid: 'id',
    show: {
      footer: true,
      toolbar: false,
    },
    columns: [
      { field: 'id', text: 'ID', size: '60px', hidden: true },
      { field: 'date', text: 'Date', size: '90px', render: 'date' },
      { field: 'email', text: 'E-mail', size: '200px', render: 'safe' },
      { field: 'full_name', text: 'Full name', size: '200px', render: 'safe' },
    ],
    onDelete: event => event.preventDefault(),
  })

  w2popup.open({
    title: 'Visits',
    width: 900,
    height: 600,
    showMax: true,
    resizable: true,
    body: '<div id="visit-grid" class="w-full h-full"></div>'
  })
    .then(() => visitGrid.render('#visit-grid'))
    .close(() => visitGrid.destroy())
}

