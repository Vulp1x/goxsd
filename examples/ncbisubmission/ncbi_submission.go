// generated by goxsd; DO NOT EDIT

package ncbisubmission

import "time"

type (
	AnyURIXML             string
	BooleanXML            bool
	DateTimeXSD           time.Time
	DateXSD               time.Time
	DecimalXML            float64
	DurationXML           time.Duration
	Float64XML            float64
	IntXML                int
	IntegerXML            int
	LanguageXML           string
	LongIntXML            int64
	NameXML               string
	NonNegativeIntegerXML uint
	NormalizedStringXML   string
	PositiveIntegerXML    uint
	ShortIntXML           int16
	StringXML             string
	TimeXSD               time.Time
	TokenXML              string
	UnsignedShortIntXML   uint16
)

// Submission is generated from an XSD element.
type Submission struct {
	SchemaVersion StringXML   `xml:"schema_version,omitempty,attr"`
	ResubmitOf    StringXML   `xml:"resubmit_of,omitempty,attr"`
	Submitted     DateXSD     `xml:"submitted,omitempty,attr"`
	LastUpdate    DateXSD     `xml:"last_update,omitempty,attr"`
	Status        StringXML   `xml:"status,omitempty,attr"`
	SubmissionID  TokenXML    `xml:"submission_id,attr"`
	Description   Description `xml:"Description"`
	Action        []Action    `xml:"Action"`
}

// Description is generated from an XSD element.
type Description struct {
	Comment            *StringXML          `xml:"Comment,omitempty"`
	Submitter          *TypeAccount        `xml:"Submitter,omitempty"`
	Organization       []TypeOrganization  `xml:"Organization"`
	Hold               *Hold               `xml:"Hold,omitempty"`
	SubmissionSoftware *SubmissionSoftware `xml:"SubmissionSoftware,omitempty"`
}

// TypeAccount is generated from an XSD element.
type TypeAccount struct {
	AccountID StringXML        `xml:"account_id,omitempty,attr"`
	UserName  StringXML        `xml:"user_name,omitempty,attr"`
	Authority StringXML        `xml:"authority,omitempty,attr"`
	Contact   *TypeContactInfo `xml:"Contact,omitempty"`
}

// TypeContactInfo is generated from an XSD element.
type TypeContactInfo struct {
	Email    StringXML    `xml:"email,attr"`
	SecEmail StringXML    `xml:"sec_email,omitempty,attr"`
	Phone    StringXML    `xml:"phone,attr"`
	Fax      StringXML    `xml:"fax,attr"`
	Address  *TypeAddress `xml:"Address,omitempty"`
	Name     *TypeName    `xml:"Name,omitempty"`
}

// TypeAddress is generated from an XSD element.
type TypeAddress struct {
	PostalCode  StringXML  `xml:"postal_code,omitempty,attr"`
	Department  *StringXML `xml:"Department,omitempty"`
	Institution *StringXML `xml:"Institution,omitempty"`
	Street      *StringXML `xml:"Street,omitempty"`
	City        StringXML  `xml:"City"`
	Sub         *StringXML `xml:"Sub,omitempty"`
	Country     StringXML  `xml:"Country"`
}

// TypeName is generated from an XSD element.
type TypeName struct {
	First  *StringXML `xml:"First,omitempty"`
	Last   StringXML  `xml:"Last"`
	Middle *StringXML `xml:"Middle,omitempty"`
	Suffix *StringXML `xml:"Suffix,omitempty"`
}

// TypeOrganization is generated from an XSD element.
type TypeOrganization struct {
	Type    StringXML          `xml:"type,attr"`
	Role    StringXML          `xml:"role,omitempty,attr"`
	OrgID   PositiveIntegerXML `xml:"org_id,omitempty,attr"`
	URL     StringXML          `xml:"url,omitempty,attr"`
	GroupID StringXML          `xml:"group_id,omitempty,attr"`
	Name    Name               `xml:"Name"`
}

// Name is generated from an XSD element.
type Name struct {
	Abbr StringXML `xml:"abbr,omitempty,attr"`
	Name StringXML `xml:",cdata"`
}

// Hold is generated from an XSD element.
type Hold struct {
	ReleaseDate DateXSD `xml:"release_date,attr"`
}

