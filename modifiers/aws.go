package modifiers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type AWS struct{}

func NewAWS() *AWS {
	e := AWS{}

	return &e
}

// http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html
func (e *AWS) Get(s string) (*string, error) {
	query := fmt.Sprintf("http://169.254.169.254/latest/meta-data/%s", s)
	response, err := http.Get(query)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	val := string(body)

	return &val, err
}
