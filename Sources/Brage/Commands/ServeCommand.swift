import ArgumentParser
import Foundation

extension Brage {
    /// Command to serve a site on a simple webserver.
    struct ServeCommand: ParsableCommand {
        static let configuration = CommandConfiguration(commandName: "serve")

        @Argument(help: "Site directory to serve")
        var source: String = "."

        @Option(help: "Port to serve the site on")
        var port: Int = 8080

        func run() {
            let sourceDirectory = URL(fileURLWithPath: source, isDirectory: true)
            let site = try! Site(siteFromDirectory: sourceDirectory)

            print("SERVE \(site) ON \(port)")
        }
    }
}
