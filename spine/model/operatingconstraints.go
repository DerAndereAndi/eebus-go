package model

type OperatingConstraintsInterruptDataType struct {
	SequenceId                  *PowerSequenceIdType `json:"sequenceId,omitempty"`
	IsPausable                  *bool                `json:"isPausable,omitempty"`
	IsStoppable                 *bool                `json:"isStoppable,omitempty"`
	NotInterruptibleAtHighPower *bool                `json:"notInterruptibleAtHighPower,omitempty"`
	MaxCyclesPerDay             *uint                `json:"maxCyclesPerDay,omitempty"`
}

type OperatingConstraintsInterruptDataElementsType struct {
	SequenceId                  *ElementTagType `json:"sequenceId,omitempty"`
	IsPausable                  *ElementTagType `json:"isPausable,omitempty"`
	IsStoppable                 *ElementTagType `json:"isStoppable,omitempty"`
	NotInterruptibleAtHighPower *ElementTagType `json:"notInterruptibleAtHighPower,omitempty"`
	MaxCyclesPerDay             *ElementTagType `json:"maxCyclesPerDay,omitempty"`
}

type OperatingConstraintsInterruptListDataType struct {
	OperatingConstraintsInterruptData []OperatingConstraintsInterruptDataType `json:"operatingConstraintsInterruptData,omitempty"`
}

type OperatingConstraintsInterruptListDataSelectorsType struct {
	SequenceId *PowerSequenceIdType `json:"sequenceId,omitempty"`
}

type OperatingConstraintsDurationDataType struct {
	SequenceId           *PowerSequenceIdType `json:"sequenceId,omitempty"`
	ActiveDurationMin    *string              `json:"activeDurationMin,omitempty"`
	ActiveDurationMax    *string              `json:"activeDurationMax,omitempty"`
	PauseDurationMin     *string              `json:"pauseDurationMin,omitempty"`
	PauseDurationMax     *string              `json:"pauseDurationMax,omitempty"`
	ActiveDurationSumMin *string              `json:"activeDurationSumMin,omitempty"`
	ActiveDurationSumMax *string              `json:"activeDurationSumMax,omitempty"`
}

type OperatingConstraintsDurationDataElementsType struct {
	SequenceId           *ElementTagType `json:"sequenceId,omitempty"`
	ActiveDurationMin    *ElementTagType `json:"activeDurationMin,omitempty"`
	ActiveDurationMax    *ElementTagType `json:"activeDurationMax,omitempty"`
	PauseDurationMin     *ElementTagType `json:"pauseDurationMin,omitempty"`
	PauseDurationMax     *ElementTagType `json:"pauseDurationMax,omitempty"`
	ActiveDurationSumMin *ElementTagType `json:"activeDurationSumMin,omitempty"`
	ActiveDurationSumMax *ElementTagType `json:"activeDurationSumMax,omitempty"`
}

type OperatingConstraintsDurationListDataType struct {
	OperatingConstraintsDurationData []OperatingConstraintsDurationDataType `json:"operatingConstraintsDurationData,omitempty"`
}

type OperatingConstraintsDurationListDataSelectorsType struct {
	SequenceId *PowerSequenceIdType `json:"sequenceId,omitempty"`
}

type OperatingConstraintsPowerDescriptionDataType struct {
	SequenceId              *PowerSequenceIdType   `json:"sequenceId,omitempty"`
	PositiveEnergyDirection *EnergyDirectionType   `json:"positiveEnergyDirection,omitempty"`
	PowerUnit               *UnitOfMeasurementType `json:"powerUnit,omitempty"`
	EnergyUnit              *UnitOfMeasurementType `json:"energyUnit,omitempty"`
	Description             *DescriptionType       `json:"description,omitempty"`
}

type OperatingConstraintsPowerDescriptionDataElementsType struct {
	SequenceId              *ElementTagType `json:"sequenceId,omitempty"`
	PositiveEnergyDirection *ElementTagType `json:"positiveEnergyDirection,omitempty"`
	PowerUnit               *ElementTagType `json:"powerUnit,omitempty"`
	EnergyUnit              *ElementTagType `json:"energyUnit,omitempty"`
	Description             *ElementTagType `json:"description,omitempty"`
}

