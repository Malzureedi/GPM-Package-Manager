# Define executable names
GPM_BINARY=gpm
AUTOBUILD_BINARY=autobuild/autobuild
LOG_FILE=build.log

# Default target (build everything)
all: clean deps build install

# Ensure dependencies are set up
deps:
	@echo "[1] Checking dependencies..." | tee -a $(LOG_FILE)
	go mod tidy | tee -a $(LOG_FILE)

# Build GPM
build:
	@echo "[2] Building GPM..." | tee -a $(LOG_FILE)
	go build -o $(GPM_BINARY) | tee -a $(LOG_FILE)
	@echo "[3] Building Autobuild..." | tee -a $(LOG_FILE)
	go build -o $(AUTOBUILD_BINARY) autobuild/autobuild.go | tee -a $(LOG_FILE)

# Install binaries to system paths
install:
	@echo "[4] Installing GPM..." | tee -a $(LOG_FILE)
	mv $(GPM_BINARY) /usr/local/bin/ | tee -a $(LOG_FILE)
	mv $(AUTOBUILD_BINARY) /usr/local/bin/ | tee -a $(LOG_FILE)

# Clean compiled files
clean:
	@echo "[CLEAN] Cleaning old builds..." | tee -a $(LOG_FILE)
	rm -rf $(GPM_BINARY) $(AUTOBUILD_BINARY) $(LOG_FILE)

# Run GPM
run:
	@echo "[RUN-GPM]Running GPM..." | tee -a $(LOG_FILE)
	./$(GPM_BINARY) | tee -a $(LOG_FILE)
