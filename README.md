# mc-server
��goʵ�ֵ�������֤��������ʵ�ֹ淶��[Yggdrasil �����](https://github.com/yushijinhun/authlib-injector/wiki)��

## ��ĿĿ��
������һ����minecraftʱ��һ�����ѵ������˺�ͻȻ�޷�ͨ��microsoft��֤��
Ϊ�˼̳�ԭ�е����ݣ���ͨ��������֤�������ķ�ʽ�ƹ��ٷ�У�飬���ﵽ����͡������桰��ͬ�����Ŀ�ġ�

## ��ʵ��API
- POST /authserver/authenticate
- POST /authserver/refresh
- POST /authserver/validate
- POST /authserver/invalidate
- POST /authserver/signout
- POST /sessionserver/session/minecraft/join
- GET /sessionserver/session/minecraft/hasJoined?username={username}&serverId={serverId}&ip={ip}
- GET /sessionserver/session/minecraft/profile/{uuid}?unsigned={unsigned}

## ʹ��
### һ������������֤������
��ȡ����

�޸�config/userid.go�е�UserInfo��

��ʽΪ`
"minecraft�ڵĽ�ɫ��"��"uuid"
`

PS:�����û�UUID��ͨ�������Ӳ�ѯ��https://mcuuid.net/

ִ������

    go build
    
    ./mc-server.exe

Ĭ��������http://127.0.0.1:8899/
### ��������mc������
����authlib-injector��https://authlib-injector.yushi.moe/

���Ƚ������������ļ���һ��Ϊ server.properties���� online-mode һ���ֵ��Ϊ true��Ȼ������ķ���˵�����ָ��� -jar ����ǰ������²�����

-javaagent:authlib-injector-1.1.44.jar=http://127.0.0.1:8899/

�ο���[�� Minecraft �����ʹ�� authlib injector](https://github.com/yushijinhun/authlib-injector/wiki/%E5%9C%A8-Minecraft-%E6%9C%8D%E5%8A%A1%E7%AB%AF%E4%BD%BF%E7%94%A8-authlib-injector)

### ��������mc������

1. �� HMCL 3 �ġ��½���Ϸ�˻������棬ѡ���¼��ʽΪ�����õ�¼��authlib-injector������
2. �������֤���������ԵļӺš�+�������� ������������֤������ �� API ��ַ(http://127.0.0.1:8899/ )���������һ�������ȴ�ʶ��� ��������Ϣ ��������ɡ�
3. �ڡ���֤���������˵���ѡ�� �ո���ӵķ���������дminecraft�Ľ�ɫ�������루������
4. ѡ�������Ϸ��ɫ�������ֻ��һ����ɫ����ô HMCL 3 ��Ĭ��ѡ���Ǹ�Ψһ�Ľ�ɫ��
