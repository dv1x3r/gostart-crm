const path = require('path')

module.exports = {
  entry: {
    main: ['./src/index.js'],
    admin: ['./src/index.js', './src/admin.js'],
  },
  output: {
    filename: '[name].bundle.js',
    path: path.resolve(__dirname, 'dist'),
  },
  module: {
    rules: [
      {
        test: /\.css$/,
        use: ['style-loader', 'css-loader']
      }
    ]
  }
}

