COVERAGE_DIR ?= .coverage

# cp from: https://github.com/yyle88/mutexmap/blob/842d6f3d77bba067fd85e355b5b4ab896f712070/Makefile#L4
test:
	@-rm -r $(COVERAGE_DIR)
	@mkdir $(COVERAGE_DIR)
	make test-with-flags TEST_FLAGS='-v -race -covermode atomic -coverprofile $$(COVERAGE_DIR)/combined.txt -bench=. -benchmem -timeout 20m'

test-with-flags:
	@go test $(TEST_FLAGS) ./...
