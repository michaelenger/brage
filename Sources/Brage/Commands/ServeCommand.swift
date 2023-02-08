import ArgumentParser

extension Brage {
	/// Command for serving a site on localhost.
	struct ServeCommand: ParsableCommand {
		static let configuration = CommandConfiguration(
			commandName: "serve",
			abstract: "Serve a site.")

		@Argument(help: "Path of the site configuration.")
		var sitePath: String

		@Option(name: .shortAndLong, help: "Port to serve the site on.")
		var port = 8080

		/// Run the command.
		mutating func run() {
			print("TODO: \(sitePath) \(port)")
		}
	}
}
