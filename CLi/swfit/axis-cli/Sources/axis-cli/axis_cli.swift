@main
public struct axis_cli {
    public private(set) var text = "Hello, World!"

    public static func main() {
        print(axis_cli().text)
    }
}
