import Foundation
import Yams

/// Errors thrown when loading a site.
enum SiteError: Error, Equatable {
    case configFileInvalid(String)
    case configFileMissing
    case siteDirectoryInvalid
}

extension SiteError: CustomStringConvertible {
    public var description: String {
        switch self {
        case .configFileInvalid(let reason):
            return "Invalid config file: \(reason)"
        case .configFileMissing:
            return "Missing config file."
        case .siteDirectoryInvalid:
            return "Provided site path is not a readable directory."
        }
    }
}

/// Site config as read from the YAML file.
struct SiteConfig {
    let title: String
    let description: String?
    let image: String?
    let rootUrl: String?
    let data: [String:Any]
}

/// Site which can be built or served.
struct Site {
    let sourceDirectory: URL
    
    var config: SiteConfig {
        get throws {
            let configContents = try String(
                contentsOf: sourceDirectory.appendingPathComponent("config.yaml"),
                encoding: .utf8)
            let potentialConfig = try Yams.load(yaml: configContents) as? [String: Any]
            
            if potentialConfig == nil {
                throw SiteError.configFileInvalid("Unable to parse YAML file")
            }
            let config = potentialConfig!
            
            if config["title"] == nil {
                throw SiteError.configFileInvalid("Missing title")
            }
            
            return SiteConfig(
                title: config["title"] as! String,
                description: config["description"] as? String,
                image: config["image"] as? String,
                rootUrl: config["root_url"] as? String,
                data: config["data"] as? [String:Any] ?? [:])
        }
    }

    /// Initialise the site based on a source directory, validating that the necessary files exist.
    init(siteFromDirectory directory: URL) throws {
        sourceDirectory = directory

        var isDirectory = ObjCBool(false)
        let exists = FileManager.default.fileExists(atPath: sourceDirectory.path, isDirectory: &isDirectory)
        if !exists || !isDirectory.boolValue {
            throw SiteError.siteDirectoryInvalid
        }

         let configFilePath = sourceDirectory.appendingPathComponent("config.yaml")
         if !FileManager.default.fileExists(atPath: configFilePath.path) {
             throw SiteError.configFileMissing
         }
    }
}