// SubmissionSoftware is generated from an XSD element.
type SubmissionSoftware struct {
	Version string `xml:"version,omitempty,attr"`
}

// Action is generated from an XSD element.
type Action struct {
	ActionID            TokenXML      `xml:"action_id,attr"`
	SubmitterTrackingID StringXML     `xml:"submitter_tracking_id,attr"`
	AddFiles            *AddFiles     `xml:"AddFiles,omitempty"`
	AddData             *AddData      `xml:"AddData,omitempty"`
	ChangeStatus        *ChangeStatus `xml:"ChangeStatus,omitempty"`
}

// AddFiles is generated from an XSD element.
type AddFiles struct {
	File           []File                  `xml:"File"`
	Status         *TypeReleaseStatus      `xml:"Status,omitempty"`
	IDentifier     *TypeIDentifier         `xml:"Identifier,omitempty"`
	Attribute      *Attribute              `xml:"Attribute,omitempty"`
	Meta           *TypeInlineData         `xml:"Meta,omitempty"`
	AttributeRefID *TypeFileAttributeRefID `xml:"AttributeRefId,omitempty"`
	SequenceData   *TypeSequenceData       `xml:"SequenceData,omitempty"`
}

// File is generated from an XSD element.
type File struct {
	FilePath      StringXML `xml:"file_path,omitempty,attr"`
	FileID        StringXML `xml:"file_id,omitempty,attr"`
	Md5           StringXML `xml:"md5,omitempty,attr"`
	Crc32         StringXML `xml:"crc32,attr"`
	ContentType   StringXML `xml:"content_type,attr"`
	TargetDbLabel string    `xml:"target_db_label,omitempty,attr"`
	DataType      StringXML `xml:"DataType"`
}

// TypeReleaseStatus is generated from an XSD element.
type TypeReleaseStatus struct {
	Release        *Release        `xml:"Release,omitempty"`
	SetReleaseDate *SetReleaseDate `xml:"SetReleaseDate,omitempty"`
}

// Release is generated from an XSD element.
type Release struct{}

// SetReleaseDate is generated from an XSD element.
type SetReleaseDate struct {
	ReleaseDate DateXSD `xml:"release_date,attr"`
}

// TypeIDentifier is generated from an XSD element.
type TypeIDentifier struct {
	PrimaryID *PrimaryID `xml:"PrimaryId,omitempty"`
	SPUID     *SPUID     `xml:"SPUID,omitempty"`
	LocalID   *LocalID   `xml:"LocalId,omitempty"`
}

// PrimaryID is generated from an XSD element.
type PrimaryID struct {
	Db        StringXML  `xml:"db,omitempty,attr"`
	ID        IntegerXML `xml:"id,omitempty,attr"`
	PrimaryID StringXML  `xml:",cdata"`
}

// SPUID is generated from an XSD element.
type SPUID struct {
	SubmitterID    StringXML `xml:"submitter_id,omitempty,attr"`
	SpuidNamespace StringXML `xml:"spuid_namespace,omitempty,attr"`
	SPUID          StringXML `xml:",cdata"`
}

// LocalID is generated from an XSD element.
type LocalID struct {
	SubmissionID StringXML `xml:"submission_id,omitempty,attr"`
	LocalID      StringXML `xml:",cdata"`
}

// Attribute is generated from an XSD element.
type Attribute struct {
	Name      StringXML `xml:"name,attr"`
	Attribute StringXML `xml:",cdata"`
}

// TypeInlineData is generated from an XSD element.
type TypeInlineData struct {
	Name            StringXML      `xml:"name,omitempty,attr"`
	DataModel       StringXML      `xml:"data_model,omitempty,attr"`
	ContentType     StringXML      `xml:"content_type,attr"`
	ContentEncoding StringXML      `xml:"content_encoding,omitempty,attr"`
	XMLContent      *XMLContent    `xml:"XmlContent,omitempty"`
	DataContent     *StringXML     `xml:"DataContent,omitempty"`
	Project         *TypeProject   `xml:"Project,omitempty"`
	BioSample       *TypeBioSample `xml:"BioSample,omitempty"`
	Genome          *TypeGenome    `xml:"Genome,omitempty"`
}

