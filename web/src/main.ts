import '@/style/style.css'




import { createApp } from 'vue'

import App from './App.vue'
import router from '@/router/index'
import i18n from '@/lang/index'
import pinia from "@/stores";
import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css'
import Fit2CloudPlus from 'fit2cloud-ui-plus';
import * as Icons from '@element-plus/icons-vue';


const app = createApp(App)

app.use(pinia)
app.use(i18n)
app.use(router)
app.use(ElementPlus)
app.use(Fit2CloudPlus, { locale: i18n.global.messages.value[localStorage.getItem('lang') || 'zh'] });

// 这段代码的含义是遍历 Icons 对象的所有键
// 并通过循环将每个键对应的值作为组件注册到 Vue 应用程序中
Object.keys(Icons).forEach((key) => {
    app.component(key, Icons[key as keyof typeof Icons]);
});


app.mount('#app')
