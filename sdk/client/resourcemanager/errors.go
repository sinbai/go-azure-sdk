package resourcemanager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var _ error = Error{}

type Error struct {
	ActivityId string
	Code       string
	Message    string
	Status     string

	FullHttpBody string
}

func (e Error) Error() string {
	return fmt.Sprintf(`the Azure API returned the following error:

Status: %q
Code: %q
Message: %q
Activity Id: %q

---

API Response:

----[start]----
%s
-----[end]-----
`, e.Status, e.Code, e.Message, e.ActivityId, e.FullHttpBody)
}

// parseErrorFromApiResponse parses the error from the API Response
// into an Error type, which allows for better surfacing of errors
func parseErrorFromApiResponse(response http.Response) (*Error, error) {
	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing response body: %+v", err)
	}
	response.Body.Close()

	respBody = bytes.TrimPrefix(respBody, []byte("\xef\xbb\xbf"))
	response.Body = io.NopCloser(bytes.NewBuffer(respBody))

	// there's a number of internal Azure error types, we should attempt unmarshalling into each
	// for now we're implementing the simple case we can add to
	var err1 lroErrorType1
	if err = json.Unmarshal(respBody, &err1); err == nil && err1.Error.Code != "" && err1.Error.Message != "" {
		e := Error{
			Code:         err1.Error.Code,
			Message:      err1.Error.Message,
			Status:       err1.Status,
			FullHttpBody: string(respBody),
		}

		// given inconsistencies between different API's, this isn't pretty, but avoids crashes whilst best-efforting it
		for _, info := range err1.Error.AdditionalInfo {
			additionalInfo, ok := info.(map[string]interface{})
			if !ok {
				continue
			}

			typeVal, ok := additionalInfo["type"].(string)
			if !ok {
				continue
			}

			if strings.EqualFold(typeVal, "ActivityId") {
				infoBlock, ok := additionalInfo["info"].(map[string]interface{})
				if !ok {
					continue
				}

				for k, v := range infoBlock {
					if strings.EqualFold(k, "ActivityId") {
						if val, ok := v.(string); ok {
							e.ActivityId = val
						}
					}
				}
			}
		}

		return &e, nil
	}

	var err2 resourceManagerErrorType1
	if err = json.Unmarshal(respBody, &err2); err == nil && err2.Code != "" && err2.Message != "" {
		return &Error{
			Code:         err2.Code,
			Message:      err2.Message,
			Status:       "Unknown",
			FullHttpBody: string(respBody),
		}, nil
	}

	return &Error{
		Code:         "internal-error",
		Message:      "Couldn't parse Azure API Response into a friendly error - please see the original HTTP Response for more details (and file a bug so we can fix this!).",
		Status:       "Unknown",
		FullHttpBody: string(respBody),
	}, nil
}

type lroErrorType1 struct {
	Error struct {
		Code           string        `json:"code"`
		Message        string        `json:"message"`
		AdditionalInfo []interface{} `json:"additionalInfo"`
	} `json:"error"`
	Status string `json:"status"`
}

type resourceManagerErrorType1 struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details []struct {
		Message string `json:"message,omitempty"`
		Code    string `json:"code,omitempty"`
	} `json:"details"`
}
