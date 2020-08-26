export function getDiv(id: string): HTMLDivElement {
    var res = <HTMLDivElement> document.getElementById(id)
    if (res == null)
        throw new Error("Failed to obtain UI element")
    return res
}