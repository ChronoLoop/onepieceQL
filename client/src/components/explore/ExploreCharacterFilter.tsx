import { SidebarHeader } from '../../components/Sidebar';
import { SetStoreFunction } from 'solid-js/store';
import { CharacterArgs } from '../../services/characters';
import Input from '../../components/Input';
import { ExploreSidebarListMenu, ExploreSidebarSection } from '.';

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

type ExploreCharacterFilterProps = {
    filter: CharacterArgs['filter'];
    setFilter: SetStoreFunction<CharacterArgs['filter']>;
};

export const ExploreCharacterFilter = (props: ExploreCharacterFilterProps) => {
    const onItemClick =
        (name: keyof CharacterArgs['filter']) => (item: string | undefined) =>
            props.setFilter(name, item);

    return (
        <>
            <ExploreSidebarSection>
                <SidebarHeader>Name</SidebarHeader>
                <Input
                    type="text"
                    placeholder="ex: Roronoa Zoro"
                    name="name"
                    id="name"
                    onInput={(e) => props.setFilter('name', e.target.value)}
                    value={props.filter.name}
                />
            </ExploreSidebarSection>
            <ExploreSidebarListMenu
                header="Origin"
                items={ORIGINS}
                selectedItem={props.filter.origin}
                onItemClick={onItemClick('origin')}
            />
            <ExploreSidebarListMenu
                header="Birth Month"
                items={MONTHS}
                selectedItem={props.filter.birthday}
                onItemClick={onItemClick('birthday')}
            />
            <ExploreSidebarListMenu
                header="Blood Type"
                items={BLOOD_TYPES}
                selectedItem={props.filter.bloodType}
                onItemClick={onItemClick('bloodType')}
            />
        </>
    );
};
