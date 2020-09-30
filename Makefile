
.PHONY: fmt
fmt:
	@echo "Formatting files..."
	@gofmt -w -l -s .
	@echo "Done formatting files."
