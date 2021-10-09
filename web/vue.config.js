const path = require('path')


function resolve(dir) {
  return path.join(__dirname, dir)
}

module.exports = {
  publicPath: "./",
  lintOnSave: process.env.NODE_ENV === 'development',
  productionSourceMap: process.env.NODE_ENV !== 'production', // 生产环境不生成 sourceMap 文件
  devServer: {
    port: 8989,
    open: false,
    overlay: {
      warnings: true,
      errors: true
    }
  },
  chainWebpack: config => {
    config.resolve.alias.set('@', resolve('src'))
    config.resolve.alias.set('@/_c', resolve('src/components'))

  }
}
