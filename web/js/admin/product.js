import { w2confirm, w2form, w2grid, w2layout, w2popup, w2utils } from 'w2ui/dist/w2ui.es6'
import * as utils from './utils'
import * as admin from '../admin'
import * as category from './category'

export function createProductGrid() {
  return new w2grid({
    name: 'productGrid',
    url: {
      get: '/product/category/:categoryID/catalog',
      save: '/product/save',
      remove: '/product/delete',
    },
    httpHeaders: { 'X-CSRF-Token': utils.getCsrfToken() },
    routeData: { categoryID: category.categorySidebar.selected ?? 0 },
    recid: 'id',
    multiSearch: true,
    reorderRows: false,
    show: {
      footer: true,
      toolbar: true,
      toolbarAdd: true,
      toolbarEdit: true,
      toolbarDelete: true,
      toolbarSave: true,
      toolbarSearch: true,
      toolbarReload: true,
      searchSave: false,
    },
    columns: [
      { field: 'id', text: 'ID', size: '60px', hidden: true },
      { field: 'thumbnail_url', text: x => x.field ? '' : 'Image', size: '36px', render: 'thumb-slim', hidden: true },
      { field: 'code', text: 'Code', size: '100px', render: 'safe', sortable: true, clipboardCopy: true },
      {
        field: 'name', text: 'Name', size: '385px', render: 'hover', sortable: true,
        info: {
          // showOn: 'mouseenter',
          render: row => {
            let s = '<table cellpadding="0" cellspacing="0"><tbody>'
            s += `<tr><td>Code</td><td>${w2utils.encodeTags(row.code)} (${w2utils.encodeTags(row.supplier.text)})</td></tr>`
            s += `<tr><td>Name</td><td>${w2utils.encodeTags(row.name.length > 64 ? row.name.slice(0, 64) + '...' : row.name)}</td></tr>`
            s += `<tr><td>Category</td><td>${w2utils.encodeTags(row.category.text)}</td></tr>`
            s += `<tr><td>Brand</td><td>${w2utils.encodeTags(row.brand.text)}</td></tr>`
            s += row.attributes?.filter(x => x.in_box && x.value.text)
              .map(x => `<tr><td>${w2utils.encodeTags(x.name)}</td><td>${w2utils.encodeTags(x.value.text)}</td></tr>`)
              .join('') ?? ''
            s += '</tbody></table>'
            return s
          }
        },
      },
      { field: 'category', text: 'Category', size: '387px', render: 'drophover', editable: utils.getSelectOptions('/category/dropdown?leafs=1', 'list'), sortable: false, hidden: true },
      { field: 'supplier', text: 'Supplier', size: '88px', render: 'dropdown', editable: utils.getSelectOptions('/supplier/dropdown', 'list'), sortable: true },
      { field: 'brand', text: 'Brand', size: '88px', render: 'dropdown', editable: utils.getSelectOptions('/brand/dropdown', 'list'), sortable: true },
      { field: 'status', text: 'Status', size: '88px', render: 'tag', editable: utils.getSelectOptions('/product/status/dropdown', 'list'), sortable: true },
      { field: 'quantity', text: 'Quantity', size: '88px', render: 'float:2', editable: { type: 'float', precision: 2 }, sortable: true },
      { field: 'price', text: 'Price', size: '88px', render: 'float:2', editable: { type: 'float', precision: 2 }, sortable: true },
      { field: 'is_published', text: 'Is Pub', size: '60px', editable: { type: 'checkbox' }, tooltip: 'Show this product', sortable: true },
      { field: 'created_at', text: 'Created at', size: '135px', render: 'datetime', sortable: true },
      { field: 'updated_at', text: 'Updated at', size: '135px', render: 'datetime', sortable: true },
    ],
    searches: [
      { field: 'code', label: 'Code', type: 'text', _all: true },
      { field: 'name', label: 'Name', type: 'text', _all: true },
      { field: 'category', label: 'Category', type: 'enum', options: utils.getSelectOptions('/category/dropdown') },
      { field: 'supplier', label: 'Supplier', type: 'enum', options: utils.getSelectOptions('/supplier/dropdown') },
      { field: 'brand', label: 'Brand', type: 'enum', options: utils.getSelectOptions('/brand/dropdown') },
      { field: 'status', label: 'Status', type: 'enum', options: utils.getSelectOptions('/product/status/dropdown') },
      { field: 'is_published', label: 'Is Pub', type: 'enum', options: utils.getSelectOptionsBool() },
      { field: 'quantity', label: 'Quantity', type: 'float' },
      { field: 'price', label: 'Price', type: 'float' },
    ],
    defaultOperator: {
      'text': 'contains',
    },
    sortData: [
      { field: 'updated_at', direction: 'desc' },
    ],
    onSearch: event => utils.gridSearchAllowedAll(event),
    onAdd: async event => await openProductPageTab(event),
    onEdit: async event => await openProductPageTab(event),
    onDblClick: async event => await utils.gridDblClickNonEditable(event, openProductPageTab),
    onSave: async event => await utils.gridShowSaved(event),
    onDelete: async event => {
      await utils.gridShowDeleted(event)
      await category.categorySidebar._reload() // update counters
    },
    onLoad: async event => await utils.gridMarkRows(event, x => !x.is_published || x.quantity == 0 ? 'darkgrey' : null),
  })
}

