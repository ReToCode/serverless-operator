// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// ConsoleQuickStartTaskApplyConfiguration represents an declarative configuration of the ConsoleQuickStartTask type for use
// with apply.
type ConsoleQuickStartTaskApplyConfiguration struct {
	Title       *string                                         `json:"title,omitempty"`
	Description *string                                         `json:"description,omitempty"`
	Review      *ConsoleQuickStartTaskReviewApplyConfiguration  `json:"review,omitempty"`
	Summary     *ConsoleQuickStartTaskSummaryApplyConfiguration `json:"summary,omitempty"`
}

// ConsoleQuickStartTaskApplyConfiguration constructs an declarative configuration of the ConsoleQuickStartTask type for use with
// apply.
func ConsoleQuickStartTask() *ConsoleQuickStartTaskApplyConfiguration {
	return &ConsoleQuickStartTaskApplyConfiguration{}
}

// WithTitle sets the Title field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Title field is set to the value of the last call.
func (b *ConsoleQuickStartTaskApplyConfiguration) WithTitle(value string) *ConsoleQuickStartTaskApplyConfiguration {
	b.Title = &value
	return b
}

// WithDescription sets the Description field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Description field is set to the value of the last call.
func (b *ConsoleQuickStartTaskApplyConfiguration) WithDescription(value string) *ConsoleQuickStartTaskApplyConfiguration {
	b.Description = &value
	return b
}

// WithReview sets the Review field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Review field is set to the value of the last call.
func (b *ConsoleQuickStartTaskApplyConfiguration) WithReview(value *ConsoleQuickStartTaskReviewApplyConfiguration) *ConsoleQuickStartTaskApplyConfiguration {
	b.Review = value
	return b
}

// WithSummary sets the Summary field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Summary field is set to the value of the last call.
func (b *ConsoleQuickStartTaskApplyConfiguration) WithSummary(value *ConsoleQuickStartTaskSummaryApplyConfiguration) *ConsoleQuickStartTaskApplyConfiguration {
	b.Summary = value
	return b
}
