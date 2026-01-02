# Bluebell - 高性能社区论坛后端 API 🐹

![Gin](https://img.shields.io/badge/Gin-Web_Framework-ff69b4?style=flat&logo=go)
![GORM](https://img.shields.io/badge/GORM-MySQL_ORM-blue?style=flat)
![License](https://img.shields.io/badge/License-MIT-green.svg)

**Bluebell** 是一个基于 Go 语言（Golang）构建的轻量级社区论坛后端服务。项目采用了现代化的分层架构（Controller-Logic-DAO），实现了用户注册登录、社区管理、帖子发布与分页查询等核心功能。

本项目旨在展示标准的企业级 Go Web 项目结构与开发流程。

## 🛠 技术栈 (Tech Stack)

*   **编程语言**: Go (Golang)
*   **Web 框架**: [Gin](https://github.com/gin-gonic/gin) - 高性能 HTTP Web 框架
*   **数据库**: MySQL 8.0
*   **ORM**: [GORM](https://gorm.io/) - 强大的数据对象映射库
*   **鉴权**: JWT (JSON Web Token) - 无状态认证
*   **配置管理**: [Viper](https://github.com/spf13/viper) - 支持多格式配置文件
*   **日志**: [Zap](https://github.com/uber-go/zap) + Lumberjack (日志切割)
*   **工具**:
    *   **Snowflake**: 推特雪花算法生成分布式唯一 ID
    *   **Bcrypt**: 密码加盐哈希存储

## ✨ 功能特性 (Features)

*   **用户系统**:
    *   用户注册（密码加密存储）
    *   用户登录（基于 JWT 颁发 Token）
    *   Token 认证中间件
*   **社区模块**:
    *   社区列表查询
    *   社区详情查询
*   **帖子模块**:
    *   发布帖子（使用 Snowflake 生成 ID）
    *   帖子详情（关联查询作者与社区信息）
    *   帖子列表（支持分页查询、按时间倒序）
*   **工程化**:
    *   规范的项目目录结构 (CLD 分层)
    *   统一的错误处理与响应封装
    *   结构化日志记录

## 📂 目录结构 (Project Layout)

```text
bluebell/
├── conf/          # 配置文件目录
├── controllers/   # 控制层：处理请求参数与响应
├── dao/           # 数据访问层：MySQL 数据库操作
├── logic/         # 业务逻辑层：处理核心业务
├── middlewares/   # 中间件：JWT 认证、日志等
├── models/        # 数据模型：定义结构体与数据库表
├── pkg/           # 公共工具包 (JWT, Snowflake, Encrypt)
├── routes/        # 路由层：注册路由
├── logger/        # 日志初始化
└── main.go        # 项目入口
```

## 🚀 快速开始 (Getting Started)

### 1. 环境要求
*   Go 1.18+
*   MySQL 8.0+

### 2. 克隆项目
```bash
git clone https://github.com/linshule/bluebell.git
cd bluebell
```

### 3. 配置数据库
1.  在 MySQL 中创建数据库：
    ```sql
    CREATE DATABASE bluebell DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
    ```
2.  修改 `conf/config.yaml` 文件，填入你的数据库账号密码：
    ```yaml
    mysql:
      host: "127.0.0.1"
      port: 3306
      user: "root"
      password: "your_password"
      dbname: "bluebell"
    ```

### 4. 运行项目
```bash
# 下载依赖
go mod tidy

# 运行
go run main.go
```
启动成功后，服务默认运行在 `http://localhost:8080`。

## 📝 API 接口示例

### 1. 用户注册
*   **URL**: `POST /signup`
*   **Body**:
    ```json
    {
        "username": "test_user",
        "password": "password123",
        "re_password": "password123"
    }
    ```

### 2. 用户登录
*   **URL**: `POST /login`
*   **Response**: 返回 JWT Token

### 3. 发布帖子 (需要 Token)
*   **URL**: `POST /api/post`
*   **Header**: `Authorization: Bearer <你的Token>`
*   **Body**:
    ```json
    {
        "community_id": 1,
        "title": "Hello Bluebell",
        "content": "This is my first post."
    }
    ```

---
**Author**: linshule