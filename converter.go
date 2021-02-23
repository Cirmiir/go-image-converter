package main

/*Converter struct contains the encode and decode transformation*/
type Converter struct {
	Encode func([]byte) ([]byte, error)
	Decode func([]byte) ([]byte, error)
}

/*Convert function to perform convertion*/
func (converter *Converter) Convert(data []byte, direction bool) ([]byte, error) {
	var transform func([]byte) ([]byte, error)
	if direction {
		transform = converter.Encode
	} else {
		transform = converter.Decode
	}
	return transform(data)
}
