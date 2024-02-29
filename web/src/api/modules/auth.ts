import {Auth} from "@/api/interface/auth";
import service from "@/api";


export const userLoginWithUsername = (params:Auth.ReqLoginWithUserName):Promise<Auth.RespLogin>=>{
    return  service.post('/api/v1/pub/auth/login/username',params)
}