CC=gcc
CFLAGS=-Wall -g -pthread
PROGRAM_NAME=raceCondition
SRC_FILES=*.c

build: $(SRC_FILES)
	$(CC) $(CFLAGS) -o $(PROGRAM_NAME) $(SRC_FILES)

clean:
	rm -f $(PROGRAM_NAME)

.PHONY: build clean