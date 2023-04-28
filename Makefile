all: compile package

compile:
	@echo "Compile"
	go build -o gosfs

package:
	@echo "Build docker image snoxe/gosfs:latest"
	docker build -t snoxe/gosfs:latest .

push:
	@echo "Push docker image to docker.io/snoxe"
	docker push snoxe/gosfs