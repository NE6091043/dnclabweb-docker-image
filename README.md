## : 
#### docker build(within the "Dockerfile" directory)
  `docker build -t 'dcnlabweb' .`
#### docker push、build、pull
  * docker tag: `docker tag dcnlabweb ne6091043/dcnlabweb`
  * docker login: `docker login`
  * docker push: `docker push ne6091043/dcnlabweb`
  * docker pull: `docker pull ne6091043/dcnlabweb`
#### docker run
  * docker run -p 8888:8888 ne6091043/dcnlabweb  (cotainer port : host post)
