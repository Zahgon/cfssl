package ubiquity

import (
	"crypto/x509"
	"time"
)

type DeprecationSeverity int

const (
	None DeprecationSeverity = iota

	Low

	Medium

	High
)

type SHA1DeprecationPolicy struct {
	Platform string `json:"platform"`

	Severity DeprecationSeverity `json:"severity"`

	Description string `json:"description"`

	EffectiveDate time.Time `json:"effective_date"`

	ExpiryDeadline time.Time `json:"expiry_deadline"`

	NeverIssueAfter time.Time `json:"never_issue_after"`
}

var SHA1DeprecationPolicys = []SHA1DeprecationPolicy{

	{
		Platform:       "Google Chrome",
		Description:    "shows the SSL connection has minor problems",
		Severity:       Medium,
		ExpiryDeadline: time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
	},

	{
		Platform:       "Google Chrome",
		Description:    "shows the SSL connection is untrusted",
		Severity:       High,
		ExpiryDeadline: time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC),
	},

	{
		Platform:       "Mozilla Firefox",
		Description:    "gives warning in the developer console",
		Severity:       Low,
		ExpiryDeadline: time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC),
	},

	{
		Platform:        "Mozilla Firefox",
		Description:     "shows the SSL connection is untrusted",
		Severity:        Medium,
		EffectiveDate:   time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
		NeverIssueAfter: time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
	},

	{
		Platform:       "Mozilla Firefox",
		Description:    "shows the SSL connection is untrusted",
		Severity:       High,
		EffectiveDate:  time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC),
		ExpiryDeadline: time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC),
	},

	{
		Platform:       "Microsoft Windows Vista and later",
		Description:    "shows the SSL connection is untrusted",
		Severity:       High,
		EffectiveDate:  time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC),
		ExpiryDeadline: time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
}

func (p SHA1DeprecationPolicy) Flag(chain []*x509.Certificate) bool {
	_ = "STUB: not implemented"
	return false
}

func SHA1DeprecationMessages(chain []*x509.Certificate) []string {
	_ = "STUB: not implemented"
	return nil
}
