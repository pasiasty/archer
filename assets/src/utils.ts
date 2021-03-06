export function getElement(id: string): HTMLElement {
    var res = document.getElementById(id)
    if (res == null)
        throw new Error("Failed to obtain element: " + id)
    return res
}

export function setCookie(name: string, val: string) {
    const date = new Date();
    const value = val;
    // Set it expire in 7 days
    date.setTime(date.getTime() + (7 * 24 * 60 * 60 * 1000));
    // Set it
    document.cookie = name + "=" + value + "; expires=" + date.toUTCString() + "; path=/";
}

export function getCookie(name: string): string {
    const value = "; " + document.cookie;
    const parts = value.split("; " + name + "=");

    if (parts.length == 2) {
        var shifted = parts.pop()?.split(";").shift();
        return <string>shifted
    }
    return ""
}

export function deleteCookie(name: string) {
    const date = new Date();
    // Set it expire in -1 days
    date.setTime(date.getTime() + (-1 * 24 * 60 * 60 * 1000));
    // Set it
    document.cookie = name + "=; expires=" + date.toUTCString() + "; path=/";
}

export function copyToClipboard(val: string) {
    const selBox = document.createElement('textarea');
    selBox.style.position = 'fixed';
    selBox.style.left = '0';
    selBox.style.top = '0';
    selBox.style.opacity = '0';
    selBox.value = val;
    document.body.appendChild(selBox);
    selBox.focus();
    selBox.select();
    document.execCommand('copy');
    document.body.removeChild(selBox);
}

export function isHost(): boolean {
    return getCookie("is_host") == "true"
}

export function optimalViewport(): ex.ScreenDimension {
    var height = $(window).height()
    var width = $(window).width()

    if (height == null || width == null) {
        throw new Error("failed to get browser width or height")
    }

    if (width * 9 == height * 16)
        return { width: width, height: height }
    if (width * 9 > height * 16)
        return { width: height * 16 / 9, height: height }
    return { width: width, height: width * 9 / 16 }
}