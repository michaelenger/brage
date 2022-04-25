import Foundation

/// Errors thrown when building a site.
enum BuilderError: Error, Equatable {
    case invalidTargetDirectory(String)
}

extension BuilderError: CustomStringConvertible {
    public var description: String {
        switch self {
        case .invalidTargetDirectory(let reason):
            return "Invalid target path: \(reason)"
        }
    }
}

/// The builder which outputs the complete site to a location.
struct Builder {
    let site: Site
    
    func build(to targetDirectory: URL) throws {
        var isDirectory = ObjCBool(false)
        let exists = FileManager.default.fileExists(atPath: targetDirectory.path, isDirectory: &isDirectory)
        
        if exists && !isDirectory.boolValue {
            throw BuilderError.invalidTargetDirectory("Path is a file")
        }
        
        // Copy assets (and delete existing?)
        
        // Build pages
    }
}
