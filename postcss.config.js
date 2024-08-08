module.exports = {
  map: false,
  plugins: [
    require('@csstools/postcss-sass')({ sourceMap: false, includePaths: ['node_modules'] }),
    require('tailwindcss'),
    require('autoprefixer'),
    require('cssnano'),
  ]
}
