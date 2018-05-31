# AwsBillingNotifyGo
AWS料金を取得しLINE通知を行うバッチ

## 事前設定
AWS-KEYを作成し、下記コマンドを実行してください。
```
export AWS_ACCESS_KEY=hoge
export AWS_SECRET_ACCESS_KEY=hoge
```

下記URLからLINE Notify APIのTOKENを発行してください。
```
https://notify-bot.line.me/my/
```

その後、下記コマンドを実行してください。
```
export LINE_NOTIFY_API_TOKEN=hoge
```
