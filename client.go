package gomutlucell

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type MutluCellClient struct {
	userName   string
	password   string
	originator string
	charset    string // turkish,unicode
}

type Message struct {
	Message string `xml:"metin"`
	Numbers string `xml:"nums"`
}

type request struct {
	XMLName    struct{}  `xml:"smspack"`
	Username   string    `xml:"ka,attr"`
	Password   string    `xml:"pwd,attr"`
	Originator string    `xml:"org,attr"`
	Charset    string    `xml:"charset,attr"`
	Message    []Message `xml:"mesaj"`
}

func NewMutluCellClient(userName, password, originator, charset string) *MutluCellClient {
	return &MutluCellClient{
		userName:   userName,
		password:   password,
		originator: originator,
		charset:    charset,
	}
}

func (mutluCell *MutluCellClient) SendSms(message Message) error {
	return mutluCell.request(request{
		Message: []Message{message},
	}, "$")
}

func (mutluCell *MutluCellClient) request(request request, expectedResponsePrefix string) error {
	request.Username = mutluCell.userName
	request.Password = mutluCell.password
	request.Originator = mutluCell.originator
	request.Charset = mutluCell.charset

	body, err := xml.Marshal(request)
	if err != nil {
		return err
	}
	fmt.Printf("mutlucell request body %s\n", string(body))

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://smsgw.mutlucell.com/smsgw-ws/sndblkex", bytes.NewBuffer([]byte(body)))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/xml; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("response status not ok, expected 200 given %d", resp.StatusCode)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	responseBodyStr := string(responseBody)
	fmt.Printf("mutlucecell response body %s\n", responseBodyStr)

	if expectedResponsePrefix != "" && responseBodyStr[:len(expectedResponsePrefix)] != expectedResponsePrefix {
		return fmt.Errorf("response prefix could not matched expected %s given %s", expectedResponsePrefix, responseBodyStr[:len(expectedResponsePrefix)])
	}

	return nil
}
