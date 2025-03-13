// This code was AUTOGENERATED using the library.
// Please DO NOT EDIT THIS FILE.

package token_metadata

import (
	"errors"

	binary "github.com/gagliardetto/binary"
	common "github.com/gagliardetto/solana-go/programs/common"
	format "github.com/gagliardetto/solana-go/text/format"
	treeout "github.com/gagliardetto/treeout"
)

// Initialize Instruction
type Initialize struct {
	// Longer name of the token
	Name *string
	// Shortened symbol of the token
	Symbol *string
	// URI pointing to more metadata (image, video, etc.)
	Uri *string
	// [0] = [WRITE] metadata `Metadata`
	// [1] = [] updateAuthority `Update authority`
	// [2] = [] mint `Mint`
	// [3] = [SIGNER] mintAuthority `Mint authority`
	common.AccountMetaSlice `bin:"-"`
	_programId              *common.PublicKey
}

// NewInitializeInstructionBuilder creates a new `Initialize` instruction builder.
func NewInitializeInstructionBuilder() *Initialize {
	return &Initialize{
		AccountMetaSlice: make(common.AccountMetaSlice, 4),
	}
}

// NewInitializeInstruction
//
// Parameters:
//
//	name: Longer name of the token
//	symbol: Shortened symbol of the token
//	uri: URI pointing to more metadata (image, video, etc.)
//	metadata: Metadata
//	updateAuthority: Update authority
//	mint: Mint
//	mintAuthority: Mint authority
func NewInitializeInstruction(
	name string,
	symbol string,
	uri string,
	metadata common.PublicKey,
	updateAuthority common.PublicKey,
	mint common.PublicKey,
	mintAuthority common.PublicKey,
) *Initialize {
	return NewInitializeInstructionBuilder().
		SetName(name).
		SetSymbol(symbol).
		SetUri(uri).
		SetMetadataAccount(metadata).
		SetUpdateAuthorityAccount(updateAuthority).
		SetMintAccount(mint).
		SetMintAuthorityAccount(mintAuthority)
}

// SetName sets the "name" parameter.
func (obj *Initialize) SetName(name string) *Initialize {
	obj.Name = &name
	return obj
}

// SetSymbol sets the "symbol" parameter.
func (obj *Initialize) SetSymbol(symbol string) *Initialize {
	obj.Symbol = &symbol
	return obj
}

// SetUri sets the "uri" parameter.
func (obj *Initialize) SetUri(uri string) *Initialize {
	obj.Uri = &uri
	return obj
}

// SetMetadataAccount sets the "metadata" parameter.
// Metadata
func (obj *Initialize) SetMetadataAccount(metadata common.PublicKey) *Initialize {
	obj.AccountMetaSlice[0] = common.Meta(metadata).WRITE()
	return obj
}

// GetMetadataAccount gets the "metadata" parameter.
// Metadata
func (obj *Initialize) GetMetadataAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(0)
}

// SetUpdateAuthorityAccount sets the "updateAuthority" parameter.
// Update authority
func (obj *Initialize) SetUpdateAuthorityAccount(updateAuthority common.PublicKey) *Initialize {
	obj.AccountMetaSlice[1] = common.Meta(updateAuthority)
	return obj
}

// GetUpdateAuthorityAccount gets the "updateAuthority" parameter.
// Update authority
func (obj *Initialize) GetUpdateAuthorityAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(1)
}

// SetMintAccount sets the "mint" parameter.
// Mint
func (obj *Initialize) SetMintAccount(mint common.PublicKey) *Initialize {
	obj.AccountMetaSlice[2] = common.Meta(mint)
	return obj
}

// GetMintAccount gets the "mint" parameter.
// Mint
func (obj *Initialize) GetMintAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(2)
}

