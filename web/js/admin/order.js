import { w2form, w2grid, w2layout, w2popup, w2utils, query } from 'w2ui/dist/w2ui.es6'
import * as utils from './utils'
import * as admin from '../admin'

export async function updateOrderCounter() {
  const res = await fetch('/order/counter', { method: 'GET', headers: { 'X-CSRF-Token': utils.getCsrfToken() } })
  if (res.status != 200) {
    await utils.fetchShowError('Failed to fetch order counter', res)
    return
  }

  try {
    const data = await res.json()
    if (data.data != null) {
      const count = data.data
      admin.mainLayout.get('top').toolbar.get('orders').count = count == 0 ? null : count
      admin.mainLayout.get('top').toolbar.refresh()
    }
  } catch {
    w2utils.notify(`Failed to read order counter`, { timeout: 4000, error: true })
  }
}

export function createOrderLayout() {
  const orderGrid = new w2grid({
    name: 'orderGrid',
    url: {
      get: '/order',
      remove: '/order/delete',
    },
    httpHeaders: { 'X-CSRF-Token': utils.getCsrfToken() },
    recid: 'id',
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
      { field: 'id', text: 'ID', size: '60px', sortable: true },
      { field: 'created_at', text: 'Created at', size: '90px', render: 'date', sortable: true },
      { field: 'updated_at', text: 'Updated at', size: '90px', render: 'date', sortable: true, hidden: true },
      { field: 'full_name', text: 'Full name', size: '170px', render: 'safe', sortable: true },
      { field: 'email', text: 'E-mail', size: '170px', render: 'safe', sortable: true },
      { field: 'status', text: 'Status', size: '100px', render: 'dropdown', sortable: true },
      { field: 'total', text: 'Total', size: '80px', render: 'float:2', sortable: true },
    ],
    searches: [
      { field: 'id', label: 'Order ID', type: 'text', _all: true },
      { field: 'full_name', label: 'Full name', type: 'text', _all: true },
      { field: 'email', label: 'Email', type: 'text', _all: true },
      { field: 'status', label: 'Status', type: 'enum', options: utils.getSelectOptions('/order/status/dropdown') },
    ],
    defaultOperator: {
      'text': 'contains',
    },
    sortData: [
      { field: 'id', direction: 'desc' },
    ],
    onDelete: async event => {
      await utils.gridShowDeleted(event)
      orderForm.recid = null
      orderForm.clear()
      orderLineGrid.clear()
    },
    onSelect: async event => {
      await utils.gridSetFormRecord(event, orderForm)
      const selectionOrderID = event.owner.getSelection()
      if (selectionOrderID.length == 1 && selectionOrderID[0] != 0) {
        orderLineGrid.routeData.orderID = selectionOrderID[0]
        await orderLineGrid.reload()
      } else {
        orderLineGrid.clear()
      }
    },
    onSearch: event => utils.gridSearchAllowedAll(event),
    onLoad: async event => {
      await utils.gridMarkRows(event, x => `#${x.status.color}`)
      await updateOrderCounter() // update toolbar counters
    },
  })

  const orderLineGrid = new w2grid({
    name: 'orderLineGrid',
    url: {
      get: '/order/:orderID/line',
    },
    httpHeaders: { 'X-CSRF-Token': utils.getCsrfToken() },
    routeData: { orderID: 0 },
    recid: 'id',
    limit: 1000,
    recordHeight: 30,
    show: {
      footer: false,
      toolbar: false,
      expandColumn: true,
    },
    columns: [
      { field: 'id', text: 'ID', size: '60px', hidden: true },
      { field: 'code', text: 'Code', size: '115px', render: 'safe', clipboardCopy: true },
      { field: 'product', text: 'Product', size: '100%', render: 'safe' },
      { field: 'quantity', text: 'Qty', size: '50px', render: 'float:2' },
      { field: 'price', text: 'Price', size: '80px', render: 'float:2' },
      { field: 'total', text: 'Total', size: '80px', render: 'float:2' },
    ],
    onDelete: event => event.preventDefault(),
    onExpand: event => {
      const row = event.owner.get(event.detail.recid)
      const safeContent = w2utils.encodeTags(`${row.code} | ${row.product}`)
      query('#' + event.detail.box_id).html(`
      <div style="padding: 5px;">
        <textarea style="width: 100%; height: 100%; resize: none; field-sizing: content;" readonly>${safeContent}</textarea>
      </div>
    `)
    },
  })

  const orderForm = new w2form({
    name: `orderForm`,
    url: '/order',
    httpHeaders: { 'X-CSRF-Token': utils.getCsrfToken() },
    style: 'min-height: calc(100vh - 68px); max-height: calc(100vh - 68px);',
    focus: 'email',
    fields: [
      { html: { html: `<div id="order-line-grid" style="height: 200px;"></div>`, span: -1, column: 'before', group: 'Order Lines' }, type: 'html' },
      { field: 'email', html: { label: 'E-mail', span: 4, group: 'Contact Details', attr: 'style="width: 100%;' }, type: 'text', required: true },
      { field: 'first_name', html: { label: 'First name', span: 4, attr: 'style="width: 100%;' }, type: 'text', required: true },
      { field: 'last_name', html: { label: 'Last name', span: 4, attr: 'style="width: 100%;' }, type: 'text', required: true },
      { field: 'phone_number', html: { label: 'Phone nr.', span: 4, attr: 'style="width: 100%;' }, type: 'text', required: true },
      { field: 'status', html: { label: 'Status', span: 4, column: 1, group: 'Status and Payment' }, type: 'list', options: utils.getSelectOptions('/order/status/dropdown'), required: true },
      { field: 'payment', html: { label: 'Payment', span: 4, column: 1 }, type: 'list', options: utils.getSelectOptions('/order/payment/dropdown'), required: true },
      { field: 'created_at', html: { label: 'Created at', span: 4, column: 1, attr: 'readonly' }, type: 'datetime' },
      { field: 'updated_at', html: { label: 'Updated at', span: 4, column: 1, attr: 'readonly' }, type: 'datetime' },
      { field: 'delivery_address', html: { label: '', span: 0, column: 1, group: 'Delivery Address', attr: 'style="width: 100%; height: 56px; resize: none; color: black;" readonly' }, type: 'textarea' },
      { field: 'comment', html: { label: '', span: 0, column: 1, group: 'Order Comment', attr: 'style="width: 100%; height: 57px; resize: none; color: black;" readonly' }, type: 'textarea' },
    ],
    actions: {
      async Save() {
        if (orderForm.recid == null || orderForm.recid == 0) {
          return
        }

        try {
          await orderForm.save()
        } catch { return }

        await orderGrid.reload()
        utils.fetchShowSuccess('Order has been successfully saved!')
      },
    },
    onRender: async event => {
      await event.complete
      orderLineGrid.render(`#order-line-grid`)
    },
    onDestroy: () => {
      orderLineGrid.destroy()
    }
  })

  return new w2layout({
    name: 'orderLayout',
    panels: [
      { type: 'left', size: '50%', style: 'margin-right: 5px;', html: orderGrid, resizable: true },
      { type: 'main', size: '50%', style: 'margin-left: 5px;', html: orderForm },
    ],
    onDestroy: () => {
      orderGrid.destroy()
      orderForm.destroy()
    }
  })
}

