// swift-tools-version: 5.7
// The swift-tools-version declares the minimum version of Swift required to build this package.

import PackageDescription

let package = Package(
    name: "brage",
    dependencies: [
        .package(url: "https://github.com/apple/swift-argument-parser.git", from: "1.2.1"),
        .package(url: "https://github.com/jpsim/Yams.git", from: "5.0.4")
    ],
    targets: [
        // Targets are the basic building blocks of a package. A target can define a module or a test suite.
        // Targets can depend on other targets in this package, and on products in packages this package depends on.
        .target(
            name: "BrageCore",
            dependencies: [
                "Yams"
            ]),
        .executableTarget(
            name: "Brage",
            dependencies: [
                .product(name: "ArgumentParser", package: "swift-argument-parser"),
                "BrageCore",
            ]),
        .testTarget(
            name: "BrageTests",
            dependencies: ["BrageCore"]),
    ]
)
