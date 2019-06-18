package apiv3

import (
	"net/mail"
)

// Email holds email name and address info
type Email struct {
	Name    string `json:"name,omitempty"`
	Address string `json:"email,omitempty"`
}

// NewEmail ...
func NewEmail(name, address string) *Email {
	return &Email{
		Name:    name,
		Address: address,
	}
}

// ParseEmail parses a string that contains an rfc822 formatted email address
// and returns an instance of *Email.
func ParseEmail(emailInfo string) (*Email, error) {
	e, err := mail.ParseAddress(emailInfo)
	if err != nil {
		return nil, err
	}

	return &Email{
		Name:    e.Name,
		Address: e.Address,
	}, nil
}

// MailV3 represents the mail structure for v3 API
type MailV3 struct {
	// Required

	Personalizations []*Personalization `json:"personalizations,omitempty"`
	From             *Email             `json:"from,omitempty"`
	Subject          string             `json:"subject,omitempty"`
	Content          []*Content         `json:"content,omitempty"`

	// Optional

	ReplyTo          *Email            `json:"reply_to,omitempty"`
	Attachments      []*Attachment     `json:"attachments,omitempty"`
	TemplateID       string            `json:"template_id,omitempty"`
	Sections         map[string]string `json:"sections,omitempty"`
	Headers          map[string]string `json:"headers,omitempty"`
	Categories       []string          `json:"categories,omitempty"`
	CustomArgs       map[string]string `json:"custom_args,omitempty"`
	SendAt           int               `json:"send_at,omitempty"`
	BatchID          string            `json:"batch_id,omitempty"`
	ASM              *ASM              `json:"asm,omitempty"`
	IPPoolName       string            `json:"ip_pool_name,omitempty"`
	MailSettings     *MailSettings     `json:"mail_settings,omitempty"`
	TrackingSettings *TrackingSettings `json:"tracking_settings,omitempty"`
}

// NewMailV3 ...
func NewMailV3(from *Email, subject string, content []*Content, personalizations []*Personalization) *MailV3 {
	return &MailV3{
		From:             from,
		Subject:          subject,
		Content:          content,
		Personalizations: personalizations,
	}
}

// AddPersonalizations adds personalizations to the current MailV3 object
func (m *MailV3) AddPersonalizations(p ...*Personalization) *MailV3 {
	m.Personalizations = append(m.Personalizations, p...)
	return m
}

// SetFrom ...
func (m *MailV3) SetFrom(e *Email) *MailV3 {
	m.From = e
	return m
}

// AddContent ...
func (m *MailV3) AddContent(c ...*Content) *MailV3 {
	m.Content = append(m.Content, c...)
	return m
}

// SetTextContent sets text content of the message
func (m *MailV3) SetTextContent(text string) *MailV3 {
	m.AddContent(&Content{
		Type:  "text/plain",
		Value: text,
	})

	return m
}

// SetHTMLContent sets HTML content of the message
func (m *MailV3) SetHTMLContent(html string) *MailV3 {
	m.AddContent(&Content{
		Type:  "text/html",
		Value: html,
	})

	return m
}

// SetReplyTo ...
func (m *MailV3) SetReplyTo(e *Email) *MailV3 {
	m.ReplyTo = e
	return m
}

// AddAttachment ...
func (m *MailV3) AddAttachment(a ...*Attachment) *MailV3 {
	m.Attachments = append(m.Attachments, a...)
	return m
}

// SetTemplateID ...
func (m *MailV3) SetTemplateID(templateID string) *MailV3 {
	m.TemplateID = templateID
	return m
}

// AddSection ...
func (m *MailV3) AddSection(key string, value string) *MailV3 {
	if m.Sections == nil {
		m.Sections = make(map[string]string)
	}

	m.Sections[key] = value
	return m
}

// SetHeader ...
func (m *MailV3) SetHeader(key string, value string) *MailV3 {
	if m.Headers == nil {
		m.Headers = make(map[string]string)
	}

	m.Headers[key] = value
	return m
}

// AddCategories ...
func (m *MailV3) AddCategories(category ...string) *MailV3 {
	m.Categories = append(m.Categories, category...)
	return m
}

// SetCustomArg ...
func (m *MailV3) SetCustomArg(key string, value string) *MailV3 {
	if m.CustomArgs == nil {
		m.CustomArgs = make(map[string]string)
	}

	m.CustomArgs[key] = value
	return m
}

// SetSendAt ...
func (m *MailV3) SetSendAt(sendAt int) *MailV3 {
	m.SendAt = sendAt
	return m
}

// SetBatchID ...
func (m *MailV3) SetBatchID(batchID string) *MailV3 {
	m.BatchID = batchID
	return m
}

