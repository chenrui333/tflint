// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsEfsFileSystemInvalidCreationTokenRule checks the pattern is valid
type AwsEfsFileSystemInvalidCreationTokenRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsEfsFileSystemInvalidCreationTokenRule returns new rule with default attributes
func NewAwsEfsFileSystemInvalidCreationTokenRule() *AwsEfsFileSystemInvalidCreationTokenRule {
	return &AwsEfsFileSystemInvalidCreationTokenRule{
		resourceType:  "aws_efs_file_system",
		attributeName: "creation_token",
		max:           64,
		min:           1,
		pattern:       regexp.MustCompile(`^.+$`),
	}
}

// Name returns the rule name
func (r *AwsEfsFileSystemInvalidCreationTokenRule) Name() string {
	return "aws_efs_file_system_invalid_creation_token"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEfsFileSystemInvalidCreationTokenRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEfsFileSystemInvalidCreationTokenRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEfsFileSystemInvalidCreationTokenRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEfsFileSystemInvalidCreationTokenRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"creation_token must be 64 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"creation_token must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^.+$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
