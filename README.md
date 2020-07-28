# Blog Server
Golang 寫的簡單 Web Server ，實作會員與文章相關的操作 RESTful API 。

架構理念採取了 Clean Architecture 與 Domain-Driven Design 的角度設計，讓框架等基礎設施層面的功能不影響業務邏輯。

# API
提供的所有 API 路徑與功能描述：

| Method | URL                  | Desc |
| :----: | :------------------: | :--: |
| POST   | /login               | 登入 |
| POST   | /logout              | 登出 |
| POST   | /accounts            | 註冊帳號 |
| GET    | /accounts/{username} | 取得指定 Username 帳號資料 |
| PUT    | /accounts/{username} | 更新指定 Username 帳號資料 |
| DELETE | /accounts/{username} | 移除指定 Username 帳號 |
| POST   | /posts/{username}    | 新增文章 |
| GET    | /posts/{username}    | 取得指定 Username 文章資料 |
| GET    | /posts               | 取得文章資料，可藉由 Filter 過濾結果 |