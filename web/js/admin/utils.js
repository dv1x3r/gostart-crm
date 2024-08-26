import { w2utils, query } from 'w2ui/dist/w2ui.es6'

export function isLocalhost() {
  return window.location.hostname == 'localhost' || window.location.hostname == '127.0.0.1'
}

export function getCsrfToken() {
  return document.querySelector('meta[name="csrf-token"]').getAttribute('content')
}

export function getURI(path) {
  return new URL(path, window.location).href
}

export function fetchShowSuccess(successMessage) {
  w2utils.notify(successMessage, { timeout: 4000 })
}

export async function fetchShowError(errorMessage, res) {
  if (res == null) {
    w2utils.notify(errorMessage, { timeout: 4000, error: true })
    return
  }
  try {
    const data = await res.json()
    w2utils.notify(`${res.status}: ${data?.message ?? res.statusText}`, { timeout: 4000, error: true })
  } catch {
    w2utils.notify(`${errorMessage}, ${res.status}: ${res.statusText}`, { timeout: 4000, error: true })
  }
}

export async function gridShowSaved(event) {
  if (event.detail.data?.status == 'success') {
    w2utils.notify('Data has been successfully saved!', { timeout: 4000 })
    await event.owner.reload()
  }
}

export async function gridShowDeleted(event) {
  if (event.detail.data?.status == 'success') {
    w2utils.notify('Data has been successfully deleted!', { timeout: 4000 })
  }
}

export async function gridNewRowAdd(event, changes) {
  if (!event.owner.get(0)) {
    event.owner.add({ id: 0 }, true)
    event.owner.get(0).w2ui = { changes: changes }
    event.owner.refresh()
  }
  // fix: edit mode is not activated if search is applied
  setTimeout(() => event.owner.editField(0, 1), 100)
}

export async function gridNewRowChange(event) {
  await event.complete
  const newRow = event.owner.get(0)
  const changes = event.owner.getChanges([newRow])
  if (newRow && changes.length == 0) {
    event.owner.editDone(0, 1) // exit edit mode to prevent next row edit
    event.owner.remove(0)
  }
}

export async function gridDblClickNonEditable(event, fn) {
  const columnIndex = event.detail.column
  const editableFields = Object.keys(event.owner.columns[columnIndex].editable)
  const isEditable = editableFields.length > 0
  if (!isEditable) {
    await fn(event)
  }
}

export function gridSearchAllowedAll(event) {
  if (event.detail.searchField == 'all') {
    const fieldsAll = event.owner.searches.filter(x => x._all).map(x => x.field)
    event.detail.searchData = event.detail.searchData.filter(x => fieldsAll.includes(x.field))
  }
}

export async function gridMarkRows(event, fn) {
  await event.complete
  event.detail.data?.records?.forEach(x => {
    const color = fn(x)
    if (color != null) {
      const row = event.owner.get(x.id)
      row.w2ui = { style: `color: ${color} !important;` }
      event.owner.refreshRow(x.id)
    }
  })
}

export async function gridPostReorderRow(event, url) {
  if (event.owner.get(0)) {
    w2utils.notify('Reorder cancelled, please save the data first', { timeout: 4000, error: true })
    event.preventDefault()
    return
  }

  event.owner.lock('Reordering...')

  const res = await fetch(url, {
    method: 'POST',
    headers: {
      'X-CSRF-Token': getCsrfToken(),
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      'id': parseInt(event.detail.recid),
      'moveBefore': parseInt(event.detail.moveBefore == 'bottom' ? 0 : event.detail.moveBefore),
    }),
  })

  event.owner.unlock()

  if (res.status != 200) {
    await fetchShowError('Failed to reorder the data', res)
    return
  }

  fetchShowSuccess('Data has been successfully reordered!')
}

export async function gridSetFormRecord(event, form) {
  await event.complete
  const selection = event.owner.getSelection()
  const id = selection.length == 1 ? selection[0] : null
  if (id == null) {
    form.recid = null
    form.clear()
  } else {
    const gridRecord = event.owner.get(id)
    form.recid = id
    form.record = gridRecord
    form.refresh()
  }
}

export function getSelectOptions(url, type) {
  return {
    type: type,
    url: url,
    recId: 'id',
    renderDrop: value => w2utils.encodeTags(value?.text),
    match: 'contains',
    openOnFocus: true,
    cacheMax: 5000,
    minLength: 0,
  }
}

export function getSelectOptionsBool() {
  return { items: [{ id: '1', text: 'True' }, { id: '0', text: 'False' }] }
}

export function w2menuOnClick(event) {
  // execute onClick function for menu items
  if (event.target.includes(':') && event.detail.subItem?.onClick) {
    event.preventDefault()
    event.detail.subItem.onClick(event)
  }
}

export function w2contextMenuOnClick(event) {
  event.detail.menuItem.onClick(event)
}

export function w2tabOnClick(event, selector) {
  const tabID = event.detail.tab.id
  const el = query(`#${tabID}`)

  // 1. hide all div's with content
  query(`${selector} > div`).addClass('hidden')

  // 2. if element is already in the DOM, then just unhide it, otherwise render a new one
  if (el.length) {
    el.removeClass('hidden')
  } else {
    query(selector).append(`<div id="${tabID}" style="height:100%;"></div>`)
    event.detail.tab.component.render(`#${tabID}`)
  }
}

export function w2tabOnClose(event) {
  if (event.owner.lock) {
    event.preventDefault()
    return
  }

  event.owner.lock = true
  event.onComplete = () => {
    event.owner.lock = false
  }

  const activeTabID = event.owner.active
  const closedTabID = event.target
  const closedTabIX = event.owner.tabs.findIndex(x => x.id == closedTabID)
  const nextTab = event.owner.tabs.at(closedTabIX + 1)
  const prevTab = event.owner.tabs.at(closedTabIX - 1)

  // 1. remove div with content and w2ui reference
  query(`#${closedTabID}`).remove()
  event.detail.tab.component.destroy()

  // 2. open the next or previous available tab
  if (closedTabID != activeTabID) {
    return // skip if closed tab is not the active one
  } else if (nextTab && nextTab.id != activeTabID) {
    event.owner.click(nextTab.id)
  } else if (prevTab && prevTab.id != activeTabID) {
    event.owner.click(prevTab.id)
  }
}

export class W2TabManager {
  constructor(tabs) {
    this.tabs = tabs
  }

  GetTabs() {
    return this.tabs
  }

  OpenTab(id, name, closable, fn, ...args) {
    if (this.tabs.lock) {
      return
    }

    if (!this.tabs.get(id)) {
      const safeName = w2utils.encodeTags(name)
      this.tabs.add({
        id: id,
        text: safeName.length > 32 ? safeName.slice(0, 32) + '...' : safeName,
        closable: closable,
        component: fn(...args)
      })
      this.tabs.refresh()
    }

    this.tabs.click(id)
  }

  CloseTab(fn) {
    const tabsSlice = this.tabs.tabs.filter(fn)
    if (tabsSlice.length) {
      this.tabs.clickClose(tabsSlice[0].id)
    }
  }

  RenameTab(name, fn) {
    const tabsSlice = this.tabs.tabs.filter(fn)
    if (tabsSlice.length) {
      const safeName = w2utils.encodeTags(name)
      tabsSlice[0].text = safeName.length > 32 ? safeName.slice(0, 32) + '...' : safeName
      this.tabs.refresh()
    }
  }
}

