package model

// MailSenderOps model about any gmail that configured truly can send mail to another mails
type MailConfig struct {
	SMTPOps       SMTPOps
	MailSenderOps MailSenderOps
}

// MailSenderOps
type MailSenderOps struct {
	MailSender  string `json:"mail-sender" validate:"required"`
	AppPassword string `json:"app-password" validate:"required"`
}

// SMTPOps is
type SMTPOps struct {
	SMTPServer string `json:"smtp-server"`
	SMTPPort   int    `json:"smtp-port"`
}
