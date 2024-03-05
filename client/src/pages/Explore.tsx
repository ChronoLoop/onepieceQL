import MainContentWrapper from '../components/MainContentWrapper';
import {
    For,
    Match,
    Show,
    Suspense,
    Switch,
    createEffect,
    createResource,
    createSignal,
    onCleanup,
} from 'solid-js';
import { createStore } from 'solid-js/store';
import { CharacterArgs, getCharacters } from '../services/characters';
import CharacterCard from '../components/CharacterCard';
import {
    ExploreLoadSpinner,
    ExploreLoadNextButton,
    ExploreSidebar,
} from '../components/explore';
import { DevilFruitArgs, getDevilFruits } from '../services/fruits';
import { Category } from '../constants/category';
import DevilFruitCard from '../components/DevilFruitCard';

type ExploreInitialDevilFruitsProps = {
    filter: DevilFruitArgs['filter'];
};

const ExploreInitialDevilFruits = (props: ExploreInitialDevilFruitsProps) => {
    const [debouncedFilter, setDebouncedFilter] = createStore<
        DevilFruitArgs['filter']
    >({
        name: props.filter.name,
        type: props.filter.type,
        previousOwner: props.filter.previousOwner,
        currentOwner: props.filter.currentOwner,
    });

    createEffect(() => {
        const { name, type, currentOwner, previousOwner } = props.filter;

        const timeoutId = setTimeout(() => {
            setDebouncedFilter({
                name,
                type,
                currentOwner,
                previousOwner,
            });
        }, 300);

        onCleanup(() => {
            if (timeoutId) clearTimeout(timeoutId);
        });
    });

    const [devilFruitsData] = createResource(
        () => ({
            filter: {
                name: debouncedFilter.name,
                type: debouncedFilter.type,
                previousOwner: debouncedFilter.previousOwner,
                currentOwner: debouncedFilter.currentOwner,
            },
        }),
        (args) => {
            return getDevilFruits({ page: 1, filter: args.filter });
        }
    );
    return (
        <>
            <For each={devilFruitsData()?.results}>
                {(devilFruit) => <DevilFruitCard devilFruit={devilFruit} />}
            </For>
            <Show when={devilFruitsData()?.info.next}>
                {(nextPage) => (
                    <LoadNextDevilFruits
                        nextPage={nextPage()}
                        filter={props.filter}
                    />
                )}
            </Show>
        </>
    );
};

type LoadNextDevilFruitsProps = {
    filter: DevilFruitArgs['filter'];
    nextPage: number;
};
const LoadNextDevilFruits = (props: LoadNextDevilFruitsProps) => {
    const [loadNext, setLoadNext] = createSignal(false);
    const [devilFruitsData] = createResource(
        () => ({
            filter: {
                name: props.filter.name,
                type: props.filter.type,
                previousOwner: props.filter.previousOwner,
                currentOwner: props.filter.currentOwner,
            },
            loadNext: loadNext(),
        }),
        (args) => {
            return getDevilFruits(
                { page: props.nextPage, filter: args.filter },
                args.loadNext ? 'cache-first' : 'cache-only'
            );
        }
    );

    return (
        <Show
            when={loadNext()}
            fallback={
                <ExploreLoadNextButton onLoadNext={() => setLoadNext(true)} />
            }
        >
            <Show
                when={!devilFruitsData.loading}
                fallback={<ExploreLoadSpinner />}
            >
                <For each={devilFruitsData()?.results}>
                    {(devilFruit) => <DevilFruitCard devilFruit={devilFruit} />}
                </For>
                <Show when={devilFruitsData()?.info.next}>
                    {(nextPage) => (
                        <LoadNextDevilFruits
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
    const [debouncedFilter, setDebouncedFilter] = createStore({
        name: props.filter.name,
        origin: props.filter.origin,
        bloodType: props.filter.bloodType,
        birthday: props.filter.birthday,
    });

    createEffect(() => {
        const { name, origin, bloodType, birthday } = props.filter;

        const timeoutId = setTimeout(() => {
            setDebouncedFilter({
                name,
                origin,
                bloodType,
                birthday,
            });
        }, 300);

        onCleanup(() => {
            if (timeoutId) clearTimeout(timeoutId);
        });
    });

    const [characterData] = createResource(
        () => ({
            filter: {
                name: debouncedFilter.name,
                origin: debouncedFilter.origin,
                bloodType: debouncedFilter.bloodType,
                birthday: debouncedFilter.birthday,
            },
        }),
        (args) => {
            return getCharacters({ page: 1, filter: args.filter });
        }
    );
    return (
        <>
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
        </>
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
                <ExploreLoadNextButton onLoadNext={() => setLoadNext(true)} />
            }
        >
            <Show
                when={!characterData.loading}
                fallback={<ExploreLoadSpinner />}
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

const Explore = () => {
    const [isOpen, setIsOpen] = createSignal(true);
    const [category, setCategory] = createSignal<Category>('Character');
    const [characterFilter, setCharacterFilter] = createStore<
        CharacterArgs['filter']
    >({
        name: '',
    });
    const [devilFruitFilter, setDevilFruitFilter] = createStore<
        DevilFruitArgs['filter']
    >({
        name: '',
    });

    return (
        <>
            <ExploreSidebar
                setCategory={setCategory}
                isOpen={isOpen()}
                setIsOpen={setIsOpen}
                characterFilter={characterFilter}
                setCharacterFilter={setCharacterFilter}
                category={category}
                devilFruitFilter={devilFruitFilter}
                setDevilFruitFilter={setDevilFruitFilter}
            />
            <MainContentWrapper sidebarOpen={isOpen()}>
                <div class="grid grid-cols-[repeat(auto-fit,_minmax(20rem,_1fr))] auto-rows-auto gap-4">
                    <Switch>
                        <Match when={category() === 'Character'}>
                            <Suspense fallback={<ExploreLoadSpinner />}>
                                <ExploreInitialCharacters
                                    filter={characterFilter}
                                />
                            </Suspense>
                        </Match>
                        <Match when={category() === 'Devil Fruit'}>
                            <Suspense fallback={<ExploreLoadSpinner />}>
                                <ExploreInitialDevilFruits
                                    filter={devilFruitFilter}
                                />
                            </Suspense>
                        </Match>
                    </Switch>
                </div>
            </MainContentWrapper>
        </>
    );
};

export default Explore;
