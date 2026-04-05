// Mock 文档数据库 —— 覆盖多个技术主题，确保搜索有丰富结果
export const mockDocuments = [
  {
    id: 1,
    title: 'Golang Web 开发入门指南',
    content:
      'Go 语言是 Google 开发的开源编程语言，专为构建简单、可靠和高效的软件而设计。Gin 是目前最流行的 Go Web 框架之一，性能极高，路由速度接近原生 net/http。本文从零开始讲解如何用 Golang 和 Gin 框架构建 RESTful API，包括路由分组、参数绑定、中间件开发和统一响应封装。',
    url: 'https://golang.org/doc/tutorial/web-service-gin',
  },
  {
    id: 2,
    title: 'Vue3 Composition API 深度解析',
    content:
      'Vue3 是一个用于构建用户界面的渐进式 JavaScript 框架。Composition API 是 Vue3 最重要的新特性，通过 setup() 函数将响应式逻辑组合在一起，比 Options API 更灵活、更易测试。ref、reactive、computed、watch 是最常用的 Composition API。结合 Vite 构建工具，热更新速度极快。',
    url: 'https://vuejs.org/guide/extras/composition-api-faq',
  },
  {
    id: 3,
    title: 'MySQL 索引原理与查询优化',
    content:
      'MySQL 使用 B+ Tree 作为 InnoDB 引擎的默认索引结构。理解索引的最左前缀原则是写出高效 SQL 的关键。EXPLAIN 命令可以分析查询执行计划，观察 type 字段从 ALL→range→ref→eq_ref→const 代表性能逐步提升。覆盖索引可以避免回表查询，大幅降低 I/O 次数。分库分表是应对大数据量的常见方案。',
    url: 'https://dev.mysql.com/doc/refman/8.0/en/optimization-indexes.html',
  },
  {
    id: 4,
    title: '搜索引擎原理：倒排索引与 TF-IDF',
    content:
      '搜索引擎的核心是倒排索引（Inverted Index）技术。倒排索引将文档中的每个词映射到包含该词的文档 ID 列表，实现从关键词到文档的快速检索。TF-IDF（词频-逆文档频率）是经典的相关性评分算法：TF 表示词在文档中的出现频率，IDF 表示该词在所有文档中的稀有程度，TF×IDF 得分越高代表文档与查询越相关。现代搜索引擎还会结合 BM25、PageRank 等算法。',
    url: 'https://en.wikipedia.org/wiki/Inverted_index',
  },
  {
    id: 5,
    title: 'Redis 缓存架构与最佳实践',
    content:
      'Redis 是一个高性能的内存键值数据库，支持 String、Hash、List、Set、ZSet 五种基本数据结构。在搜索系统中，Redis 可缓存热门查询结果，将 QPS 从数百提升到数万。缓存穿透（查不存在的 key）可用布隆过滤器解决；缓存击穿（热点 key 过期）可用互斥锁或永不过期策略；缓存雪崩（大量 key 同时过期）可用随机 TTL 分散。',
    url: 'https://redis.io/docs/manual/patterns/',
  },
  {
    id: 6,
    title: '微服务架构设计与 Go 实践',
    content:
      '微服务架构将单体应用拆分为多个独立服务，每个服务围绕单一业务能力设计。Go 语言凭借其轻量级 goroutine、低内存占用和快速编译，成为微服务的首选语言。服务发现可用 Consul 或 Kubernetes DNS，链路追踪用 Jaeger，熔断限流用 Sentinel。gRPC 基于 HTTP/2 和 Protobuf，性能比 REST JSON 高 5-10 倍，适合内部服务通信。',
    url: 'https://microservices.io/patterns/index.html',
  },
  {
    id: 7,
    title: 'Docker 与 Kubernetes 容器化实战',
    content:
      'Docker 通过容器技术实现应用与环境的隔离，解决了"在我机器上能跑"的经典问题。Dockerfile 定义镜像构建步骤，docker-compose.yml 编排多服务。Kubernetes（K8s）是容器编排的事实标准，提供服务发现、负载均衡、自动扩缩容和滚动升级能力。Pod 是 K8s 最小调度单元，Deployment 管理无状态应用，StatefulSet 管理有状态服务。',
    url: 'https://kubernetes.io/docs/concepts/overview/',
  },
  {
    id: 8,
    title: '深度学习与神经网络基础',
    content:
      '深度学习是机器学习的子领域，通过多层神经网络自动提取数据特征。卷积神经网络（CNN）擅长图像识别，循环神经网络（RNN）和 Transformer 擅长自然语言处理。PyTorch 和 TensorFlow 是最流行的深度学习框架。反向传播算法通过梯度下降更新权重，Adam 优化器是目前最常用的优化算法。大语言模型（LLM）如 GPT、BERT 基于 Transformer 架构。',
    url: 'https://www.deeplearning.ai/',
  },
  {
    id: 9,
    title: '前端性能优化 20 个核心技巧',
    content:
      '前端性能优化需要从加载性能和运行时性能两个维度入手。加载优化：代码分割（Code Splitting）、懒加载（Lazy Load）、Tree Shaking、图片 WebP 格式、HTTP/2 多路复用、CDN 静态资源加速、Service Worker 离线缓存。运行时优化：虚拟列表渲染海量数据、防抖节流事件处理、Web Worker 计算密集任务、requestAnimationFrame 动画、避免强制回流（Layout Thrashing）。',
    url: 'https://web.dev/performance/',
  },
  {
    id: 10,
    title: 'RESTful API 设计规范与最佳实践',
    content:
      'REST（Representational State Transfer）是一种 API 设计风格。核心原则：资源用名词而非动词（/users 而非 /getUsers）、使用 HTTP 方法语义（GET 查询、POST 创建、PUT 替换、PATCH 部分更新、DELETE 删除）、用 HTTP 状态码表达结果（200/201/204/400/401/403/404/500）。版本管理推荐 URL 路径方式（/api/v1/）。分页用 page+size 参数，过滤和排序作为 query string。',
    url: 'https://restfulapi.net/',
  },
  {
    id: 11,
    title: 'TypeScript 高级类型系统详解',
    content:
      'TypeScript 为 JavaScript 添加了静态类型检查，大幅提升大型项目的可维护性。高级类型包括：泛型（Generics）实现类型复用、联合类型（Union Types）和交叉类型（Intersection Types）、条件类型（Conditional Types）、映射类型（Mapped Types）、模板字面量类型。Utility Types 如 Partial、Required、Pick、Omit、Record 是日常开发必备工具。',
    url: 'https://www.typescriptlang.org/docs/handbook/2/types-from-types.html',
  },
  {
    id: 12,
    title: 'Nginx 反向代理与高性能配置',
    content:
      'Nginx 是世界上最流行的 Web 服务器和反向代理，以极低的内存占用处理高并发连接著称。核心配置包括：upstream 负载均衡（轮询/ip_hash/least_conn）、proxy_pass 反向代理后端服务、gzip 压缩响应、静态文件缓存（expires 指令）、SSL/TLS 终止。location 块匹配优先级：精确匹配（=）> 前缀匹配（^~）> 正则匹配（~/~*）> 普通前缀匹配。',
    url: 'https://nginx.org/en/docs/',
  },
  {
    id: 13,
    title: 'Git 工作流与团队协作规范',
    content:
      'Git Flow 是一种经典的分支管理策略，包含 main、develop、feature、release、hotfix 五种分支。GitHub Flow 更简洁，只有 main 和 feature 分支，通过 Pull Request 进行代码审查。Conventional Commits 规范提交信息格式：feat（新功能）、fix（修复）、docs（文档）、refactor（重构）、test（测试）。git rebase 让提交历史更线性，git cherry-pick 精准地将特定提交复制到其他分支。',
    url: 'https://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow',
  },
  {
    id: 14,
    title: 'WebSocket 实时通信原理与实践',
    content:
      'WebSocket 是一种在单个 TCP 连接上进行全双工通信的协议，适合实时聊天、协同编辑、在线游戏等场景。与 HTTP 轮询相比，WebSocket 延迟更低、服务器压力更小。Go 语言使用 gorilla/websocket 库处理 WebSocket 连接，每个连接分配独立 goroutine 读写。消息广播可用 channel + Hub 模式实现。生产环境需处理心跳保活、断线重连和消息队列堆积问题。',
    url: 'https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API',
  },
  {
    id: 15,
    title: 'PostgreSQL vs MySQL 深度对比',
    content:
      'PostgreSQL 和 MySQL 是最流行的两款开源关系型数据库。PostgreSQL 支持更丰富的数据类型（数组、JSON、地理信息），事务隔离级别更完整，支持窗口函数、递归 CTE、物化视图等高级 SQL 特性，更适合复杂分析查询。MySQL（InnoDB 引擎）读写性能在简单 OLTP 场景更优，生态更成熟，运维工具更丰富。选型建议：需要复杂查询选 PostgreSQL，简单高并发读写选 MySQL。',
    url: 'https://www.postgresql.org/about/featurematrix/',
  },
  {
    id: 16,
    title: 'Vite 构建工具原理与配置详解',
    content:
      'Vite 是新一代前端构建工具，开发模式利用浏览器原生 ESM 实现按需编译，冷启动比 Webpack 快 100 倍以上。生产构建基于 Rollup，输出体积更小。核心概念：插件（Plugin）扩展功能，环境变量通过 import.meta.env 访问，代理（proxy）解决开发跨域。vite.config.js 支持 TypeScript，alias 配置路径别名，optimizeDeps 预构建依赖。热更新（HMR）精准到模块级别，Vue/React 组件状态保持不丢失。',
    url: 'https://vitejs.dev/guide/',
  },
  {
    id: 17,
    title: 'Linux 系统调优与性能排查',
    content:
      'Linux 性能分析工具链：top/htop 查看 CPU 和内存使用，vmstat 监控虚拟内存和 CPU 调度，iostat 分析磁盘 I/O，netstat/ss 查看网络连接，strace 追踪系统调用，perf 做 CPU 性能剖析。高并发服务器调优：增大文件描述符限制（ulimit -n），优化 TCP 参数（tcp_fin_timeout、tcp_max_syn_backlog、somaxconn），使用 epoll 替代 select/poll。',
    url: 'https://www.brendangregg.com/linuxperf.html',
  },
  {
    id: 18,
    title: '前端安全：XSS、CSRF 与 CSP',
    content:
      'XSS（跨站脚本攻击）分为存储型、反射型和 DOM 型，防御措施：对用户输入进行 HTML 实体编码，使用 CSP（Content Security Policy）限制脚本来源，设置 HttpOnly Cookie 阻止 JS 读取。CSRF（跨站请求伪造）防御：验证 Origin/Referer 请求头，使用 CSRF Token，SameSite Cookie 属性。SQL 注入防御：使用预编译语句（Prepared Statement），绝不拼接 SQL 字符串。',
    url: 'https://owasp.org/www-project-top-ten/',
  },
  {
    id: 19,
    title: 'gRPC 与 Protocol Buffers 实战',
    content:
      'gRPC 是 Google 开源的高性能 RPC 框架，基于 HTTP/2 和 Protocol Buffers。相比 REST JSON，gRPC 序列化效率高 3-10 倍，支持双向流式传输（Streaming），天然支持多语言（Go、Java、Python、Node.js 等）。定义 .proto 文件描述服务接口，protoc 编译器生成客户端和服务端代码。拦截器（Interceptor）相当于 HTTP 中间件，可用于日志、认证、链路追踪。',
    url: 'https://grpc.io/docs/what-is-grpc/introduction/',
  },
  {
    id: 20,
    title: 'Elasticsearch 全文检索原理',
    content:
      'Elasticsearch 是分布式搜索和分析引擎，基于 Apache Lucene 构建。核心概念：Index（索引）≈ 数据库，Document（文档）≈ 行，Field（字段）≈ 列。倒排索引是 ES 快速全文检索的基础。分词器（Analyzer）将文本切分为 Token，中文常用 IK Analyzer。Query DSL 提供丰富查询：match、term、range、bool（must/should/must_not）。聚合（Aggregation）类似 SQL GROUP BY，支持桶聚合和度量聚合。',
    url: 'https://www.elastic.co/guide/en/elasticsearch/reference/current/index.html',
  },
]
