import XCTest
@testable import BrageCore

final class SiteTests: XCTestCase {
	private var targetDir = FileManager.default.temporaryDirectory.appendingPathComponent("BrageCoreTests", isDirectory: true)

	override func setUpWithError() throws {
		if FileManager.default.fileExists(atPath: targetDir.path) {
			try FileManager.default.removeItem(atPath: targetDir.path)
		}
		try FileManager.default.createDirectory(atPath: targetDir.path, withIntermediateDirectories: true, attributes: nil)

		let configContent =
		"""
		---
		title: My Site
		rootUrl: https://example.org
		description: This is my Brage site.
		image: dog.png
		redirects:
		  /example: https://example.org/
		data:
		  words:
		    - banana
		    - happy
		    - explosion
		"""
		FileManager.default.createFile(
			atPath: targetDir.appendingPathComponent("config.yaml").path,
			contents: Data(configContent.utf8),
			attributes: nil)
	}

	func testInitFromDirectory() throws {
		let site = try Site(fromDirectory: targetDir)

		XCTAssertEqual(site.config.title, "My Site")
		XCTAssertEqual(site.config.description, "This is my Brage site.")
		XCTAssertEqual(site.config.image, "dog.png")
		XCTAssertEqual(site.config.rootUrl, "https://example.org")
		XCTAssertEqual(site.config.redirects, ["/example": "https://example.org/"])
		XCTAssertEqual(site.config.data["words"] as! [String], ["banana", "happy", "explosion"])
	}
}