// SetASM ...
func (m *MailV3) SetASM(asm *ASM) *MailV3 {
	m.ASM = asm
	return m
}

// SetIPPoolName ...
func (m *MailV3) SetIPPoolName(IPPoolName string) *MailV3 {
	m.IPPoolName = IPPoolName
	return m
}

// SetMailSettings ...
func (m *MailV3) SetMailSettings(mailSettings *MailSettings) *MailV3 {
	m.MailSettings = mailSettings
	return m
}

// SetTrackingSettings ...
func (m *MailV3) SetTrackingSettings(trackingSettings *TrackingSettings) *MailV3 {
	m.TrackingSettings = trackingSettings
	return m
}

// Personalization holds mail body struct
type Personalization struct {
	To                  []*Email               `json:"to,omitempty"`
	CC                  []*Email               `json:"cc,omitempty"`
	BCC                 []*Email               `json:"bcc,omitempty"`
	Subject             string                 `json:"subject,omitempty"`
	Headers             map[string]string      `json:"headers,omitempty"`
	Substitutions       map[string]string      `json:"substitutions,omitempty"`
	CustomArgs          map[string]string      `json:"custom_args,omitempty"`
	SendAt              int                    `json:"send_at,omitempty"`
	DynamicTemplateData map[string]interface{} `json:"dynamic_template_data,omitempty"`
	Categories          []string               `json:"categories,omitempty"`
}

// NewPersonalization creates new personalization
func NewPersonalization() *Personalization {
	return &Personalization{
		To:                  make([]*Email, 0),
		CC:                  make([]*Email, 0),
		BCC:                 make([]*Email, 0),
		Headers:             make(map[string]string),
		Substitutions:       make(map[string]string),
		CustomArgs:          make(map[string]string),
		DynamicTemplateData: make(map[string]interface{}),
		Categories:          make([]string, 0),
	}
}

// AddTos adds 'to' addresses to the personalization
func (p *Personalization) AddTos(to ...*Email) {
	p.To = append(p.To, to...)
}

// AddCCs adds 'cc' addresses to the personalization
func (p *Personalization) AddCCs(cc ...*Email) {
	p.CC = append(p.CC, cc...)
}

// AddBCCs adds 'bcc' addresses to the personalization
func (p *Personalization) AddBCCs(bcc ...*Email) {
	p.BCC = append(p.BCC, bcc...)
}

// SetHeader sets headers to the personalization
func (p *Personalization) SetHeader(key string, value string) {
	p.Headers[key] = value
}

// SetSubstitution sets substitutions to the personalization
func (p *Personalization) SetSubstitution(key string, value string) {
	p.Substitutions[key] = value
}

// SetCustomArg sets custom args to the personalization
func (p *Personalization) SetCustomArg(key string, value string) {
	p.CustomArgs[key] = value
}

// SetDynamicTemplateData sets dynamic template data to the personalization
func (p *Personalization) SetDynamicTemplateData(key string, value interface{}) {
	p.DynamicTemplateData[key] = value
}

// SetSendAt sets SendAt to the personalization for schedule send
// sendAt should be a unix timestamp in seconds
func (p *Personalization) SetSendAt(sendAt int) {
	p.SendAt = sendAt
}

// Content defines content of the mail body
type Content struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

// NewContent ...
func NewContent(contentType string, value string) *Content {
	return &Content{
		Type:  contentType,
		Value: value,
	}
}

// Attachment holds attachement information
type Attachment struct {
	Content     string `json:"content,omitempty"`
	Type        string `json:"type,omitempty"`
	Name        string `json:"name,omitempty"`
	Filename    string `json:"filename,omitempty"`
	Disposition string `json:"disposition,omitempty"`
	ContentID   string `json:"content_id,omitempty"`
}

// NewAttachment ...
func NewAttachment() *Attachment {
	return &Attachment{}
}

// SetContent ...
func (a *Attachment) SetContent(content string) *Attachment {
	a.Content = content
	return a
}

// SetType ...
func (a *Attachment) SetType(contentType string) *Attachment {
	a.Type = contentType
	return a
}

// SetFilename ...
func (a *Attachment) SetFilename(filename string) *Attachment {
	a.Filename = filename
	return a
}

// SetDisposition ...
func (a *Attachment) SetDisposition(disposition string) *Attachment {
	a.Disposition = disposition
	return a
}

// SetContentID ...
func (a *Attachment) SetContentID(contentID string) *Attachment {
	a.ContentID = contentID
	return a
}

// ASM contains Grpip Id and int array of groups ID
type ASM struct {
	GroupID         int   `json:"group_id,omitempty"`
	GroupsToDisplay []int `json:"groups_to_display,omitempty"`
}

