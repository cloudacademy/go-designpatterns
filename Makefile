all: singleton

# ======================
# creational

.PHONY: singleton prototype

singleton:
	cd ./creational/singleton && go run .

prototype:
	cd ./creational/prototype && go run .