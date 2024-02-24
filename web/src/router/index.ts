import Nprogress from "@/config/nprogress";
import {GlobalStore} from "@/stores";
import router from "@/router/router";

router.beforeEach((to,from,next)=>{
    Nprogress.start();

    const globalStore = GlobalStore();

    // 匹配到login入口
    if(to.name==='Login' && globalStore.isLogin){
        globalStore.setLoginStatus(false)
        next({
            name:'Login',
        })
        Nprogress.done();
        return
    }

    // 权限校验
    if(to.matched.some((record)=>{
        record.meta.requiresAuth
    })){
        next({
            name:'Login'
        })
        return;
    }else{
        return next();
    }

})


router.afterEach(()=>{
    Nprogress.done();
})


export default router