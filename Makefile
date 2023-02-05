all: build install
build:
	@go build -o /usr/local/bin/siri-ipmi-web -ldflags "-X 'main.USER=$(USER)' -X 'main.USER=$(USER)' -X 'main.PASSWORD=$(PASSWORD)' -X 'main.IPADDRESS=$(IPADDRESS)' -X 'main.TOKEN=$(TOKEN)'" .

install:
	@cp siri-ipmi-web.service /etc/systemd/system/siri-ipmi-web.service
	systemctl enable siri-ipmi-web.service
	systemctl start siri-ipmi-web.service
