export function getElement(id: string): HTMLElement {
    var res = document.getElementById(id)
    if (res == null)
        throw new Error("Failed to obtain element: " + id)
    return res
}

// Given a cookie key `name`, returns the value of
// the cookie or "", if the key is not found.
function getCookie(name: string): string {
    const nameLenPlus = (name.length + 1);
    return document.cookie
        .split(';')
        .map(c => c.trim())
        .filter(cookie => {
            return cookie.substring(0, nameLenPlus) === `${name}=`;
        })
        .map(cookie => {
            return decodeURIComponent(cookie.substring(nameLenPlus));
        })[0] || "";
}