// SetMintAuthorityAccount sets the "mintAuthority" parameter.
// Mint authority
func (obj *Initialize) SetMintAuthorityAccount(mintAuthority common.PublicKey, multiSigners ...common.PublicKey) *Initialize {
	if len(multiSigners) > 0 {
		obj.AccountMetaSlice[3] = common.Meta(mintAuthority)
		for _, value := range multiSigners {
			obj.AccountMetaSlice.Append(common.NewAccountMeta(value, false, true))
		}
	} else {
		obj.AccountMetaSlice[3] = common.Meta(mintAuthority).SIGNER()
	}
	return obj
}

// GetMintAuthorityAccount gets the "mintAuthority" parameter.
// Mint authority
func (obj *Initialize) GetMintAuthorityAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(3)
}

func (obj *Initialize) SetProgramId(programId *common.PublicKey) *Initialize {
	obj._programId = programId
	return obj
}

func (obj *Initialize) Build() *Instruction {
	return &Instruction{
		BaseVariant: binary.BaseVariant{
			Impl:   obj,
			TypeID: binary.TypeIDFromBytes(Instruction_Initialize[:]),
		},
		programId: obj._programId,
		typeIdLen: 8,
	}
}

func (obj *Initialize) Validate() error {
	if obj.Name == nil {
		return errors.New("[Initialize] name param is not set")
	}
	if obj.Symbol == nil {
		return errors.New("[Initialize] symbol param is not set")
	}
	if obj.Uri == nil {
		return errors.New("[Initialize] uri param is not set")
	}

	if obj.AccountMetaSlice[0] == nil {
		return errors.New("[Initialize] accounts.metadata is not set")
	}
	if obj.AccountMetaSlice[1] == nil {
		return errors.New("[Initialize] accounts.updateAuthority is not set")
	}
	if obj.AccountMetaSlice[2] == nil {
		return errors.New("[Initialize] accounts.mint is not set")
	}
	if obj.AccountMetaSlice[3] == nil {
		return errors.New("[Initialize] accounts.mintAuthority is not set")
	}
	return nil
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (obj *Initialize) ValidateAndBuild() (*Instruction, error) {
	if err := obj.Validate(); err != nil {
		return nil, err
	}
	return obj.Build(), nil
}

func (obj *Initialize) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	if err = encoder.WriteBytes([]byte(*obj.Name), false); err != nil {
		return err
	}
	if err = encoder.WriteBytes([]byte(*obj.Symbol), false); err != nil {
		return err
	}
	if err = encoder.WriteBytes([]byte(*obj.Uri), false); err != nil {
		return err
	}
	return nil
}

func (obj *Initialize) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	if err = decoder.Decode(&obj.Name); err != nil {
		return err
	}
	if err = decoder.Decode(&obj.Symbol); err != nil {
		return err
	}
	if err = decoder.Decode(&obj.Uri); err != nil {
		return err
	}
	return nil
}

func (obj *Initialize) EncodeToTree(parent treeout.Branches) {
	parent.Child(format.Program(ProgramName, common.As(ProgramID))).
		ParentFunc(func(programBranch treeout.Branches) {
			programBranch.Child(format.Instruction("Initialize")).
				ParentFunc(func(instructionBranch treeout.Branches) {
					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch treeout.Branches) {
						paramsBranch.Child(format.Param("  Name", *obj.Name))
						paramsBranch.Child(format.Param("Symbol", *obj.Symbol))
						paramsBranch.Child(format.Param("   Uri", *obj.Uri))
					})
					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch treeout.Branches) {
						accountsBranch.Child(common.FormatMeta("       metadata", obj.AccountMetaSlice.Get(0)))
						accountsBranch.Child(common.FormatMeta("updateAuthority", obj.AccountMetaSlice.Get(1)))
						accountsBranch.Child(common.FormatMeta("           mint", obj.AccountMetaSlice.Get(2)))
						accountsBranch.Child(common.FormatMeta("  mintAuthority", obj.AccountMetaSlice.Get(3)))
					})
				})
		})
}

