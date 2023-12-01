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

export function safeRender(value) {
  return w2utils.encodeTags(value) ?? ""
}
