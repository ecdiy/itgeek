### chrome 插件方式开发时，不能在html中定义组件


npm install --save vue-draggable-resizable


Lua

## 

```
window.PluginInstall = function (th, msg) {
    var t = Cookies.get("liveTime");
    if (t && (new Date()).getTime() - Number(t) < 10000) {
        return true
    }
    th.$Modal.info({
        content: '<a v-if="pluginFlag==2" target="_blank" href="http://www.ecdiy.cn/lib/EcAssistant/EcdiyAssistant.crx">' +
        msg + '<br><h2>点击安装</h2> </a> ', okText: "取消"
    });
    return false
};

```

1.项目里安装vux 
  npm install vux --save 
  
2.安装vux-loader（vux2必须结合vux-loader使用，本文也只是介绍vux2的配置） 
  npm install vux-loader --save-dev 
