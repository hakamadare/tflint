package aws

import (
	"fmt"
	"strings"

	"github.com/hashicorp/hcl/hcl/ast"
	"github.com/wata727/tflint/issue"
)

func (d *AwsDetector) DetectAwsInstancePreviousType() []*issue.Issue {
	var issues = []*issue.Issue{}

	for _, item := range d.List.Filter("resource", "aws_instance").Items {
		instanceTypeToken := item.Val.(*ast.ObjectType).List.Filter("instance_type").Items[0].Val.(*ast.LiteralType).Token
		instanceTypeKey := d.EvalConfig.Eval(strings.Trim(instanceTypeToken.Text, "\""))

		if PreviousInstanceType[instanceTypeKey] {
			issue := &issue.Issue{
				Type:    "NOTICE",
				Message: fmt.Sprintf("\"%s\" is previous generation instance type.", instanceTypeKey),
				Line:    instanceTypeToken.Pos.Line,
				File:    d.File,
			}
			issues = append(issues, issue)
		}
	}

	return issues
}

var PreviousInstanceType = map[string]bool{
	"t1.micro":    true,
	"m1.small":    true,
	"m1.medium":   true,
	"m1.large":    true,
	"m1.xlarge":   true,
	"c1.medium":   true,
	"c1.xlarge":   true,
	"cc2.8xlarge": true,
	"cg1.4xlarge": true,
	"m2.xlarge":   true,
	"m2.2xlarge":  true,
	"m2.4xlarge":  true,
	"cr1.8xlarge": true,
	"hi1.4xlarge": true,
	"hs1.8xlarge": true,
}
