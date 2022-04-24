import XCTest
@testable import Brage

final class SiteTests: XCTestCase {
    let temporarySitePath = URL(fileURLWithPath: NSTemporaryDirectory()).appendingPathComponent("BrageSiteTests")

    override func setUpWithError() throws {
        try FileManager.default.createDirectory(
            at: temporarySitePath,
            withIntermediateDirectories: true)
        FileManager.default.createFile(
            atPath: temporarySitePath.appendingPathComponent("config.yaml").path,
            contents: """
            ---
            title: Test Site
            description: This is just a test.
            image: lol.png
            root_url: https://example.org
            data:
              one: 1
              two:
                - 2
                - 3
            """.data(using: .utf8))
        FileManager.default.createFile(
            atPath: temporarySitePath.appendingPathComponent("layout.html").path,
            contents: """
            <html>
            TODO
            </html>
            """.data(using: .utf8))
        try FileManager.default.createDirectory(
            at: temporarySitePath.appendingPathComponent("pages"),
            withIntermediateDirectories: false)
    }

    override func tearDownWithError() throws {
        if FileManager.default.fileExists(atPath: temporarySitePath.path) {
            try FileManager.default.removeItem(at: temporarySitePath)
        }
    }
    
    func testConfig() throws {
        let site = try Site(siteFromDirectory: temporarySitePath)
        
        let config = try site.config
        
        XCTAssertEqual(config.title, "Test Site")
        XCTAssertEqual(config.description, "This is just a test.")
        XCTAssertEqual(config.image, "lol.png")
        XCTAssertEqual(config.rootUrl, "https://example.org")
        XCTAssertEqual(config.data["one"] as! Int, 1)
        XCTAssertEqual(config.data["two"] as! [Int], [2, 3])
    }

    func testInit() throws {
        let site = try Site(siteFromDirectory: temporarySitePath)
        XCTAssertEqual(site.sourceDirectory, temporarySitePath)
    }

    func testInitMissingSite() throws {
        try FileManager.default.removeItem(at: temporarySitePath)

        XCTAssertThrowsError(try Site(siteFromDirectory: temporarySitePath)) { (errorThrown) in
            XCTAssertEqual(errorThrown as? SiteError, SiteError.siteDirectoryInvalid)
        }
    }
    
    func testInitMissingConfigFile() throws {
        try FileManager.default.removeItem(at: temporarySitePath.appendingPathComponent("config.yaml"))

        XCTAssertThrowsError(try Site(siteFromDirectory: temporarySitePath)) { (errorThrown) in
            XCTAssertEqual(errorThrown as? SiteError, SiteError.configFileMissing)
        }
    }
    
    func testInitMissingLayoutFile() throws {
        try FileManager.default.removeItem(at: temporarySitePath.appendingPathComponent("layout.html"))

        XCTAssertThrowsError(try Site(siteFromDirectory: temporarySitePath)) { (errorThrown) in
            XCTAssertEqual(errorThrown as? SiteError, SiteError.layoutFileMissing)
        }
    }
    
    func testInitMissingPagesDirectory() throws {
        try FileManager.default.removeItem(at: temporarySitePath.appendingPathComponent("pages"))

        XCTAssertThrowsError(try Site(siteFromDirectory: temporarySitePath)) { (errorThrown) in
            XCTAssertEqual(errorThrown as? SiteError, SiteError.pagesDirectoryMissing)
        }
    }
}
