Redis-Ping
==

## 疎通確認ツール
- 本番でredis-cliがインストールされていない場合に、Elasticacheの疎通確認するためのpingツール。

## 使い方

`redis-ping (ipアドレス)

```
% go build
% ./redis-ping 172.31.0.0
2020/05/29 12:18:52 PONG
```

---

## Linux用build

```
% GOOS=linux go build
```

できた`redis-ping` を `/tmp` 等にアップし、実行。