async function openProductPageTab(event) {
  if (!event.detail.recid) {
    admin.tabManager.OpenTab(`product-${Date.now()}-tab`, '[DRAFT] Product', true, createProductPageLayout)
    return
  }

  // if tab with the same product is already opened, then just open the existing one (compare by recid in the details form)
  const fn = x => x.component.name.startsWith('productPageLayout_') && x.component._detailsForm.recid == event.detail.recid
  const tabs = admin.tabManager.GetTabs().tabs.filter(fn)
  if (tabs.length) {
    admin.tabManager.OpenTab(tabs[0].id)
  } else {
    const row = event.owner.get(event.detail.recid)
    admin.tabManager.OpenTab(`product-${Date.now()}-tab`, `${row.code} - ${row.name}`, true, createProductPageLayout, row.id)
  }
}

function createProductPageLayout(id) {
  const tabID = Date.now()
  const attributeGrid = createProductAttributeGrid(tabID, id)
  const mediaGrid = createProductMediaGrid(tabID, id)
  const detailsForm = createProductDetailsForm(tabID, id, attributeGrid, mediaGrid)

  return new w2layout({
    name: `productPageLayout_${tabID}`,
    panels: [{ type: 'main', html: `<div id="product-details-form-${tabID}"></div>` }],
    _detailsForm: detailsForm, // needed for openProductPageTab to check if tab is already opened
    onRender: async event => {
      await event.complete
      detailsForm.render(`#product-details-form-${tabID}`)
      attributeGrid.render(`#product-attribute-grid-${tabID}`)
      mediaGrid.render(`#product-media-grid-${tabID}`)
    },
    onDestroy: () => {
      detailsForm.destroy()
      attributeGrid.destroy()
      mediaGrid.destroy()
    }
  })
}

