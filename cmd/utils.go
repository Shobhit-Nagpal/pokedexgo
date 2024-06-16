package cmd

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func getLocation(url string) (MapResponse, error) {
  data := MapResponse{}
  res, err := http.Get(url)
  if err != nil {
    log.Fatal(err)
    return data, err
  }

  body, err := io.ReadAll(res.Body)
  res.Body.Close()
  if err != nil {
    log.Fatal(err)
    return data, err
  }

  err = json.Unmarshal(body, &data)
  if err != nil {
    log.Fatal(err)
    return data, err
  }

  return data, nil
}
