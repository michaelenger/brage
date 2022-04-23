import Foundation

/// Errors thrown when loading a site.
enum SiteError: Error {
    case configFileMissing
    case siteDirectoryInvalid
}

extension SiteError: CustomStringConvertible {
    public var description: String {
        switch self {
        case .configFileMissing:
            return "Missing config file."
        case .siteDirectoryInvalid:
            return "Provided site path is not a readable directory."
        }
    }
}

/// Site which can be built or served.
struct Site {
    let sourceDirectory: URL

    /// Initialise the site based on a source directory, validating that the
    /// necessary files exist.
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
