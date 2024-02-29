const  message = {
    commons:{
        city: "city"
    },
    menu:{
        home:"Home",
        topic:"Topic",
        profile:"profile",
        notification:"notification",
        collection:"collection",
        setting:"setting",
        contacts:"contacts",
    },
    other:{
        login: 'login',
        register: 'register'
    },
    variable:{
        username: 'username',
        password: 'password',
        cancel: 'cancel',
    },
    validate:{
        loginRule:{
            username:{
                required: "please input username"
            },
            password:{
                required:"please input password"
            }
        }
    },
    error:{
        require: "failure to require"
    },
    event:{
        loginSuccess: "login succeeded"
    }

};

import fit2cloudEnLocale from 'fit2cloud-ui-plus/src/locale/lang/en';

export default{
    ...message,
    ...fit2cloudEnLocale,
}