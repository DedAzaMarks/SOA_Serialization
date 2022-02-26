package formats_comparison

import (
	"encoding/xml"
	"fmt"
)

type StringMap map[string]string

func (s *StringMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	tokens := []xml.Token{start}

	for key, value := range *s {
		t := xml.StartElement{Name: xml.Name{Local: key}}
		tokens = append(tokens, t, xml.CharData(value), t.End())
	}

	tokens = append(tokens, xml.EndElement{Name: start.Name})

	for _, t := range tokens {
		err := e.EncodeToken(t)
		if err != nil {
			return err
		}
	}

	// flush to ensure tokens are written
	err := e.Flush()
	if err != nil {
		return err
	}
	return nil
}

func (s *StringMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*s = StringMap{}
	for {
		// check if map is ended
		token, err := d.Token()
		if err != nil {
			return fmt.Errorf("read token error: %w", err)
		}
		if end, ok := token.(xml.EndElement); ok && end == start.End() {
			break
		}

		// parse open tag
		key := token.(xml.StartElement).Copy().Name.Local

		// parse content
		switch token, err := d.Token(); token.(type) {
		case xml.CharData:
			value := string(token.(xml.CharData).Copy())
			(*s)[key] = value
			if _, err := d.Token(); err != nil {
				return fmt.Errorf("read token error: %w", err)
			}
		case xml.EndElement:
			(*s)[key] = ""
		default:
			if err != nil {
				return fmt.Errorf("read token error: %w", err)
			}
			return fmt.Errorf("expected xml.CharData or xml.EndElement, got %T", token)
		}
	}
	return nil
}

type xmlSerializer struct{}

func (xmlSerializer) Marshal(t *Test) ([]byte, error) {
	return xml.Marshal(t)
}

func (xmlSerializer) Unmarshal(b []byte, t *Test) error {
	return xml.Unmarshal(b, t)
}
