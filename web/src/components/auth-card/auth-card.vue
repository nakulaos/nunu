<template>
<el-dialog
    class="auth-card"
    v-model="globalStore.$state.authModalShow"
    :show-close="true"
    width="350px"


>
  <div class="auth-wrap">
    <template v-if="globalStore.$state.profile.allowUserRegister">
      <el-tabs v-model="globalStore.$state.authModalTab"  stretch class="auth-warp-tabs">
        <el-tab-pane name="signin">
          <template #label>
            <span>{{ $t('other.login') }}</span>
          </template>
          <el-card>
            <el-form
                :label-position="'left'"
                :model="loginForm"
                label-width="70px"
                :rules="LoginRules"
                class="auth-warp-tabs-form"
                ref = "loginRef"

            >
              <el-form-item prop="username">
                <template #label>
                  <span>{{ $t('variable.username') }}</span>
                </template>

                <el-input
                    v-model="loginForm.username"
                    :placeholder="loginFormUserNamePlaceHolder"
                  >
                </el-input>
              </el-form-item>
              <el-form-item prop="password">
                <template #label>
                  <span>{{ $t('variable.password') }}</span>
                </template>

                <el-input
                    v-model="loginForm.password"
                    :placeholder="loginFormPasswordPlaceHolder"
                    show-password
                >
                </el-input>
              </el-form-item>
              <el-form-item class="auth-warp-tabs-form__button" >
                <el-button type="primary" @click="handleLogin">
                  {{ $t('other.login')}}
                </el-button>
                <el-button type="info" @click="handleCancel">
                  {{ $t('other.cancel')}}
                </el-button>
              </el-form-item>




            </el-form>
          </el-card>

        </el-tab-pane>

        <el-tab-pane name="signup">
          <template #label >
            <span>{{ $t('other.register') }}</span>
          </template>
        </el-tab-pane>
      </el-tabs>
    </template>

  </div>



</el-dialog>
</template>

<script setup lang="ts">

import {GlobalStore} from "@/stores";
import {reactive,ref} from "vue";
import {ElMessage, FormRules} from "element-plus";
import type { FormInstance } from "element-plus";
import i18n from "@/lang";
import {userLoginWithUsername} from "@/api/modules/auth";

interface LoginForm {
  username: string,
  password: string,
}

const globalStore = GlobalStore()

const loginForm = reactive<LoginForm>({
  username:'',
  password:'',
})

const loginRef = ref<FormInstance>()

const loginFormUserNamePlaceHolder= i18n.global.t('validate.loginRule.username.required')


const loginFormPasswordPlaceHolder= i18n.global.t('validate.loginRule.password.required')


const LoginRules = reactive<FormRules<LoginForm>>({
  username:[
    {
      required:true,
      message: i18n.global.t('validate.loginRule.username.required')
    }
  ],
  password:[
    {
      required:true,
      message: i18n.global.t('validate.loginRule.password.required')
    }
  ]
})


const handleLogin = async ()=>{
  await loginRef.value.validate(
      (valid, fields) =>{
        if(valid){
          userLoginWithUsername({
            username: loginForm.username,
            password:loginForm.password,
          })
              .then((res)=>{
                const token = res?.token || '';
                globalStore.setToken(token)
                ElMessage.success(i18n.global.t('event.loginSuccess'))

                // TODO: 获取用户信息



              })

        }
      }
  )

};

const handleCancel = ()=>{

}


</script>


<style scoped lang="scss">
.auth-card::-webkit-scrollbar {
  // 隐藏滚动条
  width: 0;
  height: 0;
}

.auth-card{



  .auth-wrap{

    .auth-warp-tabs {

      .auth-warp-tabs-form{

        .auth-warp-tabs-form__button{

        }

      }


    }




  }
}






.el-card{
  border: 0;
  padding: 0 0 0 0;

}








</style>