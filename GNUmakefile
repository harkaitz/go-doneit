.POSIX:
.SUFFIXES:
.PHONY: all clean install check

PROJECT   =doneit
VERSION   =1.0.0
PREFIX    =/usr/local
BUILDDIR ?=.build
EXE      ?=$(shell uname -s | awk '/Windows/ || /MSYS/ || /CYG/ { print ".exe" }')

all:
clean:
install:
check:
## -- BLOCK:go --
.PHONY: all-go install-go clean-go $(BUILDDIR)/run-only-once$(EXE)
all: all-go
install: install-go
clean: clean-go
all-go: $(BUILDDIR)/run-only-once$(EXE)
install-go:
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp  $(BUILDDIR)/run-only-once$(EXE) $(DESTDIR)$(PREFIX)/bin
clean-go:
	rm -f $(BUILDDIR)/run-only-once$(EXE)
##
$(BUILDDIR)/run-only-once$(EXE): $(GO_DEPS)
	mkdir -p $(BUILDDIR)
	go build -o $@ $(GO_CONF) ./cmd/run-only-once
## -- BLOCK:go --
