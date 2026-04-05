CREATE DATABASE IF NOT EXISTS searchengine DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE searchengine;

CREATE TABLE IF NOT EXISTS documents (
    id         BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    title      VARCHAR(512)  NOT NULL        COMMENT '文档标题',
    content    TEXT          NOT NULL        COMMENT '文档正文',
    url        VARCHAR(1024) DEFAULT ''      COMMENT '来源链接',
    created_at DATETIME      DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME      DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_created (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文档表';

-- 示例数据
INSERT INTO documents (title, content, url) VALUES
('Golang Web 开发入门',
 'Go 语言是 Google 开发的开源编程语言，专为构建简单、可靠和高效的软件而设计。Go 在 Web 开发领域有广泛应用，Gin 是目前最流行的 Go Web 框架之一。本文介绍如何用 Golang 和 Gin 框架构建 RESTful API，包括路由设计、中间件使用和数据库集成。',
 'https://example.com/golang-web'),

('Vue3 前端框架详解',
 'Vue3 是一个用于构建用户界面的渐进式 JavaScript 框架。相比 Vue2，Vue3 引入了 Composition API、更好的 TypeScript 支持和更快的渲染性能。结合 Vite 构建工具，可以实现极速的热更新开发体验，大幅提升前端开发效率。',
 'https://example.com/vue3-guide'),

('MySQL 数据库优化技巧',
 'MySQL 是世界上最流行的关系型数据库管理系统之一。良好的索引设计可以将查询性能提升数十倍。本文深入讲解 MySQL 索引原理、B+Tree 结构、查询优化器工作方式，以及慢查询分析和分区表等高级特性。',
 'https://example.com/mysql-optimization'),

('搜索引擎原理与倒排索引',
 '搜索引擎的核心是倒排索引技术。倒排索引将文档中的每个词映射到包含该词的文档列表，实现了从关键词到文档的快速检索。TF-IDF 是经典的文档相关性评分算法，TF 表示词频，IDF 表示逆文档频率，两者相乘得到最终权重，广泛应用于信息检索领域。',
 'https://example.com/search-engine-intro'),

('Redis 缓存最佳实践',
 'Redis 是一个高性能的内存键值存储数据库，支持字符串、哈希、列表、集合等多种数据结构，常用于缓存、消息队列和实时排行榜等场景。在搜索系统中，Redis 可以缓存热门查询结果，显著降低后端计算压力，将响应时间从毫秒级降低到微秒级。',
 'https://example.com/redis-practice'),

('微服务架构设计与实践',
 '微服务架构将单一应用程序拆分为一套小服务，每个服务运行在自己的进程中，围绕业务能力进行组织。服务间通过轻量级通信机制进行通信，通常使用 HTTP REST API 或消息队列。Golang 因其高性能、低内存占用和强大的并发特性，成为微服务开发的热门选择。',
 'https://example.com/microservices'),

('Docker 容器化部署指南',
 'Docker 是目前最流行的容器化技术，通过将应用及其依赖打包成镜像，实现"一次构建，到处运行"。通过 Docker Compose，可以轻松编排多个服务。本文介绍如何将 Golang 后端和 Vue 前端应用容器化，并使用 Nginx 进行反向代理和负载均衡。',
 'https://example.com/docker-deploy'),

('机器学习入门指南',
 '机器学习是人工智能的一个子集，使计算机能够从数据中学习而不需要显式编程。常见算法包括线性回归、逻辑回归、决策树、随机森林和神经网络等。深度学习是机器学习的子领域，通过多层神经网络处理复杂任务。Python 和 TensorFlow 是该领域最流行的工具。',
 'https://example.com/ml-guide'),

('前端性能优化实战',
 '前端性能直接影响用户体验和转化率。优化手段包括：代码分割与懒加载、图片压缩与 WebP 格式、CDN 加速静态资源、合理设置浏览器缓存策略、减少首屏渲染阻塞等。Lighthouse 是 Google 提供的开源性能分析工具，可量化衡量各项指标。',
 'https://example.com/frontend-performance'),

('RESTful API 设计规范',
 'REST 是一种软件架构风格，用于设计网络应用程序的 API 接口。良好的 API 设计应遵循：使用名词表示资源、利用 HTTP 方法表达操作语义、返回合适的 HTTP 状态码、支持分页和过滤参数。版本控制和统一错误响应格式是 API 设计的重要考量。',
 'https://example.com/restful-api');
