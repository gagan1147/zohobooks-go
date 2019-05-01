package zohobooks

import (
	"encoding/json"
)

// Contact struct represents the information of the contact
type Contact struct {
	Name         string `json:"contact_name"`
	Company      string `json:"company_name"`
	Website      string `json:"website"`
	LanguageCode string `json:"language_code"`
	ContactType  string `json:"contact_type"`
	Notes        string `json:"notes"`

	BillingAddress  billingAddress `json:"billing_address"`
	ShippingAddress billingAddress `json:"shipping_address"`

	// possible values ---> vat_registered,vat_not_registered,gcc_vat_not_registered,gcc_vat_registered,non_gcc,dz_vat_registered and dz_vat_not_registered.
	TaxTreatment string `json:"tax_treatment"`
	GstNO        string `json:"gst_no"`        // 15 digit
	GstTreatment string `json:"gst_treatment"` // Allowed values are business_gst , business_none , overseas , consumer
	TaxID        string `json:"tax_id"`
	IsTaxable    bool   `json:"is_taxable"`
	CreatedTime  string `json:"created_time"`

	LastModifiedTime string `json:"last_modified_time"`
}

type billingAddress struct {
	Attention string `json:"attention"`
	Address   string `json:"address"`
	Street2   string `json:"street2"`
	StateCode string `json:"state_code"`
	City      string `json:"city"`
	State     string `json:"state"`
	Zip       string `json:"zip"`
	Country   string `json:"country"`
	Fax       string `json:"fax"`
	Phone     string `json:"phone"`
}

// CustomerParams struct represents the information to create a contact
type CustomerParams struct {
	Name         string `json:"contact_name"`
	Company      string `json:"company_name"`
	Website      string `json:"website"`
	LanguageCode string `json:"language_code"`
	ContactType  string `json:"contact_type"`
	Notes        string `json:"notes"`

	BillingAddress  billingAddress `json:"billing_address"`
	ShippingAddress billingAddress `json:"shipping_address"`

	// possible values ---> vat_registered,vat_not_registered,gcc_vat_not_registered,gcc_vat_registered,non_gcc,dz_vat_registered and dz_vat_not_registered.
	TaxTreatment string `json:"tax_treatment"`
	GstNO        string `json:"gst_no"`        // 15 digit
	GstTreatment string `json:"gst_treatment"` // Allowed values are business_gst , business_none , overseas , consumer
	TaxID        string `json:"tax_id"`
	IsTaxable    bool   `json:"is_taxable"`
}

// New method will create a contact object and return a pointer to it
func (c *Contact) New() Resource {
	var obj = &Contact{}
	return obj
}

// Endpoint method returns the endpoint of the resource
func (c *Contact) Endpoint() string {
	return "/contacts"
}

// Create method will try to create a contact on razorpay
func (c *Contact) Create(params *CustomerParams, client *Client) (Contact, error) {
	var body, _ = json.Marshal(params)
	resp, err := client.Post(c.Endpoint(), string(body))

	respData, err := sendResp(resp, err, c)
	if err != nil {
		return *c, err
	}
	return respData.Contact, err
}

// FindOne tries to find the contact with given id
func (c *Contact) FindOne(id string, client *Client) (Contact, error) {
	resp, err := client.Get(c.Endpoint() + "/" + id)
	respData, err := sendResp(resp, err, c)
	if err != nil {
		return *c, err
	}
	return respData.Contact, err
}
