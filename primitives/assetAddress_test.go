package primitives

import (
	"testing"
)

func TestAssetAddressFromTypeAndPayload(t *testing.T) {
	testHelperForAssetAddressFromTypeAndPayload(t, 0, "tcaqyqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqxdyr05")
	testHelperForAssetAddressFromTypeAndPayload(t, 1, "tcaqyqsqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqjp4hqq")
	testHelperForAssetAddressFromTypeAndPayload(t, 2, "tcaqypqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqga755k")
}

func testHelperForAssetAddressFromTypeAndPayload(t *testing.T, addressType byte, expected string) {
	zeros := NewH160Zero()
	assetAddress, _ := AssetAddressFromTypeAndPayload(addressType, zeros, "tc")
	if assetAddress.Value != expected {
		t.Fatalf("Expected %s, but actual %s", expected, assetAddress.Value)
	}
}

func TestAssetAddressFromString(t *testing.T) {
	testHelperForAssetAddressFromString(t, 0, "tcaqyqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqxdyr05")
	testHelperForAssetAddressFromString(t, 1, "tcaqyqsqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqjp4hqq")
	testHelperForAssetAddressFromString(t, 2, "tcaqypqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqga755k")
}

func testHelperForAssetAddressFromString(t *testing.T, addressType byte, expected string) {
	assetAddress, err := AssetAddressFromString(expected)
	if err != nil {
		t.Fatal(err)
	}
	if assetAddress.AssetType != addressType {
		t.Fatalf("Expected AddressType %d, but actual %d", addressType, assetAddress.AssetType)
	}
	zeroH160 := NewH160Zero()
	if assetAddress.Payload != zeroH160 {
		t.Fatalf("Expected payload %s, but actual %s", zeroH160, assetAddress.Payload)
	}
}
