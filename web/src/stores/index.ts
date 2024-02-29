import {createPinia, defineStore} from 'pinia'
import piniaPersistConfig from "@/config/pinia-persist";
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate';
import type {GlobalState} from "@/stores/interface";


export const GlobalStore = defineStore({
        id: 'GlobalState',
        state: (): GlobalState => ({
            theme: localStorage.getItem("NUNU-THEME"),
            desktopModelShow: document.body.clientWidth > 821,
            language:"",
            isLogin:false,
            authModalShow: false,
            authModalTab:'signin',
            userInfo:{
                id:0,
                username:"",
                nickname:"",
                is_admin:false,
                avatar:"",
                token:"",
            },
            profile:{
                useFriendship:true,
                enableWallet: false,
                allowUserRegister: true,
            }
        }),
        getters: {},
        actions: {
            setLoginStatus(Login:boolean){
                this.isLogin = Login
            },
            triggerAuth(status:boolean){
                this.authModalShow = status
            },
            triggerAuthKey(key:string){
                this.authModalTab = key
            },
            setToken(token:string){
                this.userInfo.token = token
            }




        },
        persist: piniaPersistConfig('GlobalState'),

    }
)

const pinia = createPinia();
pinia.use(piniaPluginPersistedstate);

export default pinia;
