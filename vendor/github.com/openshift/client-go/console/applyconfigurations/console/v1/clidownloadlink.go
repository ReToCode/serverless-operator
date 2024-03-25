// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// CLIDownloadLinkApplyConfiguration represents an declarative configuration of the CLIDownloadLink type for use
// with apply.
type CLIDownloadLinkApplyConfiguration struct {
	Text *string `json:"text,omitempty"`
	Href *string `json:"href,omitempty"`
}

// CLIDownloadLinkApplyConfiguration constructs an declarative configuration of the CLIDownloadLink type for use with
// apply.
func CLIDownloadLink() *CLIDownloadLinkApplyConfiguration {
	return &CLIDownloadLinkApplyConfiguration{}
}

// WithText sets the Text field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Text field is set to the value of the last call.
func (b *CLIDownloadLinkApplyConfiguration) WithText(value string) *CLIDownloadLinkApplyConfiguration {
	b.Text = &value
	return b
}

// WithHref sets the Href field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Href field is set to the value of the last call.
func (b *CLIDownloadLinkApplyConfiguration) WithHref(value string) *CLIDownloadLinkApplyConfiguration {
	b.Href = &value
	return b
}
