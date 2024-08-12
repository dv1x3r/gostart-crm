import Htmx from 'htmx.org'
window.htmx = Htmx

import { w2ui, w2layout, w2tooltip, w2utils } from 'w2ui/dist/w2ui.es6'

import { createProductGrid, openProductStatusPopup } from './admin/product'
import { categorySidebar, createCategoryGrid } from './admin/category'
import { createSupplierGrid } from './admin/supplier'
import { createBrandGrid } from './admin/brand'
import { createAttributeLayout } from './admin/attribute'

import { createUserLayout } from './admin/user'
import { updateOrderCounter, createOrderLayout, openOrderStatusPopup, openPaymentMethodPopup } from './admin/order'

import * as utils from './admin/utils'

window.w2ui = w2ui
window.w2tooltip = w2tooltip

w2utils.settings.dataType = 'JSON'

w2utils.locale({
  weekStarts: 'M',
  dateFormat: 'yyyy-MM-dd',
  datetimeFormat: 'yyyy-MM-dd|hh24:mi',
  timeFormat: 'hh24:mi',
})

w2utils.formatters['safe'] = value => w2utils.encodeTags(value)
w2utils.formatters['dropdown'] = value => w2utils.encodeTags(value?.text)

w2utils.formatters['safe'] = value => w2utils.encodeTags(value)
w2utils.formatters['dropdown'] = value => w2utils.encodeTags(value?.text)

w2utils.formatters['hover'] = value => `<span onmouseenter="w2tooltip.show(this, {'html': decodeURIComponent(atob(('${btoa(encodeURIComponent(value))}'))), 'name': 'no-name'})" onmouseleave="w2tooltip.hide('no-name')">${w2utils.encodeTags(value)}</span>`
w2utils.formatters['drophover'] = value => value?.text == null ? null : `<span onmouseenter="w2tooltip.show(this, {'html': decodeURIComponent(atob(('${btoa(encodeURIComponent(value.text))}'))), 'name': 'no-name'})" onmouseleave="w2tooltip.hide('no-name')">${w2utils.encodeTags(value.text)}</span>`

w2utils.formatters['thumb'] = value => value == null ? null : `<img src="media/${value}" style="max-width: 72px; max-height: 72px; margin: auto;"/>`

w2utils.formatters['color'] = value => {
  const safeValue = w2utils.encodeTags(value)
  return value == null ? null : `#<span style="color:#${safeValue}"><b>${safeValue}</b></span>`
}

w2utils.formatters['tag'] = value => {
  const label = w2utils.encodeTags(value?.text ?? value?.label)
  const color = w2utils.encodeTags(value?.color)
  return label == null ? null : `<span style="border-radius: 35px 200px 200px 35px; padding: 4px 10px; background-color: #${color}; color: white;"><b>${label}</b></span>`
}

export const mainLayout = new w2layout({
  name: 'mainLayout',
  box: '#main-layout',
  panels: [
    {
      type: 'top', size: 40, toolbar: {
        items: [
          { type: 'button', id: 'products', text: 'Products', icon: 'fa fa-book', onClick: () => tabManager.OpenTab('products-tab', 'Products', false, createProductGrid) },
          { type: 'button', id: 'categories', text: 'Categories', icon: 'fa fa-folder-tree', onClick: () => tabManager.OpenTab('categories-tab', 'Categories', true, createCategoryGrid) },
          { type: 'break' },
          { type: 'button', id: 'suppliers', text: 'Suppliers', icon: 'fa fa-warehouse', onClick: () => tabManager.OpenTab('suppliers-tab', 'Suppliers', true, createSupplierGrid) },
          { type: 'button', id: 'brands', text: 'Brands', icon: 'fa-brands fa-apple', onClick: () => tabManager.OpenTab('brands-tab', 'Brands', true, createBrandGrid) },
          { type: 'button', id: 'attributes', text: 'Attributes', icon: 'fa fa-wrench', onClick: () => tabManager.OpenTab('attributes-tab', 'Attributes', true, createAttributeLayout) },
          { type: 'break' },
          { type: 'button', id: 'users', text: 'Users', icon: 'fa fa-users', onClick: () => tabManager.OpenTab('users-tab', 'Users', true, createUserLayout) },
          { type: 'button', id: 'orders', text: 'Orders', icon: 'fa fa-cart-shopping', onClick: () => tabManager.OpenTab('orders-tab', 'Orders', true, createOrderLayout) },
          { type: 'break' },
          {
            type: 'menu', id: 'settings', text: 'Settings', icon: 'fa fa-list-check',
            items: [
              { type: 'html', id: 'catalog-settings', text: 'Catalog settings', disabled: true },
              { type: 'button', id: 'product-status', text: 'Product Status', icon: 'fa fa-tags', onClick: () => openProductStatusPopup() },
              { type: 'break' },
              { type: 'html', id: 'order-settings', text: 'Order settings', disabled: true },
              { type: 'button', id: 'order-status', text: 'Order Status', icon: 'fa fa-envelope-circle-check', onClick: () => openOrderStatusPopup() },
              { type: 'button', id: 'payment-method', text: 'Payment Methods', icon: 'fa-solid fa-credit-card', onClick: () => openPaymentMethodPopup() },
            ],
          },
          { type: 'spacer' },
          { type: 'button', id: 'www', text: utils.isLocalhost() ? 'localhost:1323' : 'democrm.weasel.dev', icon: 'fa fa-arrow-right-to-city', onClick: () => window.open(utils.isLocalhost() ? 'http://localhost:1323' : 'https://democrm.weasel.dev', '_blank') },
          { type: 'button', id: 'logout', text: 'Log out', icon: 'fa fa-right-from-bracket', onClick: () => window.location = '/logout/' },
        ],
        onClick: event => utils.w2menuOnClick(event),
        _updateOrderCounter: updateOrderCounter,
      }
    },
    { type: 'left', size: 260, html: categorySidebar },
    {
      type: 'main',
      tabs: {
        onClick: event => utils.w2tabOnClick(event, '#layout_mainLayout_panel_main > .w2ui-panel-content'),
        onClose: event => utils.w2tabOnClose(event),
      },
    },
  ]
})

export const tabManager = new utils.W2TabManager(mainLayout.get('main').tabs)

