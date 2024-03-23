import { A, useLocation } from '@solidjs/router';
import { CgClose } from 'solid-icons/cg';
import { FaSolidAnglesRight } from 'solid-icons/fa';
import {
    For,
    JSX,
    ParentComponent,
    Setter,
    Show,
    createEffect,
} from 'solid-js';

type SidebarItemCommonProps = {
    children: JSX.Element;
    selected?: boolean;
};
type SidebarItemConditionalProps =
    | {
          href?: string;
          onClick?: never;
      }
    | {
          href?: never;
          onClick?: JSX.EventHandlerUnion<HTMLHeadingElement, MouseEvent>;
      };

type SidebarItemProps = SidebarItemCommonProps & SidebarItemConditionalProps;
type SidebarHeaderProps = SidebarItemProps;

export const SidebarHeader = (props: SidebarHeaderProps) => {
    if (!props.href)
        return (
            <h3
                class="font-bold text-xl"
                classList={{
                    ['cursor-pointer hover:text-sky-600 dark:hover:text-sky-500']:
                        !!props.onClick && !props.selected,
                    [' text-orange-500']: props.selected,
                }}
                onClick={props.onClick}
            >
                {props.children}
            </h3>
        );
    return (
        <A class="font-bold text-xl" target="_self" href={props.href}>
            {props.children}
        </A>
    );
};

export const SidebarItem = (props: SidebarItemProps) => {
    if (!props.href)
        return (
            <h4
                class="font-normal text-lg"
                classList={{
                    ['cursor-pointer hover:text-sky-600 dark:hover:text-orange-500']:
                        !!props.onClick && !props.selected,
                    ['text-sky-500 dark:text-orange-400']: props.selected,
                }}
                onClick={props.onClick}
            >
                {props.children}
            </h4>
        );
    return (
        <A class="font-normal text-base" target="_self" href={props.href}>
            {props.children}
        </A>
    );
};

type SidebarItem = {
    href?: string;
    title: string;
};

type SidebarSubSectionProps = { sidebarItems: SidebarItem[] };

const SidebarSubSection = (props: SidebarSubSectionProps) => {
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

type SidebarSectionProps = {
    header: SidebarItem;
    sidebarItems: SidebarItem[];
};

export const SidebarSection = (props: SidebarSectionProps) => {
    return (
        <li class="flex flex-col gap-3">
            <SidebarHeader href={props.header.href}>
                {props.header.title}
            </SidebarHeader>
            <SidebarSubSection sidebarItems={props.sidebarItems} />
        </li>
    );
};

type SidebarProps = {
    isOpen: boolean;
    setIsOpen: Setter<boolean>;
};

const Sidebar: ParentComponent<SidebarProps> = (props) => {
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
                    class="fixed z-50 top-14 left-2 flex justify-center items-center w-10 h-10 rounded-full hover:bg-gray-500/25 dark:hover:bg-gray-600/50"
                    aria-label="Open Sidebar"
                >
                    <FaSolidAnglesRight class="w-5 h-5" />
                </button>
            }
        >
            {/* pl-6 p-9 */}
            <nav class="fixed z-40 top-0 left-0 w-60 shrink-0 h-full border-r pt-12 border-solid border-gray-300 dark:border-slate-700 bg-white dark:bg-slate-900 max-lg:w-full ">
                <div class="relative h-full w-full pl-6 p-9 overflow-y-auto">
                    <ul class="flex flex-col gap-7">{props.children}</ul>
                    <button
                        onClick={toggle}
                        class="absolute top-1 right-1 flex justify-center items-center w-10 h-10 rounded-full hover:bg-gray-500/25 dark:hover:bg-gray-600/50"
                        aria-label="Close Sidebar"
                    >
                        <CgClose class="w-5 h-5" />
                    </button>
                </div>
            </nav>
        </Show>
    );
};

export default Sidebar;
