package generic

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/jsonpb"
	"github.com/pkg/errors"
	"github.com/stackrox/rox/central/notifiers"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/errorhelpers"
	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/pkg/retry"
	"github.com/stackrox/rox/pkg/stringutils"
	"github.com/stackrox/rox/pkg/urlfmt"
)

var (
	log = logging.LoggerForModule()
)

const (
	timeout = 5 * time.Second

	alertMessageKey         = "alert"
	auditMessageKey         = "audit"
	networkPolicyMessageKey = "networkpolicy"
)

// generic notifier plugin
type generic struct {
	*storage.Notifier

	client                 *http.Client
	fullyQualifiedEndpoint string
	extraFieldsJSONPrefix  string
}

// AlertNotify takes in an alert and generates the Slack message
func (g *generic) AlertNotify(alert *storage.Alert) error {
	return g.postMessageWithRetry(alert, alertMessageKey)
}

// YamlNotify takes in a yaml file and generates the Slack message
func (g *generic) NetworkPolicyYAMLNotify(yaml string, clusterName string) error {
	msg := &v1.NetworkPolicyNotification{
		Cluster: clusterName,
		Yaml:    yaml,
	}
	return g.postMessageWithRetry(msg, networkPolicyMessageKey)
}

func validateConfig(generic *storage.Generic) error {
	errList := errorhelpers.NewErrorList("Generic webhook validation")
	if generic.GetEndpoint() == "" {
		errList.AddString("endpoint is required")
	}
	if generic.GetUsername() != generic.GetPassword() && stringutils.AtLeastOneEmpty(generic.GetUsername(), generic.GetPassword()) {
		errList.AddString("both username and password must be defined together")
	}
	for _, f := range generic.GetHeaders() {
		if f.GetKey() == "" || f.GetValue() == "" {
			errList.AddString("all headers must have both a key and a value")
		}
	}
	for _, f := range generic.GetExtraFields() {
		if f.GetKey() == "" || f.GetValue() == "" {
			errList.AddString("all extra fields must have both a key and a value")
		}
	}
	return errList.ToError()
}

func getExtraFieldJSON(fields []*storage.KeyValuePair) (string, error) {
	fieldMap := make(map[string]string)
	for _, f := range fields {
		fieldMap[f.Key] = f.Value
	}
	data, err := json.Marshal(fieldMap)
	if err != nil {
		return "", err
	}

	// Cut off trailing '}'
	data = data[:len(data)-1]
	return string(data), nil
}

func newGeneric(notifier *storage.Notifier) (*generic, error) {
	genericConfig, ok := notifier.Config.(*storage.Notifier_Generic)

	if !ok {
		return nil, validateConfig(&storage.Generic{})
	}
	conf := genericConfig.Generic
	if err := validateConfig(conf); err != nil {
		return nil, err
	}
	fullyQualifiedEndpoint, err := urlfmt.FormatURL(conf.GetEndpoint(), urlfmt.HTTPS, urlfmt.HonorInputSlash)
	if err != nil {
		return nil, err
	}
	rootCAs, err := x509.SystemCertPool()
	if err != nil {
		rootCAs = x509.NewCertPool()
	}
	if conf.GetCaCert() != "" {
		if ok := rootCAs.AppendCertsFromPEM([]byte(conf.GetCaCert())); !ok {
			return nil, errors.New("could not add CA Cert passed in configuration")
		}
	}
	extraFieldsJSON, err := getExtraFieldJSON(conf.ExtraFields)
	if err != nil {
		return nil, err
	}

	return &generic{
		Notifier: notifier,

		client: &http.Client{
			Timeout: timeout,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: conf.GetSkipTLSVerify(),
					RootCAs:            rootCAs,
				},
			},
		},
		fullyQualifiedEndpoint: fullyQualifiedEndpoint,
		extraFieldsJSONPrefix:  extraFieldsJSON,
	}, nil
}

func (g *generic) ProtoNotifier() *storage.Notifier {
	return g.Notifier
}

func (g *generic) Test() error {
	alert := &storage.Alert{
		Id: "testalert",
		Policy: &storage.Policy{
			Name: "This is a test message created to test integration with StackRox.",
		},
	}
	return g.AlertNotify(alert)
}

func (g *generic) constructJSON(message proto.Message, msgKey string) (io.Reader, error) {
	msgStr, err := new(jsonpb.Marshaler).MarshalToString(message)
	if err != nil {
		return nil, err
	}

	var strJSON string
	// No extra fields append so that the payload is something like {"alert": {...}}
	if len(g.Notifier.GetGeneric().GetExtraFields()) == 0 {
		strJSON = fmt.Sprintf(`{"%s": %s}`, msgKey, msgStr)
	} else {
		strJSON = fmt.Sprintf(`%s,"%s": %s}`, g.extraFieldsJSONPrefix, msgKey, msgStr)
	}
	return bytes.NewBufferString(strJSON), nil
}

func (g *generic) postMessage(message proto.Message, msgKey string) error {
	body, err := g.constructJSON(message, msgKey)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, g.fullyQualifiedEndpoint, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	for _, h := range g.GetGeneric().GetHeaders() {
		req.Header.Add(h.GetKey(), h.GetValue())
	}

	if g.GetGeneric().GetUsername() != "" {
		req.SetBasicAuth(g.GetGeneric().GetUsername(), g.GetGeneric().GetUsername())
	}

	resp, err := g.client.Do(req)
	if err != nil {
		return err
	}

	return notifiers.CreateError("webhook", resp)
}

func (g *generic) postMessageWithRetry(message proto.Message, msgKey string) error {
	return retry.WithRetry(
		func() error {
			return g.postMessage(message, msgKey)
		},
		retry.OnlyRetryableErrors(),
		retry.Tries(3),
		retry.BetweenAttempts(func(previousAttempt int) {
			wait := time.Duration(previousAttempt * previousAttempt * 100)
			time.Sleep(wait * time.Millisecond)
		}),
	)
}

func (g *generic) SendAuditMessage(msg *v1.Audit_Message) error {
	if !g.AuditLoggingEnabled() {
		return nil
	}
	return g.postMessageWithRetry(msg, auditMessageKey)
}

func (g *generic) AuditLoggingEnabled() bool {
	return g.GetGeneric().GetAuditLoggingEnabled()
}

func init() {
	notifiers.Add("generic", func(notifier *storage.Notifier) (notifiers.Notifier, error) {
		g, err := newGeneric(notifier)
		return g, err
	})
}
