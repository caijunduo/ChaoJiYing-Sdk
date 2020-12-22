# Install
    go get github.com/caijunduo/chaojiying-sdk

# Example

```go
package main

import (
    cjySdk "github.com/caijunduo/chaojiying-sdk"
    "log"
)

func main () {
    // 获取用户题分
    cjy := cjySdk.NewChaoJiYing()
    userInfo, err := cjy.UserInfo()
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("tifen: %d, lock_tifen: %d", userInfo.TiFen, userInfo.TiFenLock)
}
```

# API
> 超级鹰文档：https://www.chaojiying.com/api-5.html

#### 获取用户题分
    cjySdk.NewChaoJiYing().UserInfo()

#### 识别图片
    cjySdk.NewChaoJiYing().IdentifyPic(codeType int, minLen int, imgBase64 string)

#### 提交识别错误并返回题分
    cjySdk.NewChaoJiYing().ReportError(picId string)

#### 设置HTTPS代理
    cjySdk.NewChaoJiYing().SetHttpsProxy(u string)

#### 设置超时时间(秒)
    cjySdk.NewChaoJiYing().SetTimeout(t time.Duration)

#### 设置超级鹰用户账号，如设置环境变量：`CHAOJIYING_USER` 则默认获取环境变量
    cjySdk.NewChaoJiYing().SetUser(user string)

#### 设置超级鹰用户密码，如设置环境变量：`CHAOJIYING_PASS` 则默认获取环境变量
    cjySdk.NewChaoJiYing().SetPass(pass string)

#### 设置超级鹰用户密码的MD5值，如设置环境变量：`CHAOJIYING_PASS2` 则默认获取环境变量
    cjySdk.NewChaoJiYing().SetPass2(pass2 string)

#### 设置超级鹰软件ID，如设置环境变量：`CHAOJIYING_SOFT_ID` 则默认获取环境变量
    cjySdk.NewChaoJiYing().SetSoftId(softId string)