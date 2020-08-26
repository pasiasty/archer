export function getElement(id: string): HTMLElement {
    var res = document.getElementById(id)
    if (res == null)
        throw new Error("Failed to obtain element: " + id)
    return res
}