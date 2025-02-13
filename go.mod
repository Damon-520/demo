module demoapi

go 1.22.12

require (
	github.com/elastic/go-elasticsearch/v8 v8.17.1 // Elasticsearch的官方Go客户端（开源）
	github.com/gin-gonic/gin v1.10.0 // Gin框架，用于Web开发，处理HTTP请求（开源）
	github.com/go-kratos/gin v0.1.0 // Kratos框架的Gin适配器（开源）
	github.com/go-kratos/kratos/v2 v2.8.3 // Go语言微服务框架，提供服务发现、链路跟踪等功能（开源）
	github.com/go-redis/redis/v8 v8.11.5 // Redis客户端，支持与Redis交互（开源）
	github.com/go-resty/resty/v2 v2.16.5 // HTTP客户端库，用于简化API调用（开源）
	github.com/golang-module/carbon v1.7.3 // 处理日期和时间的库，类似于Carbon（PHP）（开源）
	github.com/google/wire v0.6.0 // Go的依赖注入库（开源）
	github.com/jinzhu/copier v0.4.0 // 结构体复制工具，可以快速复制数据结构（开源）
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible // 日志文件按时间滚动的库（开源）
	github.com/mailru/easyjson v0.9.0 // JSON编码解码器，比标准库更高效（开源）
	github.com/nacos-group/nacos-sdk-go v1.1.5 // Nacos客户端，用于配置管理和服务发现（开源）
	github.com/openzipkin/zipkin-go v0.4.3 // Zipkin分布式追踪系统的Go实现（开源）
	github.com/pkg/errors v0.9.1 // 错误处理工具，提供堆栈跟踪和错误包装（开源）
	github.com/satori/go.uuid v1.2.0 // 生成UUID的库（开源）
	github.com/sirupsen/logrus v1.9.3 // 日志记录工具，支持多种日志级别和输出格式（开源）
	github.com/sony/sonyflake v1.2.0 // 唯一ID生成器，类似Snowflake（开源）
	github.com/spf13/cast v1.7.1 // 用于类型转换的工具库（开源）
	google.golang.org/protobuf v1.36.5 // Protocol Buffers（protobuf）的Go实现，用于序列化数据（开源）
	gorm.io/driver/mysql v1.5.7 // GORM MySQL数据库驱动（开源）
	gorm.io/driver/postgres v1.5.11 // GORM PostgreSQL数据库驱动（开源）
	gorm.io/gorm v1.25.12 // GORM是一个ORM框架，用于简化数据库操作（开源）
	gorm.io/plugin/dbresolver v1.5.3 // GORM插件，用于数据库读写分离（开源）
)

