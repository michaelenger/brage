import Foundation
import Yams

/// A site's to be built with Brage.
public struct Site {
	private var sourceDirectory: URL

	public private(set) var config: Config

	/// Create a new Site based on a directory.
	///
	/// - Parameter fromDirectory: Path to the directory.
	public init(fromDirectory sourceDirectory: URL) throws {
		self.sourceDirectory = sourceDirectory

		self.config = try Config(fromFile: sourceDirectory.appendingPathComponent("config.yaml"))
	}
}

extension Site {
	/// Site configuration.
	public struct Config {
		var title: String
		var rootUrl: String
		var description: String?
		var image: String?
		var redirects: [String: String]
		var data: [String: Any]

		/// Create a new configuration from a file.
		///
		/// - Parameter fromFile: Path to the file.
		public init(fromFile file: URL) throws {
			let contents = try Data(contentsOf: file)
			let config = try Yams.load(yaml: String(decoding: contents, as: UTF8.self)) as! [String: Any]

			guard config.keys.contains("title") else {
				throw SiteConfigError.missingTitle
			}

			guard config.keys.contains("rootUrl") else {
				throw SiteConfigError.missingRootUrl
			}

			self.title = config["title"] as! String
			self.rootUrl = config["rootUrl"] as! String
			self.description = config["description"] as? String
			self.image = config["image"] as? String

			if let redirects = config["redirects"] as? [String:String] {
				self.redirects = redirects
			} else {
				self.redirects = [:]
			}

			if let data = config["data"] as? [String:Any] {
				self.data = data
			} else {
				self.data = [:]
			}
		}
	}

	public enum SiteConfigError: Error {
		case missingTitle
		case missingRootUrl
	}
}