function createProductDetailsForm(tabID, id, attributeGrid, mediaGrid) {
  const selectedCategory = category.categorySidebar.get(category.categorySidebar.selected)
  const detailsForm = new w2form({
    name: `productDetailsForm_${tabID}`,
    recid: id,
    url: '/product',
    httpHeaders: { 'X-CSRF-Token': utils.getCsrfToken() },
    style: 'height:auto;',
    fields: [
      { field: 'code', html: { label: 'Code', span: 3, column: 0, group: 'Details', attr: 'style="width: 100%; min-width:203px;"' }, type: 'text', required: true },
      { field: 'name_lv', html: { label: 'Name LV', span: 3, column: 0, attr: 'style="width: 100%; min-width:203px;"' }, type: 'text', required: true },
      { field: 'name_ru', html: { label: 'Name RU', span: 3, column: 0, attr: 'style="width: 100%; min-width:203px;"' }, type: 'text' },
      { field: 'category', html: { label: 'Category', span: 3, column: 1, group: 'Catalog', attr: 'style="width: 100%; min-width:203px;"' }, type: 'list', options: utils.getSelectOptions('/category/dropdown?leafs=1'), required: true },
      { field: 'supplier', html: { label: 'Supplier', span: 3, column: 1, attr: 'style="width: 100%; min-width:203px;"' }, type: 'list', options: utils.getSelectOptions('/supplier/dropdown'), required: true },
      { field: 'brand', html: { label: 'Brand', span: 3, column: 1, attr: 'style="width: 100%; min-width:203px;"' }, type: 'list', options: utils.getSelectOptions('/brand/dropdown'), required: true },
      { field: 'status', html: { label: 'Status', span: 4, column: 2, group: 'Status' }, type: 'list', options: utils.getSelectOptions('/product/status/dropdown') },
      { field: 'status_expire_at', html: { label: 'Expire at', span: 4, column: 2, attr: 'autocomplete="off"', }, type: 'date' },
      { field: 'is_published', html: { label: 'Is Published', span: 4, column: 2, attr: 'style="height: 21px"' }, type: 'checkbox' },
      { field: 'quantity', html: { label: 'Quantity', span: 4, column: 2, group: 'Quantity and Prices' }, type: 'float', options: { precision: 2 }, required: true },
      { field: 'input_no_tax', html: { label: 'Input No Tax', span: 4, column: 2 }, type: 'float', options: { precision: 2 } },
      { field: 'fixed_actual', html: { label: 'Fixed Actual', span: 4, column: 2 }, type: 'float', options: { precision: 2 } },
      { field: 'fixed_old', html: { label: 'Fixed Old', span: 4, column: 2 }, type: 'float', options: { precision: 2 } },
      { field: 'web_actual', html: { label: 'Web Actual', span: 4, column: 2, attr: 'readonly' }, type: 'float', options: { precision: 2 } },
      { field: 'web_no_tax', html: { label: 'Web No Tax', span: 4, column: 2, attr: 'readonly' }, type: 'float', options: { precision: 2 } },
      { field: 'created_at', html: { label: 'Created at', span: 4, column: 2, attr: 'readonly', group: 'System' }, type: 'datetime' },
      { field: 'updated_at', html: { label: 'Updated at', span: 4, column: 2, attr: 'readonly' }, type: 'datetime' },
      { html: { html: `<div id="product-attribute-grid-${tabID}" style="height: 780px;"></div>`, span: -1, column: 0, group: 'Attributes' }, type: 'html' },
      { html: { html: `<div id="product-media-grid-${tabID}" style="height: 369px;"></div>`, span: -1, column: 1, group: 'Media' }, type: 'html' },
      { field: 'description_lv', html: { label: '', span: 0, column: 1, group: 'Description LV', attr: 'style="width: 100%; height: 338px; resize: none;"' }, type: 'textarea' },
      { field: 'description_ru', html: { label: '', span: 0, column: 2, group: 'Description RU', attr: 'style="width: 100%; height: 338px; resize: none;"' }, type: 'textarea' },
    ],
    record: {
      category: selectedCategory.nodes.length > 0 ? null : {
        id: selectedCategory.id,
        text: selectedCategory._hierarchy_text,
      },
      is_published: true,
    },
    toolbar: {
      items: [
        { id: 'save-close', type: 'button', text: 'Save and close', icon: 'fa-regular fa-floppy-disk', onClick: async () => await detailsForm._saveAndClose() },
        { id: 'save', type: 'button', text: 'Save', icon: 'fa-regular fa-floppy-disk', onClick: async () => await detailsForm._save() },
        { id: 'reset', type: 'button', text: 'Refresh', icon: 'fa-solid fa-rotate-left', onClick: async () => await detailsForm._reset() },
        { id: 'spacer', type: 'spacer' },
        { id: 'delete', type: 'button', text: 'Delete', icon: 'fa-solid fa-trash', onClick: async () => await detailsForm._delete() },
      ],
    },
    onLoad: async event => {
      await event.complete
      // save the attribute_group_id to compare it on next save
      event.owner._attribute_group_id = event.owner.record.category.attribute_group_id
    },
    _saveAndClose: async () => {
      await detailsForm._save()
      const fn = x => x.component.name.startsWith('productPageLayout_') && x.component._detailsForm.recid == detailsForm.recid
      admin.tabManager.CloseTab(fn)
    },
    _save: async () => {
      const _saveDetailsForm = async () => {
        const res = await detailsForm.save()
        detailsForm.recid = res.recid
        attributeGrid.routeData.productID = res.recid
        mediaGrid.routeData.productID = res.recid
        await detailsForm.reload()

        // set tab name based on code and name
        const fn = x => x.component.name.startsWith('productPageLayout_') && x.component._detailsForm.recid == detailsForm.recid
        admin.tabManager.RenameTab(`${detailsForm.record.code} - ${detailsForm.record.name}`, fn)

        await category.categorySidebar._reload() // update sidebar counters
        utils.fetchShowSuccess('Product has been successfully saved!')
      }

      // double check that user wants to change the category with different attribute group
      if (detailsForm._attribute_group_id != null && detailsForm._attribute_group_id != detailsForm.record.category.attribute_group_id) {
        w2confirm({
          msg: 'The selected product category is associated with another attribute group. The current product attributes will be deleted.<br>Continue?',
          btn_yes: { text: 'Save', class: 'w2ui-btn-blue' },
          btn_no: { text: 'Cancel' },
        }).yes(async () => {
          await _saveDetailsForm()
          await attributeGrid.reload()
          await mediaGrid.reload()
        })
      } else {
        // save attribute grid only if record was already in the db, with attribute group and there are changes
        if (detailsForm.recid != null && detailsForm._attribute_group_id != null && attributeGrid.getChanges().length) {
          await _saveDetailsForm()
          await attributeGrid.save()
          await mediaGrid.reload()
        } else {
          await _saveDetailsForm()
          await attributeGrid.reload()
          await mediaGrid.reload()
        }
      }
    },
    _reset: async () => {
      w2confirm({
        msg: 'Do you want to discard unapplied changes?',
        btn_yes: { text: 'Reset', class: 'w2ui-btn-blue' },
        btn_no: { text: 'Cancel' },
      })
        .yes(async () => {
          if (detailsForm.recid == null) {
            detailsForm.clear()
            await attributeGrid.reload()
            await mediaGrid.reload()
            utils.fetchShowSuccess('Form has been cleared!')
          } else {
            await detailsForm.reload()
            await attributeGrid.reload()
            await mediaGrid.reload()
            utils.fetchShowSuccess('Form has been restored to the initial state!')
          }
        })
    },
    _delete: async () => {
      w2confirm({
        msg: 'Are you sure you want to delete this product?',
        btn_yes: { text: 'Delete', class: 'w2ui-btn-red' },
        btn_no: { text: 'Cancel' },
      }).yes(async () => {
        if (detailsForm.recid) {
          const res = await fetch(`/product/${detailsForm.recid}`, { method: 'DELETE', headers: { 'X-CSRF-Token': utils.getCsrfToken() } })
          if (res.status != 200) {
            await utils.fetchShowError('Failed to delete the product', res)
            return
          }

          const safeName = w2utils.encodeTags(detailsForm.record.code + ' - ' + detailsForm.record.name)
          utils.fetchShowSuccess(`${safeName.length > 32 ? safeName.slice(0, 32) + '...' : safeName} has been successfully deleted!`)
          await category.categorySidebar._reload() // update counters
        }

        const fn = x => x.component.name.startsWith('productPageLayout_') && x.component._detailsForm.recid == detailsForm.recid
        admin.tabManager.CloseTab(fn)
      })
    },
  })

  return detailsForm
}

