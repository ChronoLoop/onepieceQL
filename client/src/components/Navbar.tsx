import BrandImage from '../assets/brand.png';
import { FaBrandsGithub } from 'solid-icons/fa';

import { A } from '@solidjs/router';
import { ParentComponent, mergeProps } from 'solid-js';
import ToggleThemeButton from './ToggleThemeButton';

const PLAYGROUND_LINK =
    process.env.NODE_ENV === 'development'
        ? 'http://localhost:5000/playground'
        : '/playground';

const NavbarListItem: ParentComponent<{ class?: string }> = (props) => {
    const mergedProps = mergeProps({ class: '' }, props);

    return <li class={`flex ${mergedProps.class}`}>{mergedProps.children}</li>;
};

const Navbar = () => {
    return (
        <header class="sticky top-0 bg-white dark:bg-slate-900 border-b border-solid border-gray-300 dark:border-slate-700 h-12 w-full z-50">
            <nav class="flex items-center px-5 h-full">
                <A href="/" class="mr-auto" aria-label="Home">
                    <img
                        src={BrandImage}
                        class="w-10 h-10 object-contain"
                        alt="One Piece Logo"
                    />
                </A>
                <div class="flex items-center gap-2 font-bold text-lg">
                    <ul class="flex gap-2 mx-2.5 items-center">
                        <NavbarListItem>
                            <A href={PLAYGROUND_LINK} target="blank">
                                Play
                            </A>
                        </NavbarListItem>
                        <NavbarListItem>
                            <A href="/">Docs</A>
                        </NavbarListItem>
                        <NavbarListItem>
                            <A
                                href="/explore"
                                class="bg-orange-700 rounded text-white px-3 py-0.5 inline-block"
                            >
                                Explore
                            </A>
                        </NavbarListItem>
                    </ul>
                    <ul class="flex gap-2.5 items-center">
                        <NavbarListItem>
                            <ToggleThemeButton />
                        </NavbarListItem>
                        <NavbarListItem>
                            <A
                                href="https://github.com/ChronoLoop/onepieceQL"
                                target="blank"
                                aria-label="Github Repository"
                            >
                                <FaBrandsGithub class="w-6 h-6" />
                            </A>
                        </NavbarListItem>
                    </ul>
                </div>
            </nav>
        </header>
    );
};

export default Navbar;
