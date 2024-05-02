.PHONY: start
start:
	cd app && go run main.go &

.PHONY: stop
stop:
	kill -9 `ps | grep main | grep -v grep | awk '{print $$1}'`

.PHONY: restart
restart: stop start

.PHONY: test
test:
	cd app && go test
