import ArgumentParser

extension Brage {
	struct Serve: ParsableCommand {
		static let configuration = CommandConfiguration(
			abstract: "Serve a site.")
		
		@Argument(help: "Path of the site configuration.")
		var sitePath: String
		
		@Option(name: .shortAndLong, help: "Port to serve the site on.")
		var port = 8080
		
		mutating func run() {
			print("TODO: \(sitePath) \(port)")
		}
	}
}