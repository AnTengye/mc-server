# mc-server
用go实现的外置认证服务器，实现规范见[Yggdrasil 服务端](https://github.com/yushijinhun/authlib-injector/wiki)。

## 项目目的
和朋友一起玩minecraft时，一个朋友的正版账号突然无法通过microsoft验证。
为了继承原有的数据，遂通过外置认证服务器的方式绕过官方校验，来达到正版和”假正版“共同游玩的目的。

## 已实现API
- POST /authserver/authenticate
- POST /authserver/refresh
- POST /authserver/validate
- POST /authserver/invalidate
- POST /authserver/signout
- POST /sessionserver/session/minecraft/join
- GET /sessionserver/session/minecraft/hasJoined?username={username}&serverId={serverId}&ip={ip}
- GET /sessionserver/session/minecraft/profile/{uuid}?unsigned={unsigned}

## 使用
### 一、启动外置认证服务器
拉取代码

修改config/userid.go中的UserInfo。

格式为`
"minecraft内的角色名"："uuid"
`

PS:正版用户UUID可通过该链接查询：https://mcuuid.net/

执行命令

    go build
    
    ./mc-server.exe

默认启动在http://127.0.0.1:8899/
### 二、配置mc服务器
下载authlib-injector：https://authlib-injector.yushi.moe/

请先将服务器配置文件（一般为 server.properties）中 online-mode 一项的值设为 true，然后在你的服务端的启动指令的 -jar 参数前添加如下参数：

-javaagent:authlib-injector-1.1.44.jar=http://127.0.0.1:8899/

参考：[在 Minecraft 服务端使用 authlib injector](https://github.com/yushijinhun/authlib-injector/wiki/%E5%9C%A8-Minecraft-%E6%9C%8D%E5%8A%A1%E7%AB%AF%E4%BD%BF%E7%94%A8-authlib-injector)

### 三、配置mc启动器

1. 打开 HMCL 3 的「新建游戏账户」界面，选择登录方式为「外置登录（authlib-injector）」。
2. 点击「认证服务器」旁的加号「+」，输入 上面启动的认证服务器 的 API 地址(http://127.0.0.1:8899/ )，点击「下一步」，等待识别出 服务器信息 后点击「完成」
3. 在「认证服务器」菜单中选择 刚刚添加的服务器，填写minecraft的角色名和密码（随便填）。
4. 选择你的游戏角色，如果你只有一个角色，那么 HMCL 3 会默认选择那个唯一的角色。
