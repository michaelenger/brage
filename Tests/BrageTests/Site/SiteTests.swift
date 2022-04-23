import XCTest
@testable import Brage

final class SiteTests: XCTestCase {
    let temporarySitePath = URL(fileURLWithPath: NSTemporaryDirectory()).appendingPathComponent("BrageSiteTests")

    override func setUpWithError() throws {
        try FileManager.default.createDirectory(
            at: temporarySitePath,
            withIntermediateDirectories:true)
        FileManager.default.createFile(
            atPath: temporarySitePath.appendingPathComponent("config.yaml").path,
            contents: "test".data(using: .utf8))
    }

    override func tearDownWithError() throws {
        if FileManager.default.fileExists(atPath: temporarySitePath.path) {
            try FileManager.default.removeItem(at: temporarySitePath)
        }
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
}
