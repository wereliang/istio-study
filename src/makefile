all:
	docker build --target basicinfo -t "wereliang/basicinfo:v1" -f Dockerfile .
	docker build --target query -t "wereliang/query:v1" -f Dockerfile .
	docker build --target recommend_v1 -t "wereliang/recommend:v1" -f Dockerfile .
	docker build --target recommend_v2 -t "wereliang/recommend:v2" -f Dockerfile .
	docker build --target star -t "wereliang/star:v1" -f Dockerfile .