## 😈:
#### docker build(within the "Dockerfile" directory)
  `docker build -t 'dcnlabweb' .`
#### docker push、build、pull
  * docker tag: `docker tag dcnlabweb ne6091043/dcnlabweb`
  * docker login: `docker login`
  * docker push: `docker push ne6091043/dcnlabweb`
  * docker pull: `docker pull ne6091043/dcnlabweb`
#### docker run
  * docker run -p 8888:8888 ne6091043/dcnlabweb  (cotainer port : host port)


##### [docker hub image](https://hub.docker.com/repository/docker/ne6091043/dcnlabweb)

---
## Reference: 🤝
  * [無痛使用 Golang 打造屬於自己的網頁](https://ithelp.ithome.com.tw/articles/10233981)
  * [Golang 使用 CSS 等靜態資源](https://michaelchen.tech/golang-web-programming/css/)
  * [使用 Docker 封裝與運行 Go 程式](https://ithelp.ithome.com.tw/articles/10240352)
  * [使用 Docker build 一個 Golang image](https://ithelp.ithome.com.tw/articles/10209304)
  * [把 Docker Image Push 到 Docker Hub](https://ithelp.ithome.com.tw/articles/10191139)
