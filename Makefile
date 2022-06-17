GOCMD = go
BINARY_NAME = conv
APP_TARGET = bin/conv

PLUGIN_DIR = plugins
PLUGIN_NAMES = $(shell ls -1 $(PLUGIN_DIR))
PLUGIN_TARGETS = $(PLUGIN_NAMES:%=bin/plugins/%.so)

.PHONY = all clean debug
.DEFAULT = all

all: clean $(APP_TARGET) $(PLUGIN_TARGETS)

clean:
	rm -f $(APP_TARGET) 
	rm -f $(PLUGIN_TARGETS)

$(APP_TARGET):
	$(GOCMD) build -o $@ ./cmd/conv

$(PLUGIN_TARGETS):
	$(GOCMD) build -buildmode=plugin -o $@ ./plugins/$(@:bin/plugins/%.so=%)/$(@:bin/plugins/%.so=%).go