// XMLContent is generated from an XSD element.
type XMLContent struct{}

// TypeProject is generated from an XSD element.
type TypeProject struct {
	SchemaVersion StringXML      `xml:"schema_version,attr"`
	ProjectID     TypeIDentifier `xml:"ProjectID"`
	Descriptor    Descriptor     `xml:"Descriptor"`
	ProjectType   ProjectType    `xml:"ProjectType"`
}

// Descriptor is generated from an XSD element.
type Descriptor struct {
	Title        *StringXML         `xml:"Title,omitempty"`
	Description  *TypeBlock         `xml:"Description,omitempty"`
	ExternalLink []TypeExternalLink `xml:"ExternalLink,omitempty"`
	Name         *StringXML         `xml:"Name,omitempty"`
	Grant        []Grant            `xml:"Grant,omitempty"`
	Relevance    *Relevance         `xml:"Relevance,omitempty"`
	Publication  []TypePublication  `xml:"Publication,omitempty"`
	Keyword      []StringXML        `xml:"Keyword,omitempty"`
	UserTerm     []UserTerm         `xml:"UserTerm,omitempty"`
}

// TypeBlock is generated from an XSD element.
type TypeBlock struct {
	P     *TypeInline `xml:"p,omitempty"`
	Ul    *TypeL      `xml:"ul,omitempty"`
	Ol    *TypeL      `xml:"ol,omitempty"`
	Table *TypeTable  `xml:"table,omitempty"`
}

// TypeInline is generated from an XSD element.
type TypeInline struct {
	A *TypeA `xml:"a,omitempty"`
}

// TypeA is generated from an XSD element.
type TypeA struct {
	Href AnyURIXML   `xml:"href,attr"`
	Type StringXML   `xml:"type,attr"`
	I    *TypeInline `xml:"i,omitempty"`
}

// TypeL is generated from an XSD element.
type TypeL struct {
	Li []TypeLI `xml:"li"`
}

// TypeLI is generated from an XSD element.
type TypeLI struct{}

// TypeTable is generated from an XSD element.
type TypeTable struct {
	Caption *TypeCaption `xml:"caption,omitempty"`
	Tr      []TypeTR     `xml:"tr"`
}

// TypeCaption is generated from an XSD element.
type TypeCaption struct{}

// TypeTR is generated from an XSD element.
type TypeTR struct {
	Th *TypeTH `xml:"th,omitempty"`
	Td *TypeTD `xml:"td,omitempty"`
}

// TypeTH is generated from an XSD element.
type TypeTH struct {
	Rowspan NonNegativeIntegerXML `xml:"rowspan,attr"`
	Colspan NonNegativeIntegerXML `xml:"colspan,attr"`
}

// TypeTD is generated from an XSD element.
type TypeTD struct {
	Rowspan NonNegativeIntegerXML `xml:"rowspan,attr"`
	Colspan NonNegativeIntegerXML `xml:"colspan,attr"`
}

// TypeExternalLink is generated from an XSD element.
type TypeExternalLink struct {
	Label       StringXML   `xml:"label,attr"`
	Category    StringXML   `xml:"category,omitempty,attr"`
	URL         *AnyURIXML  `xml:"URL,omitempty"`
	ExternalID  *ExternalID `xml:"ExternalID,omitempty"`
	EntrezQuery *StringXML  `xml:"EntrezQuery,omitempty"`
}

// ExternalID is generated from an XSD element.
type ExternalID struct {
	Db         StringXML  `xml:"db,omitempty,attr"`
	ID         IntegerXML `xml:"id,omitempty,attr"`
	ExternalID StringXML  `xml:",cdata"`
}

// Grant is generated from an XSD element.
type Grant struct {
	GrantID StringXML `xml:"GrantId,attr"`
	Agency  Agency    `xml:"Agency"`
	PI      []PI      `xml:"PI,omitempty"`
}

// Agency is generated from an XSD element.
type Agency struct {
	Abbr   StringXML `xml:"abbr,attr"`
	Agency StringXML `xml:",cdata"`
}

