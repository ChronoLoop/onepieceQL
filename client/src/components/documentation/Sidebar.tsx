import { For, ParentComponent, Setter, Show, createEffect } from 'solid-js';
import { A, AnchorProps, useLocation } from '@solidjs/router';
import { FaSolidAnglesRight } from 'solid-icons/fa';
import { CgClose } from 'solid-icons/cg';

type SidebarItem = {
    title: string;
    href: string;
};

const INTRODUCTION_SUBSECTION: SidebarItem[] = [
    { title: 'Contribute', href: '#contribute' },
    { title: 'Info', href: '#info' },
];

const CHARACTER_SUBSECTION: SidebarItem[] = [
    { title: 'Character Schema', href: '#character-schema' },
    { title: 'Get Single Character', href: '#single-character' },
    { title: 'Get All Characters', href: '#all-characters' },
    { title: 'Filter Characters', href: '#filter-characters' },
];

const DEVIL_FRUIT_SUBSECTION: SidebarItem[] = [
    { title: 'Devil Fruit Schema', href: '#devil-fruit-schema' },
    { title: 'Get Single Devil Fruit', href: '#single-devil-fruit' },
    { title: 'Get All Devil Fruits', href: '#all-devil-fruits' },
    { title: 'Filter Devil Fruits', href: '#filter-devil-fruits' },
];

const SidebarHeader = (props: AnchorProps) => {
    return <A class="font-bold text-xl" {...props} target="_self" />;
};

const SidebarItem = (props: AnchorProps) => {
    return <A class="font-normal text-base" {...props} target="_self" />;
};

type SidebarSectionProps = {
    header: {
        title: string;
        href: string;
    };
    sidebarItems: SidebarItem[];
};

const SidebarSection: ParentComponent<SidebarSectionProps> = (props) => {
    return (
        <li class="flex flex-col gap-3">
            <SidebarHeader href={props.header.href}>
                {props.header.title}
            </SidebarHeader>
            <SidebarSubSection sidebarItems={props.sidebarItems} />
        </li>
    );
};

const SidebarSubSection: ParentComponent<{ sidebarItems: SidebarItem[] }> = (
    props
) => {
    return (
        <ul class="flex flex-col gap-2">
            <For each={props.sidebarItems}>
                {(sidebarItem) => (
                    <li>
                        <SidebarItem href={sidebarItem.href}>
                            {sidebarItem.title}
                        </SidebarItem>
                    </li>
                )}
            </For>
        </ul>
    );
};

type SidebarProps = {
    isOpen: boolean;
    setIsOpen: Setter<boolean>;
};

const Sidebar = (props: SidebarProps) => {
    const location = useLocation();

    function toggle() {
        props.setIsOpen((prev) => !prev);
    }

    createEffect(() => {
        const media = window.matchMedia(`(max-width: 1023px)`);

        function closeIfNeeded() {
            if (!media.matches) {
                props.setIsOpen(true);
            } else {
                props.setIsOpen(false);
            }
        }
        closeIfNeeded();

        media.addEventListener('change', closeIfNeeded);
        return () => {
            media.removeEventListener('change', closeIfNeeded);
        };
    });

    createEffect(() => {
        if (location.hash) {
            const media = window.matchMedia(`(max-width: 1023px)`);
            if (media.matches) {
                props.setIsOpen(false);
            }
        }
    });

    return (
        <Show
            when={props.isOpen}
            fallback={
                <button
                    onClick={toggle}
                    class="fixed z-50 top-14 left-2 flex justify-center items-center w-10 h-10 rounded-full hover:bg-gray-600/50"
                >
                    <FaSolidAnglesRight class="w-5 h-5" />
                </button>
            }
        >
            <nav class="fixed z-50 top-12 left-0 w-60 shrink-0 h-full pl-6 p-9 border-r border-solid border-gray-300 dark:border-slate-700 bg-white dark:bg-slate-900 max-lg:w-full overflow-y-auto">
                <ul class="flex flex-col gap-7">
                    <SidebarSection
                        header={{
                            title: 'Introduction',
                            href: '#introduction',
                        }}
                        sidebarItems={INTRODUCTION_SUBSECTION}
                    />
                    <SidebarSection
                        header={{ title: 'Character', href: '#character' }}
                        sidebarItems={CHARACTER_SUBSECTION}
                    />
                    <SidebarSection
                        header={{ title: 'Devil Fruit', href: '#devil-fruit' }}
                        sidebarItems={DEVIL_FRUIT_SUBSECTION}
                    />
                </ul>
                <button
                    onClick={toggle}
                    class="absolute top-1 right-1 flex justify-center items-center w-10 h-10 rounded-full hover:bg-gray-600/50"
                >
                    <CgClose class="w-5 h-5" />
                </button>
            </nav>
        </Show>
    );
};

export default Sidebar;
