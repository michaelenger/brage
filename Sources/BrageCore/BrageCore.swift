public struct BrageCore {
	public private(set) var text = "Hello, World!"
	
	public init() {}

	public func printText() {
		print(BrageCore().text)
	}
}
