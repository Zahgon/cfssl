package config

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/asn1"
	"regexp"
	"time"

	"github.com/cloudflare/cfssl/auth"
	ocspConfig "github.com/cloudflare/cfssl/ocsp/config"

	_ "github.com/zmap/zlint/v3"
	"github.com/zmap/zlint/v3/lint"
)

type CSRWhitelist struct {
	Subject, PublicKeyAlgorithm, PublicKey, SignatureAlgorithm bool
	DNSNames, IPAddresses, EmailAddresses, URIs                bool
}

type OID asn1.ObjectIdentifier

type CertificatePolicy struct {
	ID         OID
	Qualifiers []CertificatePolicyQualifier
}

type CertificatePolicyQualifier struct {
	Type  string
	Value string
}

type AuthRemote struct {
	RemoteName  string `json:"remote"`
	AuthKeyName string `json:"auth_key"`
}

type CAConstraint struct {
	IsCA           bool `json:"is_ca"`
	MaxPathLen     int  `json:"max_path_len"`
	MaxPathLenZero bool `json:"max_path_len_zero"`
}

type SigningProfile struct {
	Usage               []string     `json:"usages"`
	IssuerURL           []string     `json:"issuer_urls"`
	OCSP                string       `json:"ocsp_url"`
	CRL                 string       `json:"crl_url"`
	CAConstraint        CAConstraint `json:"ca_constraint"`
	OCSPNoCheck         bool         `json:"ocsp_no_check"`
	ExpiryString        string       `json:"expiry"`
	BackdateString      string       `json:"backdate"`
	AuthKeyName         string       `json:"auth_key"`
	CopyExtensions      bool         `json:"copy_extensions"`
	PrevAuthKeyName     string       `json:"prev_auth_key"`
	RemoteName          string       `json:"remote"`
	NotBefore           time.Time    `json:"not_before"`
	NotAfter            time.Time    `json:"not_after"`
	NameWhitelistString string       `json:"name_whitelist"`
	AuthRemote          AuthRemote   `json:"auth_remote"`
	CTLogServers        []string     `json:"ct_log_servers"`
	AllowedExtensions   []OID        `json:"allowed_extensions"`
	CertStore           string       `json:"cert_store"`

	LintErrLevel lint.LintStatus `json:"lint_error_level"`

	ExcludeLints []string `json:"ignored_lints"`

	ExcludeLintSources []string `json:"ignored_lint_sources"`

	Policies                    []CertificatePolicy
	Expiry                      time.Duration
	Backdate                    time.Duration
	Provider                    auth.Provider
	PrevProvider                auth.Provider
	RemoteProvider              auth.Provider
	RemoteServer                string
	RemoteCAs                   *x509.CertPool
	ClientCert                  *tls.Certificate
	CSRWhitelist                *CSRWhitelist
	NameWhitelist               *regexp.Regexp
	ExtensionWhitelist          map[string]bool
	ClientProvidesSerialNumbers bool

	LintRegistry lint.Registry
}

func (oid *OID) UnmarshalJSON(data []byte) (err error) { _ = "STUB: not implemented"; return nil }

func (oid OID) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func parseObjectIdentifier(oidString string) (oid asn1.ObjectIdentifier, err error) {
	_ = "STUB: not implemented"
	return *new(asn1.ObjectIdentifier), nil
}

const timeFormat = "2006-01-02T15:04:05"

func (p *SigningProfile) populate(cfg *Config) error { _ = "STUB: not implemented"; return nil }

func (p *SigningProfile) updateRemote(remote string) error { _ = "STUB: not implemented"; return nil }

func (p *Signing) OverrideRemotes(remote string) error { _ = "STUB: not implemented"; return nil }

func (p *Signing) SetClientCertKeyPairFromFile(certFile string, keyFile string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Signing) SetRemoteCAsFromFile(caFile string) error { _ = "STUB: not implemented"; return nil }

func (p *Signing) SetRemoteCAs(remoteCAs *x509.CertPool) { _ = "STUB: not implemented"; return }

func (p *Signing) NeedsRemoteSigner() bool { _ = "STUB: not implemented"; return false }

func (p *Signing) NeedsLocalSigner() bool { _ = "STUB: not implemented"; return false }

func (p *SigningProfile) Usages() (ku x509.KeyUsage, eku []x509.ExtKeyUsage, unk []string) {
	_ = "STUB: not implemented"
	return *new(x509.KeyUsage), nil, nil
}

func (p *SigningProfile) validProfile(isDefault bool) bool { _ = "STUB: not implemented"; return false }

func (p *SigningProfile) hasLocalConfig() bool { _ = "STUB: not implemented"; return false }

func (p *Signing) warnSkippedSettings() { _ = "STUB: not implemented"; return }

type Signing struct {
	Profiles map[string]*SigningProfile `json:"profiles"`
	Default  *SigningProfile            `json:"default"`
}

type Config struct {
	Signing  *Signing           `json:"signing"`
	OCSP     *ocspConfig.Config `json:"ocsp"`
	AuthKeys map[string]AuthKey `json:"auth_keys,omitempty"`
	Remotes  map[string]string  `json:"remotes,omitempty"`
}

func (c *Config) Valid() bool { _ = "STUB: not implemented"; return false }

func (p *Signing) Valid() bool { _ = "STUB: not implemented"; return false }

var KeyUsage = map[string]x509.KeyUsage{
	"signing":            x509.KeyUsageDigitalSignature,
	"digital signature":  x509.KeyUsageDigitalSignature,
	"content commitment": x509.KeyUsageContentCommitment,
	"key encipherment":   x509.KeyUsageKeyEncipherment,
	"key agreement":      x509.KeyUsageKeyAgreement,
	"data encipherment":  x509.KeyUsageDataEncipherment,
	"cert sign":          x509.KeyUsageCertSign,
	"crl sign":           x509.KeyUsageCRLSign,
	"encipher only":      x509.KeyUsageEncipherOnly,
	"decipher only":      x509.KeyUsageDecipherOnly,
}

var ExtKeyUsage = map[string]x509.ExtKeyUsage{
	"any":              x509.ExtKeyUsageAny,
	"server auth":      x509.ExtKeyUsageServerAuth,
	"client auth":      x509.ExtKeyUsageClientAuth,
	"code signing":     x509.ExtKeyUsageCodeSigning,
	"email protection": x509.ExtKeyUsageEmailProtection,
	"s/mime":           x509.ExtKeyUsageEmailProtection,
	"ipsec end system": x509.ExtKeyUsageIPSECEndSystem,
	"ipsec tunnel":     x509.ExtKeyUsageIPSECTunnel,
	"ipsec user":       x509.ExtKeyUsageIPSECUser,
	"timestamping":     x509.ExtKeyUsageTimeStamping,
	"ocsp signing":     x509.ExtKeyUsageOCSPSigning,
	"microsoft sgc":    x509.ExtKeyUsageMicrosoftServerGatedCrypto,
	"netscape sgc":     x509.ExtKeyUsageNetscapeServerGatedCrypto,
}

type AuthKey struct {
	Type string `json:"type"`

	Key string `json:"key"`
}

func DefaultConfig() *SigningProfile { _ = "STUB: not implemented"; return nil }

func LoadFile(path string) (*Config, error) { _ = "STUB: not implemented"; return nil, nil }

func LoadConfig(config []byte) (*Config, error) { _ = "STUB: not implemented"; return nil, nil }
