import ArgumentParser

extension Brage {
	struct Init: ParsableCommand {
		static let configuration = CommandConfiguration(
			abstract: "Create a new boilerplate site in the specified location.")
		
		@Argument(help: "Where to create the site.")
		var targetPath: String
		
		@Flag(name: .shortAndLong, help: "Force the creation of the site contents, overwriting any existing files.")
		var force = false
		
		mutating func run() {
			print("TODO: \(targetPath) \(force)")
		}
	}
}