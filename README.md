# BBS 仿reddit论坛项目

## 项目架构

### 配置模块

- 使用viper配置管理器来管理配置文件，将连接MySQL用户名和密码，redis的一些连接配置都写进一个yaml文件里面，在setting文件里面利用viper库对配置进行初始化。

### 注册模块

- 利用gorm进行MySQL相关的数据库的操作，读取表单传来的用户注册信息，利用validator库对表单的数据加上限制，使用雪花算法生成用户ID，并对用户密码使用对称加密然后储存进数据库中。

### 登录模块

- 获取表单信息，获取用户名使用gorm操作数据库查询是否有这个用户，返回该用户并将密码进行解密与表单提交的明文密码进行比对看是否是该用户的密码。密码核对成功及登录成功并生成token。

### 社区列表模块

- 获取社区ID，社区名。可由社区ID查询社区详细信息并返回。

### 帖子模块

- 用户在某个社区发帖创建帖子，由雪花算法生成帖子ID，并通过redis记录帖子的得票情况，创建时间，帖子的是持久化到MySQL中的。帖子列表的获取包含两种情况，一种是通过帖子创建的时间按从新到旧排列的一定数目返回
第二种就是根据redis中存储的每个帖子的得分情况的大小来进行排序并按一定数量进行返回。

### 投票模块

- 帖子的投票算法是简化版的reddit论坛投票算法，将帖子的创建时间的时间戳和每票得分(432)相加，票数越高得分越高，时间越新得分越高。

### 中间件

- 采用jwt中间件进行登录的权限验证，包括投票，发帖等一系列论坛用户才能做的相关操作进行权限验证。

## 测试

- 本项目采用postman对该项目的路由都进行过测试，并由swagger生成接口文档。
该项目中的前端部分并未进行过前后端对接和调试，有兴趣的朋友可自行调式部署。