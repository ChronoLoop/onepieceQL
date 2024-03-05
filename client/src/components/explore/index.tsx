import Sidebar, { SidebarHeader, SidebarItem } from '../../components/Sidebar';
import {
    Accessor,
    For,
    Match,
    ParentComponent,
    Setter,
    Switch,
    createEffect,
    onCleanup,
} from 'solid-js';
import { SetStoreFunction } from 'solid-js/store';
import { CharacterArgs } from '../../services/characters';
import { createIntersectionObserver } from '../../utils/observer';
import Spinner from '../../components/Spinner';
import { ExploreCharacterFilter } from './ExploreCharacterFilter';
import { DevilFruitArgs } from '../../services/fruits';
import ExploreDevilFruitFilter from './ExploreDevilFruitFilter';
import { CATEGORIES, Category } from '../../constants/category';

export const ExploreSidebarSection: ParentComponent = (props) => {
    return <li class="flex flex-col gap-2">{props.children}</li>;
};

export const ExploreSidebarSubSection: ParentComponent = (props) => {
    return <ul class="flex flex-col gap-1">{props.children}</ul>;
};

export const ExploreLoadSpinner = () => {
    return (
        <div class="flex justify-center mt-3 col-span-full">
            <Spinner />
        </div>
    );
};

type ExploreSidebarListMenuProps = {
    header: string;
    items: string[] | readonly string[];
    selectedItem?: string;
    onItemClick: (item: string | undefined) => void;
    disableNone?: boolean;
};
export const ExploreSidebarListMenu = (props: ExploreSidebarListMenuProps) => {
    return (
        <ExploreSidebarSection>
            <SidebarHeader>{props.header}</SidebarHeader>
            <ExploreSidebarSubSection>
                {!props.disableNone && (
                    <SidebarItem
                        onClick={() => props.onItemClick(undefined)}
                        selected={props.selectedItem == undefined}
                    >
                        None
                    </SidebarItem>
                )}
                <For each={props.items}>
                    {(item) => (
                        <SidebarItem
                            onClick={() => props.onItemClick(item)}
                            selected={props.selectedItem === item}
                        >
                            {item}
                        </SidebarItem>
                    )}
                </For>
            </ExploreSidebarSubSection>
        </ExploreSidebarSection>
    );
};

type ExploreSidebarProps = {
    isOpen: boolean;
    setIsOpen: Setter<boolean>;
    category: Accessor<Category>;
    setCategory: Setter<Category>;
    setCharacterFilter: SetStoreFunction<CharacterArgs['filter']>;
    characterFilter: CharacterArgs['filter'];
    devilFruitFilter: DevilFruitArgs['filter'];
    setDevilFruitFilter: SetStoreFunction<DevilFruitArgs['filter']>;
};

export const ExploreSidebar = (props: ExploreSidebarProps) => {
    return (
        <Sidebar isOpen={props.isOpen} setIsOpen={props.setIsOpen}>
            <ExploreSidebarSection>
                <ExploreSidebarListMenu
                    header="Category"
                    items={CATEGORIES}
                    selectedItem={props.category()}
                    onItemClick={props.setCategory}
                    disableNone={true}
                />
            </ExploreSidebarSection>
            <Switch>
                <Match when={props.category() === 'Character'}>
                    <ExploreCharacterFilter
                        setFilter={props.setCharacterFilter}
                        filter={props.characterFilter}
                    />
                </Match>
                <Match when={props.category() === 'Devil Fruit'}>
                    <ExploreDevilFruitFilter
                        setFilter={props.setDevilFruitFilter}
                        filter={props.devilFruitFilter}
                    />
                </Match>
            </Switch>
        </Sidebar>
    );
};

type ExploreLoadNextButtonProps = {
    onLoadNext: () => void;
};

export const ExploreLoadNextButton = (props: ExploreLoadNextButtonProps) => {
    let buttonRef: HTMLButtonElement | undefined;

    createEffect(() => {
        if (!buttonRef) return;

        const observer = createIntersectionObserver({
            rootMargin: '400px',
            onIntersect: props.onLoadNext,
        });
        observer.observe(buttonRef);

        onCleanup(() => {
            if (buttonRef) observer.unobserve(buttonRef);
            observer.disconnect();
        });
    });

    return (
        <button class="invisible" ref={buttonRef} disabled>
            Loading more...
        </button>
    );
};
