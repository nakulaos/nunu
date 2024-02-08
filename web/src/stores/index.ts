import {createPinia, defineStore} from 'pinia'
import piniaPersistConfig from "@/config/pinia-persist";
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate';
import type {GlobalState} from "@/stores/interface";


export const GlobalStore = defineStore({
        id: 'GlobalState',
        state: (): GlobalState => ({
            theme: localStorage.getItem("NUNU-THEME"),
            desktopModelShow: document.body.clientWidth > 821,
            test:"aaa",
        }),
        getters: {},
        actions: {
            initTest(s:string){
                this.test=s;
            }
        },
        persist: piniaPersistConfig('GlobalState'),

    }
)

const pinia = createPinia();
pinia.use(piniaPluginPersistedstate);

export default pinia;
