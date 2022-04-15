import ArgumentParser

extension Brage {
    /// Command to build a site.
    struct BuildCommand: ParsableCommand {
        static let configuration = CommandConfiguration(commandName: "build")

        @Argument(help: "Site directory to build")
        var source: String?

        @Option(help: "Path to output the files to")
        var output: String?

        @Flag(help: "Delete any existing assets before building")
        var clean: Bool = false

        func run() {
            print("BUILD \(source) TO \(output) WITH CLEAN(\(clean))")
        }
    }
}
