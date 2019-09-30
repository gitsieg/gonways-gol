package server

import (
	"gonways-gol/gol"
	"net/http"
	"reflect"
	"testing"
)

func TestServer_HandleBoardCreateStructure(t *testing.T) {
	type fields struct {
		httpServe *http.Server
		board     *gol.Board
	}
	tests := []struct {
		name   string
		fields fields
		want   http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				httpServe: tt.fields.httpServe,
				board:     tt.fields.board,
			}
			if got := s.HandleBoardCreateStructure(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.HandleBoardCreateStructure() = %v, want %v", got, tt.want)
			}
		})
	}
}
