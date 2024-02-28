const  message = {
    commons:{
        city: "city"
    },
    menu:{
        home:"Home",
        topic:"Topic",
    }

};

import fit2cloudEnLocale from 'fit2cloud-ui-plus/src/locale/lang/en';

export default{
    ...message,
        ...fit2cloudEnLocale,
}