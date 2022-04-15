import ArgumentParser

/// Main application command.
@main
struct Brage: ParsableCommand {
    static let configuration = CommandConfiguration(
        abstract: "Brage is the Norwegian name for the ancient norse god Bragi, the skaldic god of poetry.",
        version: "0.1.0",
        subcommands: [
            BuildCommand.self,
            InitCommand.self,
            ServeCommand.self
        ]
    )
}
