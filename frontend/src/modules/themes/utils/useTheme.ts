import data from '../../../configs/themes/default/light.json';
const typedData: Record<string, any> = data;

function isUserPrefersDarkColorScheme(): boolean {
    return window.matchMedia("(prefers-color-scheme: dark)").matches
}

function replaceUrlPart(url: string, substring: string, newPart: string) {
    // Find the index of the substring
    const index = url.indexOf(substring);
    
    // If the substring is found, replace the part after it
    if (index !== -1) {
        // Get the part of the URL before and including the substring
        const beforeSubstring = url.slice(0, index + substring.length);
        return beforeSubstring + newPart; // Concatenate with the new part
    }
    
    // If substring is not found, return original URL
    return url;
}

export default async function useTheme(name?: string) {
    const themeName = name || "axolotl";
    // const colorScheme = isUserPrefersDarkColorScheme()? "dark": "light";

    // const response = await fetch(`../configs/themes/${themeName}/${colorScheme}.json`);
    // const data = await response.json();
    
    // // Set CSS variables
    // Object.keys(typedData).forEach(key => {
    //     document.documentElement.style.setProperty(`--${key}`, typedData[key]);
    // });

    // const themeLink = document.querySelector('link[href*="configs/themes/"]') as HTMLAnchorElement;
    // if (themeLink)
    //     themeLink.href = replaceUrlPart(themeLink.href, `configs/themes/`, `${themeName}.css`);
}