// example Go snippet from the Google SRE book busing gRPC
// This code runs from the front end server to communicate with the back end server

func exampleRpcCall(client pb.ExampleClient, request pb.Request) *pb.Response {
	// Set RPC timeout to 5 seconds.
	opts := grpc.WithTimeout(5 * time.Second)
	// Try up to 20 times to make the RPC call.
	attempts := 20
	for attempts > 0 {
		conn, err := grpc.Dial(*serverAddr, opts...)
		if err != nil {
			// Something went wrong in setting up the connection. Try again.
			attempts--
			continue
		}
		defer conn.Close()
		// Create a client stub and make the RPC call.
		client := pb.NewBackendClient(conn)
		response, err := client.MakeRequest(context.Background, request)
		if err != nil {
			// Something went wrong in making the call. Try again.
			attempts--
			continue
		}
		return response
	}
	grpclog.Fatalf("ran out of attempts")
}
