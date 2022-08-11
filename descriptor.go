package crondescriptor

// DescriptorOption is options for parsing and describing a Cron Expression
type DescriptorOption struct {
	ThrowExceptionOnParseError bool
	CasingType                 string
	Verbose                    bool
	DayOfWeekStartIndexZero    bool
	Use24HourTimeFormat        bool
	Code                       string
	Encoding                   string
	TimeFormatFor24Hour        string
}

// Descriptor converts a Cron Expression into a human readable string
type Descriptor struct {
}

// NewDescriptor creates & initializes a new instance of Descriptor
// @expression: The cron expression string
// @opt: Options to control the output description
func NewDescriptor(expression string, opt DescriptorOption, args ...string) *Descriptor {
	return &Descriptor{}
}

// GetDescription generates a humanreadable string for the Cron Expression
// @descriptionType: Which part(s) of the expression to describe
// return: The cron expression description
func (d *Descriptor) GetDescription(descriptionType string) string {
	return ""
}

// GetFullDescription generates the FULL description
// return: the FULL description
func (d *Descriptor) GetFullDescription() {}

// GetTimeOfDayDescription generates a description for only the TIMEOFDAY portion of the expression
// return: TIMEOFDAY description
func (d *Descriptor) GetTimeOfDayDescription() {}

func (d *Descriptor) GetSecondsDescription() {}

func (d *Descriptor) GetMinutesDescription() {}

func (d *Descriptor) GetHoursDescription() {}

func (d *Descriptor) GetDayOfWeekDescription() {}

func (d *Descriptor) GetMonthDescription() {}

func (d *Descriptor) GetDayOfMonthDescription() {}

func (d *Descriptor) get_year_description() {}

func (d *Descriptor) get_segment_description() {}

func (d *Descriptor) generate_between_segment_description() {}

func (d *Descriptor) format_time() {}

func (d *Descriptor) transform_verbosity() {}

func (d *Descriptor) transform_case() {}

func (d *Descriptor) number_to_day() {

}

// GetDescription generates a human readable string for the Cron Expression
func GetDescription(expression string, opts DescriptorOption) string {
	return NewDescriptor(expression, opts).GetDescription("")
}
