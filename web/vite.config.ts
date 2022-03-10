import {ConfigEnv, defineConfig, UserConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import Components from 'unplugin-vue-components/vite'
import AutoImport from 'unplugin-auto-import/vite'
import { createSvgIconsPlugin } from 'vite-plugin-svg-icons'
import PurgeIcons from 'vite-plugin-purge-icons'
import {ElementPlusResolver } from "unplugin-vue-components/resolvers";
import { resolve } from 'path'

const root = process.cwd()
function pathResolve(dir: string) {
  return resolve(root, '.', dir)
}

// https://vitejs.dev/config/
export default ({command, mode }: ConfigEnv): UserConfig => {
  return {
    build: {
      outDir: "../html"
    },
    server: {
      fs: {
        strict: false
      },
      port: 9000,
      open: false,
      proxy: {
        '/api': "http://127.0.0.1:9001/api"
      }
    },
    resolve: {
      extensions: ['.mjs', '.js', '.ts', '.jsx', '.tsx', '.json', '.less', '.css'],
      alias: {
        "@": pathResolve("src")
      },
    },
    plugins: [
      vue(),
      AutoImport({
        resolvers: [ElementPlusResolver()]
      }),
      Components({
        dts: true,
        deep: true,
        resolvers: [ElementPlusResolver({importStyle: true})]
      }),
      createSvgIconsPlugin({
        iconDirs: [pathResolve("src/assets/svgs")],
        symbolId: "icon-[dir]-[name]",
        svgoOptions: true
      }),
      PurgeIcons()
    ],
    css: {
      preprocessorOptions: {
        less: {
          additionalData: `@import "./src/styles/variables.less";`,
          javascriptEnabled: true
        }
      }
    },
    optimizeDeps: {
      include: [
        'vue',
        'vue-router',
        'vue-types',
        '@iconify/iconify',
        'axios',
        'qs'
      ]
    }
  }
}
