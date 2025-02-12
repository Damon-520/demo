### 项目说明：

1. 安装go1.8+及依赖的类库，见RANDME.md
2. 导入docs/test_activity.sql 数据库
3. 修改configs/local/api/config.yaml 配置文件
4. 项目根目录运行命令
   make all && make build && bin/demoapi
   或分开运行
   make all
   make build
   bin/demoapi
5. 测试使用mock目录 _1_live_room_list.http