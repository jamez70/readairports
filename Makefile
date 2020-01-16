.PHONY: install clean fmt
SRC := cgi.go
TARGET := cgi
INST_TARGET := /var/www/cgi-bin/testcgi

all:	$(TARGET)

$(TARGET): $(SRC)
	go build $(SRC)

clean:
	rm -f $(TARGET) $(INST_TARGET)

install:
	sudo cp $(TARGET) $(INST_TARGET)

fmt:
	go fmt $(SRC)
