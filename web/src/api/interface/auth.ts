export namespace Auth{
    export interface ReqLoginWithUserName{
        username:string
        password:string
    }

    export interface RespLogin{
        token:string
    }
}