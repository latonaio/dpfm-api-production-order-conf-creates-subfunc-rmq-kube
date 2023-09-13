# data-platform-api-production-order-conf-creates-subfunc
data-platform-api-production-order-conf-creates-subfunc は、データ連携基盤において、製造指図確認APIサービスのヘッダ登録/更新補助機能を担うマイクロサービスです。

## 動作環境
・ OS: LinuxOS  
・ CPU: ARM/AMD/Intel  

## 対象APIサービス
data-platform-api-production-order-conf-creates-subfunc の 対象APIサービスは次の通りです。

*  APIサービス URL: https://xxx.xxx.io/api/API_PROUCTION_ORDER_CONF_SRV/creates/

## 本レポジトリ が 対応する データ
data-platform-api-production-order-conf-creates-subfunc が対応する データ は、次のものです。

* ProductionOrderConfHeader（製造指図確認 - ヘッダデータ）

## Output
data-platform-api-production-order-conf-creates-subfunc では、[golang-logging-library-for-data-platform](https://github.com/latonaio/golang-logging-library-for-data-platform) により、Output として、RabbitMQ へのメッセージを JSON 形式で出力します。以下の項目のうち、"cursor" ～ "time"は、golang-logging-library-for-data-platform による 定型フォーマットの出力結果です。

```
```