// NewASM ...
func NewASM() *ASM {
	return &ASM{}
}

// SetGroupID ...
func (a *ASM) SetGroupID(groupID int) *ASM {
	a.GroupID = groupID
	return a
}

// AddGroupsToDisplay ...
func (a *ASM) AddGroupsToDisplay(groupsToDisplay ...int) *ASM {
	a.GroupsToDisplay = append(a.GroupsToDisplay, groupsToDisplay...)
	return a
}

// Setting enables the mail settings
type Setting struct {
	Enable *bool `json:"enable,omitempty"`
}

// NewSetting ...
func NewSetting(enable bool) *Setting {
	setEnable := enable
	return &Setting{Enable: &setEnable}
}

// MailSettings defines mail and spamCheck settings
type MailSettings struct {
	BCC                  *BCCSetting       `json:"bcc,omitempty"`
	BypassListManagement *Setting          `json:"bypass_list_management,omitempty"`
	Footer               *FooterSetting    `json:"footer,omitempty"`
	SandboxMode          *Setting          `json:"sandbox_mode,omitempty"`
	SpamCheckSetting     *SpamCheckSetting `json:"spam_check,omitempty"`
}

// NewMailSettings ...
func NewMailSettings() *MailSettings {
	return &MailSettings{}
}

// SetBCC ...
func (m *MailSettings) SetBCC(bcc *BCCSetting) *MailSettings {
	m.BCC = bcc
	return m
}

// SetBypassListManagement ...
func (m *MailSettings) SetBypassListManagement(bypassListManagement *Setting) *MailSettings {
	m.BypassListManagement = bypassListManagement
	return m
}

// SetFooter ...
func (m *MailSettings) SetFooter(footerSetting *FooterSetting) *MailSettings {
	m.Footer = footerSetting
	return m
}

// SetSandboxMode ...
func (m *MailSettings) SetSandboxMode(sandboxMode *Setting) *MailSettings {
	m.SandboxMode = sandboxMode
	return m
}

// SetSpamCheckSettings ...
func (m *MailSettings) SetSpamCheckSettings(spamCheckSetting *SpamCheckSetting) *MailSettings {
	m.SpamCheckSetting = spamCheckSetting
	return m
}

// TrackingSettings holds tracking settings and mail settings
type TrackingSettings struct {
	ClickTracking        *ClickTrackingSetting        `json:"click_tracking,omitempty"`
	OpenTracking         *OpenTrackingSetting         `json:"open_tracking,omitempty"`
	SubscriptionTracking *SubscriptionTrackingSetting `json:"subscription_tracking,omitempty"`
	GoogleAnalytics      *GASetting                   `json:"ganalytics,omitempty"`
	BCC                  *BCCSetting                  `json:"bcc,omitempty"`
	BypassListManagement *Setting                     `json:"bypass_list_management,omitempty"`
	Footer               *FooterSetting               `json:"footer,omitempty"`
	SandboxMode          *SandboxModeSetting          `json:"sandbox_mode,omitempty"`
}

// NewTrackingSettings ...
func NewTrackingSettings() *TrackingSettings {
	return &TrackingSettings{}
}

// SetClickTracking ...
func (t *TrackingSettings) SetClickTracking(clickTracking *ClickTrackingSetting) *TrackingSettings {
	t.ClickTracking = clickTracking
	return t

}

// SetOpenTracking ...
func (t *TrackingSettings) SetOpenTracking(openTracking *OpenTrackingSetting) *TrackingSettings {
	t.OpenTracking = openTracking
	return t
}

// SetSubscriptionTracking ...
func (t *TrackingSettings) SetSubscriptionTracking(subscriptionTracking *SubscriptionTrackingSetting) *TrackingSettings {
	t.SubscriptionTracking = subscriptionTracking
	return t
}

// SetGoogleAnalytics ...
func (t *TrackingSettings) SetGoogleAnalytics(googleAnalytics *GASetting) *TrackingSettings {
	t.GoogleAnalytics = googleAnalytics
	return t
}

// BCCSetting holds email bcc setings  to enable of disable
// default is false
type BCCSetting struct {
	Enable *bool `json:"enable,omitempty"`
	Email  string
}

// NewBCCSetting ...
func NewBCCSetting(enable bool, email string) *BCCSetting {
	setEnable := enable

	return &BCCSetting{
		Enable: &setEnable,
		Email:  email,
	}
}

// FooterSetting holds enaable/disable settings
// and the format of footer i.e HTML/Text
type FooterSetting struct {
	Enable *bool  `json:"enable,omitempty"`
	Text   string `json:"text,omitempty"`
	HTML   string `json:"html,omitempty"`
}

