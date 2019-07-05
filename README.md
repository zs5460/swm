# swm

send wechat message

swm是一个向微信企业号应用成员群发送消息的小工具。

## 配置

将config.json.sample改名为config.json，填入相关信息。

```json
{
    "appID":"",  // 企业号的corpid
    "appKey":"", // 应用的Secret
    "agentID":"" // 应用的AgentId
}
```

## 使用

```shell
swm hello
```

## Licence

Released under MIT license, see [LICENSE](LICENSE) for details.
