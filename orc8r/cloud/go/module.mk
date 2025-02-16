# Copyright (c) Facebook, Inc. and its affiliates.
# All rights reserved.
#
# This source code is licensed under the BSD-style license found in the
# LICENSE file in the root directory of this source tree.
#
SHELL := /bin/bash
.PHONY: build clean clean_gen download fmt gen lint plugin test tidy vet migration_plugin

build:: plugin
	go install ./...

clean::
	go clean ./...

clean_gen:
	find . -name "*.pb.go" | xargs --no-run-if-empty rm
	find . -name "*_swaggergen.go" | xargs --no-run-if-empty rm

download:
	go mod download

fmt:
	go fmt ./...

gen::
	go generate ./...

# The sed expression replaces '/' with '_' and gets rid of any './ in the path
copy_swagger_files:
	find . -name "swagger.yml" | xargs -I% --no-run-if-empty bash -c 'cp % $${SWAGGER_TEMP_GEN}/$$(echo % | sed "s#/#_#g; s/\._//g")'

lint:
	golint ./...

plugin::
	go build -buildmode=plugin -o $(PLUGIN_DIR)/$(PLUGIN_NAME).so .

test::
	go test ./...

tidy:
	go mod tidy

tools:: $(TOOL_DEPS)
$(TOOL_DEPS): %:
	go install $*

vet:
	go vet -composites=false ./...

COVER_FILE=$(COVER_DIR)/$(PLUGIN_NAME).gocov
cover:
	go test ./... -coverprofile $(COVER_FILE);
	# Don't measure coverage for protos and tools
	sed -i '/\.pb\.go/d; /.*\/tools\/.*/d' $(COVER_FILE);
	go tool cover -func=$(COVER_FILE)


# for configurator data migration
migration_plugin:
	if [[ -d ./tools/migrations/m003_configurator/plugin ]]; then \
		go build -buildmode=plugin -o $(PLUGIN_DIR)/migrations/m003_configurator/$(PLUGIN_NAME).so ./tools/migrations/m003_configurator/plugin; \
	fi
