// generated by goxsd; DO NOT EDIT

package ncbisubmission

import "time"

// Submission is generated from an XSD element.
type Submission struct {
	SchemaVersion string      `xml:"schema_version,attr"`
	ResubmitOf    string      `xml:"resubmit_of,attr"`
	Submitted     time.Time   `xml:"submitted,attr"`
	LastUpdate    time.Time   `xml:"last_update,attr"`
	Status        string      `xml:"status,attr"`
	SubmissionID  string      `xml:"submission_id,attr"`
	Description   Description `xml:"Description"`
	Action        []Action    `xml:"Action"`
}

// Description is generated from an XSD element.
type Description struct {
	Comment            string             `xml:"Comment"`
	Submitter          Submitter          `xml:"Submitter"`
	Organization       []Organization     `xml:"Organization"`
	Hold               Hold               `xml:"Hold"`
	SubmissionSoftware SubmissionSoftware `xml:"SubmissionSoftware"`
}

// Submitter is generated from an XSD element.
type Submitter struct {
	AccountID string  `xml:"account_id,attr"`
	UserName  string  `xml:"user_name,attr"`
	Authority string  `xml:"authority,attr"`
	Contact   Contact `xml:"Contact"`
}

// Contact is generated from an XSD element.
type Contact struct {
	Email    string  `xml:"email,attr"`
	SecEmail string  `xml:"sec_email,attr"`
	Phone    string  `xml:"phone,attr"`
	Fax      string  `xml:"fax,attr"`
	Address  Address `xml:"Address"`
	Name     Name    `xml:"Name"`
}

// Address is generated from an XSD element.
type Address struct {
	PostalCode  string `xml:"postal_code,attr"`
	Department  string `xml:"Department"`
	Institution string `xml:"Institution"`
	Street      string `xml:"Street"`
	City        string `xml:"City"`
	Sub         string `xml:"Sub"`
	Country     string `xml:"Country"`
}

// Name is generated from an XSD element.
type Name struct {
	First  string `xml:"First"`
	Last   string `xml:"Last"`
	Middle string `xml:"Middle"`
	Suffix string `xml:"Suffix"`
}

// Organization is generated from an XSD element.
type Organization struct {
	Type    string    `xml:"type,attr"`
	Role    string    `xml:"role,attr"`
	OrgID   uint      `xml:"org_id,attr"`
	URL     string    `xml:"url,attr"`
	GroupID string    `xml:"group_id,attr"`
	Name    Name      `xml:"Name"`
	Address Address   `xml:"Address"`
	Contact []Contact `xml:"Contact"`
}

// Hold is generated from an XSD element.
type Hold struct {
	ReleaseDate time.Time `xml:"release_date,attr"`
}

// SubmissionSoftware is generated from an XSD element.
type SubmissionSoftware struct {
	Version string `xml:"version,attr"`
}

// Action is generated from an XSD element.
type Action struct {
	ActionID            string `xml:"action_id,attr"`
	SubmitterTrackingID string `xml:"submitter_tracking_id,attr"`
}