function createProductAttributeGrid(tabID, id) {
  return new w2grid({
    name: `productAttributeGrid_${tabID}`,
    url: '/product/:productID/attributes',
    httpHeaders: { 'X-CSRF-Token': utils.getCsrfToken() },
    routeData: { productID: id ?? 0 },
    recid: 'id',
    recordHeight: 30,
    columns: [
      { field: 'id', text: 'ID', size: '60px', hidden: true },
      { field: 'name', text: 'Attribute', size: '100%', min: '100px', render: 'safe' },
      { field: 'value', text: 'Value', size: '100%', min: '100px', render: 'dropdown', editable: utils.getSelectOptions(null, 'list') },
    ],
    onSave: async event => {
      if (event.detail.data?.status == 'success') {
        await event.owner.reload()
      }
    },
    onEditField: event => {
      const attribute_set_id = event.owner.get(event.detail.recid).id
      event.owner.getColumn('value').editable.url = `/attribute/set/${attribute_set_id}/value/dropdown`
    },
    onDelete: event => event.preventDefault(),
    onError: _ => w2utils.notify(`Failed to save new attributes, please revalidate`, { timeout: 4000, error: true }),
  })
}

function createProductMediaGrid(tabID, id) {
  return new w2grid({
    name: `productMediaGrid_${tabID}`,
    url: {
      get: '/product/:productID/media',
      remove: '/product/media/delete',
    },
    httpHeaders: { 'X-CSRF-Token': utils.getCsrfToken() },
    routeData: { productID: id ?? 0 },
    recid: 'id',
    recordHeight: 80,
    reorderRows: true,
    show: {
      footer: false,
      toolbar: true,
      toolbarAdd: true,
      toolbarEdit: false,
      toolbarDelete: true,
      toolbarSave: false,
      toolbarSearch: false,
      toolbarReload: false,
    },
    columns: [
      { field: 'id', text: 'ID', size: '60px', hidden: true },
      { field: 'name', text: 'File name', size: '100%', min: '100px', render: 'safe', clipboardCopy: row => utils.getURI(`/media/${row.file_url}`) },
      { field: 'thumbnail_url', text: 'Thumbnail', size: '80px', render: 'thumb' },
    ],
    onAdd: async event => {
      const inputFile = document.createElement('input')
      inputFile.type = 'file';
      inputFile.onchange = async ev => {
        const formData = new FormData()
        formData.append('file', ev.target.files[0]);
        formData.append('productID', event.owner.routeData.productID);
        const res = await fetch(`/product/media/upload`, {
          method: 'POST',
          headers: {
            'X-CSRF-Token': utils.getCsrfToken(),
          },
          body: formData,
        })
        if (res.status != 200) {
          await utils.fetchShowError('Failed to upload the file', res)
          return
        }
        utils.fetchShowSuccess('File has been successfully uploaded!')
        await event.owner.reload()
      }
      inputFile.click()
    },
    onDelete: async event => await utils.gridShowDeleted(event),
    onReorderRow: async event => await utils.gridPostReorderRow(event, `/product/media/reorder`),
    onRequest: async event => {
      if (event.owner.routeData.productID == 0) {
        event.owner.toolbar.disable('w2ui-add')
      } else {
        event.owner.toolbar.enable('w2ui-add')
      }
    },
  })
}

