import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import i18n from '@/lang/index'
import pinia from "@/stores";
import ElementPlus from 'element-plus';
import Fit2CloudPlus from 'fit2cloud-ui-plus';
import * as Icons from '@element-plus/icons-vue';


const app = createApp(App)

app.use(pinia)
app.use(i18n)
app.use(router)
app.use(ElementPlus)


app.mount('#app')
