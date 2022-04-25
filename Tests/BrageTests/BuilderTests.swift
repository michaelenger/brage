import XCTest
@testable import Brage

final class BuilderTests: XCTestCase {
    let temporarySitePath = URL(fileURLWithPath: NSTemporaryDirectory()).appendingPathComponent("BrageBuilderTests")

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
        FileManager.default.createFile(
            atPath: temporarySitePath.appendingPathComponent("pages/index.html").path,
            contents: "Text".data(using: .utf8))
    }

    override func tearDownWithError() throws {
        if FileManager.default.fileExists(atPath: temporarySitePath.path) {
            try FileManager.default.removeItem(at: temporarySitePath)
        }
    }
    
    func testBuildFileExists() throws {
        let tempFilePath = URL(fileURLWithPath: NSTemporaryDirectory()).appendingPathComponent("BrageBuilderTestsFile")
        FileManager.default.createFile(
            atPath: tempFilePath.path,
            contents: "Text".data(using: .utf8))

        let site = try Site(siteFromDirectory: temporarySitePath)
        let builder = Builder(site: site)
        
        XCTAssertThrowsError(try builder.build(to: tempFilePath)) { (errorThrown) in
            XCTAssertEqual(errorThrown as? BuilderError, BuilderError.invalidTargetDirectory("Path is a file"))
        }
        
    }
}
