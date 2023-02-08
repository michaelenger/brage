import ArgumentParser
import BrageCore
import Foundation

@main
struct Brage: ParsableCommand {
    static let configuration = CommandConfiguration(
        abstract: "Static site generator.",
        subcommands: [BuildCommand.self, NewCommand.self, ServeCommand.self])
}
