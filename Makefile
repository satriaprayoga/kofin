BINARY_NAME=kofin.exe

build:
	@go mod vendor
	@go build -o tmp/${BINARY_NAME} cmd/kofin/main.go
	@echo Kofin built!


run: build
	@echo Staring Kofin...
	@.\tmp\${BINARY_NAME} &
	@echo Kofin started!

run-react: 
	@npm start --prefix kofin-web

clean:
	@echo Cleaning...
	@go clean
	@del .\tmp\${BINARY_NAME}
	@echo Cleaned!

start_compose:
	@docker-compose up -d

stop_compose:
	@docker-compose down

test:
	@echo Testing...
	@go test ./...
	@echo Done!


start: run

stop:
	@echo Ending services...
	@taskkill /IM ${BINARY_NAME} /F
	@echo Stopped Kofin

restart: stop start