// PI is generated from an XSD element.
type PI struct {
	Auth        StringXML  `xml:"auth,omitempty,attr"`
	Userid      StringXML  `xml:"userid,omitempty,attr"`
	GrantUserID StringXML  `xml:"grant_user_id,omitempty,attr"`
	Affil       StringXML  `xml:"affil,omitempty,attr"`
	Given       *StringXML `xml:"Given,omitempty"`
}

// Relevance is generated from an XSD element.
type Relevance struct {
	Agricultural  *StringXML `xml:"Agricultural,omitempty"`
	Medical       *StringXML `xml:"Medical,omitempty"`
	Industrial    *StringXML `xml:"Industrial,omitempty"`
	Environmental *StringXML `xml:"Environmental,omitempty"`
	Evolution     *StringXML `xml:"Evolution,omitempty"`
	ModelOrganism *StringXML `xml:"ModelOrganism,omitempty"`
	Other         *StringXML `xml:"Other,omitempty"`
}

// TypePublication is generated from an XSD element.
type TypePublication struct {
	ID                 StringXML           `xml:"id,attr"`
	Date               DateTimeXSD         `xml:"date,attr"`
	Status             TokenXML            `xml:"status,attr"`
	AuthorSet          *TypeAuthorSet      `xml:"AuthorSet,omitempty"`
	Reference          *StringXML          `xml:"Reference,omitempty"`
	StructuredCitation *StructuredCitation `xml:"StructuredCitation,omitempty"`
	DbType             TokenXML            `xml:"DbType"`
}

// TypeAuthorSet is generated from an XSD element.
type TypeAuthorSet struct {
	Author []Author `xml:"Author"`
}

// Author is generated from an XSD element.
type Author struct {
	Consortium *StringXML `xml:"Consortium,omitempty"`
}

// StructuredCitation is generated from an XSD element.
type StructuredCitation struct {
	Journal *Journal `xml:"Journal,omitempty"`
}

// Journal is generated from an XSD element.
type Journal struct {
	JournalTitle StringXML  `xml:"JournalTitle"`
	Year         *StringXML `xml:"Year,omitempty"`
	Volume       *StringXML `xml:"Volume,omitempty"`
	Issue        *StringXML `xml:"Issue,omitempty"`
	PagesFrom    *StringXML `xml:"PagesFrom,omitempty"`
	PagesTo      *StringXML `xml:"PagesTo,omitempty"`
}

// UserTerm is generated from an XSD element.
type UserTerm struct {
	Term     StringXML `xml:"term,attr"`
	Category StringXML `xml:"category,omitempty,attr"`
	Units    StringXML `xml:"units,omitempty,attr"`
	UserTerm StringXML `xml:",cdata"`
}

// ProjectType is generated from an XSD element.
type ProjectType struct {
	ProjectTypeTopAdmin   *ProjectTypeTopAdmin   `xml:"ProjectTypeTopAdmin,omitempty"`
	ProjectTypeSubmission *ProjectTypeSubmission `xml:"ProjectTypeSubmission,omitempty"`
}

// ProjectTypeTopAdmin is generated from an XSD element.
type ProjectTypeTopAdmin struct {
	Subtype                 TokenXML      `xml:"subtype,attr"`
	Organism                *TypeOrganism `xml:"Organism,omitempty"`
	DescriptionSubtypeOther *StringXML    `xml:"DescriptionSubtypeOther,omitempty"`
}

// TypeOrganism is generated from an XSD element.
type TypeOrganism struct {
	TaxonomyID   IntXML     `xml:"taxonomy_id,omitempty,attr"`
	OrganismName StringXML  `xml:"OrganismName"`
	Label        *StringXML `xml:"Label,omitempty"`
	Strain       *StringXML `xml:"Strain,omitempty"`
	IsolateName  *StringXML `xml:"IsolateName,omitempty"`
	Breed        *StringXML `xml:"Breed,omitempty"`
	Cultivar     *StringXML `xml:"Cultivar,omitempty"`
}

// ProjectTypeSubmission is generated from an XSD element.
type ProjectTypeSubmission struct {
	SampleScope         TokenXML            `xml:"sample_scope,attr"`
	BioSampleSet        *BioSampleSet       `xml:"BioSampleSet,omitempty"`
	IntendedDataTypeSet IntendedDataTypeSet `xml:"IntendedDataTypeSet"`
}

