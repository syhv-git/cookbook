package cmd

type single struct {
	id          int
	name        string
	description string
}

type nested struct {
	id       int
	name     string
	children []single
}

const (
	title   = "   ┌───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐\n   │░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░│\n   │░░░░██████╗██╗░░██╗███████╗██╗██████╗░██╗░██████╗░░░░█████╗░░█████╗░░█████╗░██╗░░██╗██████╗░░█████╗░░█████╗░██╗░░██╗░░░│\n   │░░░██╔════╝██║░██╔╝██╔════╝██║██╔══██╗╚█║██╔════╝░░░██╔══██╗██╔══██╗██╔══██╗██║░██╔╝██╔══██╗██╔══██╗██╔══██╗██║░██╔╝░░░│\n   │░░░╚████╗░░█████╔╝░█████╗░░██║██║░░██║░╚╝╚████╗░░░░░██║░░╚═╝██║░░██║██║░░██║█████╔╝░██████╦╝██║░░██║██║░░██║█████╔╝░░░░│\n   │░░░░╚████╗░█████╝░░██╔══╝░░██║██║░░██║░░░░╚████╗░░░░██║░░░░░██║░░██║██║░░██║█████╝░░██████╗░██║░░██║██║░░██║█████╝░░░░░│\n   │░░░░░╚══██╗██╔═██╗░██║░░░░░██║██║░░██║░░░░░╚══██╗░░░██║░░██╗██║░░██║██║░░██║██╔═██╗░██╔══██╗██║░░██║██║░░██║██╔═██╗░░░░│\n   │░░░██████╔╝██║░╚██╗███████╗██║██████╔╝░░░██████╔╝░░░╚█████╔╝╚█████╔╝╚█████╔╝██║░╚██╗██████╔╝╚█████╔╝╚█████╔╝██║░╚██╗░░░│\n   │░░░╚═════╝░╚═╝░░╚═╝╚══════╝╚═╝╚═════╝░░░░╚═════╝░░░░░╚════╝░░╚════╝░░╚════╝░╚═╝░░╚═╝╚═════╝░░╚════╝░░╚════╝░╚═╝░░╚═╝░░░│\n   │░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░│\n┌──┴───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┴──┐\n"
	lborder = "│               "   // 95 utf8 chars between lborder and rborder or 47 utf8 char column split by 1 gap or 17-58:70:111
	rborder = "               │\n" // each line contains 128 values including new line
)

var (
	files      = single{id: 0, name: "Files", description: "Package of commands for file interactions"}
	networking = single{id: 1, name: "Networking", description: "Package of commands involving network activity"}
	help       = single{name: "Help", description: "Read the documentation for all usable commands from this package"}

	forensics = single{id: 0, name: "Forensics", description: "Package of commands for performing forensics"}
	utility   = single{id: 1, name: "Utility", description: "Package of utility commands for handling files"}
)

var (
	checksum  = single{id: 0, name: "Checksum", description: "Perform a checksum verification on a file"}
	compress  = single{id: 1, name: "Compress", description: "Create a compressed archive file over the source files"}
	enumerate = single{id: 2, name: "Enumerate", description: "Perform an enumeration on paths and display the results"}
	extract   = single{id: 3, name: "Extract", description: "Create a new file containing the concatenated contents of all source files"}
	stego     = single{id: 4, name: "Steganography", description: "Create or detect embedded archives in image files"}
)

var (
	discovery = single{id: 0, name: "Discovery", description: "Discover important networking metadata from IP and Hostnames"}
	download  = single{id: 1, name: "Download", description: "Download a file from the source URL"}
)

var (
	page        = nested{id: 0, name: "Select Type", children: []single{files, networking, help}}
	filePage    = nested{id: 0, name: "Files", children: []single{forensics, utility}}
	networkPage = nested{id: 1, name: "Networking", children: []single{forensics, utility}}

	fileForensics = nested{id: 0, name: "Forensics", children: []single{checksum, enumerate, extract, stego, help}}
	fileUtility   = nested{id: 1, name: "Utility", children: []single{compress, help}}

	networkForensics = nested{id: 0, name: "Forensics", children: []single{discovery, help}}
	networkUtility   = nested{id: 1, name: "Utility", children: []single{download, help}}
)
