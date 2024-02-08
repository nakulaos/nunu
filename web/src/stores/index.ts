import { defineStore } from 'pinia'

const theme:string|null = localStorage.getItem("NUNU-THEME")
const desktopModelShow:boolean = document.body.clientWidth > 821


export const GlobalStore = defineStore(
    'GlobalState',
    {
    state(){
        return{
            theme:theme,
            desktopModelShow:desktopModelShow
        }

    },
})
