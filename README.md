# **GDBC-数据库连接工具**

​	**用于在渗透测试工作过程中获取到数据库信息之后的验证工作，主要获取数据库表信息，后期会增加交互式sql shell输出的结果方便打工人截图放报告，不用再去找什么navicat等奇奇怪怪的东西了。**

- **使用方法:**

  ```shell
  Usage of gdbc.exe:
    -P string
          database port example:3306(optional,with default value.)
    -U string
          database url example:mysql://root:123456@127.0.0.1:3306/test
    -d string
          database type example:mysql,mssql,oracle
    -h string
          database host example:127.0.0.1
    -p string
          database password example:123456
    -u string
          database user example:root
  ```

  

![1](./img/2.png)

- **效果展示**

![1](./img/1.png)

