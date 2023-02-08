import ArgumentParser

extension Brage {
	/// Command for creating a new site with boilerplate content.
	struct NewCommand: ParsableCommand {
		static let configuration = CommandConfiguration(
			commandName: "new",
			abstract: "Create a new boilerplate site in the specified location.")

		@Argument(help: "Where to create the site.")
		var targetPath: String

		@Flag(name: .shortAndLong, help: "Force the creation of the site contents, overwriting any existing files.")
		var force = false

		/// Run the command.
		mutating func run() {
			print("TODO: \(targetPath) \(force)")
		}
	}
}
