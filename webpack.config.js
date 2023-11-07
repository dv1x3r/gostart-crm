const path = require('path')

module.exports = {
  entry: {
    main: ['./web/index.js'],
    admin: ['./web/index.js', './web/admin.js'],
  },
  output: {
    filename: '[name].bundle.js',
    path: path.resolve(__dirname, 'dist'),
  },
  module: {
    rules: [
      {
        test: /\.css$/,
        use: ['style-loader', 'css-loader', 'postcss-loader']
      },
    ],
  },
}

