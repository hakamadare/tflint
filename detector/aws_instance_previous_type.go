package detector

import (
	"fmt"

	"github.com/hashicorp/hcl/hcl/ast"
	"github.com/wata727/tflint/issue"
)

type AwsInstancePreviousTypeDetector struct {
	*Detector
	IssueType             string
	Target                string
	DeepCheck             bool
	previousInstanceTypes map[string]bool
}

func (d *Detector) CreateAwsInstancePreviousTypeDetector() *AwsInstancePreviousTypeDetector {
	return &AwsInstancePreviousTypeDetector{
		Detector:              d,
		IssueType:             issue.WARNING,
		Target:                "aws_instance",
		DeepCheck:             false,
		previousInstanceTypes: map[string]bool{},
	}
}

func (d *AwsInstancePreviousTypeDetector) PreProcess() {
	d.previousInstanceTypes = map[string]bool{
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
}

func (d *AwsInstancePreviousTypeDetector) Detect(file string, item *ast.ObjectItem, issues *[]*issue.Issue) {
	instanceTypeToken, err := hclLiteralToken(item, "instance_type")
	if err != nil {
		d.Logger.Error(err)
		return
	}
	instanceType, err := d.evalToString(instanceTypeToken.Text)
	if err != nil {
		d.Logger.Error(err)
		return
	}

	if d.previousInstanceTypes[instanceType] {
		issue := &issue.Issue{
			Type:    d.IssueType,
			Message: fmt.Sprintf("\"%s\" is previous generation instance type.", instanceType),
			Line:    instanceTypeToken.Pos.Line,
			File:    file,
		}
		*issues = append(*issues, issue)
	}
}
