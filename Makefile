.PHONY: test
test:
	go run main.go -AccountName=domainprod -Debug=true -Env=s -VpcID=vpc-xxxx -Region=ap-southeast-2