// BioSampleSet is generated from an XSD element.
type BioSampleSet struct {
	BioSample []BioSample `xml:"BioSample"`
}

// BioSample is generated from an XSD element.
type BioSample struct {
	LocusTagPrefix TokenXML `xml:"LocusTagPrefix,omitempty,attr"`
}

// IntendedDataTypeSet is generated from an XSD element.
type IntendedDataTypeSet struct{}

// TypeBioSample is generated from an XSD element.
type TypeBioSample struct {
	SchemaVersion StringXML               `xml:"schema_version,attr"`
	SampleID      TypeBioSampleIDentifier `xml:"SampleId"`
	Descriptor    TypeDescriptor          `xml:"Descriptor"`
	BioProject    []TypeRefID             `xml:"BioProject,omitempty"`
	Package       StringXML               `xml:"Package"`
	Attributes    Attributes              `xml:"Attributes"`
}

// TypeBioSampleIDentifier is generated from an XSD element.
type TypeBioSampleIDentifier struct {
	SPUID []SPUID `xml:"SPUID,omitempty"`
}

// TypeDescriptor is generated from an XSD element.
type TypeDescriptor struct{}

// TypeRefID is generated from an XSD element.
type TypeRefID struct{}

// Attributes is generated from an XSD element.
type Attributes struct {
	Attribute []Attribute `xml:"Attribute,omitempty"`
}

// TypeGenome is generated from an XSD element.
type TypeGenome struct {
	Chromosomes *Chromosomes `xml:"Chromosomes,omitempty"`
	Unplaced    *Unplaced    `xml:"Unplaced,omitempty"`
}

// Chromosomes is generated from an XSD element.
type Chromosomes struct {
	Chromosome []Chromosome `xml:"Chromosome"`
}

// Chromosome is generated from an XSD element.
type Chromosome struct {
	Type                          StringXML              `xml:"type,attr"`
	ChromosomeName                StringXML              `xml:"chromosome_name,attr"`
	TheChromosomeSequence         *TheChromosomeSequence `xml:"TheChromosomeSequence,omitempty"`
	UnlocalizedChromosomeSequence []StringXML            `xml:"UnlocalizedChromosomeSequence,omitempty"`
}

// TheChromosomeSequence is generated from an XSD element.
type TheChromosomeSequence struct {
	Circular              TokenXML  `xml:"circular,attr"`
	EndsAbut              TokenXML  `xml:"ends_abut,omitempty,attr"`
	TheChromosomeSequence StringXML `xml:",cdata"`
}

// Unplaced is generated from an XSD element.
type Unplaced struct {
	UnplacedSequence []StringXML `xml:"UnplacedSequence"`
}

// TypeFileAttributeRefID is generated from an XSD element.
type TypeFileAttributeRefID struct {
	Name  StringXML `xml:"name,attr"`
	RefID TypeRefID `xml:"RefId"`
}

// TypeSequenceData is generated from an XSD element.
type TypeSequenceData struct {
	Sequence []Sequence `xml:"Sequence,omitempty"`
}

// Sequence is generated from an XSD element.
type Sequence struct {
	ID      StringXML `xml:"id,attr"`
	Type    StringXML `xml:"type,attr"`
	OnlyOne StringXML `xml:"only_one,omitempty,attr"`
}

// AddData is generated from an XSD element.
type AddData struct {
	Data Data `xml:"Data"`
}

// Data is generated from an XSD element.
type Data struct {
	Name            StringXML `xml:"name,omitempty,attr"`
	DataModel       StringXML `xml:"data_model,omitempty,attr"`
	ContentType     StringXML `xml:"content_type,attr"`
	ContentEncoding StringXML `xml:"content_encoding,omitempty,attr"`
	TargetDbLabel   string    `xml:"target_db_label,omitempty,attr"`
}

// ChangeStatus is generated from an XSD element.
type ChangeStatus struct {
	Target     TypeRefID  `xml:"Target"`
	Suppress   *StringXML `xml:"Suppress,omitempty"`
	Withdraw   *StringXML `xml:"Withdraw,omitempty"`
	AddComment *StringXML `xml:"AddComment,omitempty"`
}
