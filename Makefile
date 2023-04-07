APP_NAME=crocmenu
BUILD_DIR=build

.PHONY: clean

all: $(APP_NAME)

$(APP_NAME): $(APP_NAME)_mac $(APP_NAME)_mac_arm64
	lipo -create -output $(BUILD_DIR)/$(APP_NAME)_mac_universal $(BUILD_DIR)/$(APP_NAME)_mac $(BUILD_DIR)/$(APP_NAME)_mac_arm64

$(APP_NAME)_mac:
	env GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)_mac

$(APP_NAME)_mac_arm64:
	env GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DIR)/$(APP_NAME)_mac_arm64

clean:
	rm -rf $(BUILD_DIR)


