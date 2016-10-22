package aws

import (
	"github.com/hashicorp/hcl/hcl/ast"
	"github.com/wata727/tflint/issue"
)

func (d *AwsDetector) DetectAwsInstanceNotSpecifiedIamProfile() []*issue.Issue {
	var issues = []*issue.Issue{}

	for _, item := range d.List.Filter("resource", "aws_instance").Items {
		instanceIAMProfile := item.Val.(*ast.ObjectType).List.Filter("iam_instance_profile")

		if len(instanceIAMProfile.Items) == 0 {
			issue := &issue.Issue{
				Type:    "NOTICE",
				Message: "\"iam_instance_profile\" is not specified. You cannot edit this value later.",
				Line:    item.Pos().Line,
				File:    d.File,
			}
			issues = append(issues, issue)
		}
	}

	return issues
}
