release:
	docker login -u $DOCKERHUB_USER -p $DOCKERHUB_PASS
	docker build -t j1fig/farol:$VERSION .

.PHONY: build
build:
	docker build .

.PHONY: clean
clean:
	rm -rf ./bin

.PHONY: deploy
deploy: clean build
	sls deploy --verbose
