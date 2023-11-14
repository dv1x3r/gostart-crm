import { w2utils } from 'w2ui/w2ui-2.0.es6'

export function getCsrfToken() {
  const cookieName = "_csrf"
  const name = cookieName + "="
  const decodedCookie = decodeURIComponent(document.cookie)
  const cookieArray = decodedCookie.split(';')

  for (let i = 0; i < cookieArray.length; i++) {
    let cookie = cookieArray[i].trim()
    if (cookie.indexOf(name) === 0) {
      return cookie.substring(name.length, cookie.length)
    }
  }

  return null
}

// W2UI Security Rules

// 1. use the following render function for ALL text fields
// row => w2utils.encodeTags(row.field)

// 2. do NOT allow inline editing for text fields
// this may lead to broken tag characters

export function safeRender(value) {
  return w2utils.encodeTags(value)
}

export function enablePreview(event) {
  event.owner.toolbar.enable('preview')
}

export async function disablePreview(event) {
  await event.complete
  if (event.owner.getChanges().length == 0) {
    event.owner.toolbar.disable('preview')
  }
}

