# 設定 docker-compose 檔案的路徑
DOCKER_COMPOSE_FILE = docker-compose.yml

.PHONY: run stop logs clean ps

stop_nginx_rtmp:
	docker compose -f $(DOCKER_COMPOSE_FILE) stop nginx-rtmp-service

run_nginx_rtmp:
	docker compose -f $(DOCKER_COMPOSE_FILE) up -d --build nginx-rtmp-service



# 使用ffmpeg進行測試推流
testPush:
	ffmpeg -re -f lavfi -i testsrc=size=1280x720:rate=30 -vcodec libx264 -preset veryfast -f flv rtmp://localhost:1935/live/test