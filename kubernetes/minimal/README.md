# Hello
公式 tutorialはやっている前提で    
https://kubernetes.io/ja/docs/tutorials/hello-minikube/

## Quick Start

## 構築手順
1. fileたちを作成
2. docker build用のterminal設定
以下のコマンドを打つと、minikubeを使うならevalしてね。というメッセージが出る。
``` 
$ minikube docker-env
# To point your shell to minikube's docker-daemon, run: # eval $(minikube -p minikube docker-env)
``` 
設定
```
$ eval $(minikube -p minikube docker-env)
```

3. docker build
```
$ skaffold dev --port-forward
```

4. 確認
```
$ minikube dashboard
$ curl http://localhost:8001
```

5. 削除
```
$ minikube stop
```
Dockerの設定を元に戻す
```
$ eval $(minikube docker-env -u)
```

## 参考
- https://takaya030.hatenablog.com/entry/2023/04/15/231222
- https://developer.mamezou-tech.com/containers/k8s/tutorial/app/minikube/