type OperatingConstraintsPowerDescriptionListDataType struct {
	OperatingConstraintsPowerDescriptionData []OperatingConstraintsPowerDescriptionDataType `json:"operatingConstraintsPowerDescriptionData,omitempty"`
}

type OperatingConstraintsPowerDescriptionListDataSelectorsType struct {
	SequenceId *PowerSequenceIdType `json:"sequenceId,omitempty"`
}

type OperatingConstraintsPowerRangeDataType struct {
	SequenceId *PowerSequenceIdType `json:"sequenceId,omitempty"`
	PowerMin   *ScaledNumberType    `json:"powerMin,omitempty"`
	PowerMax   *ScaledNumberType    `json:"powerMax,omitempty"`
	EnergyMin  *ScaledNumberType    `json:"energyMin,omitempty"`
	EnergyMax  *ScaledNumberType    `json:"energyMax,omitempty"`
}

type OperatingConstraintsPowerRangeDataElementsType struct {
	SequenceId *ElementTagType `json:"sequenceId,omitempty"`
	PowerMin   *ElementTagType `json:"powerMin,omitempty"`
	PowerMax   *ElementTagType `json:"powerMax,omitempty"`
	EnergyMin  *ElementTagType `json:"energyMin,omitempty"`
	EnergyMax  *ElementTagType `json:"energyMax,omitempty"`
}

type OperatingConstraintsPowerRangeListDataType struct {
	OperatingConstraintsPowerRangeData []OperatingConstraintsPowerRangeDataType `json:"operatingConstraintsPowerRangeData,omitempty"`
}

type OperatingConstraintsPowerRangeListDataSelectorsType struct {
	SequenceId *PowerSequenceIdType `json:"sequenceId,omitempty"`
}

type OperatingConstraintsPowerLevelDataType struct {
	SequenceId *PowerSequenceIdType `json:"sequenceId,omitempty"`
	Power      *ScaledNumberType    `json:"power,omitempty"`
}

type OperatingConstraintsPowerLevelDataElementsType struct {
	SequenceId *ElementTagType `json:"sequenceId,omitempty"`
	Power      *ElementTagType `json:"power,omitempty"`
}

type OperatingConstraintsPowerLevelListDataType struct {
	OperatingConstraintsPowerLevelData []OperatingConstraintsPowerLevelDataType `json:"operatingConstraintsPowerLevelData,omitempty"`
}

type OperatingConstraintsPowerLevelListDataSelectorsType struct {
	SequenceId *PowerSequenceIdType `json:"sequenceId,omitempty"`
}

type OperatingConstraintsResumeImplicationDataType struct {
	SequenceId            *PowerSequenceIdType   `json:"sequenceId,omitempty"`
	ResumeEnergyEstimated *ScaledNumberType      `json:"resumeEnergyEstimated,omitempty"`
	EnergyUnit            *UnitOfMeasurementType `json:"energyUnit,omitempty"`
	ResumeCostEstimated   *ScaledNumberType      `json:"resumeCostEstimated,omitempty"`
	Currency              *CurrencyType          `json:"currency,omitempty"`
}

type OperatingConstraintsResumeImplicationDataElementsType struct {
	SequenceId            *ElementTagType           `json:"sequenceId,omitempty"`
	ResumeEnergyEstimated *ScaledNumberElementsType `json:"resumeEnergyEstimated,omitempty"`
	EnergyUnit            *ElementTagType           `json:"energyUnit,omitempty"`
	ResumeCostEstimated   *ScaledNumberElementsType `json:"resumeCostEstimated,omitempty"`
	Currency              *ElementTagType           `json:"currency,omitempty"`
}

type OperatingConstraintsResumeImplicationListDataType struct {
	OperatingConstraintsResumeImplicationData []OperatingConstraintsResumeImplicationDataType `json:"operatingConstraintsResumeImplicationData,omitempty"`
}

type OperatingConstraintsResumeImplicationListDataSelectorsType struct {
	SequenceId *PowerSequenceIdType `json:"sequenceId,omitempty"`
}
