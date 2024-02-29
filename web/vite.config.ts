import { fileURLToPath, URL } from 'node:url'

import { defineConfig,loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'
import {viteElectronDev} from "./plugins/vite.electron.dev"
import {viteElectronBuild} from "./plugins/vite.electron.build"
import type {ImportMetaEnv} from './env'
// https://vitejs.dev/config/
// export default defineConfig({
//   plugins: [
//     vue(),
//     vueJsx(),
//       viteElectronDev(),
//       viteElectronBuild(),
//   ],
//   resolve: {
//     alias: {
//       '@': fileURLToPath(new URL('./src', import.meta.url))
//     }
//   },
//   server:{
//     proxy: {
//       "/api":{
//         target: ,
//         changeOrigin: true,
//       },
//     }
//   }
// })


export default defineConfig(({mode})=>{
  let env:Record<keyof ImportMetaEnv, string>=loadEnv(mode,process.cwd())
  return{
    plugins: [
      vue(),
      vueJsx(),
      viteElectronDev(),
      viteElectronBuild(),
    ],
    resolve: {
      alias: {
        '@': fileURLToPath(new URL('./src', import.meta.url))
      }
    },
    server:{
      host: "0.0.0.0",
      port: 5173,
      proxy: {
        "/api":{
          target: env.VITE_HOST ,
          changeOrigin: true,
        },
      }
    }
  }

})