// UpdateField Instruction
type UpdateField struct {
	// Field to update in the metadata
	Field *Field
	// Value to write for the field
	Value *string
	// [0] = [WRITE] metadata `Metadata`
	// [1] = [SIGNER] updateAuthority `Update authority`
	common.AccountMetaSlice `bin:"-"`
	_programId              *common.PublicKey
}

// NewUpdateFieldInstructionBuilder creates a new `UpdateField` instruction builder.
func NewUpdateFieldInstructionBuilder() *UpdateField {
	return &UpdateField{
		AccountMetaSlice: make(common.AccountMetaSlice, 2),
	}
}

// NewUpdateFieldInstruction
//
// Parameters:
//
//	field: Field to update in the metadata
//	value: Value to write for the field
//	metadata: Metadata
//	updateAuthority: Update authority
func NewUpdateFieldInstruction(
	field Field,
	value string,
	metadata common.PublicKey,
	updateAuthority common.PublicKey,
) *UpdateField {
	return NewUpdateFieldInstructionBuilder().
		SetField(field).
		SetValue(value).
		SetMetadataAccount(metadata).
		SetUpdateAuthorityAccount(updateAuthority)
}

// SetField sets the "field" parameter.
func (obj *UpdateField) SetField(field Field) *UpdateField {
	obj.Field = &field
	return obj
}

// SetValue sets the "value" parameter.
func (obj *UpdateField) SetValue(value string) *UpdateField {
	obj.Value = &value
	return obj
}

// SetMetadataAccount sets the "metadata" parameter.
// Metadata
func (obj *UpdateField) SetMetadataAccount(metadata common.PublicKey) *UpdateField {
	obj.AccountMetaSlice[0] = common.Meta(metadata).WRITE()
	return obj
}

// GetMetadataAccount gets the "metadata" parameter.
// Metadata
func (obj *UpdateField) GetMetadataAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(0)
}

// SetUpdateAuthorityAccount sets the "updateAuthority" parameter.
// Update authority
func (obj *UpdateField) SetUpdateAuthorityAccount(updateAuthority common.PublicKey, multiSigners ...common.PublicKey) *UpdateField {
	if len(multiSigners) > 0 {
		obj.AccountMetaSlice[1] = common.Meta(updateAuthority)
		for _, value := range multiSigners {
			obj.AccountMetaSlice.Append(common.NewAccountMeta(value, false, true))
		}
	} else {
		obj.AccountMetaSlice[1] = common.Meta(updateAuthority).SIGNER()
	}
	return obj
}

// GetUpdateAuthorityAccount gets the "updateAuthority" parameter.
// Update authority
func (obj *UpdateField) GetUpdateAuthorityAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(1)
}

func (obj *UpdateField) SetProgramId(programId *common.PublicKey) *UpdateField {
	obj._programId = programId
	return obj
}

func (obj *UpdateField) Build() *Instruction {
	return &Instruction{
		BaseVariant: binary.BaseVariant{
			Impl:   obj,
			TypeID: binary.TypeIDFromBytes(Instruction_UpdateField[:]),
		},
		programId: obj._programId,
		typeIdLen: 8,
	}
}

func (obj *UpdateField) Validate() error {
	if obj.Field == nil {
		return errors.New("[UpdateField] field param is not set")
	}
	if obj.Value == nil {
		return errors.New("[UpdateField] value param is not set")
	}

	if obj.AccountMetaSlice[0] == nil {
		return errors.New("[UpdateField] accounts.metadata is not set")
	}
	if obj.AccountMetaSlice[1] == nil {
		return errors.New("[UpdateField] accounts.updateAuthority is not set")
	}
	return nil
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (obj *UpdateField) ValidateAndBuild() (*Instruction, error) {
	if err := obj.Validate(); err != nil {
		return nil, err
	}
	return obj.Build(), nil
}

func (obj *UpdateField) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	if err = encoder.Encode(&obj.Field); err != nil {
		return err
	}
	if err = encoder.WriteBytes([]byte(*obj.Value), false); err != nil {
		return err
	}
	return nil
}

