package main
import ("net/http/httptest";"testing")
func TestHealth(t *testing.T){r:=httptest.NewRecorder();health(r,httptest.NewRequest("GET","/health",nil));if r.Code!=200{t.Fatal(r.Code)}}
