import ArgumentParser

extension Brage {
	/// Command for building a site.
	struct BuildCommand: ParsableCommand {
		static let configuration = CommandConfiguration(
			commandName: "build",
			abstract: "Build a site.")

		@Argument(help: "Path of the site configuration.")
		var sitePath: String

		@Option(name: [.short, .customLong("output")], help: "Path to output the site to.")
		var outputPath: String?

		@Flag(name: .shortAndLong, help: "Override the output assets directory, removing anything already in there.")
		var clean = false

		/// Run the command.
		mutating func run() {
			if outputPath != nil {
				print("TODO: \(sitePath) \(outputPath!) \(clean)")
			} else {
				print("TODO: \(sitePath) \(clean)")
			}
		}
	}
}
