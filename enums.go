package crondescriptor

// DescriptionTypeEnum
type DescriptionTypeEnum uint8

const (
	DescriptionTypeEnumFULL       DescriptionTypeEnum = 1
	DescriptionTypeEnumTIMEOFDAY  DescriptionTypeEnum = 2
	DescriptionTypeEnumSECONDS    DescriptionTypeEnum = 3
	DescriptionTypeEnumMINUTES    DescriptionTypeEnum = 4
	DescriptionTypeEnumHOURS      DescriptionTypeEnum = 5
	DescriptionTypeEnumDAYOFWEEK  DescriptionTypeEnum = 6
	DescriptionTypeEnumMONTH      DescriptionTypeEnum = 7
	DescriptionTypeEnumDAYOFMONTH DescriptionTypeEnum = 8
	DescriptionTypeEnumYEAR       DescriptionTypeEnum = 9
)

// CasingTypeEnum is enum to define the casing types for the Cron Expression description
type CasingTypeEnum uint8

const (
	CasingTypeEnumTitle     CasingTypeEnum = 1
	CasingTypeEnumSentence  CasingTypeEnum = 2
	CasingTypeEnumLowerCase CasingTypeEnum = 3
)