export function openProductStatusPopup() {
  const productStatusGrid = new w2grid({
    name: 'productStatusGrid',
    url: {
      get: '/product/status',
      save: '/product/status/save',
      remove: '/product/status/delete',
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
      toolbarReload: false,
      searchSave: false,
    },
    columns: [
      { field: 'id', text: 'ID', size: '60px', hidden: true },
      { field: 'name', text: 'Name', size: '155px', render: 'safe', editable: { type: 'text' }, searchable: true },
      { field: 'color', text: 'Color', size: '85px', render: 'color', editable: { type: 'color' } },
      { field: 'related_products', text: '# Products', size: '90px', render: 'int' },
    ],
    defaultOperator: {
      'text': 'contains',
    },
    onAdd: async event => await utils.gridNewRowAdd(event),
    onChange: async event => await utils.gridNewRowChange(event),
    onSave: async event => await utils.gridShowSaved(event),
    onDelete: async event => await utils.gridShowDeleted(event),
    onReorderRow: async event => await utils.gridPostReorderRow(event, '/product/status/reorder'),
  })

  w2popup.open({
    title: 'Product Status',
    width: 900,
    height: 600,
    showMax: true,
    resizable: true,
    body: '<div id="product-status-grid" class="w-full h-full"></div>'
  })
    .then(() => productStatusGrid.render('#product-status-grid'))
    .close(() => productStatusGrid.destroy())
}

