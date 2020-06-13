package v1alpha1

// Context holds a general purpose REST API context
type Context struct {
	// REST API endpoint
	Endpoint string `json:"endpoint" protobuf:"bytes,1,name=endpoint"`
	// Method is HTTP request method that indicates the desired action to be performed for a given resource.
	// See RFC7231 Hypertext Transfer Protocol (HTTP/1.1): Semantics and Content
	Method string `json:"method" protobuf:"bytes,2,name=method"`
	// Port on which HTTP server is listening for incoming events.
	Port string `json:"port" protobuf:"bytes,3,name=port"`
	// URL is the url of the server.
	URL string `json:"url" protobuf:"bytes,4,name=url"`
	// ServerCertPath refers the file that contains the cert.
	ServerCertPath string `json:"serverCertPath,omitempty" protobuf:"bytes,6,opt,name=serverCertPath"`
	// ServerKeyPath refers the file that contains private key
	ServerKeyPath string `json:"serverKeyPath,omitempty" protobuf:"bytes,5,opt,name=serverKeyPath"`
}
