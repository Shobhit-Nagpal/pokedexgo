package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func getLocation(url string) (MapResponse, []byte, error) {
  data := MapResponse{}
  res, err := http.Get(url)
  if err != nil {
    log.Fatal(err)
    return data, []byte{}, err
  }

  body, err := io.ReadAll(res.Body)
  res.Body.Close()
  if err != nil {
    log.Fatal(err)
    return data, []byte{}, err
  }

  err = json.Unmarshal(body, &data)
  if err != nil {
    log.Fatal(err)
    return data, []byte{}, err
  }

  return data, body, nil
}

func printLocations(locations MapResponse) {
  for _, location := range locations.Results {
   fmt.Println(location.Name) 
  }
}
