package kforce

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// all pre hooks are herepackage kforce

// all pre hooks are here

const (
	public  = "public"
	private = "private"
)

type subnet struct {
	zone   string
	ID     string // subnet id
	facing string
	natID  string
	CIDR   string
}

// VpcFacts - vpc artifacts and az list
type VpcFacts struct {
	azs []string
	vpc []subnet
	ID  string // vpc id
}

// GetherVpcFacts - fetch aws vpc artifacts
func GetherVpcFacts(vpcID string) (*VpcFacts, error) {
	vpc := &VpcFacts{}
	sess := session.Must(session.NewSession())
	svc := ec2.New(sess)
	input := &ec2.DescribeVpcAttributeInput{
		Attribute: aws.String("enableDnsHostnames"),
		VpcId:     aws.String(vpcID),
	}

	result, err := svc.DescribeVpcAttribute(input)
	fmt.Println(result, err)
	return vpc, nil
}
