/* goslacker is a quick and dirty repo to allow easy access to a number of slack messaging
types which I find myself using frequently.  It is not feature complete and is a best effort.
Provided as-is with no guarantee of anything at all

 */
package goslacker

type Attachment struct {
	Fallback string `json:"fallback"`
	Color string `json:"color"`
	Pretext string `json:"pretext"`
	AuthorName string `json:"author_name"`
	AuthorLink string `json:"author_link"`
	AuthorIcon string `json:"author_icon"`
	Title string `json:"title"`
	TitleLink string `json:"title_link"`
	Text string `json:"text"`
	Fields []SlackField `json:"fields"`
	ImageURL string `json:"image_url"`
	ThumbURL string `json:"thumb_url"`
	Footer string `json:"footer"`
	FooterIcon string `json:"footer_icon"`
	TS int `json:"ts"`
}

type SlackField struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool `json:"short"`
}


type AttachmentMessage struct {
	Attachments []Attachment `json:"attachments"`

}

type SlackIncomingWebhook struct {
	Text string `json:"text"`
}

func NewAttachment(author, text, title, titleLink string) Attachment {
	a := Attachment{
		AuthorName: author,
		Pretext: text,
		Fallback: text,
		Text: text,
		Title: title,
		TitleLink: titleLink,
		Color: "#36a64f",
	}
	return a
}

func NewAttachmentMessage(author, text, title, titleLink string) AttachmentMessage {
	return AttachmentMessage{
		Attachments: []Attachment{NewAttachment(author, text, title, titleLink)},
	}
}