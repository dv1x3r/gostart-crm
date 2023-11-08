/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./internal/**/*.templ",
    "./pkg/**/*.templ",
    "./web/**/*.{html,js}",
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}

