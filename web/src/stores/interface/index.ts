export interface GlobalState{
    theme:string|null;
    desktopModelShow:boolean;
    language: string;
    isLogin: boolean;
    userInfo:userInfo,
}


export interface userInfo {
    id:number,
    username:string,
    nickname:string,
    is_admin:boolean,
}