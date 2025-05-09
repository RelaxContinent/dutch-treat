# dutch-treat

go言語で自分ができることを記録する

## 割り勘APIを作成する

DDDで設計
gorm＋gRPC＋proto-gen-doc+go?
フロントエンドはTypeScript
バックエンドはGo
インフラはDocker
（kubeでオーケストレーションするほど他との連携はいらない）

## 要求定義

- 割り勘清算アプリ
- 誰が、何に、いくら払ったか をその時かぎりでいいので管理できること  
（将来的には残せるとうれしい。AIのチャットスレッドみたいに。）
- どっちがいくら払えば割り勘として清算できるか、計算できること

## 要件定義

- 可能ならweb公開したい
  - githubの備え付けだけでいきたい

- 対象が二人であること（以後ではA,Bとする）
- A,Bに自由に名前をつけられること
- 支払いを25件くらいは記録できること
- 支払い全体で清算できること
＜記録する支払いに関する要件＞
- 支払いの用途を自由に入力できること（フロント要件）
- 支払い金額を自由に数値で入力できること
- 支払った人をA,Bでえらべること

- 表示と計算は別の仕組みで行えること

## 基本設計

「清算」ボタン押下

- 入力内容をjsonにしてAPIに送る
- 計算結果を表示する

### フロント構想

割り勘する二人  
A B

No   |何に  |誰が   |いくら払った
-----|------|------|----------
No.1 |旅館  |A      |5000円
No.2 |朝ご飯|B      |2200円

`清算ボタン`（POST？）

`（A）`が`（B）`に`（何円）`払う

エラーだったら、形式に従ったエラーメッセージを表示する

### APIのIF  

仕様項目名 | 内容
----------|--------
プロトコル | HTTP2(gRPCに従う)  
メソッド | POST  
Content-Type | application/json  
文字タイプ | utf-8  
アクセスURL | /dutch-treat  

#### リクエストボディ

   項目名            | 必須 |  形式  |  内容  
--------------------|------|--------|--------
/payment[]          |  *1  | []     | 支払いの一覧
/payment[]/payer    |  ○   | string | 支払った人
/payment[]/amount   |  ○   | number | 金額

*1 : payment配列が空またはnullも許容する

別案（リクエストをきれいにする案？フロント側からどっちが作りやすいか次第。）

   項目名                    | 必須 |  形式  |  内容  
----------------------------|------|--------|--------
/payment[]                  |  *1  | []     | 支払いの一覧
/payment[]/payer            |  ○   | string | 支払った人
/payment[]/payer/[]amount   |  ○   | []number | 金額



#### レスポンスボディ

   項目名            | 必須 |  形式  |  内容
--------------------|------|--------|--------
/payment[]          |  -   | []     | 支払いの一覧
/payment[]/payer    |  ○   | string | 支払った人
/payment[]/amount   |  ○   | number | 金額

リクエストのpayment配列が空またはnullの場合は、amount=0

##### リクエストサンプル

```bash
curl.exe --header "Content-Type: application/json" --request POST localhost:8080/dutch-treat --data '{\"payment\": [{\"payer\": \"Alice\", \"amount\": 500},{\"payer\": \"Bob\", \"amount\": 1500}]}'
```

##### エラーレスポンス

ステータスコード | メッセージ | 内容  
---------------|------------|--------
200            | ok         | 成功
400            | bad_request | リクエストしたjsonの値が正しくない

## 詳細設計

application層  
  handler, request処理  
domain層  
  service, repository定義  
infrastructure層  
  repositoryImpl  
main  
