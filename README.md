# ictsc2022

ictsc2022 問題環境再現用リポジトリ

## 再現できる問題

logicaが作問した問題が再現できます。

- ajl (protocが... 見つからない!?)
- dkv (答えてくれPingサーバー、ここにはuserモードとsystemdと、俺がいる！)

## 再現方法

### ajl (protocが... 見つからない!?)

1. CPU5コア、メモリ4GBのUbuntuが動いているVMを用意します。
    - 想定解のビルドが大変重いので本戦はこの構成になっていましたが、マシン性能に関しては自由で構いません。
1. userユーザーを作成します。
1. Dockerをインストールします。
    - <https://docs.docker.com/engine/install/ubuntu/> を参考に。
    - userユーザーを`docker` グループに追加し、sudo無しでDockerを利用できるようにするところまで行ってください。
1. userユーザーに切り替えます。
1. <https://github.com/logica0419/ictsc2022/tree/main/protobuf-alpine> にあるファイル・フォルダ類を、`/home/user` に追加します。
    - `/home/user` 直下に `app` ディレクトリ、`builder` ディレクトリ、`deploy` ファイルがある形になります。
1. `chmod +x /home/user/deploy` をターミナルで実行し、deployファイルを実行可能にします。
1. `/home/user/deploy` をターミナルで実行すると、初期状態が再現します。

### dkv (答えてくれPingサーバー、ここにはuserモードとsystemdと、俺がいる！)

1. CPU1コア、メモリ1GBのUbuntuが動いているVMを用意します。
1. Go 1.20.1をインストールします。
    - <https://go.dev/doc/install> を参考に。
1. userユーザーを作成します。
1. 管理者として `loginctl enable-linger user` を実行します。
    - [wiki](https://wiki.archlinux.jp/index.php/Systemd/%E3%83%A6%E3%83%BC%E3%82%B6%E3%83%BC#.E3.83.AD.E3.82.B0.E3.82.A2.E3.82.A6.E3.83.88.E6.99.82.E3.81.AB.E3.83.A6.E3.83.BC.E3.82.B6.E3.83.BC.E3.83.97.E3.83.AD.E3.82.BB.E3.82.B9.E3.82.92.E7.B5.82.E4.BA.86:~:text=/.local/data%0A...-,systemd%20%E3%81%AE%E3%83%A6%E3%83%BC%E3%82%B6%E3%83%BC%E3%82%A4%E3%83%B3%E3%82%B9%E3%82%BF%E3%83%B3%E3%82%B9%E3%82%92%E8%87%AA%E5%8B%95%E8%B5%B7%E5%8B%95,-systemd%20%E3%81%AE%E3%83%A6%E3%83%BC%E3%82%B6%E3%83%BC) にある通り、userモードのsystemdをログイン時ではなくマシンの起動時に起動するように設定するコマンドです。
1. userユーザーとしてログインし直します。
    - `su` コマンドなどでユーザーを変えても、userモードのsystemdの操作はできません。
1. <https://github.com/logica0419/ictsc2022/tree/main/user-systemd> にある以下の4ファイルを、`/home/user/webapp` ディレクトリを作ってそこに入れます。
    - deploy
    - go.mod
    - main.go
    - webapp.service
1. `chmod +x /home/user/webapp/deploy` をターミナルで実行し、deployファイルを実行可能にします。
1. `mkdir -p /home/user/.config/systemd/user && ln -s /home/user/webapp/webapp.service /home/user/.config/systemd/user/webapp.service` を実行してシンボリックリンクを作る
1. `/home/user/webapp/deploy` をターミナルで実行すると、初期状態が再現します。
