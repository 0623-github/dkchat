# user db table

|      列名      |  类型  | description | 备注 |
| :------------: | :----: | :---------: | :--: |
|       id       |  uint  |             | 自增 |
| account_number | string |    账号     | 唯一 |
|      pwd       | string |    密码     |      |
|   real_name    | string |  真实姓名   |      |
|   nick_name    | string |   用户名    |      |
|      age       |  int   |    年龄     |      |
|     grade      |  int   |    等级     |      |
|     gender     |  int   |    性别     | 0/1  |
|     phone      | string |  电话号码   | 唯一 |
|    is_exist    |  int   |  是否存在   | 0/1  |

# modular

## SignUp

注册账户

需要分配account_number

* 几位数？
  * 10位吧，Phone是11位
* 自增还是随机？
  * 随机会好一点
  * 自增写起来简单hhhh
  * 自增吧

### request

需要包含注册账号时用户填写的一些信息

> pwd --- 密码
>
> RealName --- 真实姓名
>
> NickName --- 昵称
>
> Age --- 年龄
>
> Gender --- 性别
>
> Phone --- 电话

### response

> code --- 自定义code
>
> message
>
> account_number --- 生成的账号



## SignIn

登录

