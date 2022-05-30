package valuerenderer

import (
	"context"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const thousandSeparator string = "'"

// NewADR050ValueRenderer returns a new ValueRenderer based on the ADR-050
// spec.
func NewADR050ValueRenderer() ValueRenderer {
	return adr050ValueRenderer{}
}

// adr050ValueRenderer is a value renderer based on the ADR-050 spec.
type adr050ValueRenderer struct{}

var _ ValueRenderer = adr050ValueRenderer{}

func (r adr050ValueRenderer) Format(ctx context.Context, v proto.Message) ([]string, error) {
	fields := v.ProtoReflect().Descriptor().Fields()
	fieldsLen := fields.Len()

	result := []string{}
	for i := 0; i < fieldsLen; i++ {
		fd := fields.Get(i)            // Field descriptor
		fv := v.ProtoReflect().Get(fd) // Field value

		switch {
		case fd.Kind() == protoreflect.Uint32Kind ||
			fd.Kind() == protoreflect.Uint64Kind ||
			fd.Kind() == protoreflect.Int32Kind ||
			fd.Kind() == protoreflect.Int64Kind:
			{
				formatted, err := formatInteger(fv.String())
				if err != nil {
					return nil, err
				}

				result = append(result, formatted)
			}
		default:
			return nil, fmt.Errorf("value renderers cannot format value %s of type %s", fv, fd.Kind())
		}
	}

	return result, nil
}

func (r adr050ValueRenderer) Parse(context.Context, []string) (proto.Message, error) {
	panic("implement me")
}

// formatInteger formats an integer into a value-rendered string. This function
// operates with string manipulation (instead of manipulating the int or sdk.Int
// object).
func formatInteger(v string) (string, error) {
	if v[0] == '-' {
		v = v[1:]
	}
	if len(v) > 1 {
		v = strings.TrimLeft(v, "0")

	}

	startOffset := 3
	for outputIndex := len(v); outputIndex > startOffset; {
		outputIndex -= 3
		v = v[:outputIndex] + thousandSeparator + v[outputIndex:]
	}
	return v, nil
}

// formatDecimal formats a decimal into a value-rendered string. This function
// operates with string manipulation (instead of manipulating the sdk.Dec
// object).
func formatDecimal(v string) (string, error) {
	parts := strings.Split(v, ".")
	intPart, err := formatInteger(parts[0])
	if err != nil {
		return "", err
	}

	if len(parts) > 2 {
		return "", fmt.Errorf("invalid decimal %s", v)
	}

	if len(parts) == 1 {
		return intPart, nil
	}

	decPart := strings.TrimRight(parts[1], "0")
	if len(decPart) == 0 {
		return intPart, nil
	}

	return intPart + "." + decPart, nil
}

// formatDecimal formats a sdk.Coin into a value-rendered string, using the
// given metadata about the denom.
func formatCoin(coin sdk.Coin, metadata bank.Metadata) (string, error) {
	curDenom := coin.Denom
	dispDenom := metadata.Display

	// Return early if no display denom or display denom is the current coin denom.
	if dispDenom == "" || curDenom == dispDenom {
		vr, err := formatInteger(coin.Amount.String())
		return vr + " " + coin.Denom, err
	}

	// Find exponents of both denoms.
	var curExp, dispExp uint32
	foundCurExp, foundDispExp := false, false
	for _, unit := range metadata.DenomUnits {
		if curDenom == unit.Denom {
			curExp = unit.Exponent
			foundCurExp = true
		}
		if dispDenom == unit.Denom {
			dispExp = unit.Exponent
			foundDispExp = true
		}
	}

	// If we didn't find either exponent, then we return early.
	if !foundCurExp || !foundDispExp {
		vr, err := formatInteger(coin.Amount.String())
		return vr + " " + coin.Denom, err
	}

	exponentDiff := int64(curExp) - int64(dispExp)
	var dispAmount sdk.Dec
	if exponentDiff > 0 {
		dispAmount = sdk.NewDecFromInt(coin.Amount).Mul(sdk.NewDec(10).Power(uint64(exponentDiff)))
	} else {
		dispAmount = sdk.NewDecFromInt(coin.Amount).Quo(sdk.NewDec(10).Power(uint64(-exponentDiff)))
	}

	vr, err := formatDecimal(dispAmount.String())
	return vr + " " + dispDenom, err
}

// formatDecimal formats a sdk.Coins into a value-rendered string, which uses
// `formatCoin` separated by ", " (a comma and a space).
func formatCoins(coins sdk.Coins, metadata bank.Metadata) (string, error) {
	formatted := make([]string, len(coins))
	for i, coin := range coins {
		var err error
		formatted[i], err = formatCoin(coin, metadata)
		if err != nil {
			return "", err
		}
	}

	return strings.Join(formatted, ", "), nil
}
