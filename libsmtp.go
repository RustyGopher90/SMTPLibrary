package libsmtp

import (
	"encoding/base64"
	"fmt"
	"net/smtp"
	"strings"
)

type SMTPConfig struct {
	From        string
	ReplyTo     string
	SMTPServer  string
	SMTPPort    int
	SMTPSubject string
	Recipients  []string
	Message     string
	Urgent      bool
}

func SendEmail(smtpStruct *SMTPConfig) error {
	if smtpStruct.ReplyTo == "" {
		smtpStruct.ReplyTo = smtpStruct.From
	}
	err := validateSMTPConfig(smtpStruct)
	if err != nil {
		return err
	}
	serverAddr := fmt.Sprintf("%v:%v", smtpStruct.SMTPServer, smtpStruct.SMTPPort)
	var emailBody strings.Builder
	emailBody.WriteString(smtpStruct.Message)

	smtpClient, err := smtp.Dial(serverAddr)
	if err != nil {
		return err
	}

	defer smtpClient.Close()

	err = smtpClient.Mail(smtpStruct.From)
	if err != nil {
		return err
	}

	for _, rcp := range smtpStruct.Recipients {
		err := smtpClient.Rcpt(rcp)
		if err != nil {
			return err
		}
	}

	smtpWriter, err := smtpClient.Data()
	if err != nil {
		return err
	}

	priority := ""
	microsoftPriority := ""
	otherClientPriority := ""
	if smtpStruct.Urgent {
		priority = "Priority: 1 (Highest)\r\n "
		microsoftPriority = "MSMail-Priority: High\r\n"
		otherClientPriority = "Importance: High\r\n"
	}

	email := "To: " + strings.Join(smtpStruct.Recipients, ",") + "\r\n" +
		"From: " + smtpStruct.From + "\r\n" +
		"Reply-To: " + smtpStruct.ReplyTo + "\r\n" +
		"Subject: " + smtpStruct.SMTPSubject + "\r\n" +
		priority +
		microsoftPriority +
		otherClientPriority +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"Content-Transfer-Encoding: base64\r\n" +
		"\r\n" + base64.StdEncoding.EncodeToString([]byte(emailBody.String()))

	_, err = smtpWriter.Write([]byte(email))
	if err != nil {
		return err
	}

	err = smtpWriter.Close()
	if err != nil {
		return err
	}
	smtpClient.Quit()
	return nil
}

func validateSMTPConfig(smtpStruct *SMTPConfig) error {
	if smtpStruct.From == "" {
		return fmt.Errorf("missing value for From field in SMTPConfig struct of type %T", smtpStruct.From)
	} else if smtpStruct.SMTPServer == "" {
		return fmt.Errorf("missing value for SMTPServer field in SMTPConfig struct of type %T", smtpStruct.SMTPServer)
	} else if smtpStruct.SMTPPort == 0 {
		return fmt.Errorf("zero value for SMTPPort in SMTPConfig struct of type %T", smtpStruct.SMTPPort)
	} else if smtpStruct.SMTPSubject == "" {
		return fmt.Errorf("missing value for SMTPSubject in SMTPConfig struct of type %T", smtpStruct.SMTPSubject)
	} else if len(smtpStruct.Recipients) == 0 {
		return fmt.Errorf("missing value(s) for Recipients in SMTPConfig struct of type %T", smtpStruct.Recipients)
	} else if smtpStruct.Message == "" {
		return fmt.Errorf("missing value for Message in SMTPConfig struct of type %T", smtpStruct.Message)
	} else if smtpStruct.ReplyTo == "" {
		return fmt.Errorf("missing value for Message in SMTPConfig struct of type %T", smtpStruct.ReplyTo)
	} else {
		return nil
	}
}
