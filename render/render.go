package render

import (
	"errors"
	"io"
	"net/http"
)

func DecodeJSON(r io.Reader, v any) error {
	const op = "utils.render.DecodeJSON"
	b, err := io.ReadAll(r)
	if err != nil {
		return errors.New(op + "- failed to decode json")
	}

	err = json.Unmarshall(b, &v)
	if err != nil {
		return errors.New(op + "- failed to unmarshall json")
	}
	return nil
}

func JSON(w *http.ResponseWriter, v any) {}