// NewFooterSetting ...
func NewFooterSetting(enable bool, text, html string) *FooterSetting {
	setEnable := enable
	return &FooterSetting{
		Enable: &setEnable,
		Text:   text,
		HTML:   html,
	}
}

// ClickTrackingSetting ...
type ClickTrackingSetting struct {
	Enable     *bool `json:"enable,omitempty"`
	EnableText *bool `json:"enable_text,omitempty"`
}

// NewClickTrackingSetting ...
func NewClickTrackingSetting(enable, enableText bool) *ClickTrackingSetting {
	setEnable := enable
	setEnableText := enableText

	return &ClickTrackingSetting{
		Enable:     &setEnable,
		EnableText: &setEnableText,
	}
}

// OpenTrackingSetting ...
type OpenTrackingSetting struct {
	Enable          *bool  `json:"enable,omitempty"`
	SubstitutionTag string `json:"substitution_tag,omitempty"`
}

// NewOpenTrackingSetting ...
func NewOpenTrackingSetting(enable bool, subTag string) *OpenTrackingSetting {
	setEnable := enable

	return &OpenTrackingSetting{
		Enable:          &setEnable,
		SubstitutionTag: subTag,
	}
}

// SubscriptionTrackingSetting ...
type SubscriptionTrackingSetting struct {
	Enable          *bool  `json:"enable,omitempty"`
	Text            string `json:"text,omitempty"`
	HTML            string `json:"html,omitempty"`
	SubstitutionTag string `json:"substitution_tag,omitempty"`
}

// NewSubscriptionTrackingSetting ...
func NewSubscriptionTrackingSetting(enable bool, text, html, subTag string) *SubscriptionTrackingSetting {
	setEnable := enable

	return &SubscriptionTrackingSetting{
		Enable:          &setEnable,
		Text:            text,
		HTML:            html,
		SubstitutionTag: subTag,
	}
}

// SandboxModeSetting ...
type SandboxModeSetting struct {
	Enable      *bool             `json:"enable,omitempty"`
	ForwardSpam *bool             `json:"forward_spam,omitempty"`
	SpamCheck   *SpamCheckSetting `json:"spam_check,omitempty"`
}

// NewSandboxModeSetting ...
func NewSandboxModeSetting(enable bool, forwardSpam bool, spamCheck *SpamCheckSetting) *SandboxModeSetting {
	setEnable := enable
	setForwardSpam := forwardSpam
	return &SandboxModeSetting{
		Enable:      &setEnable,
		ForwardSpam: &setForwardSpam,
		SpamCheck:   spamCheck,
	}
}

// SpamCheckSetting holds spam settings and
// which can be enable or disable and
// contains spamThreshold value
type SpamCheckSetting struct {
	Enable        *bool  `json:"enable,omitempty"`
	SpamThreshold int    `json:"threshold,omitempty"`
	PostToURL     string `json:"post_to_url,omitempty"`
}

// NewSpamCheckSetting ...
func NewSpamCheckSetting(enable bool, spamThreshold int, postToURL string) *SpamCheckSetting {
	setEnable := enable

	return &SpamCheckSetting{
		Enable:        &setEnable,
		SpamThreshold: spamThreshold,
		PostToURL:     postToURL,
	}
}

// GASetting represents the google analytics settings
type GASetting struct {
	Enable          *bool  `json:"enable,omitempty"`
	CampaignSource  string `json:"utm_source,omitempty"`
	CampaignTerm    string `json:"utm_term,omitempty"`
	CampaignContent string `json:"utm_content,omitempty"`
	CampaignName    string `json:"utm_campaign,omitempty"`
	CampaignMedium  string `json:"utm_medium,omitempty"`
}

// NewGASetting ...
func NewGASetting(enable bool) *GASetting {
	setEnable := enable

	return &GASetting{Enable: &setEnable}
}

// SetCampaignSource ...
func (g *GASetting) SetCampaignSource(campaignSource string) *GASetting {
	g.CampaignSource = campaignSource
	return g
}

// SetCampaignContent ...
func (g *GASetting) SetCampaignContent(campaignContent string) *GASetting {
	g.CampaignContent = campaignContent
	return g
}

// SetCampaignTerm ...
func (g *GASetting) SetCampaignTerm(campaignTerm string) *GASetting {
	g.CampaignTerm = campaignTerm
	return g
}

// SetCampaignName ...
func (g *GASetting) SetCampaignName(campaignName string) *GASetting {
	g.CampaignName = campaignName
	return g
}

// SetCampaignMedium ...
func (g *GASetting) SetCampaignMedium(campaignMedium string) *GASetting {
	g.CampaignMedium = campaignMedium
	return g
}
