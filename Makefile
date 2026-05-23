GO_FILE     = main.go
BUILDDIR    = build
EXEC        = gota
PREFIX     ?= $(HOME)/.local
INSTALL_DIR = $(PREFIX)/bin

.PHONY: install uninstall build clean

build:
	mkdir -p $(BUILDDIR)
	go build -o $(BUILDDIR)/$(EXEC) $(GO_FILE)

clean:
	rm $(BUILDDIR)/$(EXEC)

