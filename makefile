.PHONY: types

types:
	# test if tygo is installed
	@if ! tygo --version > /dev/null 2>&1; then \
		go install github.com/gzuidhof/tygo@latest; \
	fi
	@rm -rf dist/*
	@tygo generate
	@echo "Types generated"
	@ls -la dist
