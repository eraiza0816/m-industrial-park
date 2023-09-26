# m-industrial-park
お仕事の出退勤記録をとります.  
`出退勤記録.csv`に記録を追記していくので,利用してください.  

## インストール方法
1. [最新の m-industrial-park_Windows_x86_64.zip ](https://github.com/eraiza0816/m-industrial-park/releases)を押して保存します
この時に警告など出た際は無視します。
2. Zipファイルを展開します

## 初期設定の仕方
imageフォルダの中にスクショがあるので参考にしてください.

1. イベントビューアーを開きます.

2. windowsログ システムを押下
    ![2](https://github.com/eraiza0816/m-industrial-park/blob/main/image/%E3%82%B9%E3%82%AF%E3%83%AA%E3%83%BC%E3%83%B3%E3%82%B7%E3%83%A7%E3%83%83%E3%83%88%202023-02-05%20181158.png)
3. 右側ペインにある「操作」から「現在のログをフィルター」を押下  
IDの項目に「7001-7002」を入力して「OK」を押下

4. フィルター結果のIDを見て以下を確認  
ID:7001 ログイン を表すID  
ID:7002 ログアウト を表すID  
    [4](https://github.com/eraiza0816/m-industrial-park/blob/main/image/%E3%82%B9%E3%82%AF%E3%83%AA%E3%83%BC%E3%83%B3%E3%82%B7%E3%83%A7%E3%83%83%E3%83%88%202023-02-05%20181922.png)
- ログインをしているものを選択して「このイベントにタスクを設定」を押下  
`m-industrial-park.exe` を設定する(引数無し，又は0の時は出勤)

- ログオフをしているものを選択して「このイベントにタスクを設定」を押下  
`m-industrial-park.exe` を設定する(引数が1の時は退勤)
