import ArgumentParser

extension Brage {
	struct Build: ParsableCommand {
		static let configuration = CommandConfiguration(
			abstract: "Build a site.")
		
		@Argument(help: "Path of the site configuration.")
		var sitePath: String
		
		@Option(name: [.short, .customLong("output")], help: "Path to output the site to.")
		var outputPath: String?
		
		@Flag(name: .shortAndLong, help: "Override the output assets directory, removing anything already in there.")
		var clean = false
		
		mutating func run() {
			if outputPath != nil {
				print("TODO: \(sitePath) \(outputPath!) \(clean)")
			} else {
				print("TODO: \(sitePath) \(clean)")
			}
		}
	}
}