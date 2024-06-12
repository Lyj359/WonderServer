

build-ui:
	@cd ../ziranui && npm run build
	@cd ./statics && rm -rf '!(statics.go)'
	@cp -r ../ziranui/build/* ./statics


build-project-win:
	GOOS=windows GOARCH=amd64 go build -o Ziran.exe .