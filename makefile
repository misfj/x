BINARY_NAME := core

run:
	@echo "Building $(BINARY_NAME)..."
	go build -o $(BINARY_NAME).exe core.go
	@echo "Running $(BINARY_NAME).exe..."
	./$(BINARY_NAME).exe

# run: build
# 	@echo "Running $(BINARY_NAME).exe..."
# 	./$(BINARY_NAME).exe

clean:
	rm -f $(BINARY_NAME).exe 
	rm -f $(BINARY_NAME).exe~
	rm -f  public.pub
	rm  -f private.key
	rm .idea/ -rf
	rm  -f x.json
	rm  -f x1.json

    
all: build run

.PHONY: build run clean all