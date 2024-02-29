import axios from "axios";
import {GlobalStore} from "@/stores";
import {Trigger} from "@arco-design/web-vue";
import {ElMessage} from "element-plus";
import i18n from "@/lang";

const service =axios.create({
    baseURL:'',
    timeout:30000
})

service.interceptors.request.use(
    config=>{
        const globalStore = GlobalStore()
        if(globalStore.userInfo.token!=""){
            config.headers["token"] = globalStore.userInfo.token
        }
        return config
    }

)


service.interceptors.response.use(
    response => {
        const { data = {}, code = 0 } = response?.data || {};
        if (+code === 0) {
            return data || {};
        } else {
            Promise.reject(response?.data || {});
        }
    },
    (error = {}) => {
        const { response = {} } = error || {};
        const globalStore = GlobalStore()
        // 重定向
        if (+response?.status === 401) {
            globalStore.setToken("")

            if (response?.data.code !== 10006) {
                ElMessage(
                    {
                        showClose: true,
                        message: response?.data.msg || decodeURI(response.headers.msg),
                        type: 'error'
                    }
                )
            } else {
                // 打开登录弹窗
                globalStore.triggerAuth(true)
                globalStore.triggerAuthKey('signin')
            }
        } else {
            ElMessage(
                {
                    showClose: true,
                    message: response?.data.msg || decodeURI(response.headers.msg) || i18n.global.t('error.require') ,
                    type: 'error'
                }
            )
        }
        return Promise.reject(response?.data || {});
    }
)

export default service