func (obj *UpdateField) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	if err = decoder.Decode(&obj.Field); err != nil {
		return err
	}
	if err = decoder.Decode(&obj.Value); err != nil {
		return err
	}
	return nil
}

func (obj *UpdateField) EncodeToTree(parent treeout.Branches) {
	parent.Child(format.Program(ProgramName, common.As(ProgramID))).
		ParentFunc(func(programBranch treeout.Branches) {
			programBranch.Child(format.Instruction("UpdateField")).
				ParentFunc(func(instructionBranch treeout.Branches) {
					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch treeout.Branches) {
						paramsBranch.Child(format.Param("Field", *obj.Field))
						paramsBranch.Child(format.Param("Value", *obj.Value))
					})
					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch treeout.Branches) {
						accountsBranch.Child(common.FormatMeta("       metadata", obj.AccountMetaSlice.Get(0)))
						accountsBranch.Child(common.FormatMeta("updateAuthority", obj.AccountMetaSlice.Get(1)))
					})
				})
		})
}

// RemoveKey Instruction
type RemoveKey struct {
	// If the idempotent flag is set to true, then the instruction will not,error if the key does not exist
	Idempotent *bool
	// Key to remove in the additional metadata portion
	Key *string
	// [0] = [WRITE] metadata `Metadata`
	// [1] = [SIGNER] updateAuthority `Update authority`
	common.AccountMetaSlice `bin:"-"`
	_programId              *common.PublicKey
}

// NewRemoveKeyInstructionBuilder creates a new `RemoveKey` instruction builder.
func NewRemoveKeyInstructionBuilder() *RemoveKey {
	return &RemoveKey{
		AccountMetaSlice: make(common.AccountMetaSlice, 2),
	}
}

// NewRemoveKeyInstruction
//
// Parameters:
/*
  idempotent: If the idempotent flag is set to true, then the instruction will not
  error if the key does not exist
*/
//   key: Key to remove in the additional metadata portion
//   metadata: Metadata
//   updateAuthority: Update authority
//
func NewRemoveKeyInstruction(
	idempotent bool,
	key string,
	metadata common.PublicKey,
	updateAuthority common.PublicKey,
) *RemoveKey {
	return NewRemoveKeyInstructionBuilder().
		SetIdempotent(idempotent).
		SetKey(key).
		SetMetadataAccount(metadata).
		SetUpdateAuthorityAccount(updateAuthority)
}

// SetIdempotent sets the "idempotent" parameter.
func (obj *RemoveKey) SetIdempotent(idempotent bool) *RemoveKey {
	obj.Idempotent = &idempotent
	return obj
}

// SetKey sets the "key" parameter.
func (obj *RemoveKey) SetKey(key string) *RemoveKey {
	obj.Key = &key
	return obj
}

// SetMetadataAccount sets the "metadata" parameter.
// Metadata
func (obj *RemoveKey) SetMetadataAccount(metadata common.PublicKey) *RemoveKey {
	obj.AccountMetaSlice[0] = common.Meta(metadata).WRITE()
	return obj
}

// GetMetadataAccount gets the "metadata" parameter.
// Metadata
func (obj *RemoveKey) GetMetadataAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(0)
}

// SetUpdateAuthorityAccount sets the "updateAuthority" parameter.
// Update authority
func (obj *RemoveKey) SetUpdateAuthorityAccount(updateAuthority common.PublicKey, multiSigners ...common.PublicKey) *RemoveKey {
	if len(multiSigners) > 0 {
		obj.AccountMetaSlice[1] = common.Meta(updateAuthority)
		for _, value := range multiSigners {
			obj.AccountMetaSlice.Append(common.NewAccountMeta(value, false, true))
		}
	} else {
		obj.AccountMetaSlice[1] = common.Meta(updateAuthority).SIGNER()
	}
	return obj
}

// GetUpdateAuthorityAccount gets the "updateAuthority" parameter.
// Update authority
func (obj *RemoveKey) GetUpdateAuthorityAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(1)
}

