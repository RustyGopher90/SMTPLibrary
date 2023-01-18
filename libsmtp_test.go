package libsmtp

import (
	"fmt"
	"testing"
)

var server = "someAddress@someServer.com"
var smtpconfig = SMTPConfig{
	From:        "reporting@someServer.com",
	ReplyTo:     "SomePersonalAddr@someServer.com",
	SMTPServer:  server,
	SMTPPort:    25,
	SMTPSubject: "Super Cool Libraries",
	Recipients:  []string{"SomePersonalAddr@someServer.com"},
	Message:     "Man this library is awesome!",
	Urgent:      false,
}

func TestValidateSMTPConfigGoodValues(t *testing.T) {
	got := validateSMTPConfig(&smtpconfig)
	if got != nil {
		t.Fatalf("Got : %v. Wanted %v", got, nil)
	}
}

func TestValidateSMTPConfigBadValues(t *testing.T) {
	smtpconfigCopy := smtpconfig
	smtpconfigCopy.SMTPPort = 0
	got := validateSMTPConfig(&smtpconfigCopy)
	want := fmt.Errorf("zero value for SMTPPort in SMTPConfig struct of type %T", smtpconfigCopy.SMTPPort)
	if got.Error() != want.Error() {
		t.Fatalf("Got : %v. Wanted : %v.", got, want)
	}
}

func TestValidateSMTPConfigBadValuestwo(t *testing.T) {
	smtpconfigCopy := smtpconfig
	smtpconfigCopy.Recipients = []string{}
	got := validateSMTPConfig(&smtpconfigCopy)
	want := fmt.Errorf("missing value(s) for Recipients in SMTPConfig struct of type %T", smtpconfigCopy.Recipients)
	if got.Error() != want.Error() {
		t.Fatalf("Got : %v. Wanted : %v.", got, want)
	}
}

func TestValidateSMTPConfigBadValuesthree(t *testing.T) {
	smtpconfigCopy := smtpconfig
	smtpconfigCopy.Message = ""
	got := validateSMTPConfig(&smtpconfigCopy)
	want := fmt.Errorf("missing value for Message in SMTPConfig struct of type %T", smtpconfigCopy.Message)
	if got.Error() != want.Error() {
		t.Fatalf("Got : %v. Wanted : %v.", got, want)
	}
}

func TestSendEmail(t *testing.T) {
	got := SendEmail(&smtpconfig)
	if got != nil {
		t.Fatalf("Got : %v. Wanted : %v.", got, nil)
	}
}

func TestSendUrgentEmail(t *testing.T) {
	smtpconfigCopy := smtpconfig
	smtpconfigCopy.Urgent = true
	got := SendEmail(&smtpconfigCopy)
	if got != nil {
		t.Fatalf("Got : %v. Wanted : %v.", got, nil)
	}
}

func TestSendEmailBadServer(t *testing.T) {
	smtpconfigCopy := smtpconfig
	smtpconfigCopy.SMTPServer = "badsmtpserver.com"
	got := SendEmail(&smtpconfigCopy)
	want := fmt.Errorf("dial tcp: lookup badsmtpserver.com: no such host")
	if got.Error() != want.Error() {
		t.Fatalf("Got : %v. Wanted : %v.", got, want)
	}
}

func TestSendEmailBadMessage(t *testing.T) {
	smtpconfigCopy := smtpconfig
	smtpconfigCopy.Message = ""
	got := SendEmail(&smtpconfigCopy)
	want := fmt.Errorf("missing value for Message in SMTPConfig struct of type %T", smtpconfigCopy.Message)
	if got.Error() != want.Error() {
		t.Fatalf("Got : %v. Wanted : %v.", got, want)
	}
}
