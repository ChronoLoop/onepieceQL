import MainContentWrapper from '../components/MainContentWrapper';
import Sidebar, { SidebarHeader, SidebarItem } from '../components/Sidebar';
import {
    For,
    ParentComponent,
    Setter,
    Show,
    createEffect,
    createResource,
    createSignal,
    onCleanup,
} from 'solid-js';
import { SetStoreFunction, createStore } from 'solid-js/store';
import { CharacterArgs, getCharacters } from '../services/characters';
import Input from '../components/Input';
import { createIntersectionObserver } from '../utils/observer';
import CharacterCard from '../components/CharacterCard';
import Spinner from '../components/Spinner';

const ORIGINS = [
    'North Blue',
    'East Blue',
    'South Blue',
    'West Blue',
    'Red Line',
    'Calm Belt',
    'Grand Line',
    'New World',
];

const MONTHS = [
    'January',
    'February',
    'March',
    'April',
    'May',
    'June',
    'July',
    'August',
    'September',
    'October',
    'November',
    'December',
];

const BLOOD_TYPES = ['X', 'S', 'XF', 'F'];

const ExploreSidebarSection: ParentComponent = (props) => {
    return <li class="flex flex-col gap-2">{props.children}</li>;
};

const ExploreSidebarSubSection: ParentComponent = (props) => {
    return <ul class="flex flex-col gap-1">{props.children}</ul>;
};

const ExploreLoadCharactersSpinner = () => {
    return (
        <div class="flex justify-center mt-3 col-span-full">
            <Spinner />
        </div>
    );
};

type ExploreSidebarListMenuProps = {
    header: string;
    items: string[];
    selectedItem?: string;
    onItemClick: (item: string | undefined) => void;
};
const ExploreSidebarListMenu = (props: ExploreSidebarListMenuProps) => {
    return (
        <ExploreSidebarSection>
            <SidebarHeader>{props.header}</SidebarHeader>
            <ExploreSidebarSubSection>
                <SidebarItem
                    onClick={() => props.onItemClick(undefined)}
                    selected={props.selectedItem == undefined}
                >
                    None
                </SidebarItem>
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
    setFilter: SetStoreFunction<CharacterArgs['filter']>;
    filter: CharacterArgs['filter'];
};

const ExploreSidebar = (props: ExploreSidebarProps) => {
    let timeoutId: NodeJS.Timeout | null = null;
    function onInputChange(val: string) {
        if (timeoutId) clearTimeout(timeoutId);

        timeoutId = setTimeout(() => props.setFilter('name', val), 500);
    }
    onCleanup(() => {
        if (timeoutId) clearTimeout(timeoutId);
    });
    return (
        <Sidebar isOpen={props.isOpen} setIsOpen={props.setIsOpen}>
            <ExploreSidebarSection>
                <SidebarHeader>Name</SidebarHeader>
                <Input
                    type="text"
                    placeholder="ex: Sanji"
                    name="name"
                    id="name"
                    onInput={(e) => onInputChange(e.target.value)}
                    value={props.filter.name}
                />
            </ExploreSidebarSection>
            <ExploreSidebarListMenu
                header="Origin"
                items={ORIGINS}
                selectedItem={props.filter.origin}
                onItemClick={(origin) => props.setFilter('origin', origin)}
            />
            <ExploreSidebarListMenu
                header="Birth Month"
                items={MONTHS}
                selectedItem={props.filter.birthday}
                onItemClick={(month) => props.setFilter('birthday', month)}
            />
            <ExploreSidebarListMenu
                header="Blood Type"
                items={BLOOD_TYPES}
                selectedItem={props.filter.bloodType}
                onItemClick={(bloodType) =>
                    props.setFilter('bloodType', bloodType)
                }
            />
        </Sidebar>
    );
};

type LoadNextCharactersButtonProps = {
    onLoadNext: () => void;
};

const LoadNextCharactersButton = (props: LoadNextCharactersButtonProps) => {
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

type LoadNextCharactersProps = {
    filter: CharacterArgs['filter'];
    nextPage: number;
};
const LoadNextCharacters = (props: LoadNextCharactersProps) => {
    const [loadNext, setLoadNext] = createSignal(false);
    const [characterData] = createResource(
        () => ({
            filter: {
                name: props.filter.name,
                origin: props.filter.origin,
                bloodType: props.filter.bloodType,
                birthday: props.filter.birthday,
            },
            loadNext: loadNext(),
        }),
        (args) => {
            return getCharacters(
                { page: props.nextPage, filter: args.filter },
                args.loadNext ? 'cache-first' : 'cache-only'
            );
        }
    );

    return (
        <Show
            when={loadNext()}
            fallback={
                <LoadNextCharactersButton
                    onLoadNext={() => setLoadNext(true)}
                />
            }
        >
            <Show
                when={!characterData.loading}
                fallback={<ExploreLoadCharactersSpinner />}
            >
                <For each={characterData()?.results}>
                    {(character) => <CharacterCard character={character} />}
                </For>
                <Show when={characterData()?.info.next}>
                    {(nextPage) => (
                        <LoadNextCharacters
                            nextPage={nextPage()}
                            filter={props.filter}
                        />
                    )}
                </Show>
            </Show>
        </Show>
    );
};

type ExploreInitialCharactersProps = {
    filter: CharacterArgs['filter'];
};

const ExploreInitialCharacters = (props: ExploreInitialCharactersProps) => {
    const [characterData] = createResource(
        () => ({
            filter: {
                name: props.filter.name,
                origin: props.filter.origin,
                bloodType: props.filter.bloodType,
                birthday: props.filter.birthday,
            },
        }),
        (args) => {
            return getCharacters({ page: 1, filter: args.filter });
        }
    );
    return (
        <Show
            when={!characterData.loading}
            fallback={<ExploreLoadCharactersSpinner />}
        >
            <For each={characterData()?.results}>
                {(character) => <CharacterCard character={character} />}
            </For>
            <Show when={characterData()?.info.next}>
                {(nextPage) => (
                    <LoadNextCharacters
                        nextPage={nextPage()}
                        filter={props.filter}
                    />
                )}
            </Show>
        </Show>
    );
};

const Explore = () => {
    const [isOpen, setIsOpen] = createSignal(true);
    const [filter, setFilter] = createStore<CharacterArgs['filter']>({
        name: '',
    });

    return (
        <>
            <ExploreSidebar
                isOpen={isOpen()}
                setIsOpen={setIsOpen}
                filter={filter}
                setFilter={setFilter}
            />
            <MainContentWrapper sidebarOpen={isOpen()}>
                <div class="grid grid-cols-[repeat(auto-fit,_minmax(20rem,_1fr))] auto-rows-auto gap-4">
                    <ExploreInitialCharacters filter={filter} />
                </div>
            </MainContentWrapper>
        </>
    );
};

export default Explore;
