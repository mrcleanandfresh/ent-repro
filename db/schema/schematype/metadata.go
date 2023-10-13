package schematype

type ChoiceMeta struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type Metadata struct {
	Type             string        `json:"type,omitempty"`
	Name             string        `json:"name,omitempty"`             // Essay,KeyValue,
	MaxCharacters    int           `json:"maxCharacters,omitempty"`    // Essay
	KeyPlaceholder   *string       `json:"keyPlaceholder,omitempty"`   // KeyValue
	ValuePlaceholder *string       `json:"valuePlaceholder,omitempty"` // KeyValue
	Multi            bool          `json:"multi,omitempty"`            // KeyValue
	MoreLabel        *string       `json:"moreLabel,omitempty"`        // KeyValue
	Choices          []*ChoiceMeta `json:"choices,omitempty"`          // MultipleChoice
	Placeholder      string        `json:"placeholder,omitempty"`      // Essay,NumberAnswer
}

func (Metadata) IsMeta() {}
