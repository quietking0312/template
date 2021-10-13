import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import Components from 'unplugin-vue-components/vite'
import viteSvgIcons from 'vite-plugin-svg-icons'
import {ElementPlusResolver } from "unplugin-vue-components/resolvers";
import { resolve } from 'path'

// https://vitejs.dev/config/
export default defineConfig({
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
        alias: {
            "@": resolve(__dirname, "src")
        },
    },
    plugins: [
      vue(),
      Components({
        resolvers: [ElementPlusResolver()]
      }),
        viteSvgIcons({
            iconDirs: [resolve(process.cwd(), "src/assets/icons/svg")],
            symbolId: "icon-[dir]-[name]"
        })
    ],
    css: {
        preprocessorOptions: {
            less: {
                additionalData: `@import "./src/styles/variables.less";`,
                javascriptEnabled: true
            }
        }
    }
})