export function openOrderStatusPopup() {
  const orderStatusGrid = new w2grid({
    name: 'orderStatusGrid',
    url: {
      get: '/order/status',
      save: '/order/status/save',
      remove: '/order/status/delete',
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
      { field: 'name', text: 'Name', size: '155px', render: 'safe', editable: { type: 'text' }, searchable: true },
      { field: 'color', text: 'Color', size: '85px', render: 'color', editable: { type: 'color' } },
      { field: 'related_orders', text: '# Orders', size: '90px', render: 'int' },
    ],
    defaultOperator: {
      'text': 'contains',
    },
    onAdd: async event => await utils.gridNewRowAdd(event),
    onChange: async event => await utils.gridNewRowChange(event),
    onSave: async event => await utils.gridShowSaved(event),
    onDelete: async event => await utils.gridShowDeleted(event),
    onReorderRow: async event => await utils.gridPostReorderRow(event, '/order/status/reorder'),
  })

  w2popup.open({
    title: 'Order Status',
    width: 900,
    height: 600,
    showMax: true,
    resizable: true,
    body: '<div id="order-status-grid" class="w-full h-full"></div>'
  })
    .then(() => orderStatusGrid.render('#order-status-grid'))
    .close(() => orderStatusGrid.destroy())
}

export function openPaymentMethodPopup() {
  const paymentMethodGrid = new w2grid({
    name: 'paymentMethodGrid',
    url: {
      get: '/order/payment',
      save: '/order/payment/save',
      remove: '/order/payment/delete',
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
      { field: 'name', text: 'Name', size: '155px', render: 'safe', editable: { type: 'text' }, searchable: true },
      { field: 'related_orders', text: '# Orders', size: '90px', render: 'int' },
    ],
    defaultOperator: {
      'text': 'contains',
    },
    onAdd: async event => await utils.gridNewRowAdd(event),
    onChange: async event => await utils.gridNewRowChange(event),
    onSave: async event => await utils.gridShowSaved(event),
    onDelete: async event => await utils.gridShowDeleted(event),
    onReorderRow: async event => await utils.gridPostReorderRow(event, '/order/payment/reorder'),
  })

  w2popup.open({
    title: 'Payment Method',
    width: 900,
    height: 600,
    showMax: true,
    resizable: true,
    body: '<div id="payment-method-grid" class="w-full h-full"></div>'
  })
    .then(() => paymentMethodGrid.render('#payment-method-grid'))
    .close(() => paymentMethodGrid.destroy())
}

