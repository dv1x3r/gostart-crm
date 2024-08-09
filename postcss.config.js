module.exports = {
  map: false,
  plugins: [
    require('@csstools/postcss-sass')({
      sourceMap: false,
      includePaths: ['node_modules'],
    }),
    require('postcss-import'),
    require('tailwindcss'),
    require('autoprefixer'),
    require('cssnano'),
    {
      postcssPlugin: 'rename-files',
      OnceExit(css, { result }) {
        result.opts.to = result.opts.to.replace(/\.scss$/, '.css')
      }
    }
  ]
}
