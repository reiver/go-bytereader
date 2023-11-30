package bytereader_test

import (
	"testing"

	"io"

	"sourcecode.social/reiver/go-bytereader"
)

func TestNewByteReaderFromString(t *testing.T) {

	tests := []struct{
		Value string
	}{
		{
			Value: "",
		},



		{
			Value: "apple",
		},
		{
			Value: "banana",
		},
		{
			Value: "cherry",
		},



		{
			Value: "ONCE TWICE THRICE FOURCE",
		},



		{
			Value: "۰۱۲۳۴۵۶۷۸۹",
		},



		{
			Value: "😈 👹 😡",
		},
	}

	testloop: for testNumber, test := range tests {

		var byteReader io.ByteReader = bytereader.NewByteReaderFromString(test.Value)

		var buffer [64]byte
		var p []byte = buffer[0:0]
		{
			length := len(test.Value)

			for i:=0; i<length; i++ {

				b, err := byteReader.ReadByte()
				if nil != err {
					t.Errorf("For test #%d and loop #%d, did not expect an error but actually got one.", testNumber, i)
					t.Logf("ERROR: (%T) %s", err, err)
					t.Logf("VALUE: %q", test.Value)
					continue testloop
				}

				p = append(p, b)
			}
		}

		{
			_, err := byteReader.ReadByte()
			if nil == err {
				t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
				t.Logf("VALUE: %q", test.Value)
				continue testloop
			}

			{
				expected := io.EOF
				actual   := err

				if expected != actual {
					t.Errorf("For test #%d, expected an io.EOF error but actually got a different error.", testNumber)
					t.Logf("ERROR: (%T) %s", err, err)
					t.Logf("VALUE: %q", test.Value)
					continue testloop
				}
			}
		}

		{
			expected := test.Value
			actual   := string(p)

			if expected != actual {
				t.Errorf("For test #%d, the actual read value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				continue testloop
			}
		}
	}
}
