// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package wormhole

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CreateWrapped is the `createWrapped` instruction.
type CreateWrapped struct {

	// [0] = [WRITE, SIGNER] payer
	//
	// [1] = [] config
	//
	// [2] = [] endpoint
	//
	// [3] = [] vaa
	//
	// [4] = [WRITE] claim
	//
	// [5] = [WRITE] mint
	//
	// [6] = [WRITE] wrappedMeta
	//
	// [7] = [WRITE] splMetadata
	//
	// [8] = [] mintAuthority
	//
	// [9] = [] rent
	//
	// [10] = [] systemProgram
	//
	// [11] = [] tokenProgram
	//
	// [12] = [] splMetadataProgram
	//
	// [13] = [] wormholeProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCreateWrappedInstructionBuilder creates a new `CreateWrapped` instruction builder.
func NewCreateWrappedInstructionBuilder() *CreateWrapped {
	nd := &CreateWrapped{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 14),
	}
	return nd
}

// SetPayerAccount sets the "payer" account.
func (inst *CreateWrapped) SetPayerAccount(payer ag_solanago.PublicKey) *CreateWrapped {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *CreateWrapped) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetConfigAccount sets the "config" account.
func (inst *CreateWrapped) SetConfigAccount(config ag_solanago.PublicKey) *CreateWrapped {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *CreateWrapped) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetEndpointAccount sets the "endpoint" account.
func (inst *CreateWrapped) SetEndpointAccount(endpoint ag_solanago.PublicKey) *CreateWrapped {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(endpoint)
	return inst
}

// GetEndpointAccount gets the "endpoint" account.
func (inst *CreateWrapped) GetEndpointAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetVaaAccount sets the "vaa" account.
func (inst *CreateWrapped) SetVaaAccount(vaa ag_solanago.PublicKey) *CreateWrapped {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(vaa)
	return inst
}

// GetVaaAccount gets the "vaa" account.
func (inst *CreateWrapped) GetVaaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetClaimAccount sets the "claim" account.
func (inst *CreateWrapped) SetClaimAccount(claim ag_solanago.PublicKey) *CreateWrapped {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(claim).WRITE()
	return inst
}

// GetClaimAccount gets the "claim" account.
func (inst *CreateWrapped) GetClaimAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetMintAccount sets the "mint" account.
func (inst *CreateWrapped) SetMintAccount(mint ag_solanago.PublicKey) *CreateWrapped {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(mint).WRITE()
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *CreateWrapped) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetWrappedMetaAccount sets the "wrappedMeta" account.
func (inst *CreateWrapped) SetWrappedMetaAccount(wrappedMeta ag_solanago.PublicKey) *CreateWrapped {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(wrappedMeta).WRITE()
	return inst
}

// GetWrappedMetaAccount gets the "wrappedMeta" account.
func (inst *CreateWrapped) GetWrappedMetaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetSplMetadataAccount sets the "splMetadata" account.
func (inst *CreateWrapped) SetSplMetadataAccount(splMetadata ag_solanago.PublicKey) *CreateWrapped {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(splMetadata).WRITE()
	return inst
}

// GetSplMetadataAccount gets the "splMetadata" account.
func (inst *CreateWrapped) GetSplMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetMintAuthorityAccount sets the "mintAuthority" account.
func (inst *CreateWrapped) SetMintAuthorityAccount(mintAuthority ag_solanago.PublicKey) *CreateWrapped {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(mintAuthority)
	return inst
}

// GetMintAuthorityAccount gets the "mintAuthority" account.
func (inst *CreateWrapped) GetMintAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetRentAccount sets the "rent" account.
func (inst *CreateWrapped) SetRentAccount(rent ag_solanago.PublicKey) *CreateWrapped {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *CreateWrapped) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *CreateWrapped) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *CreateWrapped {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *CreateWrapped) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *CreateWrapped) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *CreateWrapped {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *CreateWrapped) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetSplMetadataProgramAccount sets the "splMetadataProgram" account.
func (inst *CreateWrapped) SetSplMetadataProgramAccount(splMetadataProgram ag_solanago.PublicKey) *CreateWrapped {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(splMetadataProgram)
	return inst
}

// GetSplMetadataProgramAccount gets the "splMetadataProgram" account.
func (inst *CreateWrapped) GetSplMetadataProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetWormholeProgramAccount sets the "wormholeProgram" account.
func (inst *CreateWrapped) SetWormholeProgramAccount(wormholeProgram ag_solanago.PublicKey) *CreateWrapped {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(wormholeProgram)
	return inst
}

// GetWormholeProgramAccount gets the "wormholeProgram" account.
func (inst *CreateWrapped) GetWormholeProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

func (inst CreateWrapped) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CreateWrapped,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CreateWrapped) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CreateWrapped) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Endpoint is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Vaa is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Claim is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.WrappedMeta is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.SplMetadata is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.MintAuthority is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.Rent is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.SplMetadataProgram is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.WormholeProgram is not set")
		}
	}
	return nil
}

func (inst *CreateWrapped) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CreateWrapped")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=14]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("             payer", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("            config", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("          endpoint", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("               vaa", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("             claim", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("              mint", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("       wrappedMeta", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("       splMetadata", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("     mintAuthority", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("              rent", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("     systemProgram", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("      tokenProgram", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("splMetadataProgram", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("   wormholeProgram", inst.AccountMetaSlice.Get(13)))
					})
				})
		})
}

func (obj CreateWrapped) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *CreateWrapped) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewCreateWrappedInstruction declares a new CreateWrapped instruction with the provided parameters and accounts.
func NewCreateWrappedInstruction(
	// Accounts:
	payer ag_solanago.PublicKey,
	config ag_solanago.PublicKey,
	endpoint ag_solanago.PublicKey,
	vaa ag_solanago.PublicKey,
	claim ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	wrappedMeta ag_solanago.PublicKey,
	splMetadata ag_solanago.PublicKey,
	mintAuthority ag_solanago.PublicKey,
	rent ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	splMetadataProgram ag_solanago.PublicKey,
	wormholeProgram ag_solanago.PublicKey) *CreateWrapped {
	return NewCreateWrappedInstructionBuilder().
		SetPayerAccount(payer).
		SetConfigAccount(config).
		SetEndpointAccount(endpoint).
		SetVaaAccount(vaa).
		SetClaimAccount(claim).
		SetMintAccount(mint).
		SetWrappedMetaAccount(wrappedMeta).
		SetSplMetadataAccount(splMetadata).
		SetMintAuthorityAccount(mintAuthority).
		SetRentAccount(rent).
		SetSystemProgramAccount(systemProgram).
		SetTokenProgramAccount(tokenProgram).
		SetSplMetadataProgramAccount(splMetadataProgram).
		SetWormholeProgramAccount(wormholeProgram)
}
