import NProgress from 'nprogress';
import 'nprogress/nprogress.css';

// 配置 NProgress 进度条的参数
NProgress.configure({
    easing: 'ease', // 进度条动画的缓动函数
    speed: 300, // 进度条动画完成的速度（毫秒）
    showSpinner: true, // 是否显示加载图标
    trickleSpeed: 200, // 自动增加进度条的速度（毫秒）
    minimum: 0.3, // 进度条显示的最小百分比
});

export default NProgress; // 导出 NProgress 实例