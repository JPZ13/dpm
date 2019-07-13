.PHONY: clean
clean:
	@echo "====== Cleaning build directory ======"
	./makecmd clean

.PHONY: binaries
binaries:
	@echo "====== Making binaries ======"
	./makecmd make-mac-binary
	./makecmd make-linux-binary

.PHONY: mac-binary
mac-binary:
	@echo "====== Making Mac binary ======"
	./makecmd make-mac-binary

.PHONY: linux-binary
linux-binary:
	@echo "====== Making Linux binary ======"
	./makecmd make-linux-binary

.PHONY: test
test:
	@echo "====== Running project tests ======"
	./makecmd test

.PHONY: fmt
fmt:
	@echo "====== Formatting project code ======"
	./makecmd fmt
