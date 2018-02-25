.PHONY: install
install:
	glide install

.PHONY: test
test:
	go run cmd/main.go -AccountName=domainprod -Debug=true -Env=s -VpcID=vpc-xxxx -Region=ap-southeast-2