func (obj *RemoveKey) SetProgramId(programId *common.PublicKey) *RemoveKey {
	obj._programId = programId
	return obj
}

func (obj *RemoveKey) Build() *Instruction {
	return &Instruction{
		BaseVariant: binary.BaseVariant{
			Impl:   obj,
			TypeID: binary.TypeIDFromBytes(Instruction_RemoveKey[:]),
		},
		programId: obj._programId,
		typeIdLen: 8,
	}
}

func (obj *RemoveKey) Validate() error {
	if obj.Idempotent == nil {
		return errors.New("[RemoveKey] idempotent param is not set")
	}
	if obj.Key == nil {
		return errors.New("[RemoveKey] key param is not set")
	}

	if obj.AccountMetaSlice[0] == nil {
		return errors.New("[RemoveKey] accounts.metadata is not set")
	}
	if obj.AccountMetaSlice[1] == nil {
		return errors.New("[RemoveKey] accounts.updateAuthority is not set")
	}
	return nil
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (obj *RemoveKey) ValidateAndBuild() (*Instruction, error) {
	if err := obj.Validate(); err != nil {
		return nil, err
	}
	return obj.Build(), nil
}

func (obj *RemoveKey) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	if err = encoder.Encode(&obj.Idempotent); err != nil {
		return err
	}
	if err = encoder.WriteBytes([]byte(*obj.Key), false); err != nil {
		return err
	}
	return nil
}

func (obj *RemoveKey) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	if err = decoder.Decode(&obj.Idempotent); err != nil {
		return err
	}
	if err = decoder.Decode(&obj.Key); err != nil {
		return err
	}
	return nil
}

func (obj *RemoveKey) EncodeToTree(parent treeout.Branches) {
	parent.Child(format.Program(ProgramName, common.As(ProgramID))).
		ParentFunc(func(programBranch treeout.Branches) {
			programBranch.Child(format.Instruction("RemoveKey")).
				ParentFunc(func(instructionBranch treeout.Branches) {
					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch treeout.Branches) {
						paramsBranch.Child(format.Param("Idempotent", *obj.Idempotent))
						paramsBranch.Child(format.Param("       Key", *obj.Key))
					})
					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch treeout.Branches) {
						accountsBranch.Child(common.FormatMeta("       metadata", obj.AccountMetaSlice.Get(0)))
						accountsBranch.Child(common.FormatMeta("updateAuthority", obj.AccountMetaSlice.Get(1)))
					})
				})
		})
}

// UpdateAuthority Instruction
type UpdateAuthority struct {
	// New authority for the token metadata, or unset if `None`
	NewAuthority *common.PublicKey
	// [0] = [WRITE] metadata `Metadata`
	// [1] = [SIGNER] updateAuthority `Update authority`
	common.AccountMetaSlice `bin:"-"`
	_programId              *common.PublicKey
}

// NewUpdateAuthorityInstructionBuilder creates a new `UpdateAuthority` instruction builder.
func NewUpdateAuthorityInstructionBuilder() *UpdateAuthority {
	return &UpdateAuthority{
		AccountMetaSlice: make(common.AccountMetaSlice, 2),
	}
}

// NewUpdateAuthorityInstruction
//
// Parameters:
//
//	newAuthority: New authority for the token metadata, or unset if `None`
//	metadata: Metadata
//	updateAuthority: Update authority
func NewUpdateAuthorityInstruction(
	newAuthority common.PublicKey,
	metadata common.PublicKey,
	updateAuthority common.PublicKey,
) *UpdateAuthority {
	return NewUpdateAuthorityInstructionBuilder().
		SetNewAuthority(newAuthority).
		SetMetadataAccount(metadata).
		SetUpdateAuthorityAccount(updateAuthority)
}

// SetNewAuthority sets the "newAuthority" parameter.
func (obj *UpdateAuthority) SetNewAuthority(newAuthority common.PublicKey) *UpdateAuthority {
	obj.NewAuthority = &newAuthority
	return obj
}

