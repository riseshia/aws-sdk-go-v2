// Code generated by smithy-go-codegen DO NOT EDIT.

package restxmlwithnamespace

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/restxmlwithnamespace/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithytesting "github.com/aws/smithy-go/testing"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/google/go-cmp/cmp/cmpopts"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestClient_SimpleScalarProperties_awsRestxmlSerialize(t *testing.T) {
	t.Skip("disabled test aws.protocoltests.restxml.xmlns#RestXmlWithNamespace aws.protocoltests.restxml.xmlns#SimpleScalarProperties")

	cases := map[string]struct {
		Params        *SimpleScalarPropertiesInput
		ExpectMethod  string
		ExpectURIPath string
		ExpectQuery   []smithytesting.QueryItem
		RequireQuery  []string
		ForbidQuery   []string
		ExpectHeader  http.Header
		RequireHeader []string
		ForbidHeader  []string
		BodyMediaType string
		BodyAssert    func(io.Reader) error
	}{
		// Serializes simple scalar properties
		"XmlNamespaceSimpleScalarProperties": {
			Params: &SimpleScalarPropertiesInput{
				Foo:               ptr.String("Foo"),
				StringValue:       ptr.String("string"),
				TrueBooleanValue:  ptr.Bool(true),
				FalseBooleanValue: ptr.Bool(false),
				ByteValue:         ptr.Int8(1),
				ShortValue:        ptr.Int16(2),
				IntegerValue:      ptr.Int32(3),
				LongValue:         ptr.Int64(4),
				FloatValue:        ptr.Float32(5.5),
				DoubleValue:       ptr.Float64(6.5),
				Nested: &types.NestedWithNamespace{
					AttrField: ptr.String("nestedAttrValue"),
				},
			},
			ExpectMethod:  "PUT",
			ExpectURIPath: "/SimpleScalarProperties",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/xml"},
				"X-Foo":        []string{"Foo"},
			},
			BodyMediaType: "application/xml",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareXMLReaderBytes(actual, []byte(`<SimpleScalarPropertiesInputOutput xmlns="https://example.com">
			    <stringValue>string</stringValue>
			    <trueBooleanValue>true</trueBooleanValue>
			    <falseBooleanValue>false</falseBooleanValue>
			    <byteValue>1</byteValue>
			    <shortValue>2</shortValue>
			    <integerValue>3</integerValue>
			    <longValue>4</longValue>
			    <floatValue>5.5</floatValue>
			    <DoubleDribble>6.5</DoubleDribble>
			    <Nested xmlns:xsi="https://example.com" xsi:someName="nestedAttrValue"></Nested>
			</SimpleScalarPropertiesInputOutput>
			`))
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			var actualReq *http.Request
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				actualReq = r.Clone(r.Context())
				if len(actualReq.URL.RawPath) == 0 {
					actualReq.URL.RawPath = actualReq.URL.Path
				}
				if v := actualReq.ContentLength; v != 0 {
					actualReq.Header.Set("Content-Length", strconv.FormatInt(v, 10))
				}
				var buf bytes.Buffer
				if _, err := io.Copy(&buf, r.Body); err != nil {
					t.Errorf("failed to read request body, %v", err)
				}
				actualReq.Body = ioutil.NopCloser(&buf)

				w.WriteHeader(200)
			}))
			defer server.Close()
			url := server.URL
			client := New(Options{
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = url
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				HTTPClient: awshttp.NewBuildableClient(),
				Region:     "us-west-2",
			})
			result, err := client.SimpleScalarProperties(context.Background(), c.Params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if e, a := c.ExpectMethod, actualReq.Method; e != a {
				t.Errorf("expect %v method, got %v", e, a)
			}
			if e, a := c.ExpectURIPath, actualReq.URL.RawPath; e != a {
				t.Errorf("expect %v path, got %v", e, a)
			}
			queryItems := smithytesting.ParseRawQuery(actualReq.URL.RawQuery)
			smithytesting.AssertHasQuery(t, c.ExpectQuery, queryItems)
			smithytesting.AssertHasQueryKeys(t, c.RequireQuery, queryItems)
			smithytesting.AssertNotHaveQueryKeys(t, c.ForbidQuery, queryItems)
			smithytesting.AssertHasHeader(t, c.ExpectHeader, actualReq.Header)
			smithytesting.AssertHasHeaderKeys(t, c.RequireHeader, actualReq.Header)
			smithytesting.AssertNotHaveHeaderKeys(t, c.ForbidHeader, actualReq.Header)
			if c.BodyAssert != nil {
				if err := c.BodyAssert(actualReq.Body); err != nil {
					t.Errorf("expect body equal, got %v", err)
				}
			}
		})
	}
}

func TestClient_SimpleScalarProperties_awsRestxmlDeserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *SimpleScalarPropertiesOutput
	}{
		// Serializes simple scalar properties
		"XmlNamespaceSimpleScalarProperties": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/xml"},
				"X-Foo":        []string{"Foo"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<SimpleScalarPropertiesInputOutput xmlns="https://example.com">
			    <stringValue>string</stringValue>
			    <trueBooleanValue>true</trueBooleanValue>
			    <falseBooleanValue>false</falseBooleanValue>
			    <byteValue>1</byteValue>
			    <shortValue>2</shortValue>
			    <integerValue>3</integerValue>
			    <longValue>4</longValue>
			    <floatValue>5.5</floatValue>
			    <DoubleDribble>6.5</DoubleDribble>
			    <Nested xmlns:xsi="https://example.com" xsi:someName="nestedAttrValue"></Nested>
			</SimpleScalarPropertiesInputOutput>
			`),
			ExpectResult: &SimpleScalarPropertiesOutput{
				Foo:               ptr.String("Foo"),
				StringValue:       ptr.String("string"),
				TrueBooleanValue:  ptr.Bool(true),
				FalseBooleanValue: ptr.Bool(false),
				ByteValue:         ptr.Int8(1),
				ShortValue:        ptr.Int16(2),
				IntegerValue:      ptr.Int32(3),
				LongValue:         ptr.Int64(4),
				FloatValue:        ptr.Float32(5.5),
				DoubleValue:       ptr.Float64(6.5),
				Nested: &types.NestedWithNamespace{
					AttrField: ptr.String("nestedAttrValue"),
				},
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			url := "http://localhost:8888/"
			client := New(Options{
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					headers := http.Header{}
					for k, vs := range c.Header {
						for _, v := range vs {
							headers.Add(k, v)
						}
					}
					if len(c.BodyMediaType) != 0 && len(headers.Values("Content-Type")) == 0 {
						headers.Set("Content-Type", c.BodyMediaType)
					}
					response := &http.Response{
						StatusCode: c.StatusCode,
						Header:     headers,
						Request:    r,
					}
					if len(c.Body) != 0 {
						response.ContentLength = int64(len(c.Body))
						response.Body = ioutil.NopCloser(bytes.NewReader(c.Body))
					} else {

						response.Body = http.NoBody
					}
					return response, nil
				}),
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = url
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				Region: "us-west-2",
			})
			var params SimpleScalarPropertiesInput
			result, err := client.SimpleScalarProperties(context.Background(), &params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if err := smithytesting.CompareValues(c.ExpectResult, result, cmpopts.IgnoreUnexported(middleware.Metadata{})); err != nil {
				t.Errorf("expect c.ExpectResult value match:\n%v", err)
			}
		})
	}
}
