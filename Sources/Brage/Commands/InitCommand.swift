import ArgumentParser

extension Brage {
    /// Command to initialise a site based on a template.
    struct InitCommand: ParsableCommand {
        static let configuration = CommandConfiguration(commandName: "init")

        @Argument(help: "Directory in which to initialise the site")
        var path: String?

        @Flag(help: "Overwrite anything in the path")
        var force: Bool = false

        func run() {
            print("INIT \(path) AND FORCE(\(force))")
        }
    }
}