// SetMetadataAccount sets the "metadata" parameter.
// Metadata
func (obj *UpdateAuthority) SetMetadataAccount(metadata common.PublicKey) *UpdateAuthority {
	obj.AccountMetaSlice[0] = common.Meta(metadata).WRITE()
	return obj
}

// GetMetadataAccount gets the "metadata" parameter.
// Metadata
func (obj *UpdateAuthority) GetMetadataAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(0)
}

// SetUpdateAuthorityAccount sets the "updateAuthority" parameter.
// Update authority
func (obj *UpdateAuthority) SetUpdateAuthorityAccount(updateAuthority common.PublicKey, multiSigners ...common.PublicKey) *UpdateAuthority {
	if len(multiSigners) > 0 {
		obj.AccountMetaSlice[1] = common.Meta(updateAuthority)
		for _, value := range multiSigners {
			obj.AccountMetaSlice.Append(common.NewAccountMeta(value, false, true))
		}
	} else {
		obj.AccountMetaSlice[1] = common.Meta(updateAuthority).SIGNER()
	}
	return obj
}

// GetUpdateAuthorityAccount gets the "updateAuthority" parameter.
// Update authority
func (obj *UpdateAuthority) GetUpdateAuthorityAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(1)
}

func (obj *UpdateAuthority) SetProgramId(programId *common.PublicKey) *UpdateAuthority {
	obj._programId = programId
	return obj
}

func (obj *UpdateAuthority) Build() *Instruction {
	return &Instruction{
		BaseVariant: binary.BaseVariant{
			Impl:   obj,
			TypeID: binary.TypeIDFromBytes(Instruction_UpdateAuthority[:]),
		},
		programId: obj._programId,
		typeIdLen: 8,
	}
}

func (obj *UpdateAuthority) Validate() error {
	if obj.NewAuthority == nil {
		return errors.New("[UpdateAuthority] newAuthority param is not set")
	}

	if obj.AccountMetaSlice[0] == nil {
		return errors.New("[UpdateAuthority] accounts.metadata is not set")
	}
	if obj.AccountMetaSlice[1] == nil {
		return errors.New("[UpdateAuthority] accounts.updateAuthority is not set")
	}
	return nil
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (obj *UpdateAuthority) ValidateAndBuild() (*Instruction, error) {
	if err := obj.Validate(); err != nil {
		return nil, err
	}
	return obj.Build(), nil
}

func (obj *UpdateAuthority) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	if err = encoder.Encode(&obj.NewAuthority); err != nil {
		return err
	}
	return nil
}

func (obj *UpdateAuthority) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	if err = decoder.Decode(&obj.NewAuthority); err != nil {
		return err
	}
	return nil
}

func (obj *UpdateAuthority) EncodeToTree(parent treeout.Branches) {
	parent.Child(format.Program(ProgramName, common.As(ProgramID))).
		ParentFunc(func(programBranch treeout.Branches) {
			programBranch.Child(format.Instruction("UpdateAuthority")).
				ParentFunc(func(instructionBranch treeout.Branches) {
					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch treeout.Branches) {
						paramsBranch.Child(format.Param("NewAuthority", *obj.NewAuthority))
					})
					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch treeout.Branches) {
						accountsBranch.Child(common.FormatMeta("       metadata", obj.AccountMetaSlice.Get(0)))
						accountsBranch.Child(common.FormatMeta("updateAuthority", obj.AccountMetaSlice.Get(1)))
					})
				})
		})
}

// Emit Instruction
type Emit struct {
	// Start of range of data to emit
	Start *uint64 `bin:"optional"`
	// End of range of data to emit
	End *uint64 `bin:"optional"`
	// [0] = [] metadata `Metadata`
	common.AccountMetaSlice `bin:"-"`
	_programId              *common.PublicKey
}

