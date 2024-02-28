const  message = {
    commons:{
        city: "城市"
    },
    menu:{
        home:"广场",
        topic:"话题",
    }
};

import fit2cloudEnLocale from 'fit2cloud-ui-plus/src/locale/lang/zh-cn';

// ES6 的模块语法，导出了一个默认的对象，这个对象是 message 对象的拷贝。
// export default 语法用于导出模块的默认值，
// 其他模块导入时可以直接使用 import 关键字来引入这个默认值
export default {
    ...message,
    ...fit2cloudEnLocale,
}