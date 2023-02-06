import ArgumentParser
import BrageCore

@main
struct Brage: ParsableCommand {
    static let configuration = CommandConfiguration(
        abstract: "Static site generator.",
        subcommands: [Build.self, Init.self, Serve.self])
}
