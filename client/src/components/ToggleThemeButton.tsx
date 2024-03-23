import { BsSunFill, BsMoon } from 'solid-icons/bs';

import { Show, createEffect, createSignal } from 'solid-js';

const THEME_DARK = 'dark';
const THEME_LIGHT = 'light';

function getInitalTheme() {
    const systemDarkTheme = window.matchMedia(
        '(prefers-color-scheme: dark)'
    ).matches;
    const userTheme = localStorage.getItem('theme');
    if (userTheme === THEME_LIGHT || userTheme === THEME_DARK) {
        return userTheme;
    } else if (systemDarkTheme) {
        return THEME_DARK;
    }
    return THEME_LIGHT;
}

const ToggleThemeButton = () => {
    const [theme, setTheme] = createSignal(getInitalTheme());

    function toggle() {
        setTheme((prev) => (prev === THEME_DARK ? THEME_LIGHT : THEME_DARK));
    }

    createEffect(() => {
        if (theme() === THEME_DARK) {
            document.documentElement.classList.add('dark');
            localStorage.setItem('theme', THEME_DARK);
        } else {
            document.documentElement.classList.remove('dark');
            localStorage.setItem('theme', THEME_LIGHT);
        }
    });

    return (
        <Show
            when={theme() === THEME_DARK}
            fallback={
                <button onClick={toggle} aria-label="Switch to Light Mode">
                    <BsMoon class="w-6 h-6" />
                </button>
            }
        >
            <button onClick={toggle} aria-label="Switch to Dark Mode">
                <BsSunFill class="w-6 h-6" />
            </button>
        </Show>
    );
};
export default ToggleThemeButton;
