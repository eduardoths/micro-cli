GO ?= go
GO_BUILD ?= $(GO) build
BINARY ?= microcli

#-----------------------------------------#
# Colors
#-----------------------------------------#
GREY := "\e[2;37m"
GREEN := "\e[0;32m"
BLACK := "\e[0m"
YELLOW :="\e[0;33m"
END_COLOR := "\e[0m"

VERSION := `git describe --tags`

#-----------------------------------------#
# Build
#-----------------------------------------#
.PHONY: _build/exports
_build/exports:
	@CGO_ENABLED=0 GO111MODULE=ON

.PHONY: build
build: _build/exports
	@printf $(GREY)"Starting build!"$(END_COLOR)"\n"
	@printf $(GREY)"build tag is "$(VERSION)$(END_COLOR)"\n"
	@$(GO_BUILD) \
		-installsuffix "static" \
		-ldflags "-X github.com/eduardoths/micro-cli/cmd.buildVersion=${VERSION} -s -w" \
		-o $(BINARY)
	@printf $(GREEN)"Finished build!"$(END_COLOR)"\n"

