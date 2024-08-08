/** @type {import('tailwindcss').Config} */

const defaultTheme = require('tailwindcss/defaultTheme')

module.exports = {
  content: [
    "./internal/**/*.templ",
    "./pkg/**/*.templ",
    "./web/**/*.{html,js}",
  ],
  theme: {
    extend: {
      fontFamily: {
        'sans': ['"Inter Variable"', ...defaultTheme.fontFamily.sans],
      },
    },
  },
  plugins: [],
}

