export interface GlobalState{
    theme:string|null;
    desktopModelShow:boolean;
    language: string;
    isLogin: boolean;
    // 是否展示登录界面
    authModalShow:boolean,
    authModalTab:string
    userInfo:userInfo,
    profile:profile,
}


export interface userInfo {
    id:number,
    username:string,
    nickname:string,
    is_admin:boolean,
    avatar:string|null,
    token:string

}

export interface profile {
    useFriendship: boolean,
    enableWallet: boolean,
    allowUserRegister: boolean,
}