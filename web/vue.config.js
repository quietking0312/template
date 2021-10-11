const path = require('path')
const Components = require('unplugin-vue-components/webpack')
const { ElementPlusResolver } = require('unplugin-vue-components/resolvers')

function resolve(dir) {
  return path.join(__dirname, dir)
}

const vueConfig = {
  publicPath: "./",
  lintOnSave: process.env.NODE_ENV === 'development',
  productionSourceMap: process.env.NODE_ENV !== 'production', // 生产环境不生成 sourceMap 文件
  devServer: {
    port: 8989,
    open: false,
    overlay: {
      warnings: false,
      errors: true
    }
  },

  chainWebpack: config => {
    // 引入公用less
    const types = ['vue', 'vue-modules', 'normal-modules', 'normal']
    types.forEach((type) => addStyleResource(config.module.rule('less').oneOf(type)))

    config.resolve.alias.set('@', resolve('src'))
    config.resolve.alias.set('@c', resolve('src/components'))

  },
  configureWebpack: {
    plugins: [
      Components({
        resolve: [ElementPlusResolver()]
      })
    ],
  }

}

function addStyleResource(rule) {
  rule.use('style-resource')
      .loader('style-resources-loader')
      .options({
        patterns: [
          path.resolve(__dirname, './src/styles/variables.less') // 需要全局导入的less
        ]
      })
}

module.exports = vueConfig