// NewEmitInstructionBuilder creates a new `Emit` instruction builder.
func NewEmitInstructionBuilder() *Emit {
	return &Emit{
		AccountMetaSlice: make(common.AccountMetaSlice, 1),
	}
}

// NewEmitInstruction
//
// Parameters:
//
//	start: Start of range of data to emit
//	end: End of range of data to emit
//	metadata: Metadata
func NewEmitInstruction(
	// optional,
	start *uint64,
	// optional,
	end *uint64,
	metadata common.PublicKey,
) *Emit {
	return NewEmitInstructionBuilder().
		SetStart(start).
		SetEnd(end).
		SetMetadataAccount(metadata)
}

// SetStart sets the "start" parameter.
func (obj *Emit) SetStart(start *uint64) *Emit {
	obj.Start = start
	return obj
}

// SetEnd sets the "end" parameter.
func (obj *Emit) SetEnd(end *uint64) *Emit {
	obj.End = end
	return obj
}

// SetMetadataAccount sets the "metadata" parameter.
// Metadata
func (obj *Emit) SetMetadataAccount(metadata common.PublicKey, multiSigners ...common.PublicKey) *Emit {
	if len(multiSigners) > 0 {
		obj.AccountMetaSlice[0] = common.Meta(metadata)
		for _, value := range multiSigners {
			obj.AccountMetaSlice.Append(common.NewAccountMeta(value, false, true))
		}
	} else {
		obj.AccountMetaSlice[0] = common.Meta(metadata)
	}
	return obj
}

// GetMetadataAccount gets the "metadata" parameter.
// Metadata
func (obj *Emit) GetMetadataAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(0)
}

func (obj *Emit) SetProgramId(programId *common.PublicKey) *Emit {
	obj._programId = programId
	return obj
}

func (obj *Emit) Build() *Instruction {
	return &Instruction{
		BaseVariant: binary.BaseVariant{
			Impl:   obj,
			TypeID: binary.TypeIDFromBytes(Instruction_Emit[:]),
		},
		programId: obj._programId,
		typeIdLen: 8,
	}
}

func (obj *Emit) Validate() error {

	if obj.AccountMetaSlice[0] == nil {
		return errors.New("[Emit] accounts.metadata is not set")
	}
	return nil
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (obj *Emit) ValidateAndBuild() (*Instruction, error) {
	if err := obj.Validate(); err != nil {
		return nil, err
	}
	return obj.Build(), nil
}

func (obj *Emit) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	if err = encoder.WriteBool(obj.Start != nil); err != nil {
		return err
	}
	if obj.Start != nil {
		if err = encoder.Encode(obj.Start); err != nil {
			return err
		}
	}
	if err = encoder.WriteBool(obj.End != nil); err != nil {
		return err
	}
	if obj.End != nil {
		if err = encoder.Encode(obj.End); err != nil {
			return err
		}
	}
	return nil
}

func (obj *Emit) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	if ok, err := decoder.ReadBool(); err != nil {
		return err
	} else if ok {
		if err = decoder.Decode(&obj.Start); err != nil {
			return err
		}
	}
	if ok, err := decoder.ReadBool(); err != nil {
		return err
	} else if ok {
		if err = decoder.Decode(&obj.End); err != nil {
			return err
		}
	}
	return nil
}

func (obj *Emit) EncodeToTree(parent treeout.Branches) {
	parent.Child(format.Program(ProgramName, common.As(ProgramID))).
		ParentFunc(func(programBranch treeout.Branches) {
			programBranch.Child(format.Instruction("Emit")).
				ParentFunc(func(instructionBranch treeout.Branches) {
					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch treeout.Branches) {
						paramsBranch.Child(format.Param("Start (OPT)", obj.Start))
						paramsBranch.Child(format.Param("  End (OPT)", obj.End))
					})
					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=1]").ParentFunc(func(accountsBranch treeout.Branches) {
						accountsBranch.Child(common.FormatMeta("metadata", obj.AccountMetaSlice.Get(0)))
					})
				})
		})
}
