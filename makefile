

build-ui:
	@cd ../ziranui && npm run build
	@cd ./statics && rm -rf '!(statics.go)'
	@cp -r ../ziranui/build/* ./statics
