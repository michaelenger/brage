// swift-tools-version:5.5
// The swift-tools-version declares the minimum version of Swift required to build this package.

import PackageDescription

let package = Package(
    name: "Brage",
    dependencies: [
        .package(url: "https://github.com/apple/swift-argument-parser", from: "1.1.2"),
        .package(url: "https://github.com/jpsim/Yams", from: "5.0.0"),
    ],
    targets: [
        .executableTarget(name: "Brage", dependencies: [
            .product(name: "ArgumentParser", package: "swift-argument-parser"),
            "Yams",
        ]),
        .testTarget(
            name: "BrageTests",
            dependencies: ["Brage"]),
    ]
)
