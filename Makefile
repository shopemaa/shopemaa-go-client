.PHONY: install_plugin gen_client

all: install_plugin gen_client

install_plugin:
	go get -u github.com/Yamashou/gqlgenc
	go install github.com/Yamashou/gqlgenc

gen_client:
	../../../../bin/gqlgenc generate --configdir .
