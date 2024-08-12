import { w2grid, w2sidebar, w2utils } from 'w2ui/dist/w2ui.es6'
import * as utils from './utils'
import * as admin from '../admin'

function normalizeString(str) {
  // Normalize the string to ensure consistent comparison
  return str.normalize("NFD").replace(/[\u0300-\u036f]/g, "");
}

window.searchCategory = function(value) {
  categorySidebar.expandAll()
  categorySidebar.search(value, (str, node) => {
    const str1 = normalizeString(str.toLowerCase())
    const str2 = normalizeString(node.text.toLowerCase())
    return str2.indexOf(str1) != -1
  })
}

export const categorySidebar = new w2sidebar({
  name: 'categorySidebar',
  levelPadding: 8,
  topHTML: '<div style="height: 36px; padding: 3px 5px;"><input id="category-search" style="width: 100%" class="w2ui-input" placeholder="Search categories..." onkeyup="searchCategory(this.value)"></div>',
  onRender: async event => {
    await event.owner._reload()
    event.owner.select(0)
  },
  onClick: async event => {
    await event.complete

    const tabs = admin.tabManager.GetTabs()
    if (tabs.active != 'categories-tab' && tabs.active != 'products-tab') {
      tabs.click('products-tab')
    }

    const categoryGrid = tabs.get('categories-tab')?.component
    if (categoryGrid) {
      categoryGrid.routeData.parentID = categorySidebar.selected
      await categoryGrid.reload()
    }

    const productGrid = tabs.get('products-tab')?.component
    if (productGrid) {
      productGrid.routeData.categoryID = categorySidebar.selected
      await productGrid.reload()
    }
  },
  _reload: async () => {
    function treeToNodes(data) {
      return data.map(x => ({
        id: x.id,
        text: w2utils.encodeTags(x.name),
        icon: w2utils.encodeTags(x.icon) ?? 'dummy-icon',
        count: x.related_products,
        expanded: true,
        style: !x.is_published && x.id != 0 ? 'color: darkgrey;' : null,
        nodes: treeToNodes(x.children)
      }))
    }

    const previousSelection = categorySidebar.selected

    const res = await fetch('/category/tree', { method: 'GET', headers: { 'X-CSRF-Token': utils.getCsrfToken() } })
    if (res.status != 200) {
      await utils.fetchShowError('Failed to fetch categories', res)
      return
    }

    try {
      const data = await res.json()
      categorySidebar.remove(...categorySidebar.nodes.map(x => x.id))
      if (data.data != null) {
        categorySidebar.add(treeToNodes([data.data]))
        categorySidebar.select(previousSelection)
      }
    } catch {
      w2utils.notify(`Failed to read categories`, { timeout: 4000, error: true })
    }
  }
})

export function createCategoryGrid() {
  return new w2grid({
    name: 'categoryGrid',
    url: {
      get: '/category/:parentID/children',
      save: '/category/save',
      remove: '/category/delete',
    },
    httpHeaders: { 'X-CSRF-Token': utils.getCsrfToken() },
    routeData: { parentID: categorySidebar.selected ?? 0 },
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
      { field: 'name', text: 'Name', size: '240px', render: 'safe', editable: { type: 'text' } },
      { field: 'icon', text: 'Icon', size: '200px', render: 'safe', editable: { type: 'text' }, info: { render: () => 'Use any free icon from <u><a href="https://fontawesome.com/search?o=r&m=free" target="_blank">Font Awesome</a></u>' } },
      { field: 'attribute_group', text: 'Attributes', size: '200px', render: 'dropdown', editable: utils.getSelectOptions('/attribute/group/dropdown', 'list') },
      { field: 'parent', text: 'Parent', size: '200px', render: 'dropdown', editable: utils.getSelectOptions('/category/dropdown', 'list') },
      { field: 'is_published', text: 'Is Pub', size: '60px', editable: { type: 'checkbox' }, tooltip: 'Show this category' },
    ],
    searches: [
      { field: 'name', label: 'Name', type: 'text', _all: true },
      { field: 'icon', label: 'Icon', type: 'text', _all: true },
      { field: 'attribute_group', label: 'Attributes', type: 'enum', options: utils.getSelectOptions('/attribute/group/dropdown') },
      { field: 'no_tax', label: 'No Tax', type: 'enum', options: utils.getSelectOptionsBool() },
      { field: 'is_published', label: 'Is Published', type: 'enum', options: utils.getSelectOptionsBool() },
    ],
    defaultOperator: {
      'text': 'contains',
    },
    onSearch: event => utils.gridSearchAllowedAll(event),
    onAdd: async event => {
      const id = event.owner.routeData.parentID
      const initial = { is_published: true, parent: { id: id, text: '«« CURRENT »»' } }
      await utils.gridNewRowAdd(event, initial)
    },
    onChange: async event => await utils.gridNewRowChange(event),
    onSave: async event => {
      await utils.gridShowSaved(event)
      await categorySidebar._reload()
      categorySidebar.select(event.owner.routeData.parentID)
    },
    onDelete: async event => {
      await utils.gridShowDeleted(event)
      await categorySidebar._reload()
      categorySidebar.select(event.owner.routeData.parentID)
    },
    onReorderRow: async event => {
      await utils.gridPostReorderRow(event, '/category/reorder')
      await categorySidebar._reload()
      categorySidebar.select(event.owner.routeData.parentID)
    },
    onLoad: async event => await utils.gridMarkRows(event, x => !x.is_published ? 'darkgrey' : null),
  })
}

