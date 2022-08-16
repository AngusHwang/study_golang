gen:
	go run -mod=mod ./ent/entc.go
	go run -mod=mod github.com/99designs/gqlgen

docker-build:
	aws ecr get-login-password --region ap-northeast-2 | docker login --username AWS --password-stdin 790152711045.dkr.ecr.ap-northeast-2.amazonaws.com
	docker build --platform linux/amd64 -t odin-api .
	docker tag odin-api:latest 790152711045.dkr.ecr.ap-northeast-2.amazonaws.com/odin-api:latest
	docker push 790152711045.dkr.ecr.ap-northeast-2.amazonaws.com/odin-api:latest