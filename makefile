BINARY_NAME := main
TARGET_FILE := config.yaml

.PHONY: all build run clean conf

all: clean  conf build run

build:
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME).exe main.go || exit 1

run: build
	@echo "Running $(BINARY_NAME).exe..."
	@./$(BINARY_NAME).exe || exit 1

clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME).exe 
	rm -f public.pub 
	rm -f private.key 
	rm -f  config.yaml
	rm -rf .idea/
	rm  -rf water.log

conf:
	@echo "Generating $(TARGET_FILE)..."
	@rm -f $(TARGET_FILE)
	@echo "db:" >> $(TARGET_FILE)
	@echo "  host: \"127.0.0.1\"" >> $(TARGET_FILE)
	@echo "  name: \"te\"" >> $(TARGET_FILE)
	@echo "  port: 3306" >> $(TARGET_FILE)
	@echo "  user: \"root\"" >> $(TARGET_FILE)
	@echo "  password: \"root\"" >> $(TARGET_FILE)
	@echo "  maximum-pool-size: 20" >> $(TARGET_FILE)
	@echo "  maximum-idle-size: 5" >> $(TARGET_FILE)
	@echo "jwt:" >> $(TARGET_FILE)
	@echo "  exp: 90" >> $(TARGET_FILE)
	@echo "  iss: \"cored\"" >> $(TARGET_FILE)
	@echo "  sub: \"cored\"" >> $(TARGET_FILE)
	@echo "  secret: \"cored\"" >> $(TARGET_FILE)
	@echo "cache:" >> $(TARGET_FILE)
	@echo "  expire: 3" >> $(TARGET_FILE)
	@echo "  use_ssl: false" >> $(TARGET_FILE)
	@echo "  access_id: \"minio_admin\"" >> $(TARGET_FILE)
	@echo "  end_point: \"10.10.1.40:9000\"" >> $(TARGET_FILE)
	@echo "  secret_accessKey: \"9ijnBHU*@123\"" >> $(TARGET_FILE)
	@echo "store:" >> $(TARGET_FILE)
	@echo "  expire: 3" >> $(TARGET_FILE)
	@echo "  use_ssl: false" >> $(TARGET_FILE)
	@echo "  access_id: \"minio_admin\"" >> $(TARGET_FILE)
	@echo "  end_point: \"10.10.1.40:9000\"" >> $(TARGET_FILE)
	@echo "  default_size: 500" >> $(TARGET_FILE)
	@echo "  secret_accessKey: \"9ijnBHU*@123\"" >> $(TARGET_FILE)
	@echo "backup:" >> $(TARGET_FILE)
	@echo "  addr: \"http://10.10.254.13:5001\"" >> $(TARGET_FILE)
	@echo "logging:" >> $(TARGET_FILE)
	@echo "  file: \"logs/cored.log\"" >> $(TARGET_FILE)
	@echo "  level: \"debug\"" >> $(TARGET_FILE)
	@echo "  max_age: 30" >> $(TARGET_FILE)
	@echo "  max_size: 10" >> $(TARGET_FILE)
	@echo "  max_backups: 5" >> $(TARGET_FILE)
	@echo "ws:" >> $(TARGET_FILE)
	@echo "  listen_address: " >> $(TARGET_FILE)
	@echo "dr:" >> $(TARGET_FILE)
	@echo "  del-data: \"delData\"" >> $(TARGET_FILE)
	@echo "  encrypt: \"encrypt\"" >> $(TARGET_FILE)
	@echo "  endpoint: \"https://sqwalletapi.smartdatachain.com\"" >> $(TARGET_FILE)
	@echo "  upload-data: \"uploadData\"" >> $(TARGET_FILE)
	@echo "  dr-user-create: \"drUserCreate\"" >> $(TARGET_FILE)
	@echo "  dr-user-modify: \"drUserModify\"" >> $(TARGET_FILE)
	@echo "  dr-account-info: \"drAccountInfo\"" >> $(TARGET_FILE)
	@echo "  dr-account-sync: \"drAccountSync\"" >> $(TARGET_FILE)
	@echo "  get-dr-user-info: \"getDrUserInfo\"" >> $(TARGET_FILE)
	@echo "  authentication: \"authentication\"" >> $(TARGET_FILE)
	@echo "  select-upload-data: \"selectUploadData\"" >> $(TARGET_FILE)
	@echo "api:" >> $(TARGET_FILE)
	@echo "  key-file: \"x\"" >> $(TARGET_FILE)
	@echo "  cert_file: ''" >> $(TARGET_FILE)
	@echo "  enable-tls: false" >> $(TARGET_FILE)
	@echo "  listen-address: \"0.0.0.0:8080\"" >> $(TARGET_FILE)
	@echo "water:" >> $(TARGET_FILE)
	@echo "  grid: \"security/grid\"" >> $(TARGET_FILE)
	@echo "  first: \"security/l1\"" >> $(TARGET_FILE)
	@echo "  plain: \"security/plain\"" >> $(TARGET_FILE)
	@echo "  third: \"security/l3\"" >> $(TARGET_FILE)
	@echo "  protocol: \"http\"" >> $(TARGET_FILE)
	@echo "  end_point: \"123.60.56.112:9001\"" >> $(TARGET_FILE)
	@echo "  generate-stamps: \"stamp/circle/create\"" >> $(TARGET_FILE)
	@echo "  stamp-circle-extract: \"stamp/circle/extract\"" >> $(TARGET_FILE)
	@echo "  stamp-elipse-extract: \"stamp/elipse/extract\"" >> $(TARGET_FILE)
	@echo "  generate-stamp-elipse: \"stamp/elipse/create\"" >> $(TARGET_FILE)
	@echo "  generate-bs-water-mark: \"screenBS\"" >> $(TARGET_FILE)
	@echo "  weak-water-mark-insert: \"weak\"" >> $(TARGET_FILE)
	@echo "  weak-water-mark-extract: \"weak/extract\"" >> $(TARGET_FILE)
	@echo "  electronic-watermark-bright: \"flow/plain\"" >> $(TARGET_FILE)
	@echo "  electronic-watermark-underflow: \"flow\"" >> $(TARGET_FILE)
	@echo "  electronic-watermark-underflow-extract: \"flow/extract\"" >> $(TARGET_FILE)
	@echo "manage:" >> $(TARGET_FILE)
	@echo "  listen-addr: \"0.0.0.0\"" >> $(TARGET_FILE)
	@echo "  listen-port: 6000" >> $(TARGET_FILE)
	@echo "node:" >> $(TARGET_FILE)
	@echo "  node-number: \"xxx\"" >> $(TARGET_FILE)
	@echo "  node-unit: \"xxx\"" >> $(TARGET_FILE)
	@echo "  node-deploy: \"xxx\"" >> $(TARGET_FILE)
	@echo "  node-oper: \"xxx\"" >> $(TARGET_FILE)
	@echo "  node-login-code: \"xxx\"" >> $(TARGET_FILE)
	@echo "  node-public-md5: \"\"" >> $(TARGET_FILE)
	@echo "  node-private-md5: \"\"" >> $(TARGET_FILE)
	@echo "  node-algo: \"rsa\"" >> $(TARGET_FILE)
	@echo "  node-store-key: \"1111111111111111\"" >> $(TARGET_FILE)
	@echo "  node-platform: \"ws://0.0.0.0:8083/api/v1/1\"" >> $(TARGET_FILE)
	@echo "waterx:" >> $(TARGET_FILE)
	@echo "  name: \"watermark_service.exe\"" >> $(TARGET_FILE)
	@echo "  args: \":9090 ./exe/config.json\"" >> $(TARGET_FILE)
	@echo "$(TARGET_FILE) generated successfully."
