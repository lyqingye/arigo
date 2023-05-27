package arigo_test

import (
	"testing"

	"github.com/siku2/arigo"
	"github.com/stretchr/testify/require"
)

// Dial is a convenience method which connects to an aria2 RPC interface.
// It establishes a WebSocket connection to the given url and passes it
// to the NewClient() method. The client is also started.
func TestClient(t *testing.T) {
	// arigo uses a WebSocket connection to communicate with the aria2 RPC.
	// This makes it possible to receive download events.
	// authToken is the secret string set on the aria2 server. Setting it to an empty string
	// indicates that no password should be used.
	client, err := arigo.Dial("ws://nas.lyqingye.com:9000/jsonrpc", "123456")

	// err is returned if no connection to the RPC interface could be established.
	if err != nil {
		panic(err)
	}
	opts, err := client.GetGlobalOptions()
	require.NoError(t, err)
	require.NotNil(t, opts)

	// client is now connected and can be used
	// gid, err := client.AddURI(arigo.URIs("https://file-examples.com/wp-content/storage/2017/10/file_example_JPG_100kB.jpg"), nil)
	// if err != nil {
	// 	panic(err)
	// }

	// gid isn't just a string but a GID instance.
	// this makes it possible to use all the client methods you would normally
	// pass a gid to directly on the GID instance.
	// WaitForDownload() is a special method which waits for the download to complete
	// (using the aria2 events)
	// if err = gid.WaitForDownload(); err != nil {
	// 	panic(err)
	// }

	// Get the status of the now completed download.
	// The TellStatus() method accepts the keys of the Status struct
	// which are to be requested. Please note that the keys must match
	// the aria2 keys, which may differ from the keys used in the golang
	// representation.
	// For the aria2 status documentation see: https://aria2.github.io/manual/en/html/aria2c.html#aria2.tellStatus
	// status, err := gid.TellStatus("status")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(status.Status)
}
