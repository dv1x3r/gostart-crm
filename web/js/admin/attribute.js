import { w2confirm, w2grid, w2layout, w2sidebar, w2utils } from 'w2ui/dist/w2ui.es6'
import { categorySidebar } from './category'
import * as admin from '../admin'
import * as utils from './utils'

export function createAttributeLayout() {
  const attributeGroupSidebar = new w2sidebar({
    name: 'attributeGroupSidebar',
    editable: true,
    onClick: async event => {
      await event.complete
      if (event.owner.selected != attributeSetGrid.routeData.groupID) {
        await attributeSetGrid._reload()
        await attributeValueGrid._reload()
      }
    },
    onRename: async event => {
      const res = await fetch(`/attribute/group/save`, {
        method: 'POST',
        headers: {
          'X-CSRF-Token': utils.getCsrfToken(),
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          'id': parseInt(event.target),
          'name': event.detail.text_new,
        }),
      })

      if (res.status != 200) {
        await utils.fetchShowError('Failed to update the group', res)
        return
      }

      utils.fetchShowSuccess('Group has been successfully updated!')
      await attributeGroupSidebar._reload()
      attributeSetGrid.clear()
      attributeSetGrid.toolbar.disable('w2ui-add')
      attributeValueGrid.clear()
      attributeValueGrid.toolbar.disable('w2ui-add')
    },
    _reload: async () => {
      const res = await fetch('/attribute/group', { method: 'GET', headers: { 'X-CSRF-Token': utils.getCsrfToken() } })
      if (res.status != 200) {
        await utils.fetchShowError('Failed to fetch groups', res)
        return
      }

      try {
        const data = await res.json()
        attributeGroupSidebar.remove(...attributeGroupSidebar.nodes.map(x => x.id))
        if (data.records != null) {
          data.records.forEach(x => attributeGroupSidebar.add({ id: x.id, text: w2utils.encodeTags(x.name) }))
        }
      } catch {
        w2utils.notify(`Failed to read groups`, { timeout: 4000, error: true })
      }
    },
    _delete: async (id) => {
      const res = await fetch(`/attribute/group/${id}/delete`, {
        method: 'POST',
        headers: {
          'X-CSRF-Token': utils.getCsrfToken(),
          'Content-Type': 'application/json'
        },
      })

      if (res.status != 200) {
        await utils.fetchShowError('Failed to delete the group', res)
        return
      }

      utils.fetchShowSuccess('Group has been successfully deleted!')
      await attributeGroupSidebar._reload()
      attributeSetGrid.clear()
      attributeSetGrid.toolbar.disable('w2ui-add')
      attributeValueGrid.clear()
      attributeValueGrid.toolbar.disable('w2ui-add')
    }
  })

  const attributeSetGrid = new w2grid({
    name: 'attributeSetGrid',
    url: {
      get: '/attribute/group/:groupID/set',
      save: '/attribute/group/:groupID/set/save',
      remove: '/attribute/set/delete',
    },
    httpHeaders: { 'X-CSRF-Token': utils.getCsrfToken() },
    routeData: { groupID: 0 },
    recid: 'id',
    limit: 1000,
    recordHeight: 30,
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
      { field: 'name', text: 'Attribute', size: '100%', render: 'safe', editable: { type: 'text' }, searchable: true },
      { field: 'in_box', text: 'Box', size: '50px', editable: { type: 'checkbox' }, tooltip: 'Show this attribute in the product grid box' },
      { field: 'in_filter', text: 'Filter', size: '50px', editable: { type: 'checkbox' }, tooltip: 'Allow filtering by this attribute' },
    ],
    defaultOperator: {
      'text': 'contains',
    },
    onAdd: async event => await utils.gridNewRowAdd(event),
    onChange: async event => await utils.gridNewRowChange(event),
    onSave: async event => await utils.gridShowSaved(event),
    onDelete: async event => await utils.gridShowDeleted(event),
    onReorderRow: async event => await utils.gridPostReorderRow(event, '/attribute/set/reorder'),
    onSelect: async event => {
      await event.complete
      await attributeValueGrid._reload()
    },
    onLoad: async event => {
      await event.complete
      await attributeValueGrid._reload()
    },
    _reload: async () => {
      const groupID = attributeGroupSidebar.selected
      if (groupID != undefined && groupID != 0) {
        attributeSetGrid.routeData.groupID = groupID
        await attributeSetGrid.reload()
        attributeSetGrid.toolbar.enable('w2ui-add')
      }
    },
  })

  const attributeValueGrid = new w2grid({
    name: 'attributeValueGrid',
    url: {
      get: '/attribute/set/:setID/value',
      save: '/attribute/set/:setID/value/save',
      remove: '/attribute/value/delete',
    },
    httpHeaders: { 'X-CSRF-Token': utils.getCsrfToken() },
    routeData: { setID: 0 },
    recid: 'id',
    limit: 1000,
    recordHeight: 30,
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
      { field: 'name', text: 'Value', size: '100%', render: 'safe', editable: { type: 'text' }, searchable: true },
      { field: 'published_products', text: '# Pub', size: '60px', render: 'int', tooltip: 'Number of published products' },
      { field: 'related_products', text: '# All', size: '60px', render: 'int', tooltip: 'Total number of products' },
    ],
    defaultOperator: {
      'text': 'contains',
    },
    contextMenu: [
      {
        id: 'products', type: 'button', text: 'Filter Products', icon: 'fa fa-filter', onClick: async event => {
          const selectionID = event.owner.getSelection()
          if (selectionID.length != 1) {
            return
          }

          const tabs = admin.tabManager.GetTabs()
          tabs.click('products-tab')
          categorySidebar.unselect(categorySidebar.selected)
          categorySidebar.select(0)

          const productGrid = tabs.get('products-tab').component
          productGrid.routeData.categoryID = 0
          productGrid.search([{ field: 'attribute', value: selectionID[0] }])
        }
      },
    ],
    onContextMenuClick: event => utils.w2contextMenuOnClick(event),
    onAdd: async event => await utils.gridNewRowAdd(event),
    onChange: async event => await utils.gridNewRowChange(event),
    onSave: async event => await utils.gridShowSaved(event),
    onDelete: async event => await utils.gridShowDeleted(event),
    onReorderRow: async event => await utils.gridPostReorderRow(event, '/attribute/value/reorder'),
    _reload: async () => {
      const selectionSetID = attributeSetGrid.getSelection()
      if (selectionSetID.length == 1 && selectionSetID[0] != 0) {
        attributeValueGrid.routeData.setID = selectionSetID[0]
        await attributeValueGrid.reload()
        attributeValueGrid.toolbar.enable('w2ui-add')
      } else {
        attributeValueGrid.clear()
        attributeValueGrid.toolbar.disable('w2ui-add')
      }
    },
  })

  window.addAttributeGroupSidebarItem = function() {
    if (!attributeGroupSidebar.get(0)) {
      const nodes = attributeGroupSidebar.nodes
      attributeGroupSidebar.nodes = [{ id: 0 }, ...nodes]
      attributeGroupSidebar.refresh()
    }
    attributeGroupSidebar.dblClick(0)
  }

  window.deleteAttributeGroupSidebarItem = function() {
    if (attributeGroupSidebar.selected != null) {
      w2confirm({
        msg: 'Are you sure you want to delete the selected group?',
        btn_yes: { text: 'Delete', class: 'w2ui-btn-red' },
        btn_no: { text: 'Cancel' },
      }).yes(async () => await attributeGroupSidebar._delete(attributeGroupSidebar.selected))
    }
  }

  return new w2layout({
    name: 'attributeLayout',
    panels: [
      {
        type: 'left', size: 220, style: 'margin-right: 5px;',
        html: `
        <style>
          #tb_attributeSetGrid_toolbar_item_w2ui-search .w2ui-grid-search-input,
          #tb_attributeSetGrid_toolbar_item_w2ui-search .w2ui-search-all,
          #tb_attributeValueGrid_toolbar_item_w2ui-search .w2ui-grid-search-input,
          #tb_attributeValueGrid_toolbar_item_w2ui-search .w2ui-search-all {
            width: 200px !important;
          }
        </style>
          <div class="flex flex-col h-full">
            <div class="py-2">
              <button class="w2ui-btn w2ui-btn-blue" onclick="addAttributeGroupSidebarItem()">Add Group</button>
              <button class="w2ui-btn" onclick="deleteAttributeGroupSidebarItem()">Delete Group</button>
            </div>
            <div id="attribute-group-sidebar" class="grow"></div>
          </div>
        `,
      },
      {
        type: 'main',
        html: new w2layout({
          name: 'attributeInnerLayout',
          panels: [
            { type: 'left', size: '50%', style: 'margin-right: 5px;', html: attributeSetGrid, resizable: true },
            { type: 'main', size: '50%', style: 'margin-left: 5px;', html: attributeValueGrid },
          ],
        }),
      },
    ],
    onRender: async event => {
      await event.complete
      attributeSetGrid.toolbar.disable('w2ui-add')
      attributeValueGrid.toolbar.disable('w2ui-add')
      attributeGroupSidebar.render('#attribute-group-sidebar')
      await attributeGroupSidebar._reload()
    },
    onDestroy: event => {
      event.owner.get('main').html.destroy() // inner layout
      attributeGroupSidebar.destroy()
      attributeSetGrid.destroy()
      attributeValueGrid.destroy()
      delete window.addAttributeGroupSidebarItem
      delete window.deleteAttributeGroupSidebarItem
    }
  })
}