require (
	dario.cat/mergo v1.0.1 // indirect; 深度合并结构体的库（开源）
	filippo.io/edwards25519 v1.1.0 // indirect; 高性能Edwards25519签名算法实现（开源）
	github.com/aliyun/alibaba-cloud-sdk-go v1.63.88 // indirect; 阿里云SDK，提供对阿里云服务的支持（开源）
	github.com/buger/jsonparser v1.1.1 // indirect; 高效的JSON解析库（开源）
	github.com/bytedance/sonic v1.12.8 // indirect; 高效的JSON编解码库，比标准库更快（开源）
	github.com/bytedance/sonic/loader v0.2.3 // indirect; 相关的加载器，用于加载JSON（开源）
	github.com/cespare/xxhash/v2 v2.3.0 // indirect; 高效的哈希算法，主要用于快速计算（开源）
	github.com/cloudwego/base64x v0.1.5 // indirect; 扩展的base64编码/解码库（开源）
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect; Rendezvous hashing算法实现（开源）
	github.com/elastic/elastic-transport-go/v8 v8.6.1 // indirect; Elasticsearch传输层，支持请求和响应处理（开源）
	github.com/fsnotify/fsnotify v1.8.0 // indirect; 文件系统变化监听库（开源）
	github.com/gabriel-vasile/mimetype v1.4.8 // indirect; MIME类型检测库（开源）
	github.com/gin-contrib/sse v1.0.0 // indirect; 用于在Gin中发送Server-Sent Events（SSE）（开源）
	github.com/go-errors/errors v1.5.1 // indirect; 错误处理工具，类似于`pkg/errors`（开源）
	github.com/go-kratos/aegis v0.2.0 // indirect; Kratos微服务框架的健康检查库（开源）
	github.com/go-logr/logr v1.4.2 // indirect; 日志接口库，支持多种日志实现（开源）
	github.com/go-logr/stdr v1.2.2 // indirect; Go-logr的标准日志实现（开源）
	github.com/go-playground/form/v4 v4.2.1 // indirect; 表单解析库，用于解析表单数据（开源）
	github.com/go-playground/locales v0.14.1 // indirect; 本地化和翻译支持（开源）
	github.com/go-playground/universal-translator v0.18.1 // indirect; 多语言支持库（开源）
	github.com/go-playground/validator/v10 v10.24.0 // indirect; 数据验证库，用于验证结构体字段（开源）
	github.com/go-sql-driver/mysql v1.8.1 // indirect; MySQL的Go数据库驱动（开源）
	github.com/gobuffalo/envy v1.10.2 // indirect; 环境变量读取库（开源）
	github.com/gobuffalo/packd v1.0.2 // indirect; 静态文件打包工具（开源）
	github.com/gobuffalo/packr v1.30.1 // indirect; 静态文件打包工具（与packd类似）（开源）
	github.com/goccy/go-json v0.10.5 // indirect; 快速的JSON处理库（开源）
	github.com/google/uuid v1.6.0 // indirect; 生成UUID的库（开源）
	github.com/gorilla/mux v1.8.1 // indirect; 强大的HTTP路由库（开源）
	github.com/jackc/pgpassfile v1.0.0 // indirect; PostgreSQL密码文件解析库（开源）
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect; PostgreSQL服务文件解析库（开源）
	github.com/jackc/pgx/v5 v5.7.2 // indirect; PostgreSQL数据库驱动和客户端库（开源）
	github.com/jackc/puddle/v2 v2.2.2 // indirect; 连接池管理库（开源）
	github.com/jinzhu/inflection v1.0.0 // indirect; 用于处理单复数形式的库（开源）
	github.com/jinzhu/now v1.1.5 // indirect; 日期和时间操作库（开源）
	github.com/jmespath/go-jmespath v0.4.0 // indirect; JMESPath查询语言实现（开源）
	github.com/joho/godotenv v1.5.1 // indirect; 从 `.env` 文件加载环境变量（开源）
	github.com/jonboulle/clockwork v0.4.0 // indirect; 方便的时间操作工具（开源）
	github.com/josharian/intern v1.0.0 // indirect; 字符串池实现，优化内存（开源）
	github.com/json-iterator/go v1.1.12 // indirect; 高效的JSON编解码库（开源）
	github.com/klauspost/cpuid/v2 v2.2.9 // indirect; CPU信息查询工具（开源）
	github.com/leodido/go-urn v1.4.0 // indirect; URN生成工具（开源）
	github.com/lestrrat-go/strftime v1.1.0 // indirect; 格式化时间为字符串（开源）
	github.com/mattn/go-isatty v0.0.20 // indirect; 检查标准输出是否是TTY终端（开源）
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect; 并发操作工具库（开源）
	github.com/modern-go/reflect2 v1.0.2 // indirect; 高效的反射库（开源）
	github.com/pelletier/go-toml/v2 v2.2.3 // indirect; TOML文件解析库（开源）
	github.com/rogpeppe/go-internal v1.13.1 // indirect; Go的内部工具库（开源）
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect; Golang汇编优化工具（开源）
	github.com/ugorji/go/codec v1.2.12 // indirect; 高效的编解码库，支持JSON、BSON等格式（开源）
	go.opentelemetry.io/otel v1.34.0 // indirect; OpenTelemetry的Go实现，用于分布式追踪（开源）
	go.opentelemetry.io/otel/metric v1.34.0 // indirect; OpenTelemetry的Go实现，提供度量支持（开源）
	go.opentelemetry.io/otel/trace v1.34.0 // indirect; OpenTelemetry的Go实现，提供追踪支持（开源）
	go.uber.org/atomic v1.11.0 // indirect; 轻量级的原子操作库（开源）
	go.uber.org/multierr v1.11.0 // indirect; 处理多个错误的工具库（开源）
	go.uber.org/zap v1.27.0 // indirect; 高性能的日志库（开源）
	golang.org/x/arch v0.14.0 // indirect; Go架构支持（开源）
	golang.org/x/crypto v0.33.0 // indirect; Go加密库（开源）
	golang.org/x/mod v0.23.0 // indirect; Go模块工具库（开源）
	golang.org/x/net v0.35.0 // indirect; Go的网络库（开源）
	golang.org/x/sync v0.11.0 // indirect; Go的并发工具库（开源）
	golang.org/x/sys v0.30.0 // indirect; 系统相关的Go工具库（开源）
	golang.org/x/text v0.22.0 // indirect; 文字处理库（开源）
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250212204824-5a70512c5d8b // indirect; Google API的协议缓冲库（开源）
	google.golang.org/grpc v1.70.0 // indirect; gRPC框架，用于高效的服务间通信（开源）
	gopkg.in/ini.v1 v1.67.0 // indirect; INI文件解析库（开源）
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect; 日志轮换工具（开源）
	gopkg.in/yaml.v3 v3.0.1 // indirect; YAML格式解析库（开源）
)

require github.com/golang/protobuf v1.5.4

require (
	github.com/google/subcommands v1.2.0 // indirect
	github.com/opentracing/opentracing-go v1.2.1-0.20220228012449-10b1cf09e00b // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	golang.org/x/tools v0.22.0 // indirect